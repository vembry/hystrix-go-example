// Code generated by mockery 2.9.4. DO NOT EDIT.

package mock

import (
	context "context"
	http "net/http"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Delete provides a mock function with given fields: path, headers
func (_m *HttpClient) Delete(path string, headers http.Header) (*http.Response, error) {
	ret := _m.Called(path, headers)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, http.Header) *http.Response); ok {
		r0 = rf(path, headers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Header) error); ok {
		r1 = rf(path, headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Do provides a mock function with given fields: httpMethod, path, header, body
func (_m *HttpClient) Do(httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error) {
	ret := _m.Called(httpMethod, path, header, body)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, string, http.Header, io.Reader) *http.Response); ok {
		r0 = rf(httpMethod, path, header, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, http.Header, io.Reader) error); ok {
		r1 = rf(httpMethod, path, header, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoContext provides a mock function with given fields: ctx, httpMethod, path, header, body
func (_m *HttpClient) DoContext(ctx context.Context, httpMethod string, path string, header http.Header, body io.Reader) (*http.Response, error) {
	ret := _m.Called(ctx, httpMethod, path, header, body)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, http.Header, io.Reader) *http.Response); ok {
		r0 = rf(ctx, httpMethod, path, header, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, http.Header, io.Reader) error); ok {
		r1 = rf(ctx, httpMethod, path, header, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoVanilla provides a mock function with given fields: req
func (_m *HttpClient) DoVanilla(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: path, headers
func (_m *HttpClient) Get(path string, headers http.Header) (*http.Response, error) {
	ret := _m.Called(path, headers)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, http.Header) *http.Response); ok {
		r0 = rf(path, headers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Header) error); ok {
		r1 = rf(path, headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: path, headers, body
func (_m *HttpClient) Post(path string, headers http.Header, body io.Reader) (*http.Response, error) {
	ret := _m.Called(path, headers, body)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, http.Header, io.Reader) *http.Response); ok {
		r0 = rf(path, headers, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Header, io.Reader) error); ok {
		r1 = rf(path, headers, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: path, headers, body
func (_m *HttpClient) Put(path string, headers http.Header, body io.Reader) (*http.Response, error) {
	ret := _m.Called(path, headers, body)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, http.Header, io.Reader) *http.Response); ok {
		r0 = rf(path, headers, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Header, io.Reader) error); ok {
		r1 = rf(path, headers, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
