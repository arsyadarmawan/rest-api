// Code generated by MockGen. DO NOT EDIT.
// Source: book.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"
	usecase "github.com/arsyadarmawan/rest-api/internal/app/book/usecase"

	gomock "github.com/golang/mock/gomock"
)

// MockBook is a mock of Book interface.
type MockBook struct {
	ctrl     *gomock.Controller
	recorder *MockBookMockRecorder
}

// MockBookMockRecorder is the mock recorder for MockBook.
type MockBookMockRecorder struct {
	mock *MockBook
}

// NewMockBook creates a new mock instance.
func NewMockBook(ctrl *gomock.Controller) *MockBook {
	mock := &MockBook{ctrl: ctrl}
	mock.recorder = &MockBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBook) EXPECT() *MockBookMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBook) Create(ctx context.Context, cmd usecase.BookRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBookMockRecorder) Create(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBook)(nil).Create), ctx, cmd)
}

// Delete mocks base method.
func (m *MockBook) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBook)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockBook) Get(ctx context.Context) ([]usecase.BookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]usecase.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBookMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBook)(nil).Get), ctx)
}

// GetById mocks base method.
func (m *MockBook) GetById(ctx context.Context, id string) (usecase.BookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(usecase.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBookMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBook)(nil).GetById), ctx, id)
}

// Update mocks base method.
func (m *MockBook) Update(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBookMockRecorder) Update(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBook)(nil).Update), ctx, id)
}
