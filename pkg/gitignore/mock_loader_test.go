// Code generated by mockery v1.0.0. DO NOT EDIT.

package gitignore

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockLoader is an autogenerated mock type for the Loader type
type MockLoader struct {
	mock.Mock
}

// Load provides a mock function with given fields: _a0, _a1
func (_m *MockLoader) Load(_a0 context.Context, _a1 gitignoreType) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, gitignoreType) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gitignoreType) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
