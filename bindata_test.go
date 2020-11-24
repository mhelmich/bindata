package bindata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindataBasic(t *testing.T) {
	tests := []struct {
		files []string
		err   error
	}{
		{[]string{"test-fixtures/tree1", "test-fixtures/tree2/dir1/f1.txt"}, nil},
		{[]string{"../."}, ErrNoFilesMatched},
	}

	for _, test := range tests {
		tmp := New(test.files)
		bd := tmp.(*bindata)
		err := bd.Archive()
		assert.Equal(t, test.err, err)
	}
}
