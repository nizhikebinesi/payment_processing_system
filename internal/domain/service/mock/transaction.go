// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	context "context"
	entity "payment_processing_system/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionStorage is an autogenerated mock type for the TransactionStorage type
type TransactionStorage struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, transaction
func (_m *TransactionStorage) Create(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, error) {
	ret := _m.Called(ctx, transaction)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transaction) *entity.Transaction); ok {
		r0 = rf(ctx, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Transaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TransactionStorage) GetByID(ctx context.Context, id string) (*entity.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusByID provides a mock function with given fields: ctx, id, status
func (_m *TransactionStorage) UpdateStatusByID(ctx context.Context, id string, status entity.TransactionStatus) error {
	ret := _m.Called(ctx, id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.TransactionStatus) error); ok {
		r0 = rf(ctx, id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransactionStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionStorage creates a new instance of TransactionStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionStorage(t mockConstructorTestingTNewTransactionStorage) *TransactionStorage {
	mock := &TransactionStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
