// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/DuGlaser/go-todo-server/domain"
	mock "github.com/stretchr/testify/mock"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *TodoRepository) Delete(_a0 int64) *domain.RestErr {
	ret := _m.Called(_a0)

	var r0 *domain.RestErr
	if rf, ok := ret.Get(0).(func(int64) *domain.RestErr); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RestErr)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *TodoRepository) GetAll() (domain.Todos, *domain.RestErr) {
	ret := _m.Called()

	var r0 domain.Todos
	if rf, ok := ret.Get(0).(func() domain.Todos); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Todos)
		}
	}

	var r1 *domain.RestErr
	if rf, ok := ret.Get(1).(func() *domain.RestErr); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.RestErr)
		}
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *TodoRepository) GetByID(_a0 int64) (domain.Todo, *domain.RestErr) {
	ret := _m.Called(_a0)

	var r0 domain.Todo
	if rf, ok := ret.Get(0).(func(int64) domain.Todo); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Todo)
	}

	var r1 *domain.RestErr
	if rf, ok := ret.Get(1).(func(int64) *domain.RestErr); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.RestErr)
		}
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *TodoRepository) Store(_a0 *domain.Todo) *domain.RestErr {
	ret := _m.Called(_a0)

	var r0 *domain.RestErr
	if rf, ok := ret.Get(0).(func(*domain.Todo) *domain.RestErr); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RestErr)
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *TodoRepository) Update(_a0 *domain.Todo) *domain.RestErr {
	ret := _m.Called(_a0)

	var r0 *domain.RestErr
	if rf, ok := ret.Get(0).(func(*domain.Todo) *domain.RestErr); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RestErr)
		}
	}

	return r0
}
