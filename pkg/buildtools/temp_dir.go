package buildtools

import (
	"io/ioutil"
	"os"
	"sync"
)

// TempDir simplifies working with a temporary directory by cleaning it up
// once we no longer need it.
type TempDir struct {
	path string
	once *sync.Once
}

// NewTempDir wraps the initialization of ioutil.TempDir
func NewTempDir(dir, prefix string) (TempDir, error) {
	path, err := ioutil.TempDir(dir, prefix)
	if err != nil {
		return TempDir{}, err
	}
	return TempDir{
		path: path,
		once: &sync.Once{},
	}, nil
}

// Path returns the path of the temp directory
func (td TempDir) Path() string {
	return td.path
}

// Do calls the given function with the temp path as an argument. It returns
// an error if one occured for the function. Currently, path cleanup is a
// best effort. If the user of the program has sudo permissions and runs a
// subprocess that creates files or directories as root, this may not clean up
// that directory.
func (td TempDir) Do(f func(path string) error) error {
	var err error
	td.once.Do(func() {
		defer func() { os.RemoveAll(td.path) }()
		err = f(td.path)
	})
	return err
}
