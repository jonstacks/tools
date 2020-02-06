package utils

import (
	"os/exec"
	"runtime"
)

// OpenInBrowser opens the given URL in the browser by shelling out to the
// system command for opening the browser
func OpenInBrowser(url string) error {
	var openCmd string

	switch runtime.GOOS {
	case "darwin":
		openCmd = "open"
	case "linux":
		openCmd = "xdg-open"
	}

	open := exec.Command(openCmd, url)
	return open.Run()
}
