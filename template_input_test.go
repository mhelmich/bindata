package bindata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateInputGetPackageNameSupplyPackageName(t *testing.T) {
	bd := &bindata{
		opts: &Options{
			PackageName: "mypackagename",
			OutputFile:  "dir/file.go",
		},
	}
	pkgName := bd.getPackageName()
	assert.Equal(t, "mypackagename", pkgName)
}

func TestTemplateInputGetPackageNameDerivePackageName(t *testing.T) {
	bd := &bindata{
		opts: &Options{
			OutputFile: "dir/file.go",
		},
	}
	pkgName := bd.getPackageName()
	assert.Equal(t, "dir", pkgName)
}
