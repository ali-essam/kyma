// Code generated by mockery v1.0.0. DO NOT EDIT.
package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// gqlServiceClassConverter is an autogenerated mock type for the gqlServiceClassConverter type
type gqlServiceClassConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlServiceClassConverter) ToGQL(in *v1beta1.ServiceClass) (*gqlschema.ServiceClass, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ServiceClass
	if rf, ok := ret.Get(0).(func(*v1beta1.ServiceClass) *gqlschema.ServiceClass); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.ServiceClass) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlServiceClassConverter) ToGQLs(in []*v1beta1.ServiceClass) ([]gqlschema.ServiceClass, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ServiceClass
	if rf, ok := ret.Get(0).(func([]*v1beta1.ServiceClass) []gqlschema.ServiceClass); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1beta1.ServiceClass) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
