// Code generated by counterfeiter. DO NOT EDIT.
package peasfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/peas"
	"code.cloudfoundry.org/lager"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeVolumizer struct {
	CreateStub        func(log lager.Logger, spec garden.ContainerSpec) (specs.Spec, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		log  lager.Logger
		spec garden.ContainerSpec
	}
	createReturns struct {
		result1 specs.Spec
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 specs.Spec
		result2 error
	}
	DestroyStub        func(log lager.Logger, handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumizer) Create(log lager.Logger, spec garden.ContainerSpec) (specs.Spec, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		log  lager.Logger
		spec garden.ContainerSpec
	}{log, spec})
	fake.recordInvocation("Create", []interface{}{log, spec})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(log, spec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeVolumizer) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVolumizer) CreateArgsForCall(i int) (lager.Logger, garden.ContainerSpec) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].log, fake.createArgsForCall[i].spec
}

func (fake *FakeVolumizer) CreateReturns(result1 specs.Spec, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumizer) CreateReturnsOnCall(i int, result1 specs.Spec, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 specs.Spec
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumizer) Destroy(log lager.Logger, handle string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Destroy", []interface{}{log, handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyReturns.result1
}

func (fake *FakeVolumizer) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolumizer) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].log, fake.destroyArgsForCall[i].handle
}

func (fake *FakeVolumizer) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumizer) DestroyReturnsOnCall(i int, result1 error) {
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumizer) recordInvocation(key string, args []interface{}) {
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

var _ peas.Volumizer = new(FakeVolumizer)