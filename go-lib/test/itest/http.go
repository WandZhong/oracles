package itest

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/go-ozzo/ozzo-routing"
)

// NewRoutingPostCtx returns new POST request with parameters encoded using url-encoding
func NewRoutingPostCtx(values url.Values) *routing.Context {
	// req := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
	var req = http.Request{
		Method:   http.MethodPost,
		PostForm: values}
	return routing.NewContext(httptest.NewRecorder(), &req)
}
