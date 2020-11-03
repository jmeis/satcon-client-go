// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"net/http"
	"sync"

	"github.com/IBM/satcon-client-go/client/auth"
)

type FakeAuthClient struct {
	AuthenticateStub        func(*http.Request) error
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		arg1 *http.Request
	}
	authenticateReturns struct {
		result1 error
	}
	authenticateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthClient) Authenticate(arg1 *http.Request) error {
	fake.authenticateMutex.Lock()
	ret, specificReturn := fake.authenticateReturnsOnCall[len(fake.authenticateArgsForCall)]
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	stub := fake.AuthenticateStub
	fakeReturns := fake.authenticateReturns
	fake.recordInvocation("Authenticate", []interface{}{arg1})
	fake.authenticateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthClient) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeAuthClient) AuthenticateCalls(stub func(*http.Request) error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = stub
}

func (fake *FakeAuthClient) AuthenticateArgsForCall(i int) *http.Request {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	argsForCall := fake.authenticateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthClient) AuthenticateReturns(result1 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthClient) AuthenticateReturnsOnCall(i int, result1 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	if fake.authenticateReturnsOnCall == nil {
		fake.authenticateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.authenticateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthClient) recordInvocation(key string, args []interface{}) {
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

var _ auth.AuthClient = new(FakeAuthClient)