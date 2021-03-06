// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/roles_v1.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorDescription struct {
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Category             string            `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Code                 string            `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	CorrelationId        string            `protobuf:"bytes,4,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Status               string            `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Message              string            `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Cause                string            `protobuf:"bytes,7,opt,name=cause,proto3" json:"cause,omitempty"`
	StackTrace           string            `protobuf:"bytes,8,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Details              map[string]string `protobuf:"bytes,9,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ErrorDescription) Reset()         { *m = ErrorDescription{} }
func (m *ErrorDescription) String() string { return proto.CompactTextString(m) }
func (*ErrorDescription) ProtoMessage()    {}
func (*ErrorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{0}
}

func (m *ErrorDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDescription.Unmarshal(m, b)
}
func (m *ErrorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDescription.Marshal(b, m, deterministic)
}
func (m *ErrorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDescription.Merge(m, src)
}
func (m *ErrorDescription) XXX_Size() int {
	return xxx_messageInfo_ErrorDescription.Size(m)
}
func (m *ErrorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDescription proto.InternalMessageInfo

func (m *ErrorDescription) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ErrorDescription) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ErrorDescription) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ErrorDescription) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *ErrorDescription) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ErrorDescription) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorDescription) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

func (m *ErrorDescription) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

func (m *ErrorDescription) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

type PagingParams struct {
	Skip                 int64    `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take                 int32    `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	Total                bool     `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagingParams) Reset()         { *m = PagingParams{} }
func (m *PagingParams) String() string { return proto.CompactTextString(m) }
func (*PagingParams) ProtoMessage()    {}
func (*PagingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{1}
}

func (m *PagingParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingParams.Unmarshal(m, b)
}
func (m *PagingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingParams.Marshal(b, m, deterministic)
}
func (m *PagingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingParams.Merge(m, src)
}
func (m *PagingParams) XXX_Size() int {
	return xxx_messageInfo_PagingParams.Size(m)
}
func (m *PagingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingParams.DiscardUnknown(m)
}

var xxx_messageInfo_PagingParams proto.InternalMessageInfo

func (m *PagingParams) GetSkip() int64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *PagingParams) GetTake() int32 {
	if m != nil {
		return m.Take
	}
	return 0
}

func (m *PagingParams) GetTotal() bool {
	if m != nil {
		return m.Total
	}
	return false
}

type UserRoles struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdateTime           string   `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Roles                []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRoles) Reset()         { *m = UserRoles{} }
func (m *UserRoles) String() string { return proto.CompactTextString(m) }
func (*UserRoles) ProtoMessage()    {}
func (*UserRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{2}
}

func (m *UserRoles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRoles.Unmarshal(m, b)
}
func (m *UserRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRoles.Marshal(b, m, deterministic)
}
func (m *UserRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRoles.Merge(m, src)
}
func (m *UserRoles) XXX_Size() int {
	return xxx_messageInfo_UserRoles.Size(m)
}
func (m *UserRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRoles.DiscardUnknown(m)
}

var xxx_messageInfo_UserRoles proto.InternalMessageInfo

