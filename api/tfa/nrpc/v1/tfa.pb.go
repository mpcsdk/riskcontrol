// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: tfa/v1/tfa.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/nats-rpc/nrpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// /
type TfaTxReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`         // v: required
	RiskSerial string `protobuf:"bytes,2,opt,name=RiskSerial,proto3" json:"RiskSerial,omitempty"` // v: required
}

func (x *TfaTxReq) Reset() {
	*x = TfaTxReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfaTxReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfaTxReq) ProtoMessage() {}

func (x *TfaTxReq) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfaTxReq.ProtoReflect.Descriptor instead.
func (*TfaTxReq) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{0}
}

func (x *TfaTxReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TfaTxReq) GetRiskSerial() string {
	if x != nil {
		return x.RiskSerial
	}
	return ""
}

type TfaTxRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kinds []string `protobuf:"bytes,1,rep,name=Kinds,proto3" json:"Kinds,omitempty"`
}

func (x *TfaTxRes) Reset() {
	*x = TfaTxRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfaTxRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfaTxRes) ProtoMessage() {}

func (x *TfaTxRes) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfaTxRes.ProtoReflect.Descriptor instead.
func (*TfaTxRes) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{1}
}

func (x *TfaTxRes) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

// //
type TFAReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"` // v: required
	Token  string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`   // v: required
}

func (x *TFAReq) Reset() {
	*x = TFAReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFAReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFAReq) ProtoMessage() {}

func (x *TFAReq) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFAReq.ProtoReflect.Descriptor instead.
func (*TFAReq) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{2}
}

func (x *TFAReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TFAReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TFARes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone       string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	UpPhoneTime string `protobuf:"bytes,2,opt,name=UpPhoneTime,proto3" json:"UpPhoneTime,omitempty"`
	Mail        string `protobuf:"bytes,3,opt,name=Mail,proto3" json:"Mail,omitempty"`
	UpMailTime  string `protobuf:"bytes,4,opt,name=UpMailTime,proto3" json:"UpMailTime,omitempty"`
	UserId      string `protobuf:"bytes,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *TFARes) Reset() {
	*x = TFARes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFARes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFARes) ProtoMessage() {}

func (x *TFARes) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFARes.ProtoReflect.Descriptor instead.
func (*TFARes) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{3}
}

func (x *TFARes) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TFARes) GetUpPhoneTime() string {
	if x != nil {
		return x.UpPhoneTime
	}
	return ""
}

func (x *TFARes) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *TFARes) GetUpMailTime() string {
	if x != nil {
		return x.UpMailTime
	}
	return ""
}

func (x *TFARes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SmsCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiskSerial string `protobuf:"bytes,1,opt,name=RiskSerial,proto3" json:"RiskSerial,omitempty"` // v: required
	Token      string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`           // v: required
}

func (x *SmsCodeReq) Reset() {
	*x = SmsCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsCodeReq) ProtoMessage() {}

func (x *SmsCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsCodeReq.ProtoReflect.Descriptor instead.
func (*SmsCodeReq) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{4}
}

func (x *SmsCodeReq) GetRiskSerial() string {
	if x != nil {
		return x.RiskSerial
	}
	return ""
}

func (x *SmsCodeReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SmsCodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  int32  `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SmsCodeRes) Reset() {
	*x = SmsCodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsCodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsCodeRes) ProtoMessage() {}

func (x *SmsCodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsCodeRes.ProtoReflect.Descriptor instead.
func (*SmsCodeRes) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{5}
}

func (x *SmsCodeRes) GetOk() int32 {
	if x != nil {
		return x.Ok
	}
	return 0
}

func (x *SmsCodeRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type MailCodekReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiskSerial string `protobuf:"bytes,1,opt,name=RiskSerial,proto3" json:"RiskSerial,omitempty"` // v: required
	Token      string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`           // v: required
}

func (x *MailCodekReq) Reset() {
	*x = MailCodekReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailCodekReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailCodekReq) ProtoMessage() {}

func (x *MailCodekReq) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailCodekReq.ProtoReflect.Descriptor instead.
func (*MailCodekReq) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{6}
}

func (x *MailCodekReq) GetRiskSerial() string {
	if x != nil {
		return x.RiskSerial
	}
	return ""
}

func (x *MailCodekReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type MailCodekRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  int32  `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *MailCodekRes) Reset() {
	*x = MailCodekRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailCodekRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailCodekRes) ProtoMessage() {}

