// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lushenle/plam/pkg/db (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/lushenle/plam/pkg/db"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateIncome mocks base method.
func (m *MockStore) CreateIncome(arg0 context.Context, arg1 db.CreateIncomeParams) (db.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncome", arg0, arg1)
	ret0, _ := ret[0].(db.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncome indicates an expected call of CreateIncome.
func (mr *MockStoreMockRecorder) CreateIncome(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncome", reflect.TypeOf((*MockStore)(nil).CreateIncome), arg0, arg1)
}

// CreateLoan mocks base method.
func (m *MockStore) CreateLoan(arg0 context.Context, arg1 db.CreateLoanParams) (db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoan", arg0, arg1)
	ret0, _ := ret[0].(db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoan indicates an expected call of CreateLoan.
func (mr *MockStoreMockRecorder) CreateLoan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoan", reflect.TypeOf((*MockStore)(nil).CreateLoan), arg0, arg1)
}

// CreatePayOut mocks base method.
func (m *MockStore) CreatePayOut(arg0 context.Context, arg1 db.CreatePayOutParams) (db.PayOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayOut", arg0, arg1)
	ret0, _ := ret[0].(db.PayOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayOut indicates an expected call of CreatePayOut.
func (mr *MockStoreMockRecorder) CreatePayOut(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayOut", reflect.TypeOf((*MockStore)(nil).CreatePayOut), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockStore) CreateProject(arg0 context.Context, arg1 db.CreateProjectParams) (db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockStoreMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockStore)(nil).CreateProject), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteIncome mocks base method.
func (m *MockStore) DeleteIncome(arg0 context.Context, arg1 string) (db.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIncome", arg0, arg1)
	ret0, _ := ret[0].(db.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIncome indicates an expected call of DeleteIncome.
func (mr *MockStoreMockRecorder) DeleteIncome(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIncome", reflect.TypeOf((*MockStore)(nil).DeleteIncome), arg0, arg1)
}

// DeleteLoan mocks base method.
func (m *MockStore) DeleteLoan(arg0 context.Context, arg1 string) (db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoan", arg0, arg1)
	ret0, _ := ret[0].(db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoan indicates an expected call of DeleteLoan.
func (mr *MockStoreMockRecorder) DeleteLoan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoan", reflect.TypeOf((*MockStore)(nil).DeleteLoan), arg0, arg1)
}

// DeletePayOut mocks base method.
func (m *MockStore) DeletePayOut(arg0 context.Context, arg1 string) (db.PayOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePayOut", arg0, arg1)
	ret0, _ := ret[0].(db.PayOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePayOut indicates an expected call of DeletePayOut.
func (mr *MockStoreMockRecorder) DeletePayOut(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePayOut", reflect.TypeOf((*MockStore)(nil).DeletePayOut), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockStore) DeleteProject(arg0 context.Context, arg1 string) (db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockStoreMockRecorder) DeleteProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockStore)(nil).DeleteProject), arg0, arg1)
}

// GetIncome mocks base method.
func (m *MockStore) GetIncome(arg0 context.Context, arg1 string) (db.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncome", arg0, arg1)
	ret0, _ := ret[0].(db.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncome indicates an expected call of GetIncome.
func (mr *MockStoreMockRecorder) GetIncome(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncome", reflect.TypeOf((*MockStore)(nil).GetIncome), arg0, arg1)
}

// GetLoan mocks base method.
func (m *MockStore) GetLoan(arg0 context.Context, arg1 string) (db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoan", arg0, arg1)
	ret0, _ := ret[0].(db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoan indicates an expected call of GetLoan.
func (mr *MockStoreMockRecorder) GetLoan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoan", reflect.TypeOf((*MockStore)(nil).GetLoan), arg0, arg1)
}

// GetPayOut mocks base method.
func (m *MockStore) GetPayOut(arg0 context.Context, arg1 string) (db.PayOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayOut", arg0, arg1)
	ret0, _ := ret[0].(db.PayOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayOut indicates an expected call of GetPayOut.
func (mr *MockStoreMockRecorder) GetPayOut(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayOut", reflect.TypeOf((*MockStore)(nil).GetPayOut), arg0, arg1)
}

// GetProject mocks base method.
func (m *MockStore) GetProject(arg0 context.Context, arg1 string) (db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockStoreMockRecorder) GetProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockStore)(nil).GetProject), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListIncomes mocks base method.
func (m *MockStore) ListIncomes(arg0 context.Context, arg1 db.ListIncomesParams) ([]db.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIncomes", arg0, arg1)
	ret0, _ := ret[0].([]db.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIncomes indicates an expected call of ListIncomes.
func (mr *MockStoreMockRecorder) ListIncomes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncomes", reflect.TypeOf((*MockStore)(nil).ListIncomes), arg0, arg1)
}

// ListLoans mocks base method.
func (m *MockStore) ListLoans(arg0 context.Context, arg1 db.ListLoansParams) ([]db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLoans", arg0, arg1)
	ret0, _ := ret[0].([]db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoans indicates an expected call of ListLoans.
func (mr *MockStoreMockRecorder) ListLoans(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoans", reflect.TypeOf((*MockStore)(nil).ListLoans), arg0, arg1)
}

// ListPayOuts mocks base method.
func (m *MockStore) ListPayOuts(arg0 context.Context, arg1 db.ListPayOutsParams) ([]db.PayOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPayOuts", arg0, arg1)
	ret0, _ := ret[0].([]db.PayOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPayOuts indicates an expected call of ListPayOuts.
func (mr *MockStoreMockRecorder) ListPayOuts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPayOuts", reflect.TypeOf((*MockStore)(nil).ListPayOuts), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockStore) ListProjects(arg0 context.Context, arg1 db.ListProjectsParams) ([]db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].([]db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockStoreMockRecorder) ListProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockStore)(nil).ListProjects), arg0, arg1)
}

// SearchIncomes mocks base method.
func (m *MockStore) SearchIncomes(arg0 context.Context, arg1 db.SearchIncomesParams) ([]db.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIncomes", arg0, arg1)
	ret0, _ := ret[0].([]db.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIncomes indicates an expected call of SearchIncomes.
func (mr *MockStoreMockRecorder) SearchIncomes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIncomes", reflect.TypeOf((*MockStore)(nil).SearchIncomes), arg0, arg1)
}

// SearchLoans mocks base method.
func (m *MockStore) SearchLoans(arg0 context.Context, arg1 db.SearchLoansParams) ([]db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLoans", arg0, arg1)
	ret0, _ := ret[0].([]db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLoans indicates an expected call of SearchLoans.
func (mr *MockStoreMockRecorder) SearchLoans(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLoans", reflect.TypeOf((*MockStore)(nil).SearchLoans), arg0, arg1)
}

// SearchPayOuts mocks base method.
func (m *MockStore) SearchPayOuts(arg0 context.Context, arg1 db.SearchPayOutsParams) ([]db.PayOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPayOuts", arg0, arg1)
	ret0, _ := ret[0].([]db.PayOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPayOuts indicates an expected call of SearchPayOuts.
func (mr *MockStoreMockRecorder) SearchPayOuts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPayOuts", reflect.TypeOf((*MockStore)(nil).SearchPayOuts), arg0, arg1)
}

// SearchProjects mocks base method.
func (m *MockStore) SearchProjects(arg0 context.Context, arg1 db.SearchProjectsParams) ([]db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProjects", arg0, arg1)
	ret0, _ := ret[0].([]db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProjects indicates an expected call of SearchProjects.
func (mr *MockStoreMockRecorder) SearchProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProjects", reflect.TypeOf((*MockStore)(nil).SearchProjects), arg0, arg1)
}
