// This file was generated by counterfeiter
package fake_distclient

import (
	"io"
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/distclient"
	"github.com/docker/distribution/digest"
	"github.com/pivotal-golang/lager"
)

type FakeConn struct {
	GetManifestStub        func(logger lager.Logger, tag string) (*distclient.Manifest, error)
	getManifestMutex       sync.RWMutex
	getManifestArgsForCall []struct {
		logger lager.Logger
		tag    string
	}
	getManifestReturns struct {
		result1 *distclient.Manifest
		result2 error
	}
	GetBlobReaderStub        func(logger lager.Logger, d digest.Digest) (io.Reader, error)
	getBlobReaderMutex       sync.RWMutex
	getBlobReaderArgsForCall []struct {
		logger lager.Logger
		d      digest.Digest
	}
	getBlobReaderReturns struct {
		result1 io.Reader
		result2 error
	}
}

func (fake *FakeConn) GetManifest(logger lager.Logger, tag string) (*distclient.Manifest, error) {
	fake.getManifestMutex.Lock()
	fake.getManifestArgsForCall = append(fake.getManifestArgsForCall, struct {
		logger lager.Logger
		tag    string
	}{logger, tag})
	fake.getManifestMutex.Unlock()
	if fake.GetManifestStub != nil {
		return fake.GetManifestStub(logger, tag)
	} else {
		return fake.getManifestReturns.result1, fake.getManifestReturns.result2
	}
}

func (fake *FakeConn) GetManifestCallCount() int {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return len(fake.getManifestArgsForCall)
}

func (fake *FakeConn) GetManifestArgsForCall(i int) (lager.Logger, string) {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return fake.getManifestArgsForCall[i].logger, fake.getManifestArgsForCall[i].tag
}

func (fake *FakeConn) GetManifestReturns(result1 *distclient.Manifest, result2 error) {
	fake.GetManifestStub = nil
	fake.getManifestReturns = struct {
		result1 *distclient.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) GetBlobReader(logger lager.Logger, d digest.Digest) (io.Reader, error) {
	fake.getBlobReaderMutex.Lock()
	fake.getBlobReaderArgsForCall = append(fake.getBlobReaderArgsForCall, struct {
		logger lager.Logger
		d      digest.Digest
	}{logger, d})
	fake.getBlobReaderMutex.Unlock()
	if fake.GetBlobReaderStub != nil {
		return fake.GetBlobReaderStub(logger, d)
	} else {
		return fake.getBlobReaderReturns.result1, fake.getBlobReaderReturns.result2
	}
}

func (fake *FakeConn) GetBlobReaderCallCount() int {
	fake.getBlobReaderMutex.RLock()
	defer fake.getBlobReaderMutex.RUnlock()
	return len(fake.getBlobReaderArgsForCall)
}

func (fake *FakeConn) GetBlobReaderArgsForCall(i int) (lager.Logger, digest.Digest) {
	fake.getBlobReaderMutex.RLock()
	defer fake.getBlobReaderMutex.RUnlock()
	return fake.getBlobReaderArgsForCall[i].logger, fake.getBlobReaderArgsForCall[i].d
}

func (fake *FakeConn) GetBlobReaderReturns(result1 io.Reader, result2 error) {
	fake.GetBlobReaderStub = nil
	fake.getBlobReaderReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

var _ distclient.Conn = new(FakeConn)
