package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mafredri/goodspeaker/goodspeaker"
)

var (
	key = "4efgvbn m546Uy7kolKrftgbn =-0u&~"
	iv  = "54eRty@hkL,;/y9U"
)

func main() {
	addr := flag.String("addr", "", "Address to the speaker")
	port := flag.Int("port", 9741, "Port of the speaker")
	flag.StringVar(&key, "key", key, "AES key for encryption")
	flag.StringVar(&iv, "iv", iv, "IV for encryption")

	flag.Parse()

	if *addr == "" {
		fmt.Print("error: speaker address must be provided\n\n")
		flag.Usage()
		os.Exit(1)
	}

	ctx := context.Background()

	address := fmt.Sprintf("%s:%d", *addr, *port)
	if err := run(ctx, address, key, iv); err != nil {
		panic(err)
	}
}

func run(ctx context.Context, addr, key, iv string) error {
	log.Printf("Connecting to %s...", addr)

	aes, err := goodspeaker.WithAES([]byte(key), []byte(iv))
	if err != nil {
		return err
	}

	conn, err := dial(ctx, addr)
	if err != nil {
		return err
	}

	r := goodspeaker.NewReader(conn, aes)
	w := goodspeaker.NewWriter(conn, aes)

	go func() {
		dec := json.NewDecoder(r)
		for {
			var v interface{}
			if err := dec.Decode(&v); err != nil {
				if err == io.EOF {
					log.Println("Connection lost")
					return
				}
				log.Println(err)
				continue
			}

			log.Printf("=> %v", v)
		}
	}()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		in := s.Bytes()

		fmt.Println("Sending:", string(in))

		n, err := w.Write(in)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection lost")
				break
			}
			log.Println(err)
			continue
		}
		log.Printf("Wrote %d/%d", n, len(in))
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	return nil
}
