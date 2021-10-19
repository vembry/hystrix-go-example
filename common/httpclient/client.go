package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// Client is a wrapper around net/http
// which have circuit-breaker alike functionality
type Client struct {
	client *http.Client // http client, using native golang net's http
	config Config       // configs
}

// CircuitBreakerConfig is the circuit breaker's configuration implemented
// inside the HttpClient wrapper
type CircuitBreakerConfig struct {
	SleepWindow    int                                // in ms, to wait after a circuit opens before testing for recovery
	ErrorThreshold int                                // minimum number of requests needed before a circuit can be tripped
	Timeout        int                                // in ms, how long to wait for command to complete
	Fallback       func(context.Context, error) error // custom fallback function
}

// default values for CircuitBreakerConfig
const (
	defautCbSleepWindow     = 5000
	defaultCbErrorThreshold = 20
	defaultCbTimeout        = 10000
)

// Config is the HttpClient configuration
type Config struct {
	Host                  string                    // external host
	Timeout               time.Duration             // http request timeout
	RetryCount            int                       // failing http request retry
	OnPreRetryCallback    func(*http.Request) error // callback called on every pre-retry
	IsUsingCircuitBreaker bool                      // flag to use circuit breaker, true = on, false = off
	CbConfig              CircuitBreakerConfig      // custom config for circuit breaker
}

// default values for Config
const (
	defautTimeout = 10 * time.Second
)

// NewHttpClient initialises the Client handle that is used to wrap net/http
func NewHttpClient(config Config) HttpClient {
	// force delete '/' at host's value suffix
	config.Host = strings.TrimSuffix(config.Host, "/")

	// set default value if default config not defined
	if config.Timeout == 0 {
		config.Timeout = defautTimeout
	}

	// set default value if default config not defined
	if config.OnPreRetryCallback == nil {
		config.OnPreRetryCallback = func(r *http.Request) error {
			// deliberately putting empty function so no need if validation on actual execution
			return nil
		}
	}

	// configure circuit breaker
	if config.IsUsingCircuitBreaker {
		// set default value if not defined
		if config.CbConfig.SleepWindow == 0 {
			config.CbConfig.SleepWindow = defautCbSleepWindow
		}

		// set default value if not defined
		if config.CbConfig.ErrorThreshold == 0 {
			config.CbConfig.ErrorThreshold = defaultCbErrorThreshold
		}

		// set default value if not defined
		if config.CbConfig.Fallback == nil {
			config.CbConfig.Fallback = func(ctx context.Context, e error) error {
				// deliberately putting empty function so no need if validation on actual execution
				return e
			}
		}

		// set default value if not defined
		if config.CbConfig.Timeout == 0 {
			config.CbConfig.Timeout = defaultCbTimeout
		}

		// initialize circuit breaker
		// using afex/hystrix-go lib
		// please check hystrix-go lib for further usage
		hystrix.ConfigureCommand(config.Host, hystrix.CommandConfig{
			Timeout:                config.CbConfig.Timeout,
			SleepWindow:            config.CbConfig.SleepWindow,
			RequestVolumeThreshold: config.CbConfig.ErrorThreshold,
		})
	}

	return &Client{
		config: config,
		client: &http.Client{Timeout: config.Timeout},
	}
}

// Get executes an http request with method GET
func (hc *Client) Get(path string, headers http.Header) (*http.Response, error) {
	return hc.Do(http.MethodGet, path, headers, nil)
}

// Post executes an http request with method POST
func (hc *Client) Post(path string, headers http.Header, body io.Reader) (*http.Response, error) {
	return hc.Do(http.MethodPost, path, headers, body)
}

// Put executes an http request with method PUT
func (hc *Client) Put(path string, headers http.Header, body io.Reader) (*http.Response, error) {
	return hc.Do(http.MethodPut, path, headers, body)
}

// Delete executes an http request with method DELETE
func (hc *Client) Delete(path string, headers http.Header) (*http.Response, error) {
	return hc.Do(http.MethodDelete, path, headers, nil)
}

// generateFullUrl is a helper to combine url + path
func generateFullUrl(url string, path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", url, path)
}

// Do is an in-house "native" httpclient which executes an http request to designated url/host
// wrapped with circuit breaker functionality and retry mechanism
func (hc *Client) Do(httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error) {
	return hc.DoContext(context.Background(), httpMethod, path, header, body)
}

// DoContext is an in-house "native" httpclient which executes an http request to designated url/host
// wrapped with circuit breaker functionality and retry mechanism
// with context in args
func (hc *Client) DoContext(ctx context.Context, httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error) {
	// get full url, combination of host and path
	fullUrl := generateFullUrl(hc.config.Host, path)

	// initialize request
	req, errReq := http.NewRequestWithContext(ctx, httpMethod, fullUrl, body)
	if errReq != nil {
		return nil, errReq
	}

	return hc.DoVanilla(req)
}

// CAUTION: USE THIS AT YOUR OWN RISK
// DoVanilla is a vanilla version of in-house "native" httpclient which executes an http request to designated url/host
// wrapped with circuit breaker functionality and retry mechanism
// with fully customized http.Request param
func (hc *Client) DoVanilla(req *http.Request) (*http.Response, error) {
	// execute request
	res, errRes := hc.doActual(req)

	// request validation
	if errRes != nil {
		// retry mechanism
		for i := 0; i < hc.config.RetryCount; i++ {

			// pre-retry callback
			errRetryCallback := hc.config.OnPreRetryCallback(req)
			if errRetryCallback != nil {
				// failing on pre-retry callback will stop the retry mechanism
				errRes = errRetryCallback
				break
			}

			// exponential wait time for retry
			time.Sleep(time.Duration(i+1) * time.Second)

			// re-execute request
			res, errRes = hc.doActual(req)

			//success retry will break the loop
			if errRes == nil {
				break
			}
		}
	}

	return res, errRes

}

// doActual is an in-house "native" httpclient which executes an http request to designated url/host
// wrapped with circuit breaker functionality
func (hc *Client) doActual(req *http.Request) (*http.Response, error) {
	// executes without circuit breaker
	if !hc.config.IsUsingCircuitBreaker {
		return hc.client.Do(req)
	}

	// executes with circuit breaker
	var response *http.Response
	err := hystrix.DoC(req.Context(), req.URL.String(), func(ctx context.Context) error {
		var errResponse error
		response, errResponse = hc.client.Do(req)
		return errResponse
	}, func(ctx context.Context, e error) error {

		// im defining a return here for readability
		return hc.config.CbConfig.Fallback(ctx, e)
	})

	return response, err
}
