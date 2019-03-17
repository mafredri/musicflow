package net

import (
	"syscall/js"
)

func typedArrayToByte(v js.Value) []byte {
	// TODO(maf): Verify input is typed array, or panic.
	b := make([]byte, v.Length())
	tmp := js.TypedArrayOf(b)
	tmp.Call("set", v)
	tmp.Release()
	return b
}
