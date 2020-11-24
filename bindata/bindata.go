package bindata

//
// Table of Contents
// test-fixtures/tree1/dir1/dir11/f111.txt
// test-fixtures/tree1/dir1/f11.txt
// test-fixtures/tree1/dir2/f21.txt
// test-fixtures/tree2/dir1/f1.txt
//

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type _binDataTreeElement struct {
	Name          string
	RemainingPath string
	Getter        func() string
	Children      map[string]*_binDataTreeElement
}

var _binDataTree = map[string]*_binDataTreeElement{
	"test-fixtures": {
		Name: "test-fixtures",
		Children: map[string]*_binDataTreeElement{
			"tree1": {
				Name: "tree1",
				Children: map[string]*_binDataTreeElement{
					"dir1": {
						Name: "dir1",
						Children: map[string]*_binDataTreeElement{
							"dir11": {
								Name: "dir11",
								Children: map[string]*_binDataTreeElement{
									"f111.txt": {
										Name:          "f111.txt",
										Getter:        _binDataGetVarArchive,
										RemainingPath: "test-fixtures/tree1/dir1/dir11/f111.txt",
									},
								},
							},
							"f11.txt": {
								Name:          "f11.txt",
								Getter:        _binDataGetVarArchive,
								RemainingPath: "test-fixtures/tree1/dir1/f11.txt",
							},
						},
					},
					"dir2": {
						Name: "dir2",
						Children: map[string]*_binDataTreeElement{
							"f21.txt": {
								Name:          "f21.txt",
								Getter:        _binDataGetVarArchive,
								RemainingPath: "test-fixtures/tree1/dir2/f21.txt",
							},
						},
					},
				},
			},
			"tree2": {
				Name: "tree2",
				Children: map[string]*_binDataTreeElement{
					"dir1": {
						Name: "dir1",
						Children: map[string]*_binDataTreeElement{
							"f1.txt": {
								Name:          "f1.txt",
								Getter:        _binDataGetVarArchive,
								RemainingPath: "test-fixtures/tree2/dir1/f1.txt",
							},
						},
					},
				},
			},
		},
	},
}

var (
	_bindataPathSeparator = fmt.Sprintf("%c", os.PathSeparator)

	_binDataAllFilesNames = []string{
		"test-fixtures/tree1/dir1/dir11/f111.txt",
		"test-fixtures/tree1/dir1/f11.txt",
		"test-fixtures/tree1/dir2/f21.txt",
		"test-fixtures/tree2/dir1/f1.txt",
	}
)

func _binDataGetVarArchive() string {
	return _binDataVarArchive
}

var _binDataVarArchive = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd5\x61\xaa\x83\x30\x0c\xc0\xf1\x1e\xe5\x5d\xe0\xbd\x26\xb1\xb6\xe7\x91\xb7\x06\xfc\x30\x06\x6d\x04\x8f\x3f\xa6\x6e\xb0\x0d\x9d\xa0\x28\xc3\xfc\x3e\x18\x44\x41\x45\xfe\x44\x62\x96\x5f\xae\x5b\x69\x52\xcc\x56\x52\x8c\x68\x4f\x75\xea\x0f\x68\x19\x11\xff\xa4\x15\xb3\x04\x00\x80\x77\xae\x9b\xc1\x97\xdd\x04\xea\xcf\x7b\xce\x60\x11\xca\x00\x01\x4b\xe7\x0d\x10\x10\xa2\xf9\x81\x45\x4f\x9d\xa9\xc9\x52\x25\x03\x70\xae\xd2\xff\x65\xe2\xbe\x2c\x15\xf3\xc4\xf5\xe1\x4b\x1e\xf3\x4b\xdc\x7e\xf0\xde\xef\xa0\xf6\x33\xda\x3f\xaf\x10\xfe\xe0\x73\xff\xc5\x73\xff\x18\x08\x82\xf6\xbf\x05\xd6\xfc\x0f\x6d\xa4\x7f\xb2\x4c\xbb\xf6\xaf\xfb\x7f\x1b\x4c\xda\xff\x91\xbd\xf7\x4f\xf7\xfd\xbf\x56\xfe\x33\xfa\xa7\xd7\xfe\x91\xbc\xf6\xbf\x05\xd6\xfc\x95\x52\xea\x90\xae\x01\x00\x00\xff\xff\xde\x03\x21\xcb\x00\x14\x00\x00"

// FileNames returns a list of all files in this bindata file.
func FileNames() []string {
	return _binDataAllFilesNames
}

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
func ReadFile(path string) ([]byte, error) {
	getter, remainingPath := _binDataFindPathInTree(path)
	return _binDataReadFile(remainingPath, getter)
}

var (
	_binDataInMemoryFileCache      = map[string][]byte{}
	_binDataInMemoryFileCacheMutex = &sync.Mutex{}
)

func _binDataReadFile(name string, getter func() string) ([]byte, error) {
	if name == "" || getter == nil {
		return nil, os.ErrNotExist
	}

	data, ok := _binDataInMemoryFileCache[name]
	if ok {
		return data, nil
	}

	_binDataInMemoryFileCacheMutex.Lock()
	defer _binDataInMemoryFileCacheMutex.Unlock()

	data, ok = _binDataInMemoryFileCache[name]
	if ok {
		return data, nil
	}

	gz, err := gzip.NewReader(strings.NewReader(getter()))
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", name, err.Error())
	}

	defer func() { _ = gz.Close() }()
	tr := tar.NewReader(gz)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		var buf bytes.Buffer
		_, err = io.CopyN(&buf, tr, header.Size)
		if err != nil {
			return nil, fmt.Errorf("error decompressing file %s: %s", name, err.Error())
		}

		_binDataInMemoryFileCache[header.Name] = buf.Bytes()
	}

	data, ok = _binDataInMemoryFileCache[name]
	if !ok {
		return nil, fmt.Errorf("can't find file %s", name)
	}

	return data, nil
}

func _binDataFindPathInTree(path string) (func() string, string) {
	parts := strings.Split(path, _bindataPathSeparator)
	if strings.HasPrefix(path, _bindataPathSeparator) {
		parts = parts[1:]
	}

	m := _binDataTree
	var te *_binDataTreeElement
	var ok bool
	for _, part := range parts {
		te, ok = m[part]
		if !ok {
			return nil, ""
		}
		m = te.Children
	}

	return te.Getter, te.RemainingPath
}