func (x *MailCodekRes) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailCodekRes.ProtoReflect.Descriptor instead.
func (*MailCodekRes) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{7}
}

func (x *MailCodekRes) GetOk() int32 {
	if x != nil {
		return x.Ok
	}
	return 0
}

func (x *MailCodekRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// ///
type VerifyCodekReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiskSerial string `protobuf:"bytes,1,opt,name=RiskSerial,proto3" json:"RiskSerial,omitempty"` // v: required
	Token      string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`           // v: required
	PhoneCode  string `protobuf:"bytes,3,opt,name=PhoneCode,proto3" json:"PhoneCode,omitempty"`   // v: required
	MailCode   string `protobuf:"bytes,4,opt,name=MailCode,proto3" json:"MailCode,omitempty"`     // v: required
}

func (x *VerifyCodekReq) Reset() {
	*x = VerifyCodekReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodekReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodekReq) ProtoMessage() {}

func (x *VerifyCodekReq) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodekReq.ProtoReflect.Descriptor instead.
func (*VerifyCodekReq) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{8}
}

func (x *VerifyCodekReq) GetRiskSerial() string {
	if x != nil {
		return x.RiskSerial
	}
	return ""
}

func (x *VerifyCodekReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *VerifyCodekReq) GetPhoneCode() string {
	if x != nil {
		return x.PhoneCode
	}
	return ""
}

func (x *VerifyCodekReq) GetMailCode() string {
	if x != nil {
		return x.MailCode
	}
	return ""
}

type VerifyCodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  int32  `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *VerifyCodeRes) Reset() {
	*x = VerifyCodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tfa_v1_tfa_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRes) ProtoMessage() {}

