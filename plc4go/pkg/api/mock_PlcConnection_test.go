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

package plc4go

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcConnection is an autogenerated mock type for the PlcConnection type
type MockPlcConnection struct {
	mock.Mock
}

type MockPlcConnection_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConnection) EXPECT() *MockPlcConnection_Expecter {
	return &MockPlcConnection_Expecter{mock: &_m.Mock}
}

// BlockingClose provides a mock function with given fields:
func (_m *MockPlcConnection) BlockingClose() {
	_m.Called()
}

// MockPlcConnection_BlockingClose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockingClose'
type MockPlcConnection_BlockingClose_Call struct {
	*mock.Call
}

// BlockingClose is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) BlockingClose() *MockPlcConnection_BlockingClose_Call {
	return &MockPlcConnection_BlockingClose_Call{Call: _e.mock.On("BlockingClose")}
}

func (_c *MockPlcConnection_BlockingClose_Call) Run(run func()) *MockPlcConnection_BlockingClose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_BlockingClose_Call) Return() *MockPlcConnection_BlockingClose_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPlcConnection_BlockingClose_Call) RunAndReturn(run func()) *MockPlcConnection_BlockingClose_Call {
	_c.Call.Return(run)
	return _c
}

// BrowseRequestBuilder provides a mock function with given fields:
func (_m *MockPlcConnection) BrowseRequestBuilder() model.PlcBrowseRequestBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BrowseRequestBuilder")
	}

	var r0 model.PlcBrowseRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcBrowseRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcBrowseRequestBuilder)
		}
	}

	return r0
}

// MockPlcConnection_BrowseRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BrowseRequestBuilder'
type MockPlcConnection_BrowseRequestBuilder_Call struct {
	*mock.Call
}

// BrowseRequestBuilder is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) BrowseRequestBuilder() *MockPlcConnection_BrowseRequestBuilder_Call {
	return &MockPlcConnection_BrowseRequestBuilder_Call{Call: _e.mock.On("BrowseRequestBuilder")}
}

func (_c *MockPlcConnection_BrowseRequestBuilder_Call) Run(run func()) *MockPlcConnection_BrowseRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_BrowseRequestBuilder_Call) Return(_a0 model.PlcBrowseRequestBuilder) *MockPlcConnection_BrowseRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_BrowseRequestBuilder_Call) RunAndReturn(run func() model.PlcBrowseRequestBuilder) *MockPlcConnection_BrowseRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockPlcConnection) Close() <-chan PlcConnectionCloseResult {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 <-chan PlcConnectionCloseResult
	if rf, ok := ret.Get(0).(func() <-chan PlcConnectionCloseResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionCloseResult)
		}
	}

	return r0
}

// MockPlcConnection_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPlcConnection_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) Close() *MockPlcConnection_Close_Call {
	return &MockPlcConnection_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockPlcConnection_Close_Call) Run(run func()) *MockPlcConnection_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_Close_Call) Return(_a0 <-chan PlcConnectionCloseResult) *MockPlcConnection_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_Close_Call) RunAndReturn(run func() <-chan PlcConnectionCloseResult) *MockPlcConnection_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockPlcConnection) Connect() <-chan PlcConnectionConnectResult {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 <-chan PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func() <-chan PlcConnectionConnectResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockPlcConnection_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockPlcConnection_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) Connect() *MockPlcConnection_Connect_Call {
	return &MockPlcConnection_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockPlcConnection_Connect_Call) Run(run func()) *MockPlcConnection_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_Connect_Call) Return(_a0 <-chan PlcConnectionConnectResult) *MockPlcConnection_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_Connect_Call) RunAndReturn(run func() <-chan PlcConnectionConnectResult) *MockPlcConnection_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockPlcConnection) ConnectWithContext(ctx context.Context) <-chan PlcConnectionConnectResult {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ConnectWithContext")
	}

	var r0 <-chan PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan PlcConnectionConnectResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockPlcConnection_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockPlcConnection_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPlcConnection_Expecter) ConnectWithContext(ctx interface{}) *MockPlcConnection_ConnectWithContext_Call {
	return &MockPlcConnection_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockPlcConnection_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockPlcConnection_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPlcConnection_ConnectWithContext_Call) Return(_a0 <-chan PlcConnectionConnectResult) *MockPlcConnection_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_ConnectWithContext_Call) RunAndReturn(run func(context.Context) <-chan PlcConnectionConnectResult) *MockPlcConnection_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetadata provides a mock function with given fields:
