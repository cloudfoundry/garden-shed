// This file was generated by counterfeiter
package fake_retainer

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
)

type FakeRetainer struct {
	RetainStub        func(id layercake.ID)
	retainMutex       sync.RWMutex
	retainArgsForCall []struct {
		id layercake.ID
	}
}

func (fake *FakeRetainer) Retain(id layercake.ID) {
	fake.retainMutex.Lock()
	fake.retainArgsForCall = append(fake.retainArgsForCall, struct {
		id layercake.ID
	}{id})
	fake.retainMutex.Unlock()
	if fake.RetainStub != nil {
		fake.RetainStub(id)
	}
}

func (fake *FakeRetainer) RetainCallCount() int {
	fake.retainMutex.RLock()
	defer fake.retainMutex.RUnlock()
	return len(fake.retainArgsForCall)
}

func (fake *FakeRetainer) RetainArgsForCall(i int) layercake.ID {
	fake.retainMutex.RLock()
	defer fake.retainMutex.RUnlock()
	return fake.retainArgsForCall[i].id
}

var _ layercake.Retainer = new(FakeRetainer)
