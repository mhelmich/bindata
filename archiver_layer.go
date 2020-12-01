package bindata

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

type archiverLayer interface {
	io.Closer
	WriteFile(string, os.FileInfo) error
}

type tarArchiver struct {
	tw         *tar.Writer
	compressor compressorLayer
}

func newArchiverLayer(a ArchiverLayer, c compressorLayer) archiverLayer {
	switch a {
	case Tar:
		return &tarArchiver{
			tw:         tar.NewWriter(c),
			compressor: c,
		}
	default:
		panic("unkown archiver layer " + a.String())
	}
}

func (ta *tarArchiver) WriteFile(path string, fileInfo os.FileInfo) error {
	header, err := tar.FileInfoHeader(fileInfo, path)
	if err != nil {
		return err
	}

	header.Name = path
	err = ta.tw.WriteHeader(header)
	if err != nil {
		return err
	}

	osFile, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}

	defer func() { _ = osFile.Close() }()
	_, err = io.Copy(ta.tw, osFile)
	if err != nil {
		return err
	}

	return nil
}

func (ta *tarArchiver) Close() error {
	err := ta.tw.Close()
	if err != nil {
		return err
	}

	return ta.compressor.Close()
}
