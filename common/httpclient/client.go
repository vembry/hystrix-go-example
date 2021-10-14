package httpclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

//httpclient
type Client struct {
	client *http.Client
	config Config
}

//Circuit Breaker config
type CircuitBreakerConfig struct {
	SleepWindow    int         // in ms, to wait after a circuit opens before testing for recovery
	ErrorThreshold int         // minimum number of requests needed before a circuit can be tripped
	Timeout        int         // in ms, how long to wait for command to complete
	Fallback       func(error) // custom fallback function
}

// default values for CircuitBreakerConfig
const (
	defautCbSleepWindow     = 5000
	defaultCbErrorThreshold = 20
	defaultCbTimeout        = 10000
)

//httpclient config
type Config struct {
	Host                  string               // external host
	Timeout               time.Duration        // http request timeout
	RetryCount            int                  // failing http request retry
	IsUsingCircuitBreaker bool                 // flag to use circuit breaker, true = on, false = off
	CbConfig              CircuitBreakerConfig // custom config for circuit breaker

	//Logger Logger // logging stuff
}

// default values for Config
const (
	defautTimeout = 10 * time.Second
)

//initialization
func NewHttpClient(config Config) HttpClient {

	// force delete '/' at host's value suffix
	config.Host = strings.TrimSuffix(config.Host, "/")

	// set default value if default config not defined
	if config.Timeout == 0 {
		config.Timeout = defautTimeout
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
			config.CbConfig.Fallback = func(error) {}
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

//preparing request Get
func (hc *Client) Get(path string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, errRequest := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", hc.config.Host, path), nil)
	if errRequest != nil {
		return response, errRequest
	}
	request.Header = headers

	return hc.Do(request)
}

//preparing request Post
func (hc *Client) Post(path string, headers http.Header, body io.Reader) (*http.Response, error) {
	var response *http.Response
	request, errRequest := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", hc.config.Host, path), body)
	if errRequest != nil {
		return response, errRequest
	}
	request.Header = headers
	return hc.Do(request)
}

//preparing request Put
func (hc *Client) Put(path string, headers http.Header, body io.Reader) (*http.Response, error) {
	var response *http.Response
	request, errRequest := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", hc.config.Host, path), body)
	if errRequest != nil {
		return response, errRequest
	}
	request.Header = headers
	return hc.Do(request)
}

//preparing request Delete
func (hc *Client) Delete(path string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, errRequest := http.NewRequest(http.MethodDelete, generateFullUrl(hc.config.Host, path), nil)
	if errRequest != nil {
		return response, errRequest
	}
	request.Header = headers
	return hc.Do(request)
}

// helper
// to combine url + path
func generateFullUrl(url string, path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", url, path)
}

// circuit breaker wrapper
func (hc *Client) Do(request *http.Request) (*http.Response, error) {
	if hc.config.IsUsingCircuitBreaker {
		//execute with circuit breaker
		var response *http.Response
		err := hystrix.Do(hc.config.Host, func() error {
			responseHys, err := hc.doActual(request)
			if err == nil {
				response = responseHys
			}
			return err
		}, func(e error) error {
			//circuit breaker fallback
			if hc.config.CbConfig.Fallback != nil {
				hc.config.CbConfig.Fallback(e)
			}
			return e
		})
		return response, err
	} else {
		//execute without circuit breaker
		return hc.doActual(request)
	}

}

//actual request execution
func (hc *Client) doActual(request *http.Request) (*http.Response, error) {
	if len(hc.config.Host) == 0 {
		return nil, errors.New("missing host")
	}

	//execute request
	response, errResponse := hc.client.Do(request)

	//request validation
	if errResponse != nil {
		var responseRetry *http.Response
		errResponseRetry := errResponse

		//retry mechanism
		for i := 0; i < hc.config.RetryCount; i++ {

			//re-execute request
			responseRetry, errResponseRetry = hc.client.Do(request)

			//success retry will break the loop
			if errResponseRetry == nil {
				response = responseRetry
				break
			}
		}

		//request validation
		if errResponseRetry != nil {
			//request failed
			return response, errResponse
		}
	}

	//request succeeded
	return response, nil
}
