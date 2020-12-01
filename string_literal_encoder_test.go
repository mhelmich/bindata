package bindata

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringLiteralEncoderBasic(t *testing.T) {
	buf := &bytes.Buffer{}
	w := newStringLiteralEncoder(buf)
	n, err := w.Write([]byte("Hello World!"))
	assert.Nil(t, err)
	assert.Equal(t, 12, n)
	assert.Equal(t, "\\x48\\x65\\x6c\\x6c\\x6f\\x20\\x57\\x6f\\x72\\x6c\\x64\\x21", buf.String())
}

func TestStringLiteralEncoderUnevenNumBytes(t *testing.T) {
	var n int
	var err error

	hexBuf := &bytes.Buffer{}
	enc := hex.NewEncoder(hexBuf)
	n, err = enc.Write([]byte("Hello World!"))
	assert.Nil(t, err)
	assert.Equal(t, 12, n)
	bites := hexBuf.Bytes()

	buf := &bytes.Buffer{}
	w := &hexStringLiteralEncoder{
		writer: buf,
	}
	n, err = w.Write(bites[:11])
	assert.Nil(t, err)
	assert.Equal(t, 11, n)
	n, err = w.Write(bites[11:])
	assert.Nil(t, err)
	assert.Equal(t, 13, n)
	assert.Equal(t, "\\x48\\x65\\x6c\\x6c\\x6f\\x20\\x57\\x6f\\x72\\x6c\\x64\\x21", buf.String())
}
