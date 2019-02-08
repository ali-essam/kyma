// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import apperrors "github.com/kyma-project/kyma/components/connector-service/internal/apperrors"
import mock "github.com/stretchr/testify/mock"

// Resolver is an autogenerated mock type for the Resolver type
type Resolver struct {
	mock.Mock
}

// Resolve provides a mock function with given fields: token, destination
func (_m *Resolver) Resolve(token string, destination interface{}) apperrors.AppError {
	ret := _m.Called(token, destination)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, interface{}) apperrors.AppError); ok {
		r0 = rf(token, destination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}