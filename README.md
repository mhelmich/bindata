# bindata

This package converts any file into go code. It is heavily inspired by the excellent [go-bindata](github.com/jteeuwen/go-bindata).

`bindata` can be used as library or as a commandline tool.

## Usage as Library

Import the following package to use bindata in your code:

`github.com/mhelmich/bindata`

### Usage

The following snippet explains how to use bindata to create a bindata file from your Go program.

```
import bd "github.com/mhelmich/bindata"

...

err = bd.New(
    []string{"path/to/files/*.json"},
    bd.PackageName("test"),
    bd.OutputFile("test/bindata_files.go"),
    bd.Archiver(bd.Tar),
    bd.Compressor(bd.Bz),
).Archive()
```

## Usage as Cmd Tool

Run the following to install bindata:

`go get -u github.com/mhelmich/bindata/cmd/bindata`

### Usage

The simplest call of `bindata` passes only file paths to it.

`$ bindata dir/file1.json dir/file2.json`

However `bindata` paths can contain wildcards/globs:

`$ bindata dir/*.json dir/**/*.json`

`bindata` can be configured to write the output file to a particular path. In this case, `bindata` automatically uses the last element of the path as package name:

`$ bindata -o mypackage/bindata_files.go dir/*.json dir/**/*.json`

If for some reason you need to configure the package name of the generated bindata file independently from the given path, use the `package` flag:

`$ bindata -o mypackage/bindata_files.go -package otherpackage dir/*.json dir/**/*.json`

## Configuration

Right now bindata supports the following configurations:

* TarBz - a bz compressed tar archive


## Accessing a File

The generated `bindata` file exports two public functions `FileNames` and `ReadFile`.

### ReadFile

`ReadFile` reads the file named by filename and returns the contents. A successful call returns err == nil, not err == EOF. Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.

It can be used like this:
```
data, err := bindata.ReadFile(an)
if err != nil {
    // err != nil if the file wasn't found in the bindata file
}
```

### FileNames

`FileNames` returns a list of all files in this bindata file.

It can be used like this:
```
for _, fileName := range bindata.FileNames() {
    // fileName contains a file name now
}
```