func (m *UserRoles) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserRoles) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *UserRoles) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type UserRolesPage struct {
	Total                int64        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*UserRoles `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserRolesPage) Reset()         { *m = UserRolesPage{} }
func (m *UserRolesPage) String() string { return proto.CompactTextString(m) }
func (*UserRolesPage) ProtoMessage()    {}
func (*UserRolesPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{3}
}

func (m *UserRolesPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRolesPage.Unmarshal(m, b)
}
func (m *UserRolesPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRolesPage.Marshal(b, m, deterministic)
}
func (m *UserRolesPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRolesPage.Merge(m, src)
}
func (m *UserRolesPage) XXX_Size() int {
	return xxx_messageInfo_UserRolesPage.Size(m)
}
func (m *UserRolesPage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRolesPage.DiscardUnknown(m)
}

var xxx_messageInfo_UserRolesPage proto.InternalMessageInfo

func (m *UserRolesPage) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *UserRolesPage) GetData() []*UserRoles {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request message containing the role page request.
type RolesPageRequest struct {
	CorrelationId        string            `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Filter               map[string]string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Paging               *PagingParams     `protobuf:"bytes,3,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RolesPageRequest) Reset()         { *m = RolesPageRequest{} }
func (m *RolesPageRequest) String() string { return proto.CompactTextString(m) }
func (*RolesPageRequest) ProtoMessage()    {}
func (*RolesPageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{4}
}

func (m *RolesPageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesPageRequest.Unmarshal(m, b)
}
func (m *RolesPageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesPageRequest.Marshal(b, m, deterministic)
}
func (m *RolesPageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesPageRequest.Merge(m, src)
}
func (m *RolesPageRequest) XXX_Size() int {
	return xxx_messageInfo_RolesPageRequest.Size(m)
}
func (m *RolesPageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesPageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RolesPageRequest proto.InternalMessageInfo

func (m *RolesPageRequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *RolesPageRequest) GetFilter() map[string]string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *RolesPageRequest) GetPaging() *PagingParams {
	if m != nil {
		return m.Paging
	}
	return nil
}

// The response message containing the role page response
type RolesPageReply struct {
	Error                *ErrorDescription `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Page                 *UserRolesPage    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RolesPageReply) Reset()         { *m = RolesPageReply{} }
func (m *RolesPageReply) String() string { return proto.CompactTextString(m) }
func (*RolesPageReply) ProtoMessage()    {}
func (*RolesPageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{5}
}

func (m *RolesPageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesPageReply.Unmarshal(m, b)
}
func (m *RolesPageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesPageReply.Marshal(b, m, deterministic)
}
func (m *RolesPageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesPageReply.Merge(m, src)
}
func (m *RolesPageReply) XXX_Size() int {
	return xxx_messageInfo_RolesPageReply.Size(m)
}
func (m *RolesPageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesPageReply.DiscardUnknown(m)
}

var xxx_messageInfo_RolesPageReply proto.InternalMessageInfo

func (m *RolesPageReply) GetError() *ErrorDescription {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RolesPageReply) GetPage() *UserRolesPage {
	if m != nil {
		return m.Page
	}
	return nil
}

// The request message containing the role id request.
type RoleIdRequest struct {
	CorrelationId        string   `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleIdRequest) Reset()         { *m = RoleIdRequest{} }
func (m *RoleIdRequest) String() string { return proto.CompactTextString(m) }
func (*RoleIdRequest) ProtoMessage()    {}
func (*RoleIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{6}
}

func (m *RoleIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleIdRequest.Unmarshal(m, b)
}
func (m *RoleIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleIdRequest.Marshal(b, m, deterministic)
}
func (m *RoleIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleIdRequest.Merge(m, src)
}
func (m *RoleIdRequest) XXX_Size() int {
	return xxx_messageInfo_RoleIdRequest.Size(m)
}
func (m *RoleIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleIdRequest proto.InternalMessageInfo

func (m *RoleIdRequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *RoleIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// The request message containing the roles request.
type RolesRequest struct {
	CorrelationId        string   `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Roles                []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolesRequest) Reset()         { *m = RolesRequest{} }
func (m *RolesRequest) String() string { return proto.CompactTextString(m) }
func (*RolesRequest) ProtoMessage()    {}
func (*RolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{7}
}

