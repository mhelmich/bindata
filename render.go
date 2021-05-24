package bindata

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func (b *bindata) renderFile(input *templateInput, outputFile string, typ Type) error {
	tmplPaths, ok := typeToTemplate[typ]
	if !ok {
		return fmt.Errorf("unknown bindata type: %d", typ)
	}

	tmpl, err := b.readTemplateFiles(tmplPaths)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, input)
	if err != nil {
		return err
	}

	bites, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	err = b.mkdir(outputFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, bites, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (b *bindata) mkdir(outputFile string) error {
	dir := filepath.Dir(outputFile)
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0700)
	} else if !stat.IsDir() {
		tokens := strings.Split(dir, string(filepath.Separator))
		p := strings.Join(tokens[:len(tokens)-1], string(filepath.Separator))
		return fmt.Errorf("can't create directory '%s' - found file with the same name '%s' in folder '%s'", filepath.Base(dir), filepath.Base(dir), p)
	}

	return nil
}

func (b *bindata) readTemplateFiles(tmplPaths []string) (*template.Template, error) {
	tmpl := template.New("root.tmpl")
	for _, templatePath := range tmplPaths {
		data, err := ReadFile(templatePath)
		if err != nil {
			return nil, err
		}

		_, err = tmpl.Parse(string(data))
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}
