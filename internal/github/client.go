package github

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/NERFTHISPLS/github-user-activity/internal/e"
)

const githubHost = "https://api.github.com"

const reqErr = "failed to do request"

type Client struct {
	basePath string
	http     *http.Client
}

func NewClient() *Client {
	return &Client{
		basePath: githubHost,
		http:     http.DefaultClient,
	}
}

func (c *Client) UserEvents(username string) ([]Event, error) {
	url, err := url.JoinPath(c.basePath, "users", username, "events")
	if err != nil {
		return nil, e.Wrap(reqErr, err)
	}

	resp, err := c.http.Get(url)
	if err != nil {
		return nil, e.Wrap(reqErr, err)
	}
	defer resp.Body.Close()

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, e.Wrap(reqErr, err)
	}

	return events, nil
}

type Event struct {
	Type string     `json:"type"`
	Repo Repository `json:"repo"`
}

type Repository struct {
	Name string `json:"name"`
}
