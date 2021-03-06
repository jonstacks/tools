package main

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/jonstacks/tools/pkg/utils"

	docopt "github.com/docopt/docopt-go"
	"github.com/jonstacks/tools/pkg/git"
	"github.com/jonstacks/tools/pkg/jira"
)

//go:embed jira-description.template
var descriptionTemplate string

func handleError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func command(arguments map[string]interface{}) string {
	if arguments["description"] == true {
		return "description"
	}
	return ""
}

type TemplateContext struct {
	Config *jira.Config
	Issue  *jira.Issue
}

func main() {
	usage := `JIRA

Usage:
  jira
  jira <issueKey>
  jira description
  jira description <issueKey>
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

	switch command(arguments) {
	case "description":
		client := jira.NewClient(&jira.ClientOpts{
			BaseURL:   config.Host,
			UserEmail: config.UserEmail,
			APIToken:  config.APIToken,
			Timeout:   20 * time.Second,
		})

		issue, err := client.GetIssue(issue.Key)
		handleError(err)

		templ, err := template.New("jira-gh-issue").Parse(descriptionTemplate)
		handleError(err)

		templ.Execute(os.Stdout, TemplateContext{
			Config: config,
			Issue:  issue,
		})
	default:
		url := issue.URL(config.Host)
		handleError(utils.OpenInBrowser(url))
	}
}
