/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package model

import (
	values "github.com/apache/plc4x/plc4go/pkg/api/values"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcReadResponse is an autogenerated mock type for the PlcReadResponse type
type MockPlcReadResponse struct {
	mock.Mock
}

type MockPlcReadResponse_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcReadResponse) EXPECT() *MockPlcReadResponse_Expecter {
	return &MockPlcReadResponse_Expecter{mock: &_m.Mock}
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcReadResponse) GetRequest() PlcReadRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRequest")
	}

	var r0 PlcReadRequest
	if rf, ok := ret.Get(0).(func() PlcReadRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcReadRequest)
		}
	}

	return r0
}

// MockPlcReadResponse_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcReadResponse_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcReadResponse_Expecter) GetRequest() *MockPlcReadResponse_GetRequest_Call {
	return &MockPlcReadResponse_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcReadResponse_GetRequest_Call) Run(run func()) *MockPlcReadResponse_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadResponse_GetRequest_Call) Return(_a0 PlcReadRequest) *MockPlcReadResponse_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadResponse_GetRequest_Call) RunAndReturn(run func() PlcReadRequest) *MockPlcReadResponse_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponseCode provides a mock function with given fields: tagName
func (_m *MockPlcReadResponse) GetResponseCode(tagName string) PlcResponseCode {
	ret := _m.Called(tagName)

	if len(ret) == 0 {
		panic("no return value specified for GetResponseCode")
	}

	var r0 PlcResponseCode
	if rf, ok := ret.Get(0).(func(string) PlcResponseCode); ok {
		r0 = rf(tagName)
	} else {
		r0 = ret.Get(0).(PlcResponseCode)
	}

	return r0
}

// MockPlcReadResponse_GetResponseCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponseCode'
type MockPlcReadResponse_GetResponseCode_Call struct {
	*mock.Call
}

// GetResponseCode is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcReadResponse_Expecter) GetResponseCode(tagName interface{}) *MockPlcReadResponse_GetResponseCode_Call {
	return &MockPlcReadResponse_GetResponseCode_Call{Call: _e.mock.On("GetResponseCode", tagName)}
}

func (_c *MockPlcReadResponse_GetResponseCode_Call) Run(run func(tagName string)) *MockPlcReadResponse_GetResponseCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcReadResponse_GetResponseCode_Call) Return(_a0 PlcResponseCode) *MockPlcReadResponse_GetResponseCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadResponse_GetResponseCode_Call) RunAndReturn(run func(string) PlcResponseCode) *MockPlcReadResponse_GetResponseCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcReadResponse) GetTagNames() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTagNames")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockPlcReadResponse_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcReadResponse_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcReadResponse_Expecter) GetTagNames() *MockPlcReadResponse_GetTagNames_Call {
	return &MockPlcReadResponse_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcReadResponse_GetTagNames_Call) Run(run func()) *MockPlcReadResponse_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadResponse_GetTagNames_Call) Return(_a0 []string) *MockPlcReadResponse_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadResponse_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcReadResponse_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// GetValue provides a mock function with given fields: tagName
func (_m *MockPlcReadResponse) GetValue(tagName string) values.PlcValue {
	ret := _m.Called(tagName)

	if len(ret) == 0 {
		panic("no return value specified for GetValue")
	}

	var r0 values.PlcValue
	if rf, ok := ret.Get(0).(func(string) values.PlcValue); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(values.PlcValue)
		}
	}

	return r0
}

// MockPlcReadResponse_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type MockPlcReadResponse_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcReadResponse_Expecter) GetValue(tagName interface{}) *MockPlcReadResponse_GetValue_Call {
	return &MockPlcReadResponse_GetValue_Call{Call: _e.mock.On("GetValue", tagName)}
}

func (_c *MockPlcReadResponse_GetValue_Call) Run(run func(tagName string)) *MockPlcReadResponse_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcReadResponse_GetValue_Call) Return(_a0 values.PlcValue) *MockPlcReadResponse_GetValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadResponse_GetValue_Call) RunAndReturn(run func(string) values.PlcValue) *MockPlcReadResponse_GetValue_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcReadResponse) IsAPlcMessage() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAPlcMessage")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcReadResponse_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcReadResponse_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcReadResponse_Expecter) IsAPlcMessage() *MockPlcReadResponse_IsAPlcMessage_Call {
	return &MockPlcReadResponse_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcReadResponse_IsAPlcMessage_Call) Run(run func()) *MockPlcReadResponse_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadResponse_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcReadResponse_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadResponse_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcReadResponse_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcReadResponse) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcReadResponse_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcReadResponse_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcReadResponse_Expecter) String() *MockPlcReadResponse_String_Call {
	return &MockPlcReadResponse_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcReadResponse_String_Call) Run(run func()) *MockPlcReadResponse_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadResponse_String_Call) Return(_a0 string) *MockPlcReadResponse_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadResponse_String_Call) RunAndReturn(run func() string) *MockPlcReadResponse_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcReadResponse creates a new instance of MockPlcReadResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcReadResponse(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcReadResponse {
	mock := &MockPlcReadResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
