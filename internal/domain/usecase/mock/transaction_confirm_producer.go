// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// ConfirmTransactionProducer is an autogenerated mock type for the ConfirmTransactionProducer type
type ConfirmTransactionProducer struct {
	mock.Mock
}

// CancelTransaction provides a mock function with given fields: id
func (_m *ConfirmTransactionProducer) CancelTransaction(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteTransaction provides a mock function with given fields: id
func (_m *ConfirmTransactionProducer) CompleteTransaction(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConfirmTransactionProducer interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfirmTransactionProducer creates a new instance of ConfirmTransactionProducer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfirmTransactionProducer(t mockConstructorTestingTNewConfirmTransactionProducer) *ConfirmTransactionProducer {
	mock := &ConfirmTransactionProducer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}