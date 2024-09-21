// Code generated by MockGen. DO NOT EDIT.
// Source: client/account_retriever.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	client "github.com/cosmos/cosmos-sdk/client"
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccount is a mock of Account interface.
type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}

// MockAccountMockRecorder is the mock recorder for MockAccount.
type MockAccountMockRecorder struct {
	mock *MockAccount
}

// NewMockAccount creates a new mock instance.
func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}

// GetAccountNumber mocks base method.
func (m *MockAccount) GetAccountNumber() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountNumber")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetAccountNumber indicates an expected call of GetAccountNumber.
func (mr *MockAccountMockRecorder) GetAccountNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountNumber", reflect.TypeOf((*MockAccount)(nil).GetAccountNumber))
}

// GetAddress mocks base method.
func (m *MockAccount) GetAddress() types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress")
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockAccountMockRecorder) GetAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockAccount)(nil).GetAddress))
}

// GetPubKey mocks base method.
func (m *MockAccount) GetPubKey() types.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKey")
	ret0, _ := ret[0].(types.PubKey)
	return ret0
}

// GetPubKey indicates an expected call of GetPubKey.
func (mr *MockAccountMockRecorder) GetPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKey", reflect.TypeOf((*MockAccount)(nil).GetPubKey))
}

// GetSequence mocks base method.
func (m *MockAccount) GetSequence() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSequence")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetSequence indicates an expected call of GetSequence.
func (mr *MockAccountMockRecorder) GetSequence() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSequence", reflect.TypeOf((*MockAccount)(nil).GetSequence))
}

// MockAccountRetriever is a mock of AccountRetriever interface.
type MockAccountRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRetrieverMockRecorder
}

// MockAccountRetrieverMockRecorder is the mock recorder for MockAccountRetriever.
type MockAccountRetrieverMockRecorder struct {
	mock *MockAccountRetriever
}

// NewMockAccountRetriever creates a new mock instance.
func NewMockAccountRetriever(ctrl *gomock.Controller) *MockAccountRetriever {
	mock := &MockAccountRetriever{ctrl: ctrl}
	mock.recorder = &MockAccountRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRetriever) EXPECT() *MockAccountRetrieverMockRecorder {
	return m.recorder
}

// EnsureExists mocks base method.
func (m *MockAccountRetriever) EnsureExists(clientCtx client.Context, addr types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureExists", clientCtx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureExists indicates an expected call of EnsureExists.
func (mr *MockAccountRetrieverMockRecorder) EnsureExists(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureExists", reflect.TypeOf((*MockAccountRetriever)(nil).EnsureExists), clientCtx, addr)
}

// GetAccount mocks base method.
func (m *MockAccountRetriever) GetAccount(clientCtx client.Context, addr types0.AccAddress) (client.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", clientCtx, addr)
	ret0, _ := ret[0].(client.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountRetrieverMockRecorder) GetAccount(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountRetriever)(nil).GetAccount), clientCtx, addr)
}

// GetAccountNumberSequence mocks base method.
func (m *MockAccountRetriever) GetAccountNumberSequence(clientCtx client.Context, addr types0.AccAddress) (uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountNumberSequence", clientCtx, addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAccountNumberSequence indicates an expected call of GetAccountNumberSequence.
func (mr *MockAccountRetrieverMockRecorder) GetAccountNumberSequence(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountNumberSequence", reflect.TypeOf((*MockAccountRetriever)(nil).GetAccountNumberSequence), clientCtx, addr)
}

// GetAccountWithHeight mocks base method.
func (m *MockAccountRetriever) GetAccountWithHeight(clientCtx client.Context, addr types0.AccAddress) (client.Account, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountWithHeight", clientCtx, addr)
	ret0, _ := ret[0].(client.Account)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAccountWithHeight indicates an expected call of GetAccountWithHeight.
func (mr *MockAccountRetrieverMockRecorder) GetAccountWithHeight(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountWithHeight", reflect.TypeOf((*MockAccountRetriever)(nil).GetAccountWithHeight), clientCtx, addr)
}
