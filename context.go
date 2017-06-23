package httpclient

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ClientContext provides context aware methods for HTTP requests
type ClientContext interface {
	Client
	// GetContext is the same as Get but accepts a context
	GetContext(ctx context.Context, url string) (resp *http.Response, err error)
	// HeadContext is the same as Head but accepts a context
	HeadContext(ctx context.Context, url string) (resp *http.Response, err error)
	// PostContext is the same as Post but accepts a context
	PostContext(ctx context.Context, url string, contentType string, body io.Reader) (resp *http.Response, err error)
	// PostFormContext is the same as PostForm but accepts a context
	PostFormContext(ctx context.Context, url string, data url.Values) (resp *http.Response, err error)
}

type contextClient struct {
	Client
}

var _ Client = new(contextClient)
var _ ClientContext = new(contextClient)

// NewContext wraps a Client with context aware methods
func NewContext(client Client) ClientContext {
	return &contextClient{client}
}

func (c *contextClient) GetContext(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Client.Do(req.WithContext(ctx))
}

func (c *contextClient) PostContext(ctx context.Context, url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("content-type", contentType)

	return c.Client.Do(req)
}

func (c *contextClient) HeadContext(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Client.Do(req.WithContext(ctx))
}

func (c *contextClient) PostFormContext(ctx context.Context, url string, data url.Values) (resp *http.Response, err error) {
	return c.PostContext(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