func (x *VerifyCodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_tfa_v1_tfa_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeRes.ProtoReflect.Descriptor instead.
func (*VerifyCodeRes) Descriptor() ([]byte, []int) {
	return file_tfa_v1_tfa_proto_rawDescGZIP(), []int{9}
}

func (x *VerifyCodeRes) GetOk() int32 {
	if x != nil {
		return x.Ok
	}
	return 0
}

func (x *VerifyCodeRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_tfa_v1_tfa_proto protoreflect.FileDescriptor

var file_tfa_v1_tfa_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x66, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x66, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x74, 0x66, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x42, 0x0a, 0x08, 0x54, 0x66, 0x61, 0x54, 0x78, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x08, 0x54, 0x66, 0x61, 0x54, 0x78, 0x52, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x54, 0x46, 0x41, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8c,
	0x01, 0x0a, 0x06, 0x54, 0x46, 0x41, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x55, 0x70, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x70, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x4d, 0x61, 0x69, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x70, 0x4d, 0x61, 0x69,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a,
	0x0a, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x52,
	0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x2e, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x4f, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x22, 0x44, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x6b, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x30, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x64, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a,
	0x52, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x0d,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x32,
	0xca, 0x02, 0x0a, 0x03, 0x54, 0x46, 0x41, 0x12, 0x3c, 0x0a, 0x08, 0x52, 0x70, 0x63, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x54, 0x66, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x74, 0x66, 0x61, 0x2e, 0x54, 0x46, 0x41, 0x52, 0x65, 0x71,
	0x1a, 0x0b, 0x2e, 0x74, 0x66, 0x61, 0x2e, 0x54, 0x46, 0x41, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x2a, 0x0a, 0x08, 0x52, 0x70, 0x63, 0x54, 0x66, 0x61, 0x54, 0x78, 0x12, 0x0d, 0x2e, 0x74, 0x66,
	0x61, 0x2e, 0x54, 0x66, 0x61, 0x54, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x74, 0x66, 0x61,
	0x2e, 0x54, 0x66, 0x61, 0x54, 0x78, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x52,
	0x70, 0x63, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e,
	0x74, 0x66, 0x61, 0x2e, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x74, 0x66, 0x61, 0x2e, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x0f, 0x52, 0x70, 0x63, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x66, 0x61, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x64, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x74, 0x66, 0x61, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11,
	0x52, 0x70, 0x63, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x13, 0x2e, 0x74, 0x66, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x74, 0x66, 0x61, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c,
	0x2f, 0x74, 0x66, 0x61, 0x2f, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tfa_v1_tfa_proto_rawDescOnce sync.Once
	file_tfa_v1_tfa_proto_rawDescData = file_tfa_v1_tfa_proto_rawDesc
)

func file_tfa_v1_tfa_proto_rawDescGZIP() []byte {
	file_tfa_v1_tfa_proto_rawDescOnce.Do(func() {
		file_tfa_v1_tfa_proto_rawDescData = protoimpl.X.CompressGZIP(file_tfa_v1_tfa_proto_rawDescData)
	})
	return file_tfa_v1_tfa_proto_rawDescData
}

var file_tfa_v1_tfa_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tfa_v1_tfa_proto_goTypes = []interface{}{
	(*TfaTxReq)(nil),       // 0: tfa.TfaTxReq
	(*TfaTxRes)(nil),       // 1: tfa.TfaTxRes
	(*TFAReq)(nil),         // 2: tfa.TFAReq
	(*TFARes)(nil),         // 3: tfa.TFARes
	(*SmsCodeReq)(nil),     // 4: tfa.SmsCodeReq
	(*SmsCodeRes)(nil),     // 5: tfa.SmsCodeRes
	(*MailCodekReq)(nil),   // 6: tfa.MailCodekReq
	(*MailCodekRes)(nil),   // 7: tfa.MailCodekRes
	(*VerifyCodekReq)(nil), // 8: tfa.VerifyCodekReq
	(*VerifyCodeRes)(nil),  // 9: tfa.VerifyCodeRes
	(*empty.Empty)(nil),    // 10: google.protobuf.Empty
}
var file_tfa_v1_tfa_proto_depIdxs = []int32{
	10, // 0: tfa.TFA.RpcAlive:input_type -> google.protobuf.Empty
	2,  // 1: tfa.TFA.RpcTfaInfo:input_type -> tfa.TFAReq
	0,  // 2: tfa.TFA.RpcTfaTx:input_type -> tfa.TfaTxReq
	4,  // 3: tfa.TFA.RpcSendSmsCode:input_type -> tfa.SmsCodeReq
	6,  // 4: tfa.TFA.RpcSendMailCode:input_type -> tfa.MailCodekReq
	8,  // 5: tfa.TFA.RpcSendVerifyCode:input_type -> tfa.VerifyCodekReq
	10, // 6: tfa.TFA.RpcAlive:output_type -> google.protobuf.Empty
	3,  // 7: tfa.TFA.RpcTfaInfo:output_type -> tfa.TFARes
	1,  // 8: tfa.TFA.RpcTfaTx:output_type -> tfa.TfaTxRes
	5,  // 9: tfa.TFA.RpcSendSmsCode:output_type -> tfa.SmsCodeRes
	7,  // 10: tfa.TFA.RpcSendMailCode:output_type -> tfa.MailCodekRes
	9,  // 11: tfa.TFA.RpcSendVerifyCode:output_type -> tfa.VerifyCodeRes
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_tfa_v1_tfa_proto_init() }
func file_tfa_v1_tfa_proto_init() {
	if File_tfa_v1_tfa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tfa_v1_tfa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfaTxReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfaTxRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFAReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFARes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsCodeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsCodeRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailCodekReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailCodekRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodekReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tfa_v1_tfa_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tfa_v1_tfa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tfa_v1_tfa_proto_goTypes,
		DependencyIndexes: file_tfa_v1_tfa_proto_depIdxs,
		MessageInfos:      file_tfa_v1_tfa_proto_msgTypes,
	}.Build()
	File_tfa_v1_tfa_proto = out.File
	file_tfa_v1_tfa_proto_rawDesc = nil
	file_tfa_v1_tfa_proto_goTypes = nil
	file_tfa_v1_tfa_proto_depIdxs = nil
}
