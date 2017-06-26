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

type responseTestCase struct {
	url            string
	shouldMock     bool
	returnResponse *http.Response
	returnErr      error
	expectedResp   *http.Response
	expectedErr    error
}

func TestNewContextCanWrapDefaultClient(t *testing.T) {
	clientContext := NewContext(http.DefaultClient)
	if _, ok := clientContext.(ClientContext); !ok {
		t.Error("http.DefaultClient can not be wrapped with ClientContext")
	}
}

func TestClientContextGetContextAddsContext(t *testing.T) {
	var testCases = map[string]responseTestCase{
		"malformed url": {
			url:         "://malformed",
			shouldMock:  false,
			expectedErr: fmt.Errorf("parse ://malformed: missing protocol scheme"),
		},
		"returns response": {
			url:            "http://test.com",
			shouldMock:     true,
			returnResponse: &http.Response{},
			expectedResp:   &http.Response{},
		},
		"returns error": {
			url:         "http://test.com",
			shouldMock:  true,
			returnErr:   fmt.Errorf("error"),
			expectedErr: fmt.Errorf("error"),
		},
	}

	for testName, tt := range testCases {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockClient := NewMockClient(mockCtrl)
			client := NewContext(mockClient)
			ctx := context.TODO()

			if tt.shouldMock {
				req, err := http.NewRequest("GET", tt.url, nil)
				if err != nil {
					panic(err)
				}
				req = req.WithContext(ctx)

				mockClient.EXPECT().Do(reqMatcher{req}).Return(tt.expectedResp, tt.expectedErr)
			}

			resp, err := client.GetContext(ctx, tt.url)
			if !reflect.DeepEqual(resp, tt.expectedResp) {
				t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", tt.expectedResp, resp)
			}

			if (err != nil || tt.expectedErr != nil) && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Incorrect error;\n\texpected: %s\n\t     got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestClientContextPostAddsContext(t *testing.T) {
	var testCases = map[string]responseTestCase{
		"malformed url": {
			url:         "://malformed",
			shouldMock:  false,
			expectedErr: fmt.Errorf("parse ://malformed: missing protocol scheme"),
		},
		"returns response": {
			url:            "http://test.com",
			shouldMock:     true,
			returnResponse: &http.Response{},
			expectedResp:   &http.Response{},
		},
		"returns error": {
			url:         "http://test.com",
			shouldMock:  true,
			returnErr:   fmt.Errorf("error"),
			expectedErr: fmt.Errorf("error"),
		},
	}

	for testName, tt := range testCases {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockClient := NewMockClient(mockCtrl)
			client := NewContext(mockClient)

			ctx := context.TODO()
			contentType := "text/html"
			body := bytes.NewBufferString("")

			if tt.shouldMock {
				req, err := http.NewRequest("POST", tt.url, body)
				if err != nil {
					panic(err)
				}
				req.Header.Set("content-type", contentType)
				req = req.WithContext(ctx)

				mockClient.EXPECT().Do(reqMatcher{req}).Return(tt.returnResponse, tt.returnErr)
			}

			resp, err := client.PostContext(ctx, tt.url, contentType, body)
			if !reflect.DeepEqual(resp, tt.expectedResp) {
				t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", tt.expectedResp, resp)
			}

			if (err != nil || tt.expectedErr != nil) && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Incorrect error;\n\texpected: %s\n\t     got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestClientContextHeadAddsContext(t *testing.T) {
	var testCases = map[string]responseTestCase{
		"malformed url": {
			url:         "://malformed",
			shouldMock:  false,
			expectedErr: fmt.Errorf("parse ://malformed: missing protocol scheme"),
		},
		"returns response": {
			url:            "http://test.com",
			shouldMock:     true,
			returnResponse: &http.Response{},
			expectedResp:   &http.Response{},
		},
		"returns error": {
			url:         "http://test.com",
			shouldMock:  true,
			returnErr:   fmt.Errorf("error"),
			expectedErr: fmt.Errorf("error"),
		},
	}

	for testName, tt := range testCases {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockClient := NewMockClient(mockCtrl)
			client := NewContext(mockClient)

			ctx := context.TODO()

			if tt.shouldMock {
				req, err := http.NewRequest("HEAD", tt.url, nil)
				if err != nil {
					panic(err)
				}
				req = req.WithContext(ctx)

				mockClient.EXPECT().Do(reqMatcher{req}).Return(tt.returnResponse, tt.returnErr)
			}

			resp, err := client.HeadContext(ctx, tt.url)
			if !reflect.DeepEqual(resp, tt.expectedResp) {
				t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", tt.expectedResp, resp)
			}

			if (err != nil || tt.expectedErr != nil) && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Incorrect error;\n\texpected: %s\n\t     got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestClientContextPostFormAddsContext(t *testing.T) {
	var testCases = map[string]responseTestCase{
		"malformed url": {
			url:         "://malformed",
			shouldMock:  false,
			expectedErr: fmt.Errorf("parse ://malformed: missing protocol scheme"),
		},
		"returns response": {
			url:            "http://test.com",
			shouldMock:     true,
			returnResponse: &http.Response{},
			expectedResp:   &http.Response{},
		},
		"returns error": {
			url:         "http://test.com",
			shouldMock:  true,
			returnErr:   fmt.Errorf("error"),
			expectedErr: fmt.Errorf("error"),
		},
	}

	for testName, tt := range testCases {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockClient := NewMockClient(mockCtrl)
			client := NewContext(mockClient)

			var values url.Values
			ctx := context.TODO()
			contentType := "application/x-www-form-urlencoded"
			body := strings.NewReader(values.Encode())

			if tt.shouldMock {
				req, err := http.NewRequest("POST", tt.url, body)
				if err != nil {
					panic(err)
				}
				req.Header.Set("content-type", contentType)
				req = req.WithContext(ctx)

				mockClient.EXPECT().Do(reqMatcher{req}).Return(tt.returnResponse, tt.returnErr)
			}

			resp, err := client.PostFormContext(ctx, tt.url, values)
			if !reflect.DeepEqual(resp, tt.expectedResp) {
				t.Errorf("Incorrect response;\n\texpected: %#v\n\t     got: %#v", tt.expectedResp, resp)
			}

			if (err != nil || tt.expectedErr != nil) && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Incorrect error;\n\texpected: %s\n\t     got: %s", tt.expectedErr, err)
			}
		})
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
