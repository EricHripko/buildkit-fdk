// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moby/buildkit/frontend/gateway/pb (interfaces: LLBBridgeClient,LLBBridge_ExecProcessClient)
//
// Generated by this command:
//
//	mockgen -package dtp_mock -destination mock/client.go github.com/moby/buildkit/frontend/gateway/pb LLBBridgeClient,LLBBridge_ExecProcessClient
//

// Package dtp_mock is a generated GoMock package.
package dtp_mock

import (
	context "context"
	reflect "reflect"

	moby_buildkit_v1_frontend "github.com/moby/buildkit/frontend/gateway/pb"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockLLBBridgeClient is a mock of LLBBridgeClient interface.
type MockLLBBridgeClient struct {
	ctrl     *gomock.Controller
	recorder *MockLLBBridgeClientMockRecorder
}

// MockLLBBridgeClientMockRecorder is the mock recorder for MockLLBBridgeClient.
type MockLLBBridgeClientMockRecorder struct {
	mock *MockLLBBridgeClient
}

// NewMockLLBBridgeClient creates a new mock instance.
func NewMockLLBBridgeClient(ctrl *gomock.Controller) *MockLLBBridgeClient {
	mock := &MockLLBBridgeClient{ctrl: ctrl}
	mock.recorder = &MockLLBBridgeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLLBBridgeClient) EXPECT() *MockLLBBridgeClientMockRecorder {
	return m.recorder
}

// Evaluate mocks base method.
func (m *MockLLBBridgeClient) Evaluate(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.EvaluateRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.EvaluateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Evaluate", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.EvaluateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Evaluate indicates an expected call of Evaluate.
func (mr *MockLLBBridgeClientMockRecorder) Evaluate(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evaluate", reflect.TypeOf((*MockLLBBridgeClient)(nil).Evaluate), varargs...)
}

// ExecProcess mocks base method.
func (m *MockLLBBridgeClient) ExecProcess(arg0 context.Context, arg1 ...grpc.CallOption) (moby_buildkit_v1_frontend.LLBBridge_ExecProcessClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecProcess", varargs...)
	ret0, _ := ret[0].(moby_buildkit_v1_frontend.LLBBridge_ExecProcessClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecProcess indicates an expected call of ExecProcess.
func (mr *MockLLBBridgeClientMockRecorder) ExecProcess(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecProcess", reflect.TypeOf((*MockLLBBridgeClient)(nil).ExecProcess), varargs...)
}

// Inputs mocks base method.
func (m *MockLLBBridgeClient) Inputs(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.InputsRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.InputsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Inputs", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.InputsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inputs indicates an expected call of Inputs.
func (mr *MockLLBBridgeClientMockRecorder) Inputs(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inputs", reflect.TypeOf((*MockLLBBridgeClient)(nil).Inputs), varargs...)
}

// NewContainer mocks base method.
func (m *MockLLBBridgeClient) NewContainer(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.NewContainerRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.NewContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewContainer", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.NewContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewContainer indicates an expected call of NewContainer.
func (mr *MockLLBBridgeClientMockRecorder) NewContainer(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewContainer", reflect.TypeOf((*MockLLBBridgeClient)(nil).NewContainer), varargs...)
}

// Ping mocks base method.
func (m *MockLLBBridgeClient) Ping(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.PingRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.PongResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.PongResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockLLBBridgeClientMockRecorder) Ping(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockLLBBridgeClient)(nil).Ping), varargs...)
}

// ReadDir mocks base method.
func (m *MockLLBBridgeClient) ReadDir(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.ReadDirRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.ReadDirResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadDir", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.ReadDirResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockLLBBridgeClientMockRecorder) ReadDir(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockLLBBridgeClient)(nil).ReadDir), varargs...)
}

// ReadFile mocks base method.
func (m *MockLLBBridgeClient) ReadFile(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.ReadFileRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.ReadFileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadFile", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.ReadFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockLLBBridgeClientMockRecorder) ReadFile(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockLLBBridgeClient)(nil).ReadFile), varargs...)
}

// ReleaseContainer mocks base method.
func (m *MockLLBBridgeClient) ReleaseContainer(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.ReleaseContainerRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.ReleaseContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReleaseContainer", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.ReleaseContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseContainer indicates an expected call of ReleaseContainer.
func (mr *MockLLBBridgeClientMockRecorder) ReleaseContainer(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseContainer", reflect.TypeOf((*MockLLBBridgeClient)(nil).ReleaseContainer), varargs...)
}

// ResolveImageConfig mocks base method.
func (m *MockLLBBridgeClient) ResolveImageConfig(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.ResolveImageConfigRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.ResolveImageConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResolveImageConfig", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.ResolveImageConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveImageConfig indicates an expected call of ResolveImageConfig.
func (mr *MockLLBBridgeClientMockRecorder) ResolveImageConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveImageConfig", reflect.TypeOf((*MockLLBBridgeClient)(nil).ResolveImageConfig), varargs...)
}

// Return mocks base method.
func (m *MockLLBBridgeClient) Return(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.ReturnRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.ReturnResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Return", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.ReturnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Return indicates an expected call of Return.
func (mr *MockLLBBridgeClientMockRecorder) Return(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Return", reflect.TypeOf((*MockLLBBridgeClient)(nil).Return), varargs...)
}

// Solve mocks base method.
func (m *MockLLBBridgeClient) Solve(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.SolveRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.SolveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Solve", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.SolveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Solve indicates an expected call of Solve.
func (mr *MockLLBBridgeClientMockRecorder) Solve(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Solve", reflect.TypeOf((*MockLLBBridgeClient)(nil).Solve), varargs...)
}

// StatFile mocks base method.
func (m *MockLLBBridgeClient) StatFile(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.StatFileRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.StatFileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatFile", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.StatFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatFile indicates an expected call of StatFile.
func (mr *MockLLBBridgeClientMockRecorder) StatFile(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatFile", reflect.TypeOf((*MockLLBBridgeClient)(nil).StatFile), varargs...)
}

// Warn mocks base method.
func (m *MockLLBBridgeClient) Warn(arg0 context.Context, arg1 *moby_buildkit_v1_frontend.WarnRequest, arg2 ...grpc.CallOption) (*moby_buildkit_v1_frontend.WarnResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Warn", varargs...)
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.WarnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Warn indicates an expected call of Warn.
func (mr *MockLLBBridgeClientMockRecorder) Warn(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLLBBridgeClient)(nil).Warn), varargs...)
}

// MockLLBBridge_ExecProcessClient is a mock of LLBBridge_ExecProcessClient interface.
type MockLLBBridge_ExecProcessClient struct {
	ctrl     *gomock.Controller
	recorder *MockLLBBridge_ExecProcessClientMockRecorder
}

// MockLLBBridge_ExecProcessClientMockRecorder is the mock recorder for MockLLBBridge_ExecProcessClient.
type MockLLBBridge_ExecProcessClientMockRecorder struct {
	mock *MockLLBBridge_ExecProcessClient
}

// NewMockLLBBridge_ExecProcessClient creates a new mock instance.
func NewMockLLBBridge_ExecProcessClient(ctrl *gomock.Controller) *MockLLBBridge_ExecProcessClient {
	mock := &MockLLBBridge_ExecProcessClient{ctrl: ctrl}
	mock.recorder = &MockLLBBridge_ExecProcessClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLLBBridge_ExecProcessClient) EXPECT() *MockLLBBridge_ExecProcessClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockLLBBridge_ExecProcessClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockLLBBridge_ExecProcessClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).Context))
}

// Header mocks base method.
func (m *MockLLBBridge_ExecProcessClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockLLBBridge_ExecProcessClient) Recv() (*moby_buildkit_v1_frontend.ExecMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*moby_buildkit_v1_frontend.ExecMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockLLBBridge_ExecProcessClient) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) RecvMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockLLBBridge_ExecProcessClient) Send(arg0 *moby_buildkit_v1_frontend.ExecMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockLLBBridge_ExecProcessClient) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) SendMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockLLBBridge_ExecProcessClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockLLBBridge_ExecProcessClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockLLBBridge_ExecProcessClient)(nil).Trailer))
}