func (m *RolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesRequest.Unmarshal(m, b)
}
func (m *RolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesRequest.Marshal(b, m, deterministic)
}
func (m *RolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesRequest.Merge(m, src)
}
func (m *RolesRequest) XXX_Size() int {
	return xxx_messageInfo_RolesRequest.Size(m)
}
func (m *RolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RolesRequest proto.InternalMessageInfo

func (m *RolesRequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *RolesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RolesRequest) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

// The response message containing the roles response
type RolesReply struct {
	Error                *ErrorDescription `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Roles                []string          `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RolesReply) Reset()         { *m = RolesReply{} }
func (m *RolesReply) String() string { return proto.CompactTextString(m) }
func (*RolesReply) ProtoMessage()    {}
func (*RolesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{8}
}

func (m *RolesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesReply.Unmarshal(m, b)
}
func (m *RolesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesReply.Marshal(b, m, deterministic)
}
func (m *RolesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesReply.Merge(m, src)
}
func (m *RolesReply) XXX_Size() int {
	return xxx_messageInfo_RolesReply.Size(m)
}
func (m *RolesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesReply.DiscardUnknown(m)
}

var xxx_messageInfo_RolesReply proto.InternalMessageInfo

func (m *RolesReply) GetError() *ErrorDescription {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RolesReply) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

// The response message containing the authorize response
type AuthorizeReply struct {
	Error                *ErrorDescription `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Authorized           bool              `protobuf:"varint,2,opt,name=authorized,proto3" json:"authorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuthorizeReply) Reset()         { *m = AuthorizeReply{} }
func (m *AuthorizeReply) String() string { return proto.CompactTextString(m) }
func (*AuthorizeReply) ProtoMessage()    {}
func (*AuthorizeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_14bd5e58956bb2b8, []int{9}
}

func (m *AuthorizeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeReply.Unmarshal(m, b)
}
func (m *AuthorizeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeReply.Marshal(b, m, deterministic)
}
func (m *AuthorizeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeReply.Merge(m, src)
}
func (m *AuthorizeReply) XXX_Size() int {
	return xxx_messageInfo_AuthorizeReply.Size(m)
}
func (m *AuthorizeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeReply proto.InternalMessageInfo

func (m *AuthorizeReply) GetError() *ErrorDescription {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AuthorizeReply) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func init() {
	proto.RegisterType((*ErrorDescription)(nil), "roles_v1.ErrorDescription")
	proto.RegisterMapType((map[string]string)(nil), "roles_v1.ErrorDescription.DetailsEntry")
	proto.RegisterType((*PagingParams)(nil), "roles_v1.PagingParams")
	proto.RegisterType((*UserRoles)(nil), "roles_v1.UserRoles")
	proto.RegisterType((*UserRolesPage)(nil), "roles_v1.UserRolesPage")
	proto.RegisterType((*RolesPageRequest)(nil), "roles_v1.RolesPageRequest")
	proto.RegisterMapType((map[string]string)(nil), "roles_v1.RolesPageRequest.FilterEntry")
	proto.RegisterType((*RolesPageReply)(nil), "roles_v1.RolesPageReply")
	proto.RegisterType((*RoleIdRequest)(nil), "roles_v1.RoleIdRequest")
	proto.RegisterType((*RolesRequest)(nil), "roles_v1.RolesRequest")
	proto.RegisterType((*RolesReply)(nil), "roles_v1.RolesReply")
	proto.RegisterType((*AuthorizeReply)(nil), "roles_v1.AuthorizeReply")
}

func init() { proto.RegisterFile("protos/roles_v1.proto", fileDescriptor_14bd5e58956bb2b8) }

var fileDescriptor_14bd5e58956bb2b8 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0xd3, 0x4a,
	0x10, 0x6d, 0x9c, 0xef, 0x49, 0x9a, 0x5b, 0x6d, 0x7b, 0x5b, 0xcb, 0x0f, 0xf7, 0x16, 0x4b, 0xd0,
	0x4a, 0x08, 0x43, 0xc2, 0x0b, 0x14, 0x51, 0xd4, 0xd2, 0x22, 0x15, 0xb5, 0x34, 0x32, 0xa5, 0x0f,
	0xbc, 0x58, 0x5b, 0x7b, 0x09, 0xab, 0x38, 0xb1, 0xd9, 0x5d, 0x47, 0x32, 0x7f, 0x84, 0x77, 0x7e,
	0x1b, 0xfc, 0x0f, 0xb4, 0xbb, 0xb6, 0xeb, 0x96, 0x04, 0xa9, 0x11, 0x6f, 0x3b, 0x33, 0xe7, 0x9c,
	0x9d, 0x3d, 0x33, 0x96, 0xe1, 0xdf, 0x98, 0x45, 0x22, 0xe2, 0x8f, 0x59, 0x14, 0x12, 0xee, 0xcd,
	0xfa, 0x8e, 0x8a, 0x51, 0x2b, 0x8f, 0xed, 0x9f, 0x06, 0xac, 0x1d, 0x33, 0x16, 0xb1, 0x23, 0xc2,
	0x7d, 0x46, 0x63, 0x41, 0xa3, 0x29, 0x42, 0x50, 0x13, 0x69, 0x4c, 0xcc, 0xca, 0x76, 0x65, 0xb7,
	0xed, 0xaa, 0x33, 0xb2, 0xa0, 0xe5, 0x63, 0x41, 0x46, 0x11, 0x4b, 0x4d, 0x43, 0xe5, 0x8b, 0x58,
	0xe2, 0xfd, 0x28, 0x20, 0x66, 0x55, 0xe3, 0xe5, 0x19, 0xdd, 0x87, 0x9e, 0x1f, 0x31, 0x46, 0x42,
	0x2c, 0x25, 0x3d, 0x1a, 0x98, 0x35, 0x55, 0x5d, 0x2d, 0x65, 0x4f, 0x02, 0xb4, 0x09, 0x0d, 0x2e,
	0xb0, 0x48, 0xb8, 0x59, 0x57, 0xe5, 0x2c, 0x42, 0x26, 0x34, 0x27, 0x84, 0x73, 0x3c, 0x22, 0x66,
	0x43, 0x15, 0xf2, 0x10, 0x6d, 0x40, 0xdd, 0xc7, 0x09, 0x27, 0x66, 0x53, 0xe5, 0x75, 0x80, 0xfe,
	0x87, 0x0e, 0x17, 0xd8, 0x1f, 0x7b, 0x82, 0x61, 0x9f, 0x98, 0x2d, 0x55, 0x03, 0x95, 0xba, 0x90,
	0x19, 0x74, 0x00, 0xcd, 0x80, 0x08, 0x4c, 0x43, 0x6e, 0xb6, 0xb7, 0xab, 0xbb, 0x9d, 0xc1, 0x8e,
	0x53, 0x98, 0x72, 0xdb, 0x00, 0xe7, 0x48, 0x23, 0x8f, 0xa7, 0x82, 0xa5, 0x6e, 0xce, 0xb3, 0xf6,
	0xa0, 0x5b, 0x2e, 0xa0, 0x35, 0xa8, 0x8e, 0x49, 0x9a, 0xb9, 0x24, 0x8f, 0xb2, 0xb7, 0x19, 0x0e,
	0x13, 0x92, 0x39, 0xa4, 0x83, 0x3d, 0xe3, 0x59, 0xc5, 0x3e, 0x85, 0xee, 0x10, 0x8f, 0xe8, 0x74,
	0x34, 0xc4, 0x0c, 0x4f, 0xb8, 0xb4, 0x8c, 0x8f, 0x69, 0xac, 0xc8, 0x55, 0x57, 0x9d, 0x95, 0xed,
	0x78, 0xac, 0xc9, 0x75, 0x57, 0x9d, 0xa5, 0xa2, 0x88, 0x04, 0x0e, 0x95, 0xb7, 0x2d, 0x57, 0x07,
	0xb6, 0x0b, 0xed, 0x0f, 0x9c, 0x30, 0x57, 0x3e, 0x00, 0xf5, 0xc0, 0xa0, 0x41, 0xd6, 0x85, 0x41,
	0x03, 0x69, 0x45, 0x12, 0x07, 0x58, 0x10, 0x4f, 0xd0, 0x49, 0xde, 0x0a, 0xe8, 0xd4, 0x05, 0x9d,
	0x28, 0x4d, 0xf5, 0x74, 0xb3, 0xba, 0x5d, 0x95, 0x5d, 0xaa, 0xc0, 0x7e, 0x07, 0xab, 0x85, 0xe6,
	0x30, 0x33, 0x5a, 0x5f, 0xad, 0x7b, 0xd4, 0x01, 0xda, 0x81, 0x5a, 0x80, 0x05, 0x36, 0x0d, 0x65,
	0xe2, 0xfa, 0xb5, 0x89, 0x05, 0xd9, 0x55, 0x00, 0xfb, 0x47, 0x05, 0xd6, 0x0a, 0x31, 0x97, 0x7c,
	0x49, 0x08, 0x17, 0x73, 0xb6, 0xa2, 0x32, 0x6f, 0x2b, 0xf6, 0xa1, 0xf1, 0x89, 0x86, 0x82, 0xb0,
	0xec, 0x9a, 0x07, 0xd7, 0xd7, 0xdc, 0x96, 0x74, 0xde, 0x28, 0xa0, 0x1e, 0x55, 0xc6, 0x42, 0x0e,
	0x34, 0x62, 0xe5, 0xb6, 0xb2, 0xad, 0x33, 0xd8, 0xbc, 0xe6, 0x97, 0xa7, 0xe0, 0x66, 0x28, 0xeb,
	0x39, 0x74, 0x4a, 0x32, 0x77, 0x1a, 0x6c, 0x04, 0xbd, 0x52, 0x4b, 0x71, 0x98, 0xa2, 0x27, 0x50,
	0x27, 0x72, 0xa1, 0x14, 0xbf, 0x33, 0xb0, 0x16, 0xef, 0x99, 0xab, 0x81, 0xe8, 0x21, 0xd4, 0x62,
	0xb9, 0xe9, 0x86, 0x22, 0x6c, 0xcd, 0xf1, 0x54, 0xa9, 0x2b, 0x90, 0x7d, 0x0e, 0xab, 0x32, 0x75,
	0x12, 0xdc, 0xd1, 0xd3, 0x2d, 0x68, 0x26, 0x9c, 0x30, 0x59, 0xd7, 0x8f, 0x68, 0xc8, 0xf0, 0x24,
	0xb0, 0x03, 0xe8, 0xea, 0xb9, 0xfd, 0x1d, 0xbd, 0x05, 0xeb, 0x75, 0x01, 0x90, 0xdd, 0xb2, 0x9c,
	0x47, 0x85, 0xaa, 0x51, 0x56, 0xbd, 0x82, 0xde, 0x41, 0x22, 0x3e, 0x47, 0x8c, 0x7e, 0x5d, 0xda,
	0xfd, 0xff, 0x00, 0x70, 0xae, 0xa1, 0xdf, 0xd2, 0x72, 0x4b, 0x99, 0xc1, 0xb7, 0x2a, 0xd4, 0xf5,
	0x97, 0x76, 0x06, 0xeb, 0x23, 0x22, 0x3c, 0xad, 0x78, 0x95, 0x7a, 0xd9, 0xb6, 0x59, 0x8b, 0xb7,
	0xd3, 0x32, 0xe7, 0xd6, 0xe2, 0x30, 0xb5, 0x57, 0xd0, 0x21, 0xfc, 0x73, 0x43, 0x8e, 0x06, 0x68,
	0xeb, 0x26, 0xbc, 0x18, 0xb2, 0xb5, 0x71, 0x4b, 0x27, 0xd7, 0x78, 0x01, 0x6d, 0x9e, 0x6b, 0xa0,
	0xcd, 0xdf, 0x40, 0x7f, 0x26, 0xbf, 0x84, 0xce, 0x88, 0xe1, 0xe9, 0xb2, 0xf4, 0x7d, 0xe8, 0x32,
	0x32, 0x8b, 0xc6, 0x64, 0x49, 0xfe, 0x2b, 0x68, 0x17, 0x36, 0x2f, 0x24, 0x97, 0x0c, 0xbc, 0x39,
	0x69, 0x7b, 0xe5, 0xf0, 0x2d, 0xdc, 0x8b, 0x69, 0xfc, 0x88, 0x13, 0x36, 0xa3, 0x3e, 0xe1, 0x8e,
	0x5c, 0x40, 0xae, 0xf1, 0xce, 0x8c, 0x30, 0x4e, 0xa3, 0x69, 0xff, 0x50, 0x2f, 0xf7, 0x50, 0xfe,
	0xf7, 0x2e, 0xfb, 0xc3, 0xca, 0xc7, 0x86, 0xfe, 0x25, 0x7e, 0x37, 0xba, 0xee, 0xf9, 0xe9, 0xf1,
	0x7b, 0xef, 0xf5, 0xd9, 0x91, 0x77, 0xd9, 0xbf, 0xd2, 0xe9, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0xcc, 0xff, 0xeb, 0x32, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RolesClient interface {
	GetRolesByFilter(ctx context.Context, in *RolesPageRequest, opts ...grpc.CallOption) (*RolesPageReply, error)
	GetRolesById(ctx context.Context, in *RoleIdRequest, opts ...grpc.CallOption) (*RolesReply, error)
	SetRoles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error)
	GrantRoles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error)
	RevokeRoles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error)
	Authorize(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*AuthorizeReply, error)
}

