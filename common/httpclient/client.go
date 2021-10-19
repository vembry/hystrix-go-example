package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
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

type Parameter struct {
	Path          string
	PathVariables []string
	QueryParams   map[string]string
	Header        map[string]string
	Body          interface{}
}

func addQueryString(req *http.Request, queryStrings map[string]string) {
	q := req.URL.Query()
	for key, value := range queryStrings {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
}

// generateBody is a helper to convert header map to http.Header
func generateHeaders(headersPlain map[string]string) http.Header {
	headers := http.Header{}
	for key, value := range headersPlain {
		headers.Set(key, value)
	}
	return headers
}

// generateBody is a helper to convert body object to io.Reader
func generateBody(bodyPlain interface{}) (io.Reader, error) {
	jsonBody, err := json.Marshal(bodyPlain)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer([]byte(jsonBody)), nil
}

// generateUrl is a helper to combine url + path + path-variables
func (hc *Client) generateUrl(param Parameter) string {
	fullPath := hc.config.Host
	if len(param.Path) > 0 && !strings.HasPrefix(param.Path, "/") {
		fullPath = fullPath + "/"
	}
	fullPath = fullPath + param.Path

	if len(param.PathVariables) > 0 {
		fullPath = fmt.Sprintf("%s/%s", fullPath, strings.Join(param.PathVariables, "/"))
	}

	return fullPath
}

// Get executes an http request with method GET
// wrapped with circuit breaker functionality and retry mechanism
func (hc *Client) Get(param Parameter) (*http.Response, error) {
	return hc.Do(http.MethodGet, param)
}

// Get executes an http request with method GET
// wrapped with circuit breaker functionality and retry mechanism
// with context in args
func (hc *Client) GetWithContext(ctx context.Context, param Parameter) (*http.Response, error) {
	return hc.DoContext(ctx, http.MethodGet, param)
}

// Post executes an http request with method POST
// wrapped with circuit breaker functionality and retry mechanism
func (hc *Client) Post(param Parameter) (*http.Response, error) {
	return hc.Do(http.MethodPost, param)
}

// Post executes an http request with method POST
// wrapped with circuit breaker functionality and retry mechanism
// with context in args
func (hc *Client) PostWithContext(ctx context.Context, param Parameter) (*http.Response, error) {
	return hc.DoContext(ctx, http.MethodPost, param)
}

// Put executes an http request with method PUT
// wrapped with circuit breaker functionality and retry mechanism
func (hc *Client) Put(param Parameter) (*http.Response, error) {
	return hc.Do(http.MethodPut, param)
}

// Put executes an http request with method PUT
// wrapped with circuit breaker functionality and retry mechanism
// with context in args
func (hc *Client) PutWithContext(ctx context.Context, param Parameter) (*http.Response, error) {
	return hc.DoContext(ctx, http.MethodPut, param)
}

// Delete executes an http request with method DELETE
// wrapped with circuit breaker functionality and retry mechanism
func (hc *Client) Delete(param Parameter) (*http.Response, error) {
	return hc.Do(http.MethodDelete, param)
}

// Delete executes an http request with method DELETE
// wrapped with circuit breaker functionality and retry mechanism
// with context in args
func (hc *Client) DeleteWithContext(ctx context.Context, param Parameter) (*http.Response, error) {
	return hc.DoContext(ctx, http.MethodDelete, param)
}

// Do is an in-house "native" httpclient which executes an http request to designated url/host
// wrapped with circuit breaker functionality and retry mechanism
func (hc *Client) Do(httpMethod string, param Parameter) (*http.Response, error) {
	return hc.DoContext(context.Background(), httpMethod, param)
}

// DoContext is an in-house "native" httpclient which executes an http request to designated url/host
// wrapped with circuit breaker functionality and retry mechanism
// with context in args
func (hc *Client) DoContext(ctx context.Context, httpMethod string, param Parameter) (*http.Response, error) {

	fullUrl := hc.generateUrl(param)
	headers := generateHeaders(param.Header)
	body, err := generateBody(param.Body)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, errReq := http.NewRequestWithContext(ctx, httpMethod, fullUrl, body)
	if errReq != nil {
		return nil, errReq
	}

	// add query string
	addQueryString(req, param.QueryParams)

	// add headers
	req.Header = headers

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
