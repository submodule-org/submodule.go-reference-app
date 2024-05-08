package http_test

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client struct {
	protocol string
	base     string
	port     int
	headers  http.Header
}

func (c *Client) Get(path string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("%s://%s:%d%s", c.protocol, c.base, c.port, path))
}

func (c *Client) Post(path string) (*http.Response, error) {
	return http.Post(fmt.Sprintf("%s://%s:%d%s", c.protocol, c.base, c.port, path), "application/json", nil)
}

func (c *Client) PostWithBody(path string, body []byte) (*http.Response, error) {
	return http.Post(fmt.Sprintf("%s://%s:%d%s", c.protocol, c.base, c.port, path), "application/json", bytes.NewBuffer(body))
}

func (c *Client) Delete(path string) (*http.Response, error) {
	req, e := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s://%s:%d%s", c.protocol, c.base, c.port, path),
		nil,
	)

	if e != nil {
		return nil, e
	}

	return http.DefaultClient.Do(req)
}

type ClientOpts func(Client) Client

func WithHeaderValue(header string, value string) ClientOpts {
	return func(c Client) Client {
		c.headers.Set(header, value)
		return c
	}
}

func WithPort(port int) ClientOpts {
	return func(c Client) Client {
		c.port = port
		return c
	}
}

func NewClient(opts ...ClientOpts) Client {
	c := Client{
		protocol: "http",
		base:     "localhost",
		port:     8080,
		headers:  make(http.Header),
	}

	for _, o := range opts {
		c = o(c)
	}

	return c
}
