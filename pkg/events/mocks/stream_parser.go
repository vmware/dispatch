// Code generated by mockery v1.0.0

package mocks

import events "github.com/vmware/dispatch/pkg/events"
import io "io"
import mock "github.com/stretchr/testify/mock"

// StreamParser is an autogenerated mock type for the StreamParser type
type StreamParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: _a0
func (_m *StreamParser) Parse(_a0 io.Reader) ([]events.CloudEvent, error) {
	ret := _m.Called(_a0)

	var r0 []events.CloudEvent
	if rf, ok := ret.Get(0).(func(io.Reader) []events.CloudEvent); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.CloudEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
