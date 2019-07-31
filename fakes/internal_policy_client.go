// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy_client"
)

type InternalPolicyClient struct {
	GetPoliciesStub        func() ([]*policy_client.Policy, []*policy_client.EgressPolicy, error)
	getPoliciesMutex       sync.RWMutex
	getPoliciesArgsForCall []struct {
	}
	getPoliciesReturns struct {
		result1 []*policy_client.Policy
		result2 []*policy_client.EgressPolicy
		result3 error
	}
	getPoliciesReturnsOnCall map[int]struct {
		result1 []*policy_client.Policy
		result2 []*policy_client.EgressPolicy
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *InternalPolicyClient) GetPolicies() ([]*policy_client.Policy, []*policy_client.EgressPolicy, error) {
	fake.getPoliciesMutex.Lock()
	ret, specificReturn := fake.getPoliciesReturnsOnCall[len(fake.getPoliciesArgsForCall)]
	fake.getPoliciesArgsForCall = append(fake.getPoliciesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetPolicies", []interface{}{})
	fake.getPoliciesMutex.Unlock()
	if fake.GetPoliciesStub != nil {
		return fake.GetPoliciesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getPoliciesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *InternalPolicyClient) GetPoliciesCallCount() int {
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	return len(fake.getPoliciesArgsForCall)
}

func (fake *InternalPolicyClient) GetPoliciesCalls(stub func() ([]*policy_client.Policy, []*policy_client.EgressPolicy, error)) {
	fake.getPoliciesMutex.Lock()
	defer fake.getPoliciesMutex.Unlock()
	fake.GetPoliciesStub = stub
}

func (fake *InternalPolicyClient) GetPoliciesReturns(result1 []*policy_client.Policy, result2 []*policy_client.EgressPolicy, result3 error) {
	fake.getPoliciesMutex.Lock()
	defer fake.getPoliciesMutex.Unlock()
	fake.GetPoliciesStub = nil
	fake.getPoliciesReturns = struct {
		result1 []*policy_client.Policy
		result2 []*policy_client.EgressPolicy
		result3 error
	}{result1, result2, result3}
}

func (fake *InternalPolicyClient) GetPoliciesReturnsOnCall(i int, result1 []*policy_client.Policy, result2 []*policy_client.EgressPolicy, result3 error) {
	fake.getPoliciesMutex.Lock()
	defer fake.getPoliciesMutex.Unlock()
	fake.GetPoliciesStub = nil
	if fake.getPoliciesReturnsOnCall == nil {
		fake.getPoliciesReturnsOnCall = make(map[int]struct {
			result1 []*policy_client.Policy
			result2 []*policy_client.EgressPolicy
			result3 error
		})
	}
	fake.getPoliciesReturnsOnCall[i] = struct {
		result1 []*policy_client.Policy
		result2 []*policy_client.EgressPolicy
		result3 error
	}{result1, result2, result3}
}

func (fake *InternalPolicyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *InternalPolicyClient) recordInvocation(key string, args []interface{}) {
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

var _ policy_client.InternalPolicyClient = new(InternalPolicyClient)
