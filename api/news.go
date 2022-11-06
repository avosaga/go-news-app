package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

func NewClient(client *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{client, key, pageSize}
}

func (c *Client) FetchAll(query string, page string) (*Results, error) {
	endpoint := fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=%s&apiKey=%s&pageSize=%d&page=%s&language=en&sortBy=publishedAt",
		url.QueryEscape(query),
		c.key,
		c.PageSize,
		page,
	)

	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	results := &Results{}

	return results, json.Unmarshal(body, results)
}
