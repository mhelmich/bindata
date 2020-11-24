package bindata

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

type templateInput struct {
	PackageName string
	FileNames   []string
	Files       []file
	Tree        map[string]*treeElement
}

type file struct {
	GoName  string
	Content string
}

type treeElement struct {
	GoName        string
	RemainingPath string
	Children      map[string]*treeElement
}

func (b *bindata) createTemplateInput(pais []*pathInfo, buf *bytes.Buffer) (*templateInput, error) {
	tree, err := b.buildPathTree(pais)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, pai := range pais {
		fileNames = append(fileNames, pai.path)
	}

	input := &templateInput{
		PackageName: b.getPackageName(),
		FileNames:   fileNames,
		Files: []file{
			{
				GoName:  "Archive",
				Content: buf.String(),
			},
		},
		Tree: tree,
	}

	return input, nil
}

func (b *bindata) getPackageName() string {
	if b.opts.PackageName != "" {
		return b.opts.PackageName
	}

	dir := filepath.Dir(b.opts.OutputFile)
	parts := strings.Split(dir, pathSeparator)
	return parts[len(parts)-1]
}

func (b *bindata) buildPathTree(pais []*pathInfo) (map[string]*treeElement, error) {
	tree := make(map[string]*treeElement)
	for _, pai := range pais {
		parts := b.splitPath(pai.path)
		err := b.traverseTree(parts, tree, pai.path, pai.path)
		if err != nil {
			return nil, err
		}
	}
	return tree, nil
}

func (b *bindata) traverseTree(parts []string, m map[string]*treeElement, path string, remainingPath string) error {
	item := parts[0]

	if len(parts) == 1 {
		_, ok := m[item]
		if ok {
			return fmt.Errorf("file [%s] already exists", item)
		}

		m[item] = &treeElement{
			RemainingPath: remainingPath,
			GoName:        "Archive",
		}
		return nil
	}

	n, ok := m[item]
	if ok {
		err := b.traverseTree(parts[1:], n.Children, path, remainingPath)
		if err != nil {
			return err
		}
	} else {
		n = &treeElement{
			Children: map[string]*treeElement{},
		}
		m[item] = n
		err := b.traverseTree(parts[1:], n.Children, path, remainingPath)
		if err != nil {
			return err
		}
	}

	return nil
}
