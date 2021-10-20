// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	context "context"
	http "net/http"
	"playground/common/httpclient"

	mock "github.com/stretchr/testify/mock"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Delete provides a mock function with given fields: param
func (_m *HttpClient) Delete(param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(httpclient.Parameter) *http.Response); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(httpclient.Parameter) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWithContext provides a mock function with given fields: ctx, param
func (_m *HttpClient) DeleteWithContext(ctx context.Context, param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(ctx, param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, httpclient.Parameter) *http.Response); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, httpclient.Parameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Do provides a mock function with given fields: httpMethod, param
func (_m *HttpClient) Do(httpMethod string, param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(httpMethod, param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, httpclient.Parameter) *http.Response); ok {
		r0 = rf(httpMethod, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, httpclient.Parameter) error); ok {
		r1 = rf(httpMethod, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoContext provides a mock function with given fields: ctx, httpMethod, param
func (_m *HttpClient) DoContext(ctx context.Context, httpMethod string, param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(ctx, httpMethod, param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, httpclient.Parameter) *http.Response); ok {
		r0 = rf(ctx, httpMethod, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, httpclient.Parameter) error); ok {
		r1 = rf(ctx, httpMethod, param)
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

// Get provides a mock function with given fields: param
func (_m *HttpClient) Get(param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(httpclient.Parameter) *http.Response); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(httpclient.Parameter) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithContext provides a mock function with given fields: ctx, param
func (_m *HttpClient) GetWithContext(ctx context.Context, param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(ctx, param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, httpclient.Parameter) *http.Response); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, httpclient.Parameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: param
func (_m *HttpClient) Post(param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(httpclient.Parameter) *http.Response); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(httpclient.Parameter) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostWithContext provides a mock function with given fields: ctx, param
func (_m *HttpClient) PostWithContext(ctx context.Context, param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(ctx, param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, httpclient.Parameter) *http.Response); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, httpclient.Parameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: param
func (_m *HttpClient) Put(param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(httpclient.Parameter) *http.Response); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(httpclient.Parameter) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutWithContext provides a mock function with given fields: ctx, param
func (_m *HttpClient) PutWithContext(ctx context.Context, param httpclient.Parameter) (*http.Response, error) {
	ret := _m.Called(ctx, param)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, httpclient.Parameter) *http.Response); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, httpclient.Parameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
