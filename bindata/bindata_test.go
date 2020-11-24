package bindata

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinDataBasic(t *testing.T) {
	bites, err := ReadFile("test-fixtures/tree2/dir1/f1.txt")
	assert.Nil(t, err)
	fmt.Printf("%s\n", string(bites))
}

func TestBinDataUnknownFile(t *testing.T) {
	bites, err := ReadFile("test-fixtures/tree2/doesnt/exist.txt")
	assert.Equal(t, os.ErrNotExist, err)
	assert.Nil(t, bites)
}

func TestBinDataFileNames(t *testing.T) {
	names := FileNames()
	assert.Equal(t, 4, len(names))
	assert.Equal(t, "test-fixtures/tree1/dir1/dir11/f111.txt", names[0])
	assert.Equal(t, "test-fixtures/tree1/dir1/f11.txt", names[1])
	assert.Equal(t, "test-fixtures/tree1/dir2/f21.txt", names[2])
	assert.Equal(t, "test-fixtures/tree2/dir1/f1.txt", names[3])
}
