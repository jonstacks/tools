package git

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	dir, err := ioutil.TempDir("", "timeout-test")
	assert.Nil(t, err)
	defer func() { os.RemoveAll(dir) }()

	timeout := 300 * time.Millisecond
	r := NewRepo(dir).WithTimeout(&timeout)
	_, err = r.ShallowClone("https://github.com/torvalds/linux.git")
	assert.Equal(t, err, context.DeadlineExceeded)
}
