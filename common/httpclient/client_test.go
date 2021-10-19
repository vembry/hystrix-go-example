package httpclient

import (
	"context"
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
func createTestServer() *httptest.Server {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{ "response": "ok" }`))
		if err != nil {
			fmt.Print("something went wrong while initializing dummy server")
		}
	}
	return httptest.NewServer(http.HandlerFunc(dummyHandler))
}

// this is just a helper
func createTestClient(host string) HttpClient {
	return NewHttpClient(Config{
		Host:       host,
		RetryCount: 1,
		OnPreRetryCallback: func(r *http.Request) error {
			return nil
		},
		IsUsingCircuitBreaker: true,
	})
}

// this is just a helper
func createTestParameter() Parameter {
	type Body struct {
		SomeKey string
	}
	return Parameter{
		Path: "asd",
		PathVariables: []string{
			"some-path-variable",
		},
		QueryParams: map[string]string{
			"some-query-param": "some-query-param-value",
		},
		Header: map[string]string{
			"x-some-header": "some-header-value",
		},
		Body: Body{
			SomeKey: "SomeValue",
		},
	}
}

func Test_HttpClient_Initialization(t *testing.T) {
	createTestClient(testHost)
}

func Test_Get_Success(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.Get(parameter)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_GetWithContext_Success(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.GetWithContext(context.Background(), parameter)
	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Get_Success_WithCircuitBreaker(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.Get(parameter)

	require.NoError(t, err, "should not have failed to make a GET request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

// testing without circuit breaker
func Test_Get_Failed_OnRequesting(t *testing.T) {
	client := NewHttpClient(Config{
		Host: "testHost",
	})
	parameter := createTestParameter()

	_, err := client.Get(parameter)

	require.Error(t, err, "should have failed to make a GET request: failed to request to host")
}

// testing without circuit breaker
func Test_Get_Failed_WithRetry(t *testing.T) {
	client := NewHttpClient(Config{
		Host:       "some-host",
		RetryCount: 1,
	})
	parameter := createTestParameter()

	_, err := client.Get(parameter)

	require.Error(t, err, "should have failed to make a GET request: missing host")
}

// testing with circuit breaker
func Test_Get_Failed_WithRetry_OnPreCallbackFailing(t *testing.T) {
	client := NewHttpClient(Config{
		Host:       "some-host",
		RetryCount: 1,
		OnPreRetryCallback: func(r *http.Request) error {
			return errors.New("something")
		},
		IsUsingCircuitBreaker: true,
	})
	parameter := createTestParameter()

	_, err := client.Get(parameter)

	require.Error(t, err, "should have failed to make a GET request: error pre-retry")
}

func Test_Post_Success(t *testing.T) {

	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.Post(parameter)

	require.NoError(t, err, "should not have failed to make a POST request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Post_Failed_JsonMarshal(t *testing.T) {

	client := createTestClient(testHost)
	parameter := Parameter{
		Path: "/path",
		Body: make(chan int),
	}

	_, err := client.Post(parameter)

	require.Error(t, err, "should have failed to make a GET request: failed marshal body")
}

func Test_PostWithContext_Success(t *testing.T) {

	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.PostWithContext(context.Background(), parameter)

	require.NoError(t, err, "should not have failed to make a POST request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Put_Success(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.Put(parameter)
	require.NoError(t, err, "should not have failed to make a PUT request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_PutWithContext_Success(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.PutWithContext(context.Background(), parameter)
	require.NoError(t, err, "should not have failed to make a PUT request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Delete_Success(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.Delete(parameter)
	require.NoError(t, err, "should not have failed to make a DELETE request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_DeleteWithContext_Success(t *testing.T) {
	server := createTestServer()
	defer server.Close()

	client := createTestClient(server.URL)
	parameter := createTestParameter()

	response, err := client.DeleteWithContext(context.Background(), parameter)
	require.NoError(t, err, "should not have failed to make a DELETE request")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)
	assert.Equal(t, "{ \"response\": \"ok\" }", string(body))
}

func Test_Do_Failed_OnPreparingRequest(t *testing.T) {
	server := createTestServer()
	defer server.Close()
	client := createTestClient(server.URL)
	parameter := createTestParameter()

	// sorry had to do this "Î",
	// to create a scenario which is a failure during request preparation
	_, err := client.Do("Î", parameter)
	require.Error(t, err)
}
