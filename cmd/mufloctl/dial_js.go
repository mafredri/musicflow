package main

import (
	"context"
	"io"

	net "github.com/mafredri/goodspeaker/js/net"
)

func dial(ctx context.Context, addr string) (io.ReadWriteCloser, error) {
	return net.Dial("tcp", addr)
}
