// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	domain "github.com/Alina9496/documents/internal/domain"
	dto "github.com/Alina9496/documents/internal/service/dto"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddGrant mocks base method.
func (m *MockRepository) AddGrant(ctx context.Context, grant *domain.Grant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGrant", ctx, grant)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddGrant indicates an expected call of AddGrant.
func (mr *MockRepositoryMockRecorder) AddGrant(ctx, grant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGrant", reflect.TypeOf((*MockRepository)(nil).AddGrant), ctx, grant)
}

// Authentication mocks base method.
func (m *MockRepository) Authentication(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authentication indicates an expected call of Authentication.
func (mr *MockRepositoryMockRecorder) Authentication(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockRepository)(nil).Authentication), ctx, user)
}

// CheckGrant mocks base method.
func (m *MockRepository) CheckGrant(ctx context.Context, documentID uuid.UUID, login string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckGrant", ctx, documentID, login)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckGrant indicates an expected call of CheckGrant.
func (mr *MockRepositoryMockRecorder) CheckGrant(ctx, documentID, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckGrant", reflect.TypeOf((*MockRepository)(nil).CheckGrant), ctx, documentID, login)
}

// CheckUser mocks base method.
func (m *MockRepository) CheckUser(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", ctx, user)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockRepositoryMockRecorder) CheckUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockRepository)(nil).CheckUser), ctx, user)
}

// DeleteDocument mocks base method.
func (m *MockRepository) DeleteDocument(ctx context.Context, id, userID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDocument", ctx, id, userID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDocument indicates an expected call of DeleteDocument.
func (mr *MockRepositoryMockRecorder) DeleteDocument(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDocument", reflect.TypeOf((*MockRepository)(nil).DeleteDocument), ctx, id, userID)
}

// ExecTx mocks base method.
func (m *MockRepository) ExecTx(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockRepositoryMockRecorder) ExecTx(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockRepository)(nil).ExecTx), ctx, fn)
}

// GetDocument mocks base method.
func (m *MockRepository) GetDocument(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocument", ctx, id)
	ret0, _ := ret[0].(*domain.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocument indicates an expected call of GetDocument.
func (mr *MockRepositoryMockRecorder) GetDocument(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocument", reflect.TypeOf((*MockRepository)(nil).GetDocument), ctx, id)
}

// GetDocuments mocks base method.
func (m *MockRepository) GetDocuments(ctx context.Context, filter *dto.GetDocuments) ([]domain.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocuments", ctx, filter)
	ret0, _ := ret[0].([]domain.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocuments indicates an expected call of GetDocuments.
func (mr *MockRepositoryMockRecorder) GetDocuments(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocuments", reflect.TypeOf((*MockRepository)(nil).GetDocuments), ctx, filter)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), ctx, id)
}

// GetUserID mocks base method.
func (m *MockRepository) GetUserID(ctx context.Context, token string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", ctx, token)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockRepositoryMockRecorder) GetUserID(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockRepository)(nil).GetUserID), ctx, token)
}

// LogOut mocks base method.
func (m *MockRepository) LogOut(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogOut", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogOut indicates an expected call of LogOut.
func (mr *MockRepositoryMockRecorder) LogOut(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOut", reflect.TypeOf((*MockRepository)(nil).LogOut), ctx, token)
}

// Registration mocks base method.
func (m *MockRepository) Registration(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registration", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Registration indicates an expected call of Registration.
func (mr *MockRepositoryMockRecorder) Registration(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockRepository)(nil).Registration), ctx, user)
}

// Save mocks base method.
func (m *MockRepository) Save(ctx context.Context, document *domain.Document) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, document)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(ctx, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, document)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCache) Get(k string) (any, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", k)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), k)
}

// Set mocks base method.
func (m *MockCache) Set(k string, x any, d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", k, x, d)
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(k, x, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), k, x, d)
}
