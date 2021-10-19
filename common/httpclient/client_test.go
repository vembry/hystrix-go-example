package httpclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testHost = "http://some-host"
)

// this is just a helper
func createDummyServer() *httptest.Server {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{ "response": "ok" }`))
		if err != nil {
			fmt.Print("something went wrong while initializing dummy server")
		}
	}
	return httptest.NewServer(http.HandlerFunc(dummyHandler))
}

func Test_HttpClient_Initialization(t *testing.T) {
	NewHttpClient(Config{
		Host: testHost,
		OnPreRetryCallback: func(r *http.Request) error {
			return nil
		},
		IsUsingCircuitBreaker: true,
	})
}

func Test_Get_Success(t *testing.T) {
	server := createDummyServer()
	defer server.Close()

	client := NewHttpClient(Config{
		Host: server.URL,
		OnPreRetryCallback: func(r *http.Request) error {
			return nil
		},
	})

	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	response, err := client.Get("", header)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Get_Success_WithCircuitBreaker(t *testing.T) {
	server := createDummyServer()
	defer server.Close()

	client := NewHttpClient(Config{
		Host:                  server.URL,
		IsUsingCircuitBreaker: true,
	})

	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	response, err := client.Get("", header)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Get_Failed(t *testing.T) {
	client := NewHttpClient(Config{
		Host: testHost,
	})

	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	_, err := client.Get("", header)
	require.Error(t, err, "should have failed to make a GET request: failed to request to host")
}

func Test_Get_Failed_WithCircuitBreaker(t *testing.T) {
	client := NewHttpClient(Config{
		Host:                  testHost,
		IsUsingCircuitBreaker: true,
	})
	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	_, err := client.Get("", header)
	require.Error(t, err, "should have failed to make a GET request: failed to request to host")
}

func Test_Get_Failed_OnPreparingRequest(t *testing.T) {
	client := NewHttpClient(Config{
		Host: "testHost",
	})
	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	_, err := client.Get("", header)

	require.Error(t, err, "should have failed to make a GET request: failed to request to host")
}

func Test_Get_Failed_WithRetry(t *testing.T) {
	client := NewHttpClient(Config{
		Host:       testHost,
		RetryCount: 1,
	})
	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	_, err := client.Get("", header)
	require.Error(t, err, "should have failed to make a GET request: missing host")
}

func Test_Get_Failed_WithRetry_OnPreCallbackFailing(t *testing.T) {
	client := NewHttpClient(Config{
		Host:       testHost,
		RetryCount: 1,
		OnPreRetryCallback: func(r *http.Request) error {
			return errors.New("error pre-retry")
		},
	})
	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	_, err := client.Get("", header)
	require.Error(t, err, "should have failed to make a GET request: error pre-retry")
}

func Test_Post_Success(t *testing.T) {
	server := createDummyServer()
	defer server.Close()

	client := NewHttpClient(Config{
		Host: server.URL,
	})

	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	response, err := client.Post("", header, nil)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Put_Success(t *testing.T) {
	server := createDummyServer()
	defer server.Close()

	client := NewHttpClient(Config{
		Host: server.URL,
	})

	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	response, err := client.Put("", header, nil)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Delete_Success(t *testing.T) {
	server := createDummyServer()
	defer server.Close()

	client := NewHttpClient(Config{
		Host: server.URL,
	})

	header := http.Header{}
	header.Set("x-some-header", "some-header-value")

	response, err := client.Delete("", header)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Do_Failed_OnPreparingRequest(t *testing.T) {
	client := NewHttpClient(Config{
		Host: testHost,
	})

	// sorry had to do this "Î",
	// to create a scenario which is a failure during request preparation
	_, err := client.Do("Î", "", nil, nil)
	require.Error(t, err)
}
