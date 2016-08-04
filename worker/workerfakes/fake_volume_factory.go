// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/worker"
	"github.com/concourse/baggageclaim"
)

type FakeVolumeFactory struct {
	BuildStub        func(lager.Logger, baggageclaim.Volume) (worker.Volume, bool, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 lager.Logger
		arg2 baggageclaim.Volume
	}
	buildReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeFactory) Build(arg1 lager.Logger, arg2 baggageclaim.Volume) (worker.Volume, bool, error) {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 lager.Logger
		arg2 baggageclaim.Volume
	}{arg1, arg2})
	fake.recordInvocation("Build", []interface{}{arg1, arg2})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(arg1, arg2)
	} else {
		return fake.buildReturns.result1, fake.buildReturns.result2, fake.buildReturns.result3
	}
}

func (fake *FakeVolumeFactory) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeVolumeFactory) BuildArgsForCall(i int) (lager.Logger, baggageclaim.Volume) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].arg1, fake.buildArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) BuildReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumeFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.VolumeFactory = new(FakeVolumeFactory)
