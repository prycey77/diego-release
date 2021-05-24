// Code generated by counterfeiter. DO NOT EDIT.
package monitorfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/diego-release/bbs/db/sqldb/helpers/monitor"
)

type FakeMonitor struct {
	FailedStub        func() int64
	failedMutex       sync.RWMutex
	failedArgsForCall []struct {
	}
	failedReturns struct {
		result1 int64
	}
	failedReturnsOnCall map[int]struct {
		result1 int64
	}
	MonitorStub        func(func() error) error
	monitorMutex       sync.RWMutex
	monitorArgsForCall []struct {
		arg1 func() error
	}
	monitorReturns struct {
		result1 error
	}
	monitorReturnsOnCall map[int]struct {
		result1 error
	}
	ReadAndResetDurationMaxStub        func() time.Duration
	readAndResetDurationMaxMutex       sync.RWMutex
	readAndResetDurationMaxArgsForCall []struct {
	}
	readAndResetDurationMaxReturns struct {
		result1 time.Duration
	}
	readAndResetDurationMaxReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	ReadAndResetInFlightMaxStub        func() int64
	readAndResetInFlightMaxMutex       sync.RWMutex
	readAndResetInFlightMaxArgsForCall []struct {
	}
	readAndResetInFlightMaxReturns struct {
		result1 int64
	}
	readAndResetInFlightMaxReturnsOnCall map[int]struct {
		result1 int64
	}
	SucceededStub        func() int64
	succeededMutex       sync.RWMutex
	succeededArgsForCall []struct {
	}
	succeededReturns struct {
		result1 int64
	}
	succeededReturnsOnCall map[int]struct {
		result1 int64
	}
	TotalStub        func() int64
	totalMutex       sync.RWMutex
	totalArgsForCall []struct {
	}
	totalReturns struct {
		result1 int64
	}
	totalReturnsOnCall map[int]struct {
		result1 int64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMonitor) Failed() int64 {
	fake.failedMutex.Lock()
	ret, specificReturn := fake.failedReturnsOnCall[len(fake.failedArgsForCall)]
	fake.failedArgsForCall = append(fake.failedArgsForCall, struct {
	}{})
	stub := fake.FailedStub
	fakeReturns := fake.failedReturns
	fake.recordInvocation("Failed", []interface{}{})
	fake.failedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) FailedCallCount() int {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return len(fake.failedArgsForCall)
}

func (fake *FakeMonitor) FailedCalls(stub func() int64) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = stub
}

func (fake *FakeMonitor) FailedReturns(result1 int64) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = nil
	fake.failedReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) FailedReturnsOnCall(i int, result1 int64) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = nil
	if fake.failedReturnsOnCall == nil {
		fake.failedReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.failedReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) Monitor(arg1 func() error) error {
	fake.monitorMutex.Lock()
	ret, specificReturn := fake.monitorReturnsOnCall[len(fake.monitorArgsForCall)]
	fake.monitorArgsForCall = append(fake.monitorArgsForCall, struct {
		arg1 func() error
	}{arg1})
	stub := fake.MonitorStub
	fakeReturns := fake.monitorReturns
	fake.recordInvocation("Monitor", []interface{}{arg1})
	fake.monitorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) MonitorCallCount() int {
	fake.monitorMutex.RLock()
	defer fake.monitorMutex.RUnlock()
	return len(fake.monitorArgsForCall)
}

func (fake *FakeMonitor) MonitorCalls(stub func(func() error) error) {
	fake.monitorMutex.Lock()
	defer fake.monitorMutex.Unlock()
	fake.MonitorStub = stub
}

