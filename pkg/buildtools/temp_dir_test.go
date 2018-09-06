package buildtools

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTempDirPath(t *testing.T) {
	td, err := NewTempDir("", "mytmp-")
	assert.Nil(t, err)

	dirName := filepath.Base(td.Path())
	assert.True(t, strings.HasPrefix(dirName, "mytmp-"))
}
