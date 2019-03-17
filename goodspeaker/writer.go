package goodspeaker

import (
	"encoding/binary"
	"errors"
	"io"
)

// Writer writes LG speaker messages.
type Writer struct {
	w   io.Writer
	aes *aesBlock
}

// NewWriter returns a new writer that writes LG speaker messages. The
// writer will only write encrypted messages if the encryption option is
// provided.
func NewWriter(wr io.Writer, opts ...Option) *Writer {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	w := &Writer{
		w:   wr,
		aes: o.aes,
	}
	if o.aes != nil {
		w.aes = o.aes
	}
	return w
}

// Write a message either in plain text or encrypted format (see Writer options).
func (w *Writer) Write(p []byte) (n int, err error) {
	if w.aes != nil {
		return w.writeEncrypted(p)
	}
	return w.writePlain(p)
}

// writeHeader writes a 5 byte header onto the underlying writer
// indicating type of message and length of message (excluding header).
func (w *Writer) writeHeader(format byte, size int) (n int, err error) {
	var header [5]byte
	header[0] = format
	binary.BigEndian.PutUint32(header[1:5], uint32(size))

	if n, err = w.w.Write(header[:]); err != nil {
		return n, errWrap(err, "error writing header")
	}
	if n != 5 {
		return n, errors.New("wrote incomplete header")
	}
	return n, nil
}

// writePlain writes p along with a plain text header onto the
// underlying writer.
func (w *Writer) writePlain(p []byte) (n int, err error) {
	_, err = w.writeHeader(headerPlainText, len(p))
	if err != nil {
		return 0, err
	}
	if n, err = w.w.Write(p); err != nil {
		return n, errWrap(err, "error writing message")
	}
	if n != len(p) {
		return n, errors.New("wrote incomplete message")
	}
	return n, nil
}

// writeEncrypted encrypts p and writes it along with an encryption
// header to the underlying writer.
func (w *Writer) writeEncrypted(p []byte) (n int, err error) {
	cbm := w.aes.newEncrypter()
	bs := cbm.BlockSize()

	size := len(p) + bs - len(p)%bs

	_, err = w.writeHeader(headerEncrypted, size)
	if err != nil {
		return 0, err
	}

	src := make([]byte, bs)
	dst := make([]byte, bs)

	for ; size > 0; size -= bs {
		limit := min(len(p), len(src))
		nn := copy(src, p[:limit])
		p = p[limit:]

		err := pad(src, nn) // Apply PKCS7 padding, if needed.
		if err != nil {
			return n, errWrap(err, "could not apply pkcs7 padding")
		}

		cbm.CryptBlocks(dst, src)
		_, err = w.w.Write(dst)
		if err != nil {
			return n, errWrap(err, "could not write encrypted block")
		}

		n += nn // Increment n with the length of plain text processed.
	}

	return n, nil
}

// pad applies PKCS7 padding to p.
func pad(p []byte, size int) error {
	n := len(p) - size
	if n < 0 || n > len(p) {
		return errors.New("bad input")
	}
	for i := 0; i < n; i++ {
		p[size+i] = byte(n)
	}
	return nil
}
