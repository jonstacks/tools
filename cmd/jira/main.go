package main

import (
	"fmt"
	"os"

	"github.com/jonstacks/tools/pkg/utils"

	docopt "github.com/docopt/docopt-go"
	"github.com/jonstacks/tools/pkg/git"
	"github.com/jonstacks/tools/pkg/jira"
)

func handleError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	usage := `JIRA

Usage:
  jira
  jira <issueKey>
  jira -h | --help
  jira --version

Options:
  -h --help          Show this screen.
  --version          Show version.
  `

	arguments, _ := docopt.Parse(usage, nil, true, "0.1.0", false)

	config, err := jira.GetConfig()
	handleError(err)

	var key string
	if arguments["<issueKey>"] != nil {
		key = arguments["<issueKey>"].(string)
	} else {
		repo := git.NewRepo("")
		key, err = repo.CurrentBranch()
		handleError(err)
	}

	issue, err := jira.ParseIssue(key)
	handleError(err)

	url := issue.URL(config.Host)

	handleError(utils.OpenInBrowser(url))
}
