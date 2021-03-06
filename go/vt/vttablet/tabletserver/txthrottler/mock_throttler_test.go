// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/youtube/vitess/go/vt/tabletserver/txthrottler (interfaces: ThrottlerInterface)

package txthrottler

import (
	time "time"

	gomock "github.com/golang/mock/gomock"
	discovery "github.com/youtube/vitess/go/vt/discovery"
	throttlerdata "github.com/youtube/vitess/go/vt/proto/throttlerdata"
)

// Mock of ThrottlerInterface interface
type MockThrottlerInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockThrottlerInterfaceRecorder
}

// Recorder for MockThrottlerInterface (not exported)
type _MockThrottlerInterfaceRecorder struct {
	mock *MockThrottlerInterface
}

func NewMockThrottlerInterface(ctrl *gomock.Controller) *MockThrottlerInterface {
	mock := &MockThrottlerInterface{ctrl: ctrl}
	mock.recorder = &_MockThrottlerInterfaceRecorder{mock}
	return mock
}

func (_m *MockThrottlerInterface) EXPECT() *_MockThrottlerInterfaceRecorder {
	return _m.recorder
}

func (_m *MockThrottlerInterface) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockThrottlerInterfaceRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockThrottlerInterface) GetConfiguration() *throttlerdata.Configuration {
	ret := _m.ctrl.Call(_m, "GetConfiguration")
	ret0, _ := ret[0].(*throttlerdata.Configuration)
	return ret0
}

func (_mr *_MockThrottlerInterfaceRecorder) GetConfiguration() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConfiguration")
}

func (_m *MockThrottlerInterface) MaxRate() int64 {
	ret := _m.ctrl.Call(_m, "MaxRate")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockThrottlerInterfaceRecorder) MaxRate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxRate")
}

func (_m *MockThrottlerInterface) RecordReplicationLag(_param0 time.Time, _param1 *discovery.TabletStats) {
	_m.ctrl.Call(_m, "RecordReplicationLag", _param0, _param1)
}

func (_mr *_MockThrottlerInterfaceRecorder) RecordReplicationLag(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecordReplicationLag", arg0, arg1)
}

func (_m *MockThrottlerInterface) ResetConfiguration() {
	_m.ctrl.Call(_m, "ResetConfiguration")
}

func (_mr *_MockThrottlerInterfaceRecorder) ResetConfiguration() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResetConfiguration")
}

func (_m *MockThrottlerInterface) SetMaxRate(_param0 int64) {
	_m.ctrl.Call(_m, "SetMaxRate", _param0)
}

func (_mr *_MockThrottlerInterfaceRecorder) SetMaxRate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMaxRate", arg0)
}

func (_m *MockThrottlerInterface) ThreadFinished(_param0 int) {
	_m.ctrl.Call(_m, "ThreadFinished", _param0)
}

func (_mr *_MockThrottlerInterfaceRecorder) ThreadFinished(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ThreadFinished", arg0)
}

func (_m *MockThrottlerInterface) Throttle(_param0 int) time.Duration {
	ret := _m.ctrl.Call(_m, "Throttle", _param0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockThrottlerInterfaceRecorder) Throttle(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Throttle", arg0)
}

func (_m *MockThrottlerInterface) UpdateConfiguration(_param0 *throttlerdata.Configuration, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "UpdateConfiguration", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockThrottlerInterfaceRecorder) UpdateConfiguration(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateConfiguration", arg0, arg1)
}
