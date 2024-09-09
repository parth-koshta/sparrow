// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/parth-koshta/sparrow/db/sqlc"
	mock "github.com/stretchr/testify/mock"

	pgtype "github.com/jackc/pgx/v5/pgtype"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Querier) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetUserByID(ctx context.Context, id pgtype.UUID) (db.GetUserByIDRow, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 db.GetUserByIDRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.GetUserByIDRow, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.GetUserByIDRow); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.GetUserByIDRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, arg
func (_m *Querier) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListUsersParams) ([]db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListUsersParams) []db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListUsersParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