func (_m *MockPlcConnection) GetMetadata() model.PlcConnectionMetadata {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 model.PlcConnectionMetadata
	if rf, ok := ret.Get(0).(func() model.PlcConnectionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcConnectionMetadata)
		}
	}

	return r0
}

// MockPlcConnection_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type MockPlcConnection_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) GetMetadata() *MockPlcConnection_GetMetadata_Call {
	return &MockPlcConnection_GetMetadata_Call{Call: _e.mock.On("GetMetadata")}
}

func (_c *MockPlcConnection_GetMetadata_Call) Run(run func()) *MockPlcConnection_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_GetMetadata_Call) Return(_a0 model.PlcConnectionMetadata) *MockPlcConnection_GetMetadata_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_GetMetadata_Call) RunAndReturn(run func() model.PlcConnectionMetadata) *MockPlcConnection_GetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// IsConnected provides a mock function with given fields:
func (_m *MockPlcConnection) IsConnected() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsConnected")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcConnection_IsConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConnected'
type MockPlcConnection_IsConnected_Call struct {
	*mock.Call
}

// IsConnected is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) IsConnected() *MockPlcConnection_IsConnected_Call {
	return &MockPlcConnection_IsConnected_Call{Call: _e.mock.On("IsConnected")}
}

func (_c *MockPlcConnection_IsConnected_Call) Run(run func()) *MockPlcConnection_IsConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_IsConnected_Call) Return(_a0 bool) *MockPlcConnection_IsConnected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_IsConnected_Call) RunAndReturn(run func() bool) *MockPlcConnection_IsConnected_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields:
func (_m *MockPlcConnection) Ping() <-chan PlcConnectionPingResult {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 <-chan PlcConnectionPingResult
	if rf, ok := ret.Get(0).(func() <-chan PlcConnectionPingResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionPingResult)
		}
	}

	return r0
}

// MockPlcConnection_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type MockPlcConnection_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) Ping() *MockPlcConnection_Ping_Call {
	return &MockPlcConnection_Ping_Call{Call: _e.mock.On("Ping")}
}

func (_c *MockPlcConnection_Ping_Call) Run(run func()) *MockPlcConnection_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_Ping_Call) Return(_a0 <-chan PlcConnectionPingResult) *MockPlcConnection_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_Ping_Call) RunAndReturn(run func() <-chan PlcConnectionPingResult) *MockPlcConnection_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// ReadRequestBuilder provides a mock function with given fields:
func (_m *MockPlcConnection) ReadRequestBuilder() model.PlcReadRequestBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadRequestBuilder")
	}

	var r0 model.PlcReadRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcReadRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcReadRequestBuilder)
		}
	}

	return r0
}

// MockPlcConnection_ReadRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadRequestBuilder'
type MockPlcConnection_ReadRequestBuilder_Call struct {
	*mock.Call
}

// ReadRequestBuilder is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) ReadRequestBuilder() *MockPlcConnection_ReadRequestBuilder_Call {
	return &MockPlcConnection_ReadRequestBuilder_Call{Call: _e.mock.On("ReadRequestBuilder")}
}

func (_c *MockPlcConnection_ReadRequestBuilder_Call) Run(run func()) *MockPlcConnection_ReadRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_ReadRequestBuilder_Call) Return(_a0 model.PlcReadRequestBuilder) *MockPlcConnection_ReadRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_ReadRequestBuilder_Call) RunAndReturn(run func() model.PlcReadRequestBuilder) *MockPlcConnection_ReadRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcConnection) String() string {
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

// MockPlcConnection_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcConnection_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) String() *MockPlcConnection_String_Call {
	return &MockPlcConnection_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcConnection_String_Call) Run(run func()) *MockPlcConnection_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_String_Call) Return(_a0 string) *MockPlcConnection_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_String_Call) RunAndReturn(run func() string) *MockPlcConnection_String_Call {
	_c.Call.Return(run)
	return _c
}

