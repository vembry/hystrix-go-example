//go:generate mockery --all --output=./mock --outpkg=mock
package httpclient

import (
	"context"
	"net/http"
)

// HttpClient is the interface that implements all the http request functions.
type HttpClient interface {
	// Get executes an http request with method GET
	Get(param Parameter) (*http.Response, error)

	// Get executes an http request with method GET
	// with context in args
	GetWithContext(ctx context.Context, param Parameter) (*http.Response, error)

	// Post executes an http request with method POST
	Post(param Parameter) (*http.Response, error)

	// Post executes an http request with method POST
	// with context in args
	PostWithContext(ctx context.Context, param Parameter) (*http.Response, error)

	// Put executes an http request with method PUT
	Put(param Parameter) (*http.Response, error)

	// Put executes an http request with method PUT
	// with context in args
	PutWithContext(ctx context.Context, param Parameter) (*http.Response, error)

	// Delete executes an http request with method DELETE
	Delete(param Parameter) (*http.Response, error)

	// Delete executes an http request with method DELETE
	// with context in args
	DeleteWithContext(ctx context.Context, param Parameter) (*http.Response, error)

	// Do is an in-house "native" httpclient which executes an http request to designated url/host,
	// wrapped with circuit breaker functionality and retry mechanism
	Do(httpMethod string, param Parameter) (*http.Response, error)

	// DoContext is an in-house "native" httpclient which executes an http request to designated url/host,
	// wrapped with circuit breaker functionality and retry mechanism
	// with context in args
	DoContext(ctx context.Context, httpMethod string, param Parameter) (*http.Response, error)

	// CAUTION: USE THIS AT YOUR OWN RISK
	// DoVanilla is a vanilla version of in-house "native" httpclient which executes an http request to designated url/host
	// wrapped with circuit breaker functionality and retry mechanism
	// with fully customized http.Request param
	DoVanilla(req *http.Request) (*http.Response, error)
}
