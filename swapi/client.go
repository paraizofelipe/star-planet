package swapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURLScheme = "https"
	defaultBaseURLHost   = "swapi.dev"
	defaultBasePath      = "/api/"
	defaultUserAgent     = "swapi.go"
)

type Option func(*Client)

type Client struct {
	baseURL    *url.URL
	basePath   string
	userAgent  string
	httpClient *http.Client
}

func NewClient(options ...Option) *Client {
	c := &Client{
		baseURL: &url.URL{
			Scheme: defaultBaseURLScheme,
			Host:   defaultBaseURLHost,
		},
		basePath:   defaultBasePath,
		userAgent:  defaultUserAgent,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *Client) newRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(c.basePath + s)
	if err != nil {
		return nil, err
	}

	q := rel.Query()
	q.Set("format", "json")

	rel.RawQuery = q.Encode()

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}
