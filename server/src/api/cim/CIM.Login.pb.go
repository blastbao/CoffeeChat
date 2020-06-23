// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.10.0
// source: CIM.Login.proto

package cim

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 认证请求
type CIMAuthTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0101
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// CoffeeChat不存储用户信息，消息推送时需要显示昵称
	// 基于流量考虑，昵称不放在每条消息中携带
	// 但是如果期间用户更新昵称后，消息推送显示昵称会有延迟，CoffeeChat认为是能接受的
	NickName   string        `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	UserToken  string        `protobuf:"bytes,3,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	ClientType CIMClientType `protobuf:"varint,4,opt,name=client_type,json=clientType,proto3,enum=CIM.Def.CIMClientType" json:"client_type,omitempty"`
	//optional
	ClientVersion string `protobuf:"bytes,5,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
}

func (x *CIMAuthTokenReq) Reset() {
	*x = CIMAuthTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMAuthTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMAuthTokenReq) ProtoMessage() {}

func (x *CIMAuthTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMAuthTokenReq.ProtoReflect.Descriptor instead.
func (*CIMAuthTokenReq) Descriptor() ([]byte, []int) {
	return file_CIM_Login_proto_rawDescGZIP(), []int{0}
}

func (x *CIMAuthTokenReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMAuthTokenReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CIMAuthTokenReq) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

func (x *CIMAuthTokenReq) GetClientType() CIMClientType {
	if x != nil {
		return x.ClientType
	}
	return CIMClientType_kCIM_CLIENT_TYPE_DEFAULT
}

func (x *CIMAuthTokenReq) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

type CIMAuthTokenRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0102
	ServerTime uint32       `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`                           // 服务器时间
	ResultCode CIMErrorCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=CIM.Def.CIMErrorCode" json:"result_code,omitempty"` // 验证结果
	//optional
	ResultString string `protobuf:"bytes,3,opt,name=result_string,json=resultString,proto3" json:"result_string,omitempty"` // 结果描述
	//optional
	UserInfo *CIMUserInfo `protobuf:"bytes,4,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"` // 用户信息
}

func (x *CIMAuthTokenRsp) Reset() {
	*x = CIMAuthTokenRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMAuthTokenRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMAuthTokenRsp) ProtoMessage() {}

func (x *CIMAuthTokenRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMAuthTokenRsp.ProtoReflect.Descriptor instead.
func (*CIMAuthTokenRsp) Descriptor() ([]byte, []int) {
	return file_CIM_Login_proto_rawDescGZIP(), []int{1}
}

func (x *CIMAuthTokenRsp) GetServerTime() uint32 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *CIMAuthTokenRsp) GetResultCode() CIMErrorCode {
	if x != nil {
		return x.ResultCode
	}
	return CIMErrorCode_kCIM_ERR_UNKNOWN
}

func (x *CIMAuthTokenRsp) GetResultString() string {
	if x != nil {
		return x.ResultString
	}
	return ""
}

func (x *CIMAuthTokenRsp) GetUserInfo() *CIMUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

// 登出
type CIMLogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0103
	UserId     uint64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientType CIMClientType `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=CIM.Def.CIMClientType" json:"client_type,omitempty"`
}

func (x *CIMLogoutReq) Reset() {
	*x = CIMLogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMLogoutReq) ProtoMessage() {}

