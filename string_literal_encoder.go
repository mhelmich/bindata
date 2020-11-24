package bindata

import (
	"encoding/hex"
	"io"
)

type stringLiteralEncoder interface {
	io.Writer
}

type hexStringLiteralEncoder struct {
	writer                  io.Writer
	shouldWriteTrailingByte bool
}

func newStringLiteralEncoder(writer io.Writer) stringLiteralEncoder {
	return hex.NewEncoder(&hexStringLiteralEncoder{
		writer:                  writer,
		shouldWriteTrailingByte: false,
	})
}

func (e *hexStringLiteralEncoder) Write(p []byte) (int, error) {
	written := 0

	if e.shouldWriteTrailingByte {
		_, err := e.writer.Write([]byte{p[0]})
		if err != nil {
			return written, err
		}
		p = p[1:]
		written++
	}

	idx := 0
	for ; idx < len(p)-1; idx += 2 {
		_, err := e.writer.Write([]byte{'\\', 'x', p[idx], p[idx+1]})
		if err != nil {
			return written, err
		}
		written += 2
	}

	if idx < len(p) {
		_, err := e.writer.Write([]byte{'\\', 'x', p[idx]})
		if err != nil {
			return written, err
		}
		written++
		e.shouldWriteTrailingByte = true
	} else {
		e.shouldWriteTrailingByte = false
	}

	return written, nil
}
