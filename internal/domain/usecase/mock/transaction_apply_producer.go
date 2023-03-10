// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	context "context"
	entity "payment_processing_system/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// ApplyTransactionProducer is an autogenerated mock type for the ApplyTransactionProducer type
type ApplyTransactionProducer struct {
	mock.Mock
}

// ApplyTransaction provides a mock function with given fields: ctx, transaction
func (_m *ApplyTransactionProducer) ApplyTransaction(ctx context.Context, transaction entity.Transaction) error {
	ret := _m.Called(ctx, transaction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transaction) error); ok {
		r0 = rf(ctx, transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewApplyTransactionProducer interface {
	mock.TestingT
	Cleanup(func())
}

// NewApplyTransactionProducer creates a new instance of ApplyTransactionProducer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApplyTransactionProducer(t mockConstructorTestingTNewApplyTransactionProducer) *ApplyTransactionProducer {
	mock := &ApplyTransactionProducer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
