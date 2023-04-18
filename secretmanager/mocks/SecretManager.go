// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SecretManager is an autogenerated mock type for the SecretManager type
type SecretManager struct {
	mock.Mock
}

type SecretManager_Expecter struct {
	mock *mock.Mock
}

func (_m *SecretManager) EXPECT() *SecretManager_Expecter {
	return &SecretManager_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewSecretManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewSecretManager creates a new instance of SecretManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSecretManager(t mockConstructorTestingTNewSecretManager) *SecretManager {
	mock := &SecretManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
