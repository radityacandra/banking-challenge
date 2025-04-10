// Code generated by mockery v2.52.4. DO NOT EDIT.

package useraccount

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockServerInterface is an autogenerated mock type for the ServerInterface type
type MockServerInterface struct {
	mock.Mock
}

type MockServerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServerInterface) EXPECT() *MockServerInterface_Expecter {
	return &MockServerInterface_Expecter{mock: &_m.Mock}
}

// AccountBalanceGet provides a mock function with given fields: ctx, noRekening
func (_m *MockServerInterface) AccountBalanceGet(ctx echo.Context, noRekening string) error {
	ret := _m.Called(ctx, noRekening)

	if len(ret) == 0 {
		panic("no return value specified for AccountBalanceGet")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, noRekening)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_AccountBalanceGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccountBalanceGet'
type MockServerInterface_AccountBalanceGet_Call struct {
	*mock.Call
}

// AccountBalanceGet is a helper method to define mock.On call
//   - ctx echo.Context
//   - noRekening string
func (_e *MockServerInterface_Expecter) AccountBalanceGet(ctx interface{}, noRekening interface{}) *MockServerInterface_AccountBalanceGet_Call {
	return &MockServerInterface_AccountBalanceGet_Call{Call: _e.mock.On("AccountBalanceGet", ctx, noRekening)}
}

func (_c *MockServerInterface_AccountBalanceGet_Call) Run(run func(ctx echo.Context, noRekening string)) *MockServerInterface_AccountBalanceGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockServerInterface_AccountBalanceGet_Call) Return(_a0 error) *MockServerInterface_AccountBalanceGet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_AccountBalanceGet_Call) RunAndReturn(run func(echo.Context, string) error) *MockServerInterface_AccountBalanceGet_Call {
	_c.Call.Return(run)
	return _c
}

// TransactionCreditPost provides a mock function with given fields: ctx
func (_m *MockServerInterface) TransactionCreditPost(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for TransactionCreditPost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_TransactionCreditPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactionCreditPost'
type MockServerInterface_TransactionCreditPost_Call struct {
	*mock.Call
}

// TransactionCreditPost is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockServerInterface_Expecter) TransactionCreditPost(ctx interface{}) *MockServerInterface_TransactionCreditPost_Call {
	return &MockServerInterface_TransactionCreditPost_Call{Call: _e.mock.On("TransactionCreditPost", ctx)}
}

func (_c *MockServerInterface_TransactionCreditPost_Call) Run(run func(ctx echo.Context)) *MockServerInterface_TransactionCreditPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockServerInterface_TransactionCreditPost_Call) Return(_a0 error) *MockServerInterface_TransactionCreditPost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_TransactionCreditPost_Call) RunAndReturn(run func(echo.Context) error) *MockServerInterface_TransactionCreditPost_Call {
	_c.Call.Return(run)
	return _c
}

// TransactionDebitPost provides a mock function with given fields: ctx
func (_m *MockServerInterface) TransactionDebitPost(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for TransactionDebitPost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_TransactionDebitPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactionDebitPost'
type MockServerInterface_TransactionDebitPost_Call struct {
	*mock.Call
}

// TransactionDebitPost is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockServerInterface_Expecter) TransactionDebitPost(ctx interface{}) *MockServerInterface_TransactionDebitPost_Call {
	return &MockServerInterface_TransactionDebitPost_Call{Call: _e.mock.On("TransactionDebitPost", ctx)}
}

func (_c *MockServerInterface_TransactionDebitPost_Call) Run(run func(ctx echo.Context)) *MockServerInterface_TransactionDebitPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockServerInterface_TransactionDebitPost_Call) Return(_a0 error) *MockServerInterface_TransactionDebitPost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_TransactionDebitPost_Call) RunAndReturn(run func(echo.Context) error) *MockServerInterface_TransactionDebitPost_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServerInterface creates a new instance of MockServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServerInterface {
	mock := &MockServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
