package musicflow

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/crc64"
	"io"
	"net"
	"sync"
	"time"

	errors "golang.org/x/xerrors"

	"github.com/mafredri/musicflow/api"
)

// makeClientID returns a clientID which looks similar in structure to
// the one used on Android ("and"-prefix followed by 16 hexadecimals).
func makeClientID(prefix, hashpart string) string {
	tab := crc64.MakeTable(crc64.ECMA)
	return fmt.Sprintf("%s%x", prefix, crc64.Checksum([]byte(hashpart), tab))
}

var clientID = makeClientID("gsm", "goodspeaker/musicflow v0.0.0")

// Client represents a Music Flow Player client.
type Client struct {
	conn io.ReadWriteCloser
	o    dialOptions

	waitC chan *waitFor // Concurrency guard, not buffered.
	recvC chan Response

	mu        sync.RWMutex
	broadcast func(string, []byte)
}

// NewClient returns a new Music Flow Player client that uses the
// provided connection. Use Dial for more options.
func NewClient(conn io.ReadWriteCloser) *Client {
	return newClient(conn, dialOptions{})
}

func newClient(conn io.ReadWriteCloser, o dialOptions) *Client {
	if o.logger == nil {
		o.logger = noopLogger{}
	}
	c := &Client{
		conn:  conn,
		o:     o,
		waitC: make(chan *waitFor),
		recvC: make(chan Response, 1),
	}
	go c.recv()
	go c.read()

	return c
}

func (c *Client) log() Logger {
	return c.o.logger
}

type waitFor struct {
	ctx     context.Context
	message string
	result  string
	respC   chan Response
}

func (w *waitFor) init(ctx context.Context, message string) {
	w.ctx = ctx
	w.respC = make(chan Response, 1)
	if w.message == "" {
		// By default, we expect the response to be same as request.
		w.message = message
		w.result = "OK"
	}
}

func (w *waitFor) respond(r Response) bool {
	if w.message == r.Message || r.Message == api.MessageParsingError {
		w.respC <- r
		return true
	}
	return false
}

type sendOptions struct {
	wait waitFor
}

type SendOption func(*sendOptions)

// WaitFor allows Send to wait for a different response
// message than the request. Result is usually be "OK",
// or empty for events. If the response result differs
// from this value, Send will return an error.
func WaitFor(message string, result string) SendOption {
	return func(o *sendOptions) {
		o.wait.message = message
		o.wait.result = result
	}
}

// Send a request to the Music Flow device.
func (c *Client) Send(ctx context.Context, req Request, reply interface{}, opts ...SendOption) error {
	// Clean up the sent JSON, ignore "data" key when request has no
	// additional parameters.
	if z, ok := req.Data.(interface{ IsZero() bool }); ok && z.IsZero() {
		req.Data = nil
	}

	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Add a newline to try to circumvent potential issue in
	// firmware. Not sure what the cause is but sometimes the
	// soundbar stops responding after using the JSON API.
	b = append(b, '\n')

	o := sendOptions{}
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	o.wait.init(ctx, req.Message)
	errC := make(chan error, 1)

	select {
	case c.waitC <- &o.wait: // Concurrency guard.
	case <-ctx.Done():
		return ctx.Err()
	}

	go func() {
		// Avoid blocking for a long time if the connection disappeared.
		if conn, ok := c.conn.(interface{ Conn() net.Conn }); ok {
			_ = conn.Conn().SetWriteDeadline(time.Now().Add(10 * time.Second))
		}

		c.log().Printf("<= %s", b)

		_, err = c.conn.Write(b)
		if err != nil {
			errC <- errors.Errorf("Send: write failed: %w", err)
			return
		}

		// Disable timeout.
		if conn, ok := c.conn.(interface{ Conn() net.Conn }); ok {
			_ = conn.Conn().SetWriteDeadline(time.Time{})
		}
	}()

	var resp Response
	select {
	case resp = <-o.wait.respC:
	case err = <-errC:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}

	switch {
	case resp.Message == api.MessageParsingError:
		return errors.New("Send: player could not parse the request")
	case resp.Result != o.wait.result:
		return errors.Errorf("Send: player returned unexpected result: %q != %q", o.wait.result, resp.Result)
	}

	if reply == nil {
		return nil
	}
	err = json.Unmarshal([]byte(resp.Data), reply)
	if err != nil {
		return errors.Errorf("unmarshal %s reply into %T failed: %w", req.Message, reply, err)
	}

	return nil
}

func (c *Client) read() {
	dec := json.NewDecoder(c.conn)
	for {
		var r Response
		err := dec.Decode(&r)
		if err != nil {
			defer c.close(err)

			if errors.Is(err, io.EOF) {
				c.log().Printf("Connection lost")
				return
			}
			c.log().Printf("%+v", err)
			return
		}
		if _, ok := c.log().(noopLogger); !ok {
			b, _ := json.Marshal(r)
			c.log().Printf("=> %s", string(b))
		}
		c.recvC <- r
	}
}

func (c *Client) recv() {
	var wait *waitFor
recvLoop:
	for {
		var resp Response
		if wait == nil {
			select {
			case wait = <-c.waitC:
				goto recvLoop
			case resp = <-c.recvC:
				// Broadcast.
			}
		} else {
			select {
			case <-wait.ctx.Done():
				wait = nil
				goto recvLoop
			case resp = <-c.recvC:
				// Send response or broadcast.
			}

			if wait.respond(resp) {
				wait = nil
				goto recvLoop
			}
		}

		// No wait pending, forward response broadcast.
		c.mu.RLock()
		if c.broadcast != nil {
			c.broadcast(resp.Message, resp.Data)
		}
		c.mu.RUnlock()
	}
}

func (c *Client) Close() error {
	return c.close(nil)
}

func (c *Client) close(err error) error {
	// TODO: Close recv goroutine.
	// TODO: Save the error that closed the connection.
	_ = err

	err = c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
