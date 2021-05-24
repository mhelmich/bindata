package bindata

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
)

var (
	pathSeparator = fmt.Sprintf("%c", os.PathSeparator)

	// ErrNoFilesMatched is returned if no files could be found in the given paths
	ErrNoFilesMatched = fmt.Errorf("no files matched the path provided")
)

// ArchiverLayer -
type ArchiverLayer int

var (
	archiverLayerValuesToNames = map[ArchiverLayer]string{}
)

func (a ArchiverLayer) String() string {
	return archiverLayerValuesToNames[a]
}

const (
	// NoArchiver -
	NoArchiver ArchiverLayer = iota
	// Tar -
	Tar
)

// CompressorLayer -
type CompressorLayer int

var (
	compressorLayerValuesToNames = map[CompressorLayer]string{}
)

func (c CompressorLayer) String() string {
	return compressorLayerValuesToNames[c]
}

const (
	// NoCompressor -
	NoCompressor CompressorLayer = iota
	// Bz -
	Bz
)

// Bindata -
type Bindata interface {
	Archive() error
}

type bindata struct {
	opts  *Options
	paths []string
}

// Options -
type Options struct {
	Paths       []string
	PackageName string
	OutputFile  string
	Archiver    ArchiverLayer
	Compressor  CompressorLayer
}

// Option -
type Option func(*Options)

// PackageName defines the package name that should appear in the generated file.
func PackageName(packageName string) Option {
	return func(opts *Options) {
		opts.PackageName = packageName
	}
}

// OutputFile defines the name of the output file that is being generated.
func OutputFile(output string) Option {
	return func(opts *Options) {
		opts.OutputFile = output
	}
}

// Archiver defines the archiver to use (defaults to tar).
func Archiver(a ArchiverLayer) Option {
	return func(opts *Options) {
		opts.Archiver = a
	}
}

// Compressor defines the compressor to use (default to bz).
func Compressor(c CompressorLayer) Option {
	return func(opts *Options) {
		opts.Compressor = c
	}
}

// New creates and configures a new instance of bindata.
func New(paths []string, setters ...Option) Bindata {
	opts := &Options{
		Paths:       paths,
		PackageName: "bindata",
		OutputFile:  "bindata/bindata.go",
		Archiver:    Tar,
		Compressor:  Bz,
	}

	for _, setter := range setters {
		setter(opts)
	}

	return &bindata{
		paths: paths,
		opts:  opts,
	}
}

// Archive generates the bindata archive based on the specified options.
func (b *bindata) Archive() error {
	pathsAndInfos, err := b.resolvePathsToFiles(b.paths)
	if err != nil {
		return err
	}

	if len(pathsAndInfos) <= 0 {
		glog.Info("No files have been matched")
		return ErrNoFilesMatched
	}

	buf := &bytes.Buffer{}
	err = b.compress(pathsAndInfos, buf)
	if err != nil {
		return err
	}

	input, err := b.createTemplateInput(pathsAndInfos, buf)
	if err != nil {
		return err
	}

	absOutputPath, err := filepath.Abs(b.opts.OutputFile)
	if err != nil {
		return err
	}

	err = b.renderFile(input, absOutputPath, TarBz)
	if err != nil {
		return err
	}

	return nil
}

func (b *bindata) compress(pais []*pathInfo, out io.Writer) error {
	l := newStringLiteralEncoder(out)
	c := newCompressorLayer(b.opts.Compressor, l)
	w := newArchiverLayer(b.opts.Archiver, c)
	defer func() { _ = w.Close() }()

	for _, pai := range pais {
		err := w.WriteFile(pai.path, pai.info)
		if err != nil {
			return err
		}
	}
	return nil
}

type pathInfo struct {
	path string
	info os.FileInfo
}

func (b *bindata) resolvePathsToFiles(paths []string) ([]*pathInfo, error) {
	var matchedFiles []*pathInfo
	if len(paths) > 0 {
		var matchedPaths []string
		for _, path := range paths {
			path = strings.TrimSpace(path)
			if strings.HasPrefix(path, "..") {
				glog.Infof("won't include path starting with '..'. Skipping %s", path)
				continue
			}

			matched, err := filepath.Glob(path)
			if err != nil {
				return nil, err
			}
			matchedPaths = append(matchedPaths, matched...)
		}

		for _, matchedPath := range matchedPaths {
			matchedPath = filepath.Clean(matchedPath)
			fileInfo, err := os.Stat(matchedPath)
			if err != nil {
				return nil, err
			}

			if fileInfo.IsDir() {
				err = filepath.Walk(matchedPath, func(file string, fi os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if !fi.IsDir() {
						matchedFiles = append(matchedFiles, &pathInfo{
							path: file,
							info: fi,
						})
					}

					return nil
				})

				if err != nil {
					return nil, err
				}
			} else {
				matchedFiles = append(matchedFiles, &pathInfo{
					path: matchedPath,
					info: fileInfo,
				})
			}
		}
	}

	return matchedFiles, nil
}

func (b *bindata) splitPath(path string) []string {
	parts := strings.Split(path, pathSeparator)
	if strings.HasPrefix(path, pathSeparator) {
		return parts[1:]
	}
	return parts
}
