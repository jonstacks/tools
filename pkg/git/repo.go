package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// Repo is the git repo and information we need to know about it
type Repo struct {
	path string
}

// NewRepo creates a new Repo for the given path
func NewRepo(path string) *Repo {
	return &Repo{path: path}
}

// CurrentBranch returns the current git branch if it exists, otherwise it
// returns the error
func (r *Repo) CurrentBranch() (string, error) {
	out, err := r.execute("rev-parse", "--abbrev-ref", "HEAD")
	return strings.TrimSpace(out), err
}

// CommitHash returns the commit hash for a given repo
func (r *Repo) CommitHash() (string, error) {
	out, err := r.execute("rev-parse", "HEAD")
	return strings.TrimSpace(out), err
}

func (r *Repo) execute(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	if r.path != "" {
		cmd.Dir = r.path
	}

	out, err := cmd.CombinedOutput()
	return string(out), err
}

// ShallowClone does a shallow clone of the repo to the filepath given.
func (r *Repo) ShallowClone(repo string) (string, error) {
	var cwd string
	var err error

	if r.path == "" {
		cwd, err = os.Getwd()
		if err != nil {
			return cwd, err
		}
		r.path = cwd
	}
	return r.execute("clone", "--depth", "1", repo, r.path)
}

// CheckoutCommitOnly is mainly used by a CI/CD system to build a particular
// commit. Eventually look at keeping a local copy and cloning from that for
// better performance.
func (r *Repo) CheckoutCommitOnly(repo string, commitSHA string) (string, error) {
	var buf bytes.Buffer

	commands := [][]string{
		[]string{"init"},
		[]string{"remote", "add", "origin", repo},
		[]string{"fetch", "origin", commitSHA},
		[]string{"reset", "--hard", "FETCH_HEAD"},
	}

	for _, command := range commands {
		out, err := r.execute(command...)
		buf.WriteString(out)
		if err != nil {
			return buf.String(), err
		}
	}
	return buf.String(), nil
}
