// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// UserHdl is an autogenerated mock type for the UserHdl type
type UserHdl struct {
	mock.Mock
}

// CreateGorm provides a mock function with given fields: ctx
func (_m *UserHdl) CreateGorm(ctx *gin.Context) {
	_m.Called(ctx)
}

// GetGorm provides a mock function with given fields: ctx
func (_m *UserHdl) GetGorm(ctx *gin.Context) {
	_m.Called(ctx)
}

// Login provides a mock function with given fields: ctx
func (_m *UserHdl) Login(ctx *gin.Context) {
	_m.Called(ctx)
}

// NewUserHdl creates a new instance of UserHdl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserHdl(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserHdl {
	mock := &UserHdl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
