package jira

import (
	"fmt"
	"regexp"
)

var issueRegex = regexp.MustCompile(`\w+\-\d+`)

// Issue is a struct representing a JIRA issue.
type Issue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields struct {
		Description string `json:"description"`
		Summary     string `json:"summmary"`
	} `json:"fields"`
}

// ParseIssue tries to parse a JIRA key from a string. If it
// is able to, it returns a pointer to the issue otherwise the
// issue is nil and error will be populated
func ParseIssue(key string) (*Issue, error) {
	match := issueRegex.FindString(key)
	if match == "" {
		// No match was found
		return nil, fmt.Errorf("unable to parse Issue key from '%s'", key)
	}
	return &Issue{Key: match}, nil
}

func (i Issue) String() string {
	return i.Key
}

// URL returns the URL for the JIRA issue
func (i Issue) URL(h string) string {
	return fmt.Sprintf("%s/browse/%s", h, i.Key)
}
