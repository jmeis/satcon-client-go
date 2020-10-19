// Code generated by counterfeiter. DO NOT EDIT.
package subscriptionsfakes

import (
	"sync"

	"github.com/IBM/satcon-client-go/client/actions/subscriptions"
)

type FakeSubscriptionService struct {
	SubscriptionsStub        func(string, string) (subscriptions.SubscriptionList, error)
	subscriptionsMutex       sync.RWMutex
	subscriptionsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	subscriptionsReturns struct {
		result1 subscriptions.SubscriptionList
		result2 error
	}
	subscriptionsReturnsOnCall map[int]struct {
		result1 subscriptions.SubscriptionList
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSubscriptionService) Subscriptions(arg1 string, arg2 string) (subscriptions.SubscriptionList, error) {
	fake.subscriptionsMutex.Lock()
	ret, specificReturn := fake.subscriptionsReturnsOnCall[len(fake.subscriptionsArgsForCall)]
	fake.subscriptionsArgsForCall = append(fake.subscriptionsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Subscriptions", []interface{}{arg1, arg2})
	fake.subscriptionsMutex.Unlock()
	if fake.SubscriptionsStub != nil {
		return fake.SubscriptionsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.subscriptionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSubscriptionService) SubscriptionsCallCount() int {
	fake.subscriptionsMutex.RLock()
	defer fake.subscriptionsMutex.RUnlock()
	return len(fake.subscriptionsArgsForCall)
}

func (fake *FakeSubscriptionService) SubscriptionsCalls(stub func(string, string) (subscriptions.SubscriptionList, error)) {
	fake.subscriptionsMutex.Lock()
	defer fake.subscriptionsMutex.Unlock()
	fake.SubscriptionsStub = stub
}

func (fake *FakeSubscriptionService) SubscriptionsArgsForCall(i int) (string, string) {
	fake.subscriptionsMutex.RLock()
	defer fake.subscriptionsMutex.RUnlock()
	argsForCall := fake.subscriptionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSubscriptionService) SubscriptionsReturns(result1 subscriptions.SubscriptionList, result2 error) {
	fake.subscriptionsMutex.Lock()
	defer fake.subscriptionsMutex.Unlock()
	fake.SubscriptionsStub = nil
	fake.subscriptionsReturns = struct {
		result1 subscriptions.SubscriptionList
		result2 error
	}{result1, result2}
}

func (fake *FakeSubscriptionService) SubscriptionsReturnsOnCall(i int, result1 subscriptions.SubscriptionList, result2 error) {
	fake.subscriptionsMutex.Lock()
	defer fake.subscriptionsMutex.Unlock()
	fake.SubscriptionsStub = nil
	if fake.subscriptionsReturnsOnCall == nil {
		fake.subscriptionsReturnsOnCall = make(map[int]struct {
			result1 subscriptions.SubscriptionList
			result2 error
		})
	}
	fake.subscriptionsReturnsOnCall[i] = struct {
		result1 subscriptions.SubscriptionList
		result2 error
	}{result1, result2}
}

func (fake *FakeSubscriptionService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.subscriptionsMutex.RLock()
	defer fake.subscriptionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSubscriptionService) recordInvocation(key string, args []interface{}) {
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

var _ subscriptions.SubscriptionService = new(FakeSubscriptionService)
