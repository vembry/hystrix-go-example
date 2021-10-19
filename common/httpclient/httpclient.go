//go:generate mockery --all --output=./mock --outpkg=mock
package httpclient

import (
	"context"
	"io"
	"net/http"
)

// HttpClient is the interface that implements all the http request functions.
type HttpClient interface {
	// Get executes an http request with method GET
	Get(path string, headers http.Header) (*http.Response, error)

	// Post executes an http request with method POST
	Post(path string, headers http.Header, body io.Reader) (*http.Response, error)

	// Put executes an http request with method PUT
	Put(path string, headers http.Header, body io.Reader) (*http.Response, error)

	// Delete executes an http request with method DELETE
	Delete(path string, headers http.Header) (*http.Response, error)

	// Do is an in-house "native" httpclient which executes an http request to designated url/host,
	// wrapped with circuit breaker functionality and retry mechanism
	Do(httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error)

	// DoContext is an in-house "native" httpclient which executes an http request to designated url/host,
	// wrapped with circuit breaker functionality and retry mechanism
	// with context in args
	DoContext(ctx context.Context, httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error)

	// CAUTION: USE THIS AT YOUR OWN RISK
	// DoVanilla is a vanilla version of in-house "native" httpclient which executes an http request to designated url/host
	// wrapped with circuit breaker functionality and retry mechanism
	// with fully customized http.Request param
	DoVanilla(req *http.Request) (*http.Response, error)
}
