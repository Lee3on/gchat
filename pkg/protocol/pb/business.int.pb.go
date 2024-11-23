// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: pkg/protocol/proto/business.int.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceId int64  `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_proto_business_int_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_proto_business_int_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_proto_business_int_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthReq) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *AuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds map[int64]int32 `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetUsersReq) Reset() {
	*x = GetUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_proto_business_int_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersReq) ProtoMessage() {}

func (x *GetUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_proto_business_int_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersReq.ProtoReflect.Descriptor instead.
func (*GetUsersReq) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_proto_business_int_proto_rawDescGZIP(), []int{1}
}

func (x *GetUsersReq) GetUserIds() map[int64]int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users map[int64]*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetUsersResp) Reset() {
	*x = GetUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_proto_business_int_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResp) ProtoMessage() {}

func (x *GetUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_proto_business_int_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResp.ProtoReflect.Descriptor instead.
func (*GetUsersResp) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_proto_business_int_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersResp) GetUsers() map[int64]*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_pkg_protocol_proto_business_int_proto protoreflect.FileDescriptor

var file_pkg_protocol_proto_business_int_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x55, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x1a,
	0x3a, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x85, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a,
	0x42, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protocol_proto_business_int_proto_rawDescOnce sync.Once
	file_pkg_protocol_proto_business_int_proto_rawDescData = file_pkg_protocol_proto_business_int_proto_rawDesc
)

func file_pkg_protocol_proto_business_int_proto_rawDescGZIP() []byte {
	file_pkg_protocol_proto_business_int_proto_rawDescOnce.Do(func() {
		file_pkg_protocol_proto_business_int_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protocol_proto_business_int_proto_rawDescData)
	})
	return file_pkg_protocol_proto_business_int_proto_rawDescData
}

var file_pkg_protocol_proto_business_int_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_protocol_proto_business_int_proto_goTypes = []interface{}{
	(*AuthReq)(nil),       // 0: pb.AuthReq
	(*GetUsersReq)(nil),   // 1: pb.GetUsersReq
	(*GetUsersResp)(nil),  // 2: pb.GetUsersResp
	nil,                   // 3: pb.GetUsersReq.UserIdsEntry
	nil,                   // 4: pb.GetUsersResp.UsersEntry
	(*User)(nil),          // 5: pb.User
	(*GetUserReq)(nil),    // 6: pb.GetUserReq
	(*emptypb.Empty)(nil), // 7: google.protobuf.Empty
	(*GetUserResp)(nil),   // 8: pb.GetUserResp
}
var file_pkg_protocol_proto_business_int_proto_depIdxs = []int32{
	3, // 0: pb.GetUsersReq.user_ids:type_name -> pb.GetUsersReq.UserIdsEntry
	4, // 1: pb.GetUsersResp.users:type_name -> pb.GetUsersResp.UsersEntry
	5, // 2: pb.GetUsersResp.UsersEntry.value:type_name -> pb.User
	0, // 3: pb.BusinessInt.Auth:input_type -> pb.AuthReq
	6, // 4: pb.BusinessInt.GetUser:input_type -> pb.GetUserReq
	1, // 5: pb.BusinessInt.GetUsers:input_type -> pb.GetUsersReq
	7, // 6: pb.BusinessInt.Auth:output_type -> google.protobuf.Empty
	8, // 7: pb.BusinessInt.GetUser:output_type -> pb.GetUserResp
	2, // 8: pb.BusinessInt.GetUsers:output_type -> pb.GetUsersResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_protocol_proto_business_int_proto_init() }
func file_pkg_protocol_proto_business_int_proto_init() {
	if File_pkg_protocol_proto_business_int_proto != nil {
		return
	}
	file_pkg_protocol_proto_business_ext_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_protocol_proto_business_int_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_pkg_protocol_proto_business_int_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersReq); i {
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
		file_pkg_protocol_proto_business_int_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResp); i {
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
			RawDescriptor: file_pkg_protocol_proto_business_int_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protocol_proto_business_int_proto_goTypes,
		DependencyIndexes: file_pkg_protocol_proto_business_int_proto_depIdxs,
		MessageInfos:      file_pkg_protocol_proto_business_int_proto_msgTypes,
	}.Build()
	File_pkg_protocol_proto_business_int_proto = out.File
	file_pkg_protocol_proto_business_int_proto_rawDesc = nil
	file_pkg_protocol_proto_business_int_proto_goTypes = nil
	file_pkg_protocol_proto_business_int_proto_depIdxs = nil
}