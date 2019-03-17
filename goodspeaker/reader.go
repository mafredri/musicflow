package goodspeaker

import (
	"crypto/cipher"
	"encoding/binary"
	"io"

	errors "golang.org/x/xerrors"
)

// Reader reads LG speaker messages.
type Reader struct {
	r    io.Reader
	left int // Remaining bytes.

	aes   *aesBlock
	cbm   cipher.BlockMode // Set if message is encrypted.
	buf   []byte           // Decrypt buffer.
	dleft int              // Current length of dbuf.
	dpos  int              // Reading index of dbuf.
}

// NewReader returns a new reader that reads LG speaker messages. The
// reader can read both plain text messages and decrypt encrypted
// messages if the encryption option is provided.
func NewReader(rd io.Reader, opts ...Option) *Reader {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	r := &Reader{
		r:   rd,
		aes: o.aes,
	}
	if o.aes != nil {
		bs := r.aes.BlockSize()
		r.aes = o.aes
		r.buf = make([]byte, bs)
	}
	return r
}

// Read reads the next message or blocks until one is available.
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.left == 0 {
		err = r.parseHeader()
		if err != nil {
			return n, err
		}
	}
	if r.cbm != nil {
		return r.readEncrypted(p)
	}
	return r.readPlain(p)
}

// parseHeader reads the first 5 bytes och the incoming message where the first
// byte represents the message type (plain text or encrypted) and the following
// four bytes the message length (in bytes).
func (r *Reader) parseHeader() error {
	var header [5]byte
	n, err := r.r.Read(header[:])
	if err != nil {
		return errWrap(err, "error reading header")
	}
	if n != 5 {
		return errors.Errorf("invalid header length %d/5", n)
	}

	r.left = int(binary.BigEndian.Uint32(header[1:5]))
	r.cbm = nil
	if (header[0] & 0xF0) == 0x10 {
		r.cbm = r.aes.newDecrypter()
	}

	return nil
}

// readPlain simply calls Read on the reader and tracks message length.
func (r *Reader) readPlain(p []byte) (n int, err error) {
	limit := min(len(p), r.left)
	n, err = r.r.Read(p[:limit])
	if err != nil {
		return n, err
	}
	r.left -= n
	return n, nil
}

// readEncrypted uses two temporary buffers to decrypt the message one block at
// a time, the result is written to p as a typical reader would do.
func (r *Reader) readEncrypted(p []byte) (n int, err error) {
	// Left over from previous read.
	if r.dpos > 0 && r.dpos < r.dleft {
		limit := min(len(p), r.dleft)
		n = copy(p, r.buf[r.dpos:limit])
		r.dpos += n

		p = p[n:]
	}

	bs := r.cbm.BlockSize()
	buf := make([]byte, bs)

	// Deplete as long as there is space.
	for len(p) > 0 && r.left > 0 {
		// Read the next block.
		rn, err := r.r.Read(buf)
		if err != nil {
			return 0, err
		}
		if rn != len(buf) {
			return 0, errors.New("decryption failed, incomplete read")
		}

		r.left -= rn                  // Update message length.
		r.cbm.CryptBlocks(r.buf, buf) // Decrypt.
		r.dleft = len(r.buf)          // Reset length of decrypt buffer.
		r.dpos = 0                    // Reset reading index.

		if r.left == 0 {
			// Remove potential PKCS7 padding on the last block.
			pl, err := padlen(r.buf)
			if err != nil {
				return 0, err
			}
			r.dleft -= pl
		}

		limit := min(len(p), r.dleft)     // Copy at most len(p) or dbuf bytes.
		rn = copy(p, r.buf[r.dpos:limit]) // Update number of bytes read.
		r.dpos += rn                      // Update reading index.

		p = p[rn:] // Prepare p for next block.
		n += rn    // Increment total bytes read.
	}

	return n, nil
}

// padlen returns the length of the (PKCS7) padding.
func padlen(p []byte) (int, error) {
	c := p[len(p)-1]
	n := int(c)
	if n > len(p) {
		return 0, errors.New("bad padding")
	}
	for i := 0; i < n; i++ {
		if p[len(p)-1-i] != c {
			return 0, errors.New("padding is corrupt")
		}
	}
	return n, nil
}
