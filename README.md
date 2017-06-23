[![Build Status](https://travis-ci.org/sebnow/httpclient.svg?branch=master)](https://travis-ci.org/sebnow/httpclient)
[![Coverage Status](https://coveralls.io/repos/github/sebnow/httpclient/badge.svg?branch=master)](https://coveralls.io/github/sebnow/httpclient?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sebnow/httpclient)](https://goreportcard.com/report/github.com/sebnow/httpclient)

Description
===========

The `httpclient` package provides an interface for Go's standard
`http.Client` concrete type. This makes it easier to mock HTTP clients,
or wrap them with additional functionality.

Usage
=====

```go
client := httpclient.NewContext(http.DefaultClient)
response, err := client.GetContext(ctx.TODO(), "https://golang.org")
```