func (x *CIMLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMLogoutReq.ProtoReflect.Descriptor instead.
func (*CIMLogoutReq) Descriptor() ([]byte, []int) {
	return file_CIM_Login_proto_rawDescGZIP(), []int{2}
}

func (x *CIMLogoutReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMLogoutReq) GetClientType() CIMClientType {
	if x != nil {
		return x.ClientType
	}
	return CIMClientType_kCIM_CLIENT_TYPE_DEFAULT
}

type CIMLogoutRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0104
	ResultCode uint32 `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
}

func (x *CIMLogoutRsp) Reset() {
	*x = CIMLogoutRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMLogoutRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMLogoutRsp) ProtoMessage() {}

func (x *CIMLogoutRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMLogoutRsp.ProtoReflect.Descriptor instead.
func (*CIMLogoutRsp) Descriptor() ([]byte, []int) {
	return file_CIM_Login_proto_rawDescGZIP(), []int{3}
}

func (x *CIMLogoutRsp) GetResultCode() uint32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

// 心跳
type CIMHeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CIMHeartBeat) Reset() {
	*x = CIMHeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMHeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMHeartBeat) ProtoMessage() {}

func (x *CIMHeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMHeartBeat.ProtoReflect.Descriptor instead.
func (*CIMHeartBeat) Descriptor() ([]byte, []int) {
	return file_CIM_Login_proto_rawDescGZIP(), []int{4}
}

var File_CIM_Login_proto protoreflect.FileDescriptor

var file_CIM_Login_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x49, 0x4d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x43, 0x49, 0x4d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x0d, 0x43, 0x49,
	0x4d, 0x2e, 0x44, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x0f,
	0x43, 0x49, 0x4d, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x43, 0x49, 0x4d, 0x2e,
	0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x0f, 0x43, 0x49, 0x4d, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x43, 0x49, 0x4d, 0x2e, 0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x49, 0x4d, 0x2e,
	0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x60, 0x0a, 0x0c, 0x43, 0x49, 0x4d,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x43, 0x49, 0x4d, 0x2e, 0x44, 0x65,
	0x66, 0x2e, 0x43, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2f, 0x0a, 0x0c, 0x43,
	0x49, 0x4d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x0e, 0x0a, 0x0c,
	0x43, 0x49, 0x4d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x42, 0x22, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x48, 0x03, 0x5a, 0x05, 0x2e, 0x3b, 0x63, 0x69, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CIM_Login_proto_rawDescOnce sync.Once
	file_CIM_Login_proto_rawDescData = file_CIM_Login_proto_rawDesc
)

func file_CIM_Login_proto_rawDescGZIP() []byte {
	file_CIM_Login_proto_rawDescOnce.Do(func() {
		file_CIM_Login_proto_rawDescData = protoimpl.X.CompressGZIP(file_CIM_Login_proto_rawDescData)
	})
	return file_CIM_Login_proto_rawDescData
}

var file_CIM_Login_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_CIM_Login_proto_goTypes = []interface{}{
	(*CIMAuthTokenReq)(nil), // 0: CIM.Login.CIMAuthTokenReq
	(*CIMAuthTokenRsp)(nil), // 1: CIM.Login.CIMAuthTokenRsp
	(*CIMLogoutReq)(nil),    // 2: CIM.Login.CIMLogoutReq
	(*CIMLogoutRsp)(nil),    // 3: CIM.Login.CIMLogoutRsp
	(*CIMHeartBeat)(nil),    // 4: CIM.Login.CIMHeartBeat
	(CIMClientType)(0),      // 5: CIM.Def.CIMClientType
	(CIMErrorCode)(0),       // 6: CIM.Def.CIMErrorCode
	(*CIMUserInfo)(nil),     // 7: CIM.Def.CIMUserInfo
}
var file_CIM_Login_proto_depIdxs = []int32{
	5, // 0: CIM.Login.CIMAuthTokenReq.client_type:type_name -> CIM.Def.CIMClientType
	6, // 1: CIM.Login.CIMAuthTokenRsp.result_code:type_name -> CIM.Def.CIMErrorCode
	7, // 2: CIM.Login.CIMAuthTokenRsp.user_info:type_name -> CIM.Def.CIMUserInfo
	5, // 3: CIM.Login.CIMLogoutReq.client_type:type_name -> CIM.Def.CIMClientType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_CIM_Login_proto_init() }
func file_CIM_Login_proto_init() {
	if File_CIM_Login_proto != nil {
		return
	}
	file_CIM_Def_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CIM_Login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMAuthTokenReq); i {
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
		file_CIM_Login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMAuthTokenRsp); i {
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
		file_CIM_Login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMLogoutReq); i {
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
		file_CIM_Login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMLogoutRsp); i {
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
		file_CIM_Login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMHeartBeat); i {
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
			RawDescriptor: file_CIM_Login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CIM_Login_proto_goTypes,
		DependencyIndexes: file_CIM_Login_proto_depIdxs,
		MessageInfos:      file_CIM_Login_proto_msgTypes,
	}.Build()
	File_CIM_Login_proto = out.File
	file_CIM_Login_proto_rawDesc = nil
	file_CIM_Login_proto_goTypes = nil
	file_CIM_Login_proto_depIdxs = nil
}
