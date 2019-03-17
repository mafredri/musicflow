// +build !js

package main

import (
	"context"
	"io"
	"net"
)

func dial(ctx context.Context, addr string) (io.ReadWriteCloser, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "tcp", addr)
}
