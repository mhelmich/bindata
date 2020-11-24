package bindata

import (
	"compress/gzip"
	"io"
)

type compressorLayer interface {
	io.Writer
	io.Closer
}

type bzCompressorLayer struct {
	writer *gzip.Writer
}

func newCompressorLayer(c CompressorLayer, out io.Writer) compressorLayer {
	switch c {
	case Bz:
		return &bzCompressorLayer{
			writer: gzip.NewWriter(out),
		}
	default:
		panic("unknown compressor layer " + c.String())
	}
}

func (c *bzCompressorLayer) Write(p []byte) (int, error) {
	return c.writer.Write(p)
}

func (c *bzCompressorLayer) Close() error {
	return c.writer.Close()
}
