package httpclient

import (
	"io"
	"net/http"
	"net/url"
)

// Client provides an interface for the `http.Client` concrete type
type Client interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

var _ Client = http.DefaultClient
