package git

import (
	"bytes"
	"os/exec"
	"strings"
)

// CurrentBranch returns the current git branch if it exists, otherwise it
// returns the error
func CurrentBranch() (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
