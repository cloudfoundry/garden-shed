// This file was generated by counterfeiter
package fakes

import (
	"net/url"
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
	"github.com/cloudfoundry-incubator/garden-shed/repository_fetcher"
)

type FakeRemoteImageIDFetcher struct {
	FetchIDStub        func(u *url.URL) (layercake.ID, error)
	fetchIDMutex       sync.RWMutex
	fetchIDArgsForCall []struct {
		u *url.URL
	}
	fetchIDReturns struct {
		result1 layercake.ID
		result2 error
	}
}

func (fake *FakeRemoteImageIDFetcher) FetchID(u *url.URL) (layercake.ID, error) {
	fake.fetchIDMutex.Lock()
	fake.fetchIDArgsForCall = append(fake.fetchIDArgsForCall, struct {
		u *url.URL
	}{u})
	fake.fetchIDMutex.Unlock()
	if fake.FetchIDStub != nil {
		return fake.FetchIDStub(u)
	} else {
		return fake.fetchIDReturns.result1, fake.fetchIDReturns.result2
	}
}

func (fake *FakeRemoteImageIDFetcher) FetchIDCallCount() int {
	fake.fetchIDMutex.RLock()
	defer fake.fetchIDMutex.RUnlock()
	return len(fake.fetchIDArgsForCall)
}

func (fake *FakeRemoteImageIDFetcher) FetchIDArgsForCall(i int) *url.URL {
	fake.fetchIDMutex.RLock()
	defer fake.fetchIDMutex.RUnlock()
	return fake.fetchIDArgsForCall[i].u
}

func (fake *FakeRemoteImageIDFetcher) FetchIDReturns(result1 layercake.ID, result2 error) {
	fake.FetchIDStub = nil
	fake.fetchIDReturns = struct {
		result1 layercake.ID
		result2 error
	}{result1, result2}
}

var _ repository_fetcher.RemoteImageIDFetcher = new(FakeRemoteImageIDFetcher)
