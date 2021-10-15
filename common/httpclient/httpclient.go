//go:generate mockery --all --output=./mock --outpkg=mock
package httpclient

import (
	"io"
	"net/http"
)

type HttpClient interface {

	// to do an http request GET
	Get(url string, headers http.Header) (*http.Response, error)

	// to do an http request POST
	Post(path string, headers http.Header, body io.Reader) (*http.Response, error)

	// to do an http request PUT
	Put(url string, headers http.Header, body io.Reader) (*http.Response, error)

	// to do an http request DELETE
	Delete(url string, headers http.Header) (*http.Response, error)

	// in-house "native" httpclient
	// wrapped with circuit breaker, and is active if initialized using it
	Do(httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error)
}
