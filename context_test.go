package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewContextCanWrapDefaultClient(t *testing.T) {
	clientContext := NewContext(http.DefaultClient)
	if _, ok := clientContext.(ClientContext); !ok {
		t.Error("http.DefaultClient can not be wrapped with ClientContext")
	}
}

func TestClientContextGetContextAddsContext(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)
	client := NewContext(mockClient)

	ctx := context.TODO()
	expectedResp := &http.Response{}
	req, err := http.NewRequest("GET", "http://test.com", nil)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)

	mockClient.EXPECT().Do(reqMatcher{req}).Return(expectedResp, nil)
	resp, err := client.GetContext(ctx, "http://test.com")
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", expectedResp, resp)
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("Incorrect error;\n\texpected: %#v\n\t     got: %#v", nil, err)
	}
}

func TestClientContextPostAddsContext(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)
	client := NewContext(mockClient)

	ctx := context.TODO()
	contentType := "text/html"
	expectedResp := &http.Response{}
	body := bytes.NewBufferString("")
	req, err := http.NewRequest("POST", "http://test.com", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", contentType)
	req = req.WithContext(ctx)

	mockClient.EXPECT().Do(reqMatcher{req}).Return(expectedResp, nil)
	resp, err := client.PostContext(ctx, "http://test.com", contentType, body)
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", expectedResp, resp)
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("Incorrect error;\n\texpected: %#v\n\t     got: %#v", nil, err)
	}
}

func TestClientContextHeadAddsContext(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)
	client := NewContext(mockClient)

	ctx := context.TODO()
	expectedResp := &http.Response{}
	req, err := http.NewRequest("HEAD", "http://test.com", nil)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)

	mockClient.EXPECT().Do(reqMatcher{req}).Return(expectedResp, nil)
	resp, err := client.HeadContext(ctx, "http://test.com")
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", expectedResp, resp)
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("Incorrect error;\n\texpected: %#v\n\t     got: %#v", nil, err)
	}
}

func TestClientContextPostFormAddsContext(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)
	client := NewContext(mockClient)

	var values url.Values
	ctx := context.TODO()
	contentType := "application/x-www-form-urlencoded"
	expectedResp := &http.Response{}
	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest("POST", "http://test.com", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", contentType)
	req = req.WithContext(ctx)

	mockClient.EXPECT().Do(reqMatcher{req}).Return(expectedResp, nil)
	resp, err := client.PostFormContext(ctx, "http://test.com", values)
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", expectedResp, resp)
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("Incorrect error;\n\texpected: %#v\n\t     got: %#v", nil, err)
	}
}

type reqMatcher struct {
	x *http.Request
}

var _ gomock.Matcher = reqMatcher{}

func (m reqMatcher) Matches(i interface{}) bool {
	x, ok := i.(*http.Request)
	if !ok {
		return false
	}

	r := true
	r = r && m.x.Method == x.Method
	r = r && m.x.URL.String() == x.URL.String()
	r = r && m.x.Header.Get("content-type") == x.Header.Get("content-type")
	r = r && reflect.DeepEqual(m.x.Body, x.Body)
	r = r && m.x.Context() == x.Context()
	return r
}

func (m reqMatcher) String() string {
	return fmt.Sprintf("is equal to %v", m.x)
}
