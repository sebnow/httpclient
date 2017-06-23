package httpclient

import (
	"net/http"
	"testing"
)

func TestInterfaceIsCompatibleWithDefaultClient(t *testing.T) {
	var client interface{} = &http.Client{}
	if _, ok := client.(Client); !ok {
		t.Error("Client interface is not compatible with http.DefaultClient")
	}
}
