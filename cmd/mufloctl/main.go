package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mafredri/goodspeaker"
	"github.com/mafredri/musicflow"
)

var (
	key = "4efgvbn m546Uy7kolKrftgbn =-0u&~"
	iv  = "54eRty@hkL,;/y9U"
)

func main() {
	host := flag.String("addr", "", "Host address or IP of the speaker")
	port := flag.Int("port", 9741, "Port of the speaker")
	flag.StringVar(&key, "key", key, "AES key for encryption")
	flag.StringVar(&iv, "iv", iv, "IV for encryption")
	doTest := flag.Bool("test", false, "Perform a communication test with the speaker")

	flag.Parse()

	if *host == "" {
		fmt.Print("error: speaker address must be provided\n\n")
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range ch {
			fmt.Printf("Received %s, exiting...\n", sig.String())
			cancel()
		}
	}()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	if *doTest {
		if err := testRun(ctx, addr, key, iv); err != nil && !errors.Is(err, context.Canceled) {
			panic(err)
		}
		return
	}

	if err := run(ctx, addr, key, iv); err != nil && !errors.Is(err, context.Canceled) {
		panic(err)
	}
}

func run(ctx context.Context, addr, key, iv string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Printf("Connecting to %s...", addr)

	var gsOpt []goodspeaker.Option
	if key != "" && iv != "" {
		aes, err := goodspeaker.WithAES([]byte(key), []byte(iv))
		if err != nil {
			return err
		}
		gsOpt = append(gsOpt, aes)
	}
	opt := []musicflow.DialOption{
		musicflow.WithGoodspeakerOption(gsOpt...),
		musicflow.WithLogger(log.New(os.Stderr, "[musicflow] ", log.Flags())),
	}

	c, err := musicflow.Dial(ctx, addr, opt...)
	if err != nil {
		return err
	}
	defer c.Close()

	c.OnBroadcast(func(message string, data []byte) {
		log.Printf("Broadcast: %s %s", message, data)
	})

	_, err = c.ProductInfo(ctx, time.Now(), true)
	if err != nil {
		return err
	}

	fmt.Printf("Reading stdin for JSON input (Ctrl+C to exit)...\nExample(s):\n\t%s\n\t%s\n\n",
		`{"data":{"fadetime":0,"vol":8},"msg":"VOLUME_SETTING"}`,
		`{"data":{"nightmode":false},"msg":"NIGHT_MODE_SET"}`,
	)
	go func() {
		defer cancel()

		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			req := musicflow.Request{}
			err = json.Unmarshal(s.Bytes(), &req)
			if err != nil {
				log.Println(err)
				continue
			}

			err = c.Send(ctx, req, nil)
			if err != nil {
				if err == io.EOF {
					log.Println("Connection lost")
					break
				}
				log.Println(err)
				return
			}
		}
		if err = s.Err(); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
