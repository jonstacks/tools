package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientOpts struct {
	BaseURL   string
	UserEmail string
	APIToken  string
	Timeout   time.Duration
}

type Client struct {
	opts   *ClientOpts
	client *http.Client
}

func NewClient(opts *ClientOpts) *Client {
	client := http.Client{
		Timeout: opts.Timeout,
	}
	return &Client{
		opts:   opts,
		client: &client,
	}
}

func (c *Client) GetIssue(key string) (*Issue, error) {
	url := fmt.Sprintf("issue/%s", key)
	req, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	issue := &Issue{}
	err = json.Unmarshal(body, issue)
	return issue, err
}

func (c *Client) request(method string, url string, body io.Reader) (*http.Request, error) {
	url = fmt.Sprintf("%s/rest/api/2/%s", c.opts.BaseURL, url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	authorization := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", c.opts.UserEmail, c.opts.APIToken)),
	)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authorization))
	req.Header.Add("Content-Type", "application/json")

	return req, err
}