// SubscriptionRequestBuilder provides a mock function with given fields:
func (_m *MockPlcConnection) SubscriptionRequestBuilder() model.PlcSubscriptionRequestBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SubscriptionRequestBuilder")
	}

	var r0 model.PlcSubscriptionRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcSubscriptionRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcSubscriptionRequestBuilder)
		}
	}

	return r0
}

// MockPlcConnection_SubscriptionRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscriptionRequestBuilder'
type MockPlcConnection_SubscriptionRequestBuilder_Call struct {
	*mock.Call
}

// SubscriptionRequestBuilder is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) SubscriptionRequestBuilder() *MockPlcConnection_SubscriptionRequestBuilder_Call {
	return &MockPlcConnection_SubscriptionRequestBuilder_Call{Call: _e.mock.On("SubscriptionRequestBuilder")}
}

func (_c *MockPlcConnection_SubscriptionRequestBuilder_Call) Run(run func()) *MockPlcConnection_SubscriptionRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_SubscriptionRequestBuilder_Call) Return(_a0 model.PlcSubscriptionRequestBuilder) *MockPlcConnection_SubscriptionRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_SubscriptionRequestBuilder_Call) RunAndReturn(run func() model.PlcSubscriptionRequestBuilder) *MockPlcConnection_SubscriptionRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// UnsubscriptionRequestBuilder provides a mock function with given fields:
func (_m *MockPlcConnection) UnsubscriptionRequestBuilder() model.PlcUnsubscriptionRequestBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UnsubscriptionRequestBuilder")
	}

	var r0 model.PlcUnsubscriptionRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcUnsubscriptionRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcUnsubscriptionRequestBuilder)
		}
	}

	return r0
}

// MockPlcConnection_UnsubscriptionRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnsubscriptionRequestBuilder'
type MockPlcConnection_UnsubscriptionRequestBuilder_Call struct {
	*mock.Call
}

// UnsubscriptionRequestBuilder is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) UnsubscriptionRequestBuilder() *MockPlcConnection_UnsubscriptionRequestBuilder_Call {
	return &MockPlcConnection_UnsubscriptionRequestBuilder_Call{Call: _e.mock.On("UnsubscriptionRequestBuilder")}
}

func (_c *MockPlcConnection_UnsubscriptionRequestBuilder_Call) Run(run func()) *MockPlcConnection_UnsubscriptionRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_UnsubscriptionRequestBuilder_Call) Return(_a0 model.PlcUnsubscriptionRequestBuilder) *MockPlcConnection_UnsubscriptionRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_UnsubscriptionRequestBuilder_Call) RunAndReturn(run func() model.PlcUnsubscriptionRequestBuilder) *MockPlcConnection_UnsubscriptionRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// WriteRequestBuilder provides a mock function with given fields:
func (_m *MockPlcConnection) WriteRequestBuilder() model.PlcWriteRequestBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for WriteRequestBuilder")
	}

	var r0 model.PlcWriteRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcWriteRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcWriteRequestBuilder)
		}
	}

	return r0
}

// MockPlcConnection_WriteRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteRequestBuilder'
type MockPlcConnection_WriteRequestBuilder_Call struct {
	*mock.Call
}

// WriteRequestBuilder is a helper method to define mock.On call
func (_e *MockPlcConnection_Expecter) WriteRequestBuilder() *MockPlcConnection_WriteRequestBuilder_Call {
	return &MockPlcConnection_WriteRequestBuilder_Call{Call: _e.mock.On("WriteRequestBuilder")}
}

func (_c *MockPlcConnection_WriteRequestBuilder_Call) Run(run func()) *MockPlcConnection_WriteRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnection_WriteRequestBuilder_Call) Return(_a0 model.PlcWriteRequestBuilder) *MockPlcConnection_WriteRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnection_WriteRequestBuilder_Call) RunAndReturn(run func() model.PlcWriteRequestBuilder) *MockPlcConnection_WriteRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcConnection creates a new instance of MockPlcConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcConnection {
	mock := &MockPlcConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
