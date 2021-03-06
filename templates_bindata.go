package bindata

//
// Table of Contents
// templates/root.tmpl
// templates/tarbz.tmpl
// templates/tree.tmpl
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
	"templates": {
		Name: "templates",
		Children: map[string]*_binDataTreeElement{
			"root.tmpl": {
				Name:          "root.tmpl",
				Getter:        _binDataGetVarArchive,
				RemainingPath: "templates/root.tmpl",
			},
			"tarbz.tmpl": {
				Name:          "tarbz.tmpl",
				Getter:        _binDataGetVarArchive,
				RemainingPath: "templates/tarbz.tmpl",
			},
			"tree.tmpl": {
				Name:          "tree.tmpl",
				Getter:        _binDataGetVarArchive,
				RemainingPath: "templates/tree.tmpl",
			},
		},
	},
}

var (
	_bindataPathSeparator = fmt.Sprintf("%c", os.PathSeparator)

	_binDataAllFilesNames = []string{
		"templates/root.tmpl",
		"templates/tarbz.tmpl",
		"templates/tree.tmpl",
	}
)

func _binDataGetVarArchive() string {
	return _binDataVarArchive
}

var _binDataVarArchive = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x6e\xdb\x3a\x12\xce\xad\xf8\x14\x53\x21\xed\x4a\x0b\x45\x96\xe3\xfc\x00\x01\x7c\xd1\x66\xd3\x6e\x81\x6d\xb6\x68\xba\x7b\x13\x04\x01\x2d\x8f\x6c\xc1\x92\x68\x90\x74\x9a\x44\xd5\xbb\x2f\x86\x94\x64\x29\x71\x92\x2e\xd0\xd3\x9e\x03\x98\x17\x96\x44\xce\xff\x90\xdf\x8c\xa9\x31\x5f\x66\x5c\xa3\x1a\x48\x21\x74\xa8\xf3\x65\xb6\xf3\xb3\x47\x14\x45\xd1\xd1\xc1\x81\x79\x1e\x1f\x1d\x9a\x67\xb4\x6f\xbf\xe9\x6d\xb8\x3f\xda\x19\x8e\x8e\x8f\xa2\x68\x34\x3c\x88\x0e\x77\xa2\xe1\xe1\xfe\xf0\x68\x07\xa2\x9f\x6e\xc9\x86\xb1\x52\x9a\xcb\x9d\x28\xca\xb9\x8c\xc5\x33\x74\x4a\xf3\x24\x79\x66\xbd\x76\xa6\x7d\xfe\x45\xc6\x92\xc7\x0b\x3e\x43\x28\x4b\x08\x3f\xdb\xf7\x73\x9e\x23\x54\x15\x63\x83\x01\x1b\x0c\xe0\x2b\x9f\x64\x08\x22\x81\x53\x51\x68\x2c\xb4\x62\x65\xb9\x07\x92\x17\x33\x84\xf0\x7d\x9a\x19\x7a\x45\x0c\x83\x81\x11\x43\xaf\x44\x82\xc5\xd4\xce\x32\x96\xe6\x4b\x21\x35\x78\xcc\x71\xb9\x8c\xe7\xe9\x0d\x0e\x34\x97\x2e\x73\xdc\xc9\x9d\x46\x45\x2f\xb1\xc8\x97\x12\x95\x1a\xcc\xee\xd3\x25\x4d\x24\xb9\xa6\x47\x2a\xe8\x57\x18\x1a\xa5\x65\x5a\xcc\xec\xeb\x5d\x11\xbb\xcc\x67\xac\x2c\xa1\xd9\xc3\xe0\x6a\x89\xe8\x42\xf8\x55\xa2\xf5\xe0\x86\x4b\xf0\x18\x00\xc0\xf5\x24\x2d\xa6\x5c\xf3\xcf\x5c\xcf\x2f\x70\xc9\x25\xd7\x42\xc2\x18\x92\x5c\x87\x17\x4b\x99\x16\x3a\xf1\xdc\xd7\xb1\x1b\x80\x50\x61\x8f\xc8\x67\xad\x80\x7f\x70\xcd\xdf\x66\x19\x39\xad\xac\xd7\x63\xb8\xbc\xb2\x56\x95\x86\xea\xa9\xd0\x40\x3d\xdc\x3a\x40\x6e\xd0\x92\xd7\x61\xa2\xcf\xaa\x76\xa8\x23\x41\xc1\x5e\x55\xb1\x64\x55\xc4\xad\x05\x1f\x50\xff\x97\x4b\x12\xf4\x41\xd4\xb9\xf2\x7c\xb0\x56\x80\x35\x43\xa2\x5e\xc9\xa2\xe5\x78\x48\xce\xea\xd0\x3c\xb5\x0e\x63\x6b\x68\x9d\x71\xb2\xb7\x9b\x51\x4a\xf4\xda\x3b\xab\x4b\x01\x87\x2c\x55\x9a\x36\x0a\xcf\x32\x48\x8c\xed\x69\x01\x7a\x9e\x2a\xa8\x83\x6f\x66\x43\xeb\x4d\x2b\xc0\xf3\xdb\x18\x6e\xb6\xbe\x17\x71\x66\xd5\x7f\x41\x3e\xa5\x49\x90\xc8\xa7\x0a\xf4\x1c\x8d\x6c\x28\x78\x8e\x53\x98\xdc\x99\x2f\xfa\x00\x5e\x4c\x5b\x13\x89\x2c\xae\x77\x71\x48\x62\xde\x82\x5a\xc5\x31\x2a\x95\xac\x32\x88\xc9\xee\x86\x14\xa5\x84\xf1\x18\x8a\x34\x0b\xa0\x10\xba\xf9\x3e\xfb\xf7\x7b\xc3\xf8\x0e\x63\xbe\x52\xb8\xc9\x8e\x6f\x73\x91\x59\x6b\x02\x48\x35\x4c\x05\x2a\x23\x41\x4b\xe4\x1a\x78\x41\x32\x20\x91\x22\x37\xcc\xc0\x15\xcd\xa1\x94\x42\x82\x16\x30\x21\x51\x74\x58\x70\x1a\x82\x8d\x54\xa3\xc3\x5b\x72\x3d\xaf\xf3\xec\x83\x77\x79\x45\x67\x27\xb0\xac\x7e\x1d\xba\x19\x6a\x8d\x32\x00\x89\x39\x4f\x8b\xb4\x98\xd1\x56\x86\x93\x71\x1b\xcc\xf7\x69\x31\xa5\xb9\x8f\x05\x1d\x12\x23\xd2\xdf\x14\xf4\x56\x69\x4f\x52\x50\x2b\xf0\x29\x0d\xfd\x93\xc7\xe5\xe4\xde\x35\x7b\xfb\x77\x43\xda\x76\xfc\x1f\x63\x5d\xff\x4d\x06\xff\x90\x06\xe0\x85\xfa\x3f\x1a\x1d\x1f\x36\xf5\x7f\x74\xb8\x3f\xa4\xfa\x3f\x3a\x8e\xb6\xf5\xff\x57\x0c\x82\xf5\x29\x26\x69\xb1\x3e\xc3\x54\x70\xfa\x85\x93\x00\xe1\x63\xf1\x09\x73\x21\xef\x08\x14\x4e\x79\x3c\x47\x18\x43\xce\x97\x97\x16\x8f\xae\x2c\x1a\x95\xd5\xf3\x3c\x9f\x56\x1a\x6f\x61\x0c\x6f\xa8\x7c\x87\xe6\xab\x34\x45\xaf\x57\xe0\x5a\xec\x31\x00\x6e\x15\x34\xc8\x03\x44\xd9\x96\xbb\xa7\x60\x30\x4d\x4c\x25\x20\xc4\x76\x5d\xf8\xfe\xbd\x61\xb6\x88\x5e\x13\x75\x50\xcf\xc0\xbc\x50\xe1\x99\x94\xe7\x42\x9f\xdd\xa6\x4a\xd7\x05\xd9\x3c\xa8\x76\x05\x20\x16\x5d\x20\x7d\xe4\xdb\x25\x69\xbc\x6a\xd4\x8b\xc5\x63\x2d\x56\x4c\x91\x66\x5d\xd9\xcf\xc7\x2a\xfc\x97\x88\x17\x9e\x45\xe8\x29\x26\x28\x5f\xa2\xff\x4f\x91\x59\x8e\xbe\xe1\x3f\xdf\xee\xd9\xbd\x89\x39\x85\x84\x3a\xb6\xf0\x1c\xbf\x51\xde\x50\x7a\x75\x97\xd6\x99\xb1\xd1\xf7\x7c\xdf\x6f\xd4\x10\xe7\xab\xe7\x92\x41\x7d\xd9\x19\xa5\x34\xf1\x5c\x5b\x1c\xa9\xc2\x52\x8f\x60\xea\xfc\x6b\x75\x02\xaf\x95\x1b\x98\x34\x1b\x43\x2c\xb5\x57\xab\x68\x12\x67\x42\x56\x6f\x99\x12\xae\x81\x8c\x0d\x4f\x33\xa1\xd0\xf3\xa1\xaa\x03\xab\x8d\x17\x9a\xcb\xae\xc9\xf7\x76\x2d\x11\xb2\x63\xe1\xdc\x2c\xb6\x8e\x6b\xe2\xb8\xd5\xb5\x98\x8e\x6b\xe3\x31\xa4\x22\xa4\x42\xbf\xe6\xa5\x31\x91\xc8\x17\xed\x4c\x05\x98\x29\x7c\x32\x1c\x0f\x43\x82\x52\xae\x59\x59\xfb\x4a\x47\x75\xb2\x4a\xc0\x74\xd2\xe1\xbb\x55\x92\xe0\x9a\xee\xda\xda\x6a\xcc\x39\x15\xcb\xbb\x73\xef\xcd\x64\x95\x04\xa0\x65\x50\x3b\x13\x5e\xa4\xf7\xf8\xc8\x81\x17\x8d\x79\x9c\x9f\x29\x36\x2d\xfc\x0f\x67\xe9\x81\x27\x4f\x6f\xd1\xda\x54\xea\x00\xaf\x60\x4c\xee\x86\xef\xc8\x5d\xcf\xdf\x78\x4c\x7f\x78\xb7\xbf\xda\xb4\xdd\x1f\xf9\x17\xf3\xe2\x6f\x1a\x92\xb4\x98\x36\x6e\xd5\x2e\x35\xda\xe9\x97\x6d\x3c\x30\xd5\x03\x68\xdb\xd0\x7e\xad\xa1\xac\x07\x6d\x41\x3b\x5f\x32\x67\xc9\xa5\x56\xb4\xe3\x9a\xb3\x75\xb1\xcc\x52\x6d\xd8\x83\xcd\x7f\x6d\x7c\xe6\xa4\x49\x4b\xfe\x4f\xae\x3e\x4b\x4c\xd2\xdb\x67\x59\x48\x53\xad\x6a\x0c\xe6\x79\x39\x3c\xb9\x62\x4e\xc5\x98\x93\x77\xb1\x8f\x6c\x67\x0e\x6d\x3c\x8d\xf0\xf7\xee\xec\x59\x86\x39\x16\xda\x2e\x8a\x05\x4c\x84\xc8\x98\x43\x87\xe8\x3a\x30\x22\x49\x8c\xfd\x7b\x63\x15\x91\x4a\x82\x6f\x93\xb5\xfc\x92\x26\xaf\x98\xe3\xb4\xb9\x71\x1c\xa7\x9b\x16\xd7\x65\x8e\x53\x31\xc7\xc9\x61\x0c\x1a\xc3\xd3\x79\x9a\x4d\x25\x16\xd6\xc8\x9a\x52\x63\xf8\xa1\xee\x83\x35\x86\x5f\xba\x0d\x2c\x5b\xff\x35\xdd\xdb\x76\xaa\xfd\xfe\x4f\x22\xfe\x96\xfb\x9f\xe1\xf0\x28\xa2\xfe\xef\xf0\x78\x38\x1a\x0e\x0f\x0f\xcc\xfd\x4f\xb4\xbf\xed\xff\x7e\xc5\xe8\xf5\x7f\xe6\xf6\x84\xda\x3f\x7d\xb7\x44\xd8\x70\xae\x09\x50\x56\xb1\xb6\x90\x69\x2e\x0b\xda\x61\xa1\xc6\x2c\xf4\x4e\x5c\x77\xc1\x9e\xca\x86\xa3\x07\x77\x66\xbd\x39\xcd\x76\xbd\xd3\x5d\x6e\xc4\x98\x07\x97\x18\xe6\xd2\x67\xfc\x12\x57\x09\xbd\xbf\xad\x12\xe3\x95\x54\xe9\x0d\x7e\xb5\x37\x47\xf6\x72\xa4\x03\x11\x8c\x75\x23\xf4\x80\x7c\xaf\xbe\xe9\xb2\x80\xb6\xbb\x08\x60\xf7\x86\x00\xce\x88\x71\xcb\x12\x76\x17\x50\x55\xee\x49\xc9\x28\x56\x27\xb0\x9e\x0a\x0c\x5f\x9a\xc0\x4c\x83\x97\x61\x01\xbb\x37\x2d\x96\xf9\x10\x11\x7f\xf3\x79\xf2\xa2\x4b\x46\xd6\x93\x3e\x75\x24\x1b\xef\xac\x6a\xd3\x79\x54\x15\xb3\x29\x39\x79\xee\x72\x29\x60\xbd\x84\x5a\x37\xfa\xb0\xba\x76\xa9\xbe\x22\xaa\x7a\x5f\x5b\xc8\xdd\x8e\xed\xd8\x8e\xed\xf8\xd3\x8d\xff\x05\x00\x00\xff\xff\xbc\xa6\x41\x19\x00\x1c\x00\x00"

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
