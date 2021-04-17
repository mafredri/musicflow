package musicflow

import (
	"context"

	"github.com/mafredri/goodspeaker"
	"github.com/mafredri/goodspeaker/js/net"
)

type connWrapper struct {
	c net.Conn
	*goodspeaker.Reader
	*goodspeaker.Writer
}

func (c *connWrapper) Conn() net.Conn { return c.c }
func (c *connWrapper) Close() error   { return c.c.Close() }

type dialOptions struct {
	addr   string
	gsOpts []goodspeaker.Option
	logger Logger
}

// A DialOption sets custom options for Dial.
type DialOption func(*dialOptions)

// WithGoodspeakerOption sets the option(s) passed to the
// goodspeaker package.
func WithGoodspeakerOption(opt ...goodspeaker.Option) DialOption {
	return func(o *dialOptions) {
		o.gsOpts = append(o.gsOpts, opt...)
	}
}

// Dial connects to the Music Flow speaker and returns a Client that
// implements the API.
func Dial(ctx context.Context, addr string, opts ...DialOption) (*Client, error) {
	o := dialOptions{addr: addr}
	for _, opt := range opts {
		opt(&o)
	}
	conn, err := dial(ctx, o)
	if err != nil {
		return nil, err
	}
	return newClient(conn, o), nil
}

func dial(ctx context.Context, o dialOptions) (*connWrapper, error) {
	conn, err := goodspeaker.Dial(ctx, o.addr)
	if err != nil {
		return nil, err
	}
	return &connWrapper{
		c:      conn,
		Reader: goodspeaker.NewReader(conn, o.gsOpts...),
		Writer: goodspeaker.NewWriter(conn, o.gsOpts...),
	}, nil
}