type rolesClient struct {
	cc *grpc.ClientConn
}

func NewRolesClient(cc *grpc.ClientConn) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) GetRolesByFilter(ctx context.Context, in *RolesPageRequest, opts ...grpc.CallOption) (*RolesPageReply, error) {
	out := new(RolesPageReply)
	err := c.cc.Invoke(ctx, "/roles_v1.Roles/get_roles_by_filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetRolesById(ctx context.Context, in *RoleIdRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := c.cc.Invoke(ctx, "/roles_v1.Roles/get_roles_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) SetRoles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := c.cc.Invoke(ctx, "/roles_v1.Roles/set_roles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GrantRoles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := c.cc.Invoke(ctx, "/roles_v1.Roles/grant_roles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RevokeRoles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := c.cc.Invoke(ctx, "/roles_v1.Roles/revoke_roles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Authorize(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*AuthorizeReply, error) {
	out := new(AuthorizeReply)
	err := c.cc.Invoke(ctx, "/roles_v1.Roles/authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
type RolesServer interface {
	GetRolesByFilter(context.Context, *RolesPageRequest) (*RolesPageReply, error)
	GetRolesById(context.Context, *RoleIdRequest) (*RolesReply, error)
	SetRoles(context.Context, *RolesRequest) (*RolesReply, error)
	GrantRoles(context.Context, *RolesRequest) (*RolesReply, error)
	RevokeRoles(context.Context, *RolesRequest) (*RolesReply, error)
	Authorize(context.Context, *RolesRequest) (*AuthorizeReply, error)
}

// UnimplementedRolesServer can be embedded to have forward compatible implementations.
type UnimplementedRolesServer struct {
}

func (*UnimplementedRolesServer) GetRolesByFilter(ctx context.Context, req *RolesPageRequest) (*RolesPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesByFilter not implemented")
}
func (*UnimplementedRolesServer) GetRolesById(ctx context.Context, req *RoleIdRequest) (*RolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesById not implemented")
}
func (*UnimplementedRolesServer) SetRoles(ctx context.Context, req *RolesRequest) (*RolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoles not implemented")
}
func (*UnimplementedRolesServer) GrantRoles(ctx context.Context, req *RolesRequest) (*RolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantRoles not implemented")
}
func (*UnimplementedRolesServer) RevokeRoles(ctx context.Context, req *RolesRequest) (*RolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRoles not implemented")
}
func (*UnimplementedRolesServer) Authorize(ctx context.Context, req *RolesRequest) (*AuthorizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}

func RegisterRolesServer(s *grpc.Server, srv RolesServer) {
	s.RegisterService(&_Roles_serviceDesc, srv)
}

func _Roles_GetRolesByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetRolesByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles_v1.Roles/GetRolesByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetRolesByFilter(ctx, req.(*RolesPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GetRolesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetRolesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles_v1.Roles/GetRolesById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetRolesById(ctx, req.(*RoleIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_SetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).SetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles_v1.Roles/SetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).SetRoles(ctx, req.(*RolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GrantRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GrantRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles_v1.Roles/GrantRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GrantRoles(ctx, req.(*RolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_RevokeRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).RevokeRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles_v1.Roles/RevokeRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).RevokeRoles(ctx, req.(*RolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles_v1.Roles/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Authorize(ctx, req.(*RolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roles_v1.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_roles_by_filter",
			Handler:    _Roles_GetRolesByFilter_Handler,
		},
		{
			MethodName: "get_roles_by_id",
			Handler:    _Roles_GetRolesById_Handler,
		},
		{
			MethodName: "set_roles",
			Handler:    _Roles_SetRoles_Handler,
		},
		{
			MethodName: "grant_roles",
			Handler:    _Roles_GrantRoles_Handler,
		},
		{
			MethodName: "revoke_roles",
			Handler:    _Roles_RevokeRoles_Handler,
		},
		{
			MethodName: "authorize",
			Handler:    _Roles_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/roles_v1.proto",
}