func (fake *FakeMonitor) MonitorArgsForCall(i int) func() error {
	fake.monitorMutex.RLock()
	defer fake.monitorMutex.RUnlock()
	argsForCall := fake.monitorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMonitor) MonitorReturns(result1 error) {
	fake.monitorMutex.Lock()
	defer fake.monitorMutex.Unlock()
	fake.MonitorStub = nil
	fake.monitorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMonitor) MonitorReturnsOnCall(i int, result1 error) {
	fake.monitorMutex.Lock()
	defer fake.monitorMutex.Unlock()
	fake.MonitorStub = nil
	if fake.monitorReturnsOnCall == nil {
		fake.monitorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.monitorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMonitor) ReadAndResetDurationMax() time.Duration {
	fake.readAndResetDurationMaxMutex.Lock()
	ret, specificReturn := fake.readAndResetDurationMaxReturnsOnCall[len(fake.readAndResetDurationMaxArgsForCall)]
	fake.readAndResetDurationMaxArgsForCall = append(fake.readAndResetDurationMaxArgsForCall, struct {
	}{})
	stub := fake.ReadAndResetDurationMaxStub
	fakeReturns := fake.readAndResetDurationMaxReturns
	fake.recordInvocation("ReadAndResetDurationMax", []interface{}{})
	fake.readAndResetDurationMaxMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) ReadAndResetDurationMaxCallCount() int {
	fake.readAndResetDurationMaxMutex.RLock()
	defer fake.readAndResetDurationMaxMutex.RUnlock()
	return len(fake.readAndResetDurationMaxArgsForCall)
}

func (fake *FakeMonitor) ReadAndResetDurationMaxCalls(stub func() time.Duration) {
	fake.readAndResetDurationMaxMutex.Lock()
	defer fake.readAndResetDurationMaxMutex.Unlock()
	fake.ReadAndResetDurationMaxStub = stub
}

func (fake *FakeMonitor) ReadAndResetDurationMaxReturns(result1 time.Duration) {
	fake.readAndResetDurationMaxMutex.Lock()
	defer fake.readAndResetDurationMaxMutex.Unlock()
	fake.ReadAndResetDurationMaxStub = nil
	fake.readAndResetDurationMaxReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeMonitor) ReadAndResetDurationMaxReturnsOnCall(i int, result1 time.Duration) {
	fake.readAndResetDurationMaxMutex.Lock()
	defer fake.readAndResetDurationMaxMutex.Unlock()
	fake.ReadAndResetDurationMaxStub = nil
	if fake.readAndResetDurationMaxReturnsOnCall == nil {
		fake.readAndResetDurationMaxReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.readAndResetDurationMaxReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeMonitor) ReadAndResetInFlightMax() int64 {
	fake.readAndResetInFlightMaxMutex.Lock()
	ret, specificReturn := fake.readAndResetInFlightMaxReturnsOnCall[len(fake.readAndResetInFlightMaxArgsForCall)]
	fake.readAndResetInFlightMaxArgsForCall = append(fake.readAndResetInFlightMaxArgsForCall, struct {
	}{})
	stub := fake.ReadAndResetInFlightMaxStub
	fakeReturns := fake.readAndResetInFlightMaxReturns
	fake.recordInvocation("ReadAndResetInFlightMax", []interface{}{})
	fake.readAndResetInFlightMaxMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) ReadAndResetInFlightMaxCallCount() int {
	fake.readAndResetInFlightMaxMutex.RLock()
	defer fake.readAndResetInFlightMaxMutex.RUnlock()
	return len(fake.readAndResetInFlightMaxArgsForCall)
}

func (fake *FakeMonitor) ReadAndResetInFlightMaxCalls(stub func() int64) {
	fake.readAndResetInFlightMaxMutex.Lock()
	defer fake.readAndResetInFlightMaxMutex.Unlock()
	fake.ReadAndResetInFlightMaxStub = stub
}

func (fake *FakeMonitor) ReadAndResetInFlightMaxReturns(result1 int64) {
	fake.readAndResetInFlightMaxMutex.Lock()
	defer fake.readAndResetInFlightMaxMutex.Unlock()
	fake.ReadAndResetInFlightMaxStub = nil
	fake.readAndResetInFlightMaxReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) ReadAndResetInFlightMaxReturnsOnCall(i int, result1 int64) {
	fake.readAndResetInFlightMaxMutex.Lock()
	defer fake.readAndResetInFlightMaxMutex.Unlock()
	fake.ReadAndResetInFlightMaxStub = nil
	if fake.readAndResetInFlightMaxReturnsOnCall == nil {
		fake.readAndResetInFlightMaxReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.readAndResetInFlightMaxReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) Succeeded() int64 {
	fake.succeededMutex.Lock()
	ret, specificReturn := fake.succeededReturnsOnCall[len(fake.succeededArgsForCall)]
	fake.succeededArgsForCall = append(fake.succeededArgsForCall, struct {
	}{})
	stub := fake.SucceededStub
	fakeReturns := fake.succeededReturns
	fake.recordInvocation("Succeeded", []interface{}{})
	fake.succeededMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) SucceededCallCount() int {
	fake.succeededMutex.RLock()
	defer fake.succeededMutex.RUnlock()
	return len(fake.succeededArgsForCall)
}

func (fake *FakeMonitor) SucceededCalls(stub func() int64) {
	fake.succeededMutex.Lock()
	defer fake.succeededMutex.Unlock()
	fake.SucceededStub = stub
}

func (fake *FakeMonitor) SucceededReturns(result1 int64) {
	fake.succeededMutex.Lock()
	defer fake.succeededMutex.Unlock()
	fake.SucceededStub = nil
	fake.succeededReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) SucceededReturnsOnCall(i int, result1 int64) {
	fake.succeededMutex.Lock()
	defer fake.succeededMutex.Unlock()
	fake.SucceededStub = nil
	if fake.succeededReturnsOnCall == nil {
		fake.succeededReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.succeededReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) Total() int64 {
	fake.totalMutex.Lock()
	ret, specificReturn := fake.totalReturnsOnCall[len(fake.totalArgsForCall)]
	fake.totalArgsForCall = append(fake.totalArgsForCall, struct {
	}{})
	stub := fake.TotalStub
	fakeReturns := fake.totalReturns
	fake.recordInvocation("Total", []interface{}{})
	fake.totalMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) TotalCallCount() int {
	fake.totalMutex.RLock()
	defer fake.totalMutex.RUnlock()
	return len(fake.totalArgsForCall)
}

func (fake *FakeMonitor) TotalCalls(stub func() int64) {
	fake.totalMutex.Lock()
	defer fake.totalMutex.Unlock()
	fake.TotalStub = stub
}

func (fake *FakeMonitor) TotalReturns(result1 int64) {
	fake.totalMutex.Lock()
	defer fake.totalMutex.Unlock()
	fake.TotalStub = nil
	fake.totalReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) TotalReturnsOnCall(i int, result1 int64) {
	fake.totalMutex.Lock()
	defer fake.totalMutex.Unlock()
	fake.TotalStub = nil
	if fake.totalReturnsOnCall == nil {
		fake.totalReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.totalReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMonitor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	fake.monitorMutex.RLock()
	defer fake.monitorMutex.RUnlock()
	fake.readAndResetDurationMaxMutex.RLock()
	defer fake.readAndResetDurationMaxMutex.RUnlock()
	fake.readAndResetInFlightMaxMutex.RLock()
	defer fake.readAndResetInFlightMaxMutex.RUnlock()
	fake.succeededMutex.RLock()
	defer fake.succeededMutex.RUnlock()
	fake.totalMutex.RLock()
	defer fake.totalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMonitor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ monitor.Monitor = new(FakeMonitor)
