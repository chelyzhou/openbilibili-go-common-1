// Code generated by MockGen. DO NOT EDIT.
// Source: api.pb.go

// Package api is a generated GoMock package.
package api

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAccountClient is a mock of AccountClient interface
type MockAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountClientMockRecorder
}

// MockAccountClientMockRecorder is the mock recorder for MockAccountClient
type MockAccountClientMockRecorder struct {
	mock *MockAccountClient
}

// NewMockAccountClient creates a new mock instance
func NewMockAccountClient(ctrl *gomock.Controller) *MockAccountClient {
	mock := &MockAccountClient{ctrl: ctrl}
	mock.recorder = &MockAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountClient) EXPECT() *MockAccountClientMockRecorder {
	return m.recorder
}

// Info3 mocks base method
func (m *MockAccountClient) Info3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*InfoReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info3", varargs...)
	ret0, _ := ret[0].(*InfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info3 indicates an expected call of Info3
func (mr *MockAccountClientMockRecorder) Info3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info3", reflect.TypeOf((*MockAccountClient)(nil).Info3), varargs...)
}

// Infos3 mocks base method
func (m *MockAccountClient) Infos3(ctx context.Context, in *MidsReq, opts ...grpc.CallOption) (*InfosReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Infos3", varargs...)
	ret0, _ := ret[0].(*InfosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Infos3 indicates an expected call of Infos3
func (mr *MockAccountClientMockRecorder) Infos3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infos3", reflect.TypeOf((*MockAccountClient)(nil).Infos3), varargs...)
}

// InfosByName3 mocks base method
func (m *MockAccountClient) InfosByName3(ctx context.Context, in *NamesReq, opts ...grpc.CallOption) (*InfosReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InfosByName3", varargs...)
	ret0, _ := ret[0].(*InfosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfosByName3 indicates an expected call of InfosByName3
func (mr *MockAccountClientMockRecorder) InfosByName3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfosByName3", reflect.TypeOf((*MockAccountClient)(nil).InfosByName3), varargs...)
}

// Card3 mocks base method
func (m *MockAccountClient) Card3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*CardReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Card3", varargs...)
	ret0, _ := ret[0].(*CardReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Card3 indicates an expected call of Card3
func (mr *MockAccountClientMockRecorder) Card3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Card3", reflect.TypeOf((*MockAccountClient)(nil).Card3), varargs...)
}

// Cards3 mocks base method
func (m *MockAccountClient) Cards3(ctx context.Context, in *MidsReq, opts ...grpc.CallOption) (*CardsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Cards3", varargs...)
	ret0, _ := ret[0].(*CardsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cards3 indicates an expected call of Cards3
func (mr *MockAccountClientMockRecorder) Cards3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cards3", reflect.TypeOf((*MockAccountClient)(nil).Cards3), varargs...)
}

// Profile3 mocks base method
func (m *MockAccountClient) Profile3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*ProfileReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Profile3", varargs...)
	ret0, _ := ret[0].(*ProfileReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile3 indicates an expected call of Profile3
func (mr *MockAccountClientMockRecorder) Profile3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile3", reflect.TypeOf((*MockAccountClient)(nil).Profile3), varargs...)
}

// ProfileWithStat3 mocks base method
func (m *MockAccountClient) ProfileWithStat3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*ProfileStatReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProfileWithStat3", varargs...)
	ret0, _ := ret[0].(*ProfileStatReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileWithStat3 indicates an expected call of ProfileWithStat3
func (mr *MockAccountClientMockRecorder) ProfileWithStat3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileWithStat3", reflect.TypeOf((*MockAccountClient)(nil).ProfileWithStat3), varargs...)
}

// AddExp3 mocks base method
func (m *MockAccountClient) AddExp3(ctx context.Context, in *ExpReq, opts ...grpc.CallOption) (*ExpReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddExp3", varargs...)
	ret0, _ := ret[0].(*ExpReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddExp3 indicates an expected call of AddExp3
func (mr *MockAccountClientMockRecorder) AddExp3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExp3", reflect.TypeOf((*MockAccountClient)(nil).AddExp3), varargs...)
}

// AddMoral3 mocks base method
func (m *MockAccountClient) AddMoral3(ctx context.Context, in *MoralReq, opts ...grpc.CallOption) (*MoralReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMoral3", varargs...)
	ret0, _ := ret[0].(*MoralReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMoral3 indicates an expected call of AddMoral3
func (mr *MockAccountClientMockRecorder) AddMoral3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMoral3", reflect.TypeOf((*MockAccountClient)(nil).AddMoral3), varargs...)
}

// Relation3 mocks base method
func (m *MockAccountClient) Relation3(ctx context.Context, in *RelationReq, opts ...grpc.CallOption) (*RelationReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Relation3", varargs...)
	ret0, _ := ret[0].(*RelationReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation3 indicates an expected call of Relation3
func (mr *MockAccountClientMockRecorder) Relation3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation3", reflect.TypeOf((*MockAccountClient)(nil).Relation3), varargs...)
}

// Attentions3 mocks base method
func (m *MockAccountClient) Attentions3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*AttentionsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Attentions3", varargs...)
	ret0, _ := ret[0].(*AttentionsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attentions3 indicates an expected call of Attentions3
func (mr *MockAccountClientMockRecorder) Attentions3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attentions3", reflect.TypeOf((*MockAccountClient)(nil).Attentions3), varargs...)
}

// Blacks3 mocks base method
func (m *MockAccountClient) Blacks3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*BlacksReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Blacks3", varargs...)
	ret0, _ := ret[0].(*BlacksReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Blacks3 indicates an expected call of Blacks3
func (mr *MockAccountClientMockRecorder) Blacks3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blacks3", reflect.TypeOf((*MockAccountClient)(nil).Blacks3), varargs...)
}

// Relations3 mocks base method
func (m *MockAccountClient) Relations3(ctx context.Context, in *RelationsReq, opts ...grpc.CallOption) (*RelationsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Relations3", varargs...)
	ret0, _ := ret[0].(*RelationsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relations3 indicates an expected call of Relations3
func (mr *MockAccountClientMockRecorder) Relations3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relations3", reflect.TypeOf((*MockAccountClient)(nil).Relations3), varargs...)
}

// RichRelations3 mocks base method
func (m *MockAccountClient) RichRelations3(ctx context.Context, in *RichRelationReq, opts ...grpc.CallOption) (*RichRelationsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RichRelations3", varargs...)
	ret0, _ := ret[0].(*RichRelationsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RichRelations3 indicates an expected call of RichRelations3
func (mr *MockAccountClientMockRecorder) RichRelations3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RichRelations3", reflect.TypeOf((*MockAccountClient)(nil).RichRelations3), varargs...)
}

// Vip3 mocks base method
func (m *MockAccountClient) Vip3(ctx context.Context, in *MidReq, opts ...grpc.CallOption) (*VipReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Vip3", varargs...)
	ret0, _ := ret[0].(*VipReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vip3 indicates an expected call of Vip3
func (mr *MockAccountClientMockRecorder) Vip3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vip3", reflect.TypeOf((*MockAccountClient)(nil).Vip3), varargs...)
}

// Vips3 mocks base method
func (m *MockAccountClient) Vips3(ctx context.Context, in *MidsReq, opts ...grpc.CallOption) (*VipsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Vips3", varargs...)
	ret0, _ := ret[0].(*VipsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vips3 indicates an expected call of Vips3
func (mr *MockAccountClientMockRecorder) Vips3(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vips3", reflect.TypeOf((*MockAccountClient)(nil).Vips3), varargs...)
}

// MockAccountServer is a mock of AccountServer interface
type MockAccountServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServerMockRecorder
}

// MockAccountServerMockRecorder is the mock recorder for MockAccountServer
type MockAccountServerMockRecorder struct {
	mock *MockAccountServer
}

// NewMockAccountServer creates a new mock instance
func NewMockAccountServer(ctrl *gomock.Controller) *MockAccountServer {
	mock := &MockAccountServer{ctrl: ctrl}
	mock.recorder = &MockAccountServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountServer) EXPECT() *MockAccountServerMockRecorder {
	return m.recorder
}

// Info3 mocks base method
func (m *MockAccountServer) Info3(arg0 context.Context, arg1 *MidReq) (*InfoReply, error) {
	ret := m.ctrl.Call(m, "Info3", arg0, arg1)
	ret0, _ := ret[0].(*InfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info3 indicates an expected call of Info3
func (mr *MockAccountServerMockRecorder) Info3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info3", reflect.TypeOf((*MockAccountServer)(nil).Info3), arg0, arg1)
}

// Infos3 mocks base method
func (m *MockAccountServer) Infos3(arg0 context.Context, arg1 *MidsReq) (*InfosReply, error) {
	ret := m.ctrl.Call(m, "Infos3", arg0, arg1)
	ret0, _ := ret[0].(*InfosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Infos3 indicates an expected call of Infos3
func (mr *MockAccountServerMockRecorder) Infos3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infos3", reflect.TypeOf((*MockAccountServer)(nil).Infos3), arg0, arg1)
}

// InfosByName3 mocks base method
func (m *MockAccountServer) InfosByName3(arg0 context.Context, arg1 *NamesReq) (*InfosReply, error) {
	ret := m.ctrl.Call(m, "InfosByName3", arg0, arg1)
	ret0, _ := ret[0].(*InfosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfosByName3 indicates an expected call of InfosByName3
func (mr *MockAccountServerMockRecorder) InfosByName3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfosByName3", reflect.TypeOf((*MockAccountServer)(nil).InfosByName3), arg0, arg1)
}

// Card3 mocks base method
func (m *MockAccountServer) Card3(arg0 context.Context, arg1 *MidReq) (*CardReply, error) {
	ret := m.ctrl.Call(m, "Card3", arg0, arg1)
	ret0, _ := ret[0].(*CardReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Card3 indicates an expected call of Card3
func (mr *MockAccountServerMockRecorder) Card3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Card3", reflect.TypeOf((*MockAccountServer)(nil).Card3), arg0, arg1)
}

// Cards3 mocks base method
func (m *MockAccountServer) Cards3(arg0 context.Context, arg1 *MidsReq) (*CardsReply, error) {
	ret := m.ctrl.Call(m, "Cards3", arg0, arg1)
	ret0, _ := ret[0].(*CardsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cards3 indicates an expected call of Cards3
func (mr *MockAccountServerMockRecorder) Cards3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cards3", reflect.TypeOf((*MockAccountServer)(nil).Cards3), arg0, arg1)
}

// Profile3 mocks base method
func (m *MockAccountServer) Profile3(arg0 context.Context, arg1 *MidReq) (*ProfileReply, error) {
	ret := m.ctrl.Call(m, "Profile3", arg0, arg1)
	ret0, _ := ret[0].(*ProfileReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile3 indicates an expected call of Profile3
func (mr *MockAccountServerMockRecorder) Profile3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile3", reflect.TypeOf((*MockAccountServer)(nil).Profile3), arg0, arg1)
}

// ProfileWithStat3 mocks base method
func (m *MockAccountServer) ProfileWithStat3(arg0 context.Context, arg1 *MidReq) (*ProfileStatReply, error) {
	ret := m.ctrl.Call(m, "ProfileWithStat3", arg0, arg1)
	ret0, _ := ret[0].(*ProfileStatReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileWithStat3 indicates an expected call of ProfileWithStat3
func (mr *MockAccountServerMockRecorder) ProfileWithStat3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileWithStat3", reflect.TypeOf((*MockAccountServer)(nil).ProfileWithStat3), arg0, arg1)
}

// AddExp3 mocks base method
func (m *MockAccountServer) AddExp3(arg0 context.Context, arg1 *ExpReq) (*ExpReply, error) {
	ret := m.ctrl.Call(m, "AddExp3", arg0, arg1)
	ret0, _ := ret[0].(*ExpReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddExp3 indicates an expected call of AddExp3
func (mr *MockAccountServerMockRecorder) AddExp3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExp3", reflect.TypeOf((*MockAccountServer)(nil).AddExp3), arg0, arg1)
}

// AddMoral3 mocks base method
func (m *MockAccountServer) AddMoral3(arg0 context.Context, arg1 *MoralReq) (*MoralReply, error) {
	ret := m.ctrl.Call(m, "AddMoral3", arg0, arg1)
	ret0, _ := ret[0].(*MoralReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMoral3 indicates an expected call of AddMoral3
func (mr *MockAccountServerMockRecorder) AddMoral3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMoral3", reflect.TypeOf((*MockAccountServer)(nil).AddMoral3), arg0, arg1)
}

// Relation3 mocks base method
func (m *MockAccountServer) Relation3(arg0 context.Context, arg1 *RelationReq) (*RelationReply, error) {
	ret := m.ctrl.Call(m, "Relation3", arg0, arg1)
	ret0, _ := ret[0].(*RelationReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation3 indicates an expected call of Relation3
func (mr *MockAccountServerMockRecorder) Relation3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation3", reflect.TypeOf((*MockAccountServer)(nil).Relation3), arg0, arg1)
}

// Attentions3 mocks base method
func (m *MockAccountServer) Attentions3(arg0 context.Context, arg1 *MidReq) (*AttentionsReply, error) {
	ret := m.ctrl.Call(m, "Attentions3", arg0, arg1)
	ret0, _ := ret[0].(*AttentionsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attentions3 indicates an expected call of Attentions3
func (mr *MockAccountServerMockRecorder) Attentions3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attentions3", reflect.TypeOf((*MockAccountServer)(nil).Attentions3), arg0, arg1)
}

// Blacks3 mocks base method
func (m *MockAccountServer) Blacks3(arg0 context.Context, arg1 *MidReq) (*BlacksReply, error) {
	ret := m.ctrl.Call(m, "Blacks3", arg0, arg1)
	ret0, _ := ret[0].(*BlacksReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Blacks3 indicates an expected call of Blacks3
func (mr *MockAccountServerMockRecorder) Blacks3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blacks3", reflect.TypeOf((*MockAccountServer)(nil).Blacks3), arg0, arg1)
}

// Relations3 mocks base method
func (m *MockAccountServer) Relations3(arg0 context.Context, arg1 *RelationsReq) (*RelationsReply, error) {
	ret := m.ctrl.Call(m, "Relations3", arg0, arg1)
	ret0, _ := ret[0].(*RelationsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relations3 indicates an expected call of Relations3
func (mr *MockAccountServerMockRecorder) Relations3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relations3", reflect.TypeOf((*MockAccountServer)(nil).Relations3), arg0, arg1)
}

// RichRelations3 mocks base method
func (m *MockAccountServer) RichRelations3(arg0 context.Context, arg1 *RichRelationReq) (*RichRelationsReply, error) {
	ret := m.ctrl.Call(m, "RichRelations3", arg0, arg1)
	ret0, _ := ret[0].(*RichRelationsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RichRelations3 indicates an expected call of RichRelations3
func (mr *MockAccountServerMockRecorder) RichRelations3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RichRelations3", reflect.TypeOf((*MockAccountServer)(nil).RichRelations3), arg0, arg1)
}

// Vip3 mocks base method
func (m *MockAccountServer) Vip3(arg0 context.Context, arg1 *MidReq) (*VipReply, error) {
	ret := m.ctrl.Call(m, "Vip3", arg0, arg1)
	ret0, _ := ret[0].(*VipReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vip3 indicates an expected call of Vip3
func (mr *MockAccountServerMockRecorder) Vip3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vip3", reflect.TypeOf((*MockAccountServer)(nil).Vip3), arg0, arg1)
}

// Vips3 mocks base method
func (m *MockAccountServer) Vips3(arg0 context.Context, arg1 *MidsReq) (*VipsReply, error) {
	ret := m.ctrl.Call(m, "Vips3", arg0, arg1)
	ret0, _ := ret[0].(*VipsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vips3 indicates an expected call of Vips3
func (mr *MockAccountServerMockRecorder) Vips3(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vips3", reflect.TypeOf((*MockAccountServer)(nil).Vips3), arg0, arg1)
}
