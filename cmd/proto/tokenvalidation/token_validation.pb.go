// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.12.4
// source: token_validation.proto

package tokenvalidation

import (
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

// The request message containing the token to validate
type TokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	mi := &file_token_validation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_validation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_token_validation_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// The response message
type TokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Message indicating success or failure
	Data          *UserData              `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`       // The user data if token is valid
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	mi := &file_token_validation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_validation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_token_validation_proto_rawDescGZIP(), []int{1}
}

func (x *TokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TokenResponse) GetData() *UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

// The user data returned if the token is valid
type UserData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FullName      string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserData) Reset() {
	*x = UserData{}
	mi := &file_token_validation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_token_validation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_token_validation_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_token_validation_proto protoreflect.FileDescriptor

var file_token_validation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x58, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x72, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x61, 0x0a,
	0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4e, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_validation_proto_rawDescOnce sync.Once
	file_token_validation_proto_rawDescData = file_token_validation_proto_rawDesc
)

func file_token_validation_proto_rawDescGZIP() []byte {
	file_token_validation_proto_rawDescOnce.Do(func() {
		file_token_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_validation_proto_rawDescData)
	})
	return file_token_validation_proto_rawDescData
}

var file_token_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_token_validation_proto_goTypes = []any{
	(*TokenRequest)(nil),  // 0: tokenvalidation.TokenRequest
	(*TokenResponse)(nil), // 1: tokenvalidation.TokenResponse
	(*UserData)(nil),      // 2: tokenvalidation.UserData
}
var file_token_validation_proto_depIdxs = []int32{
	2, // 0: tokenvalidation.TokenResponse.data:type_name -> tokenvalidation.UserData
	0, // 1: tokenvalidation.TokenValidation.ValidateToken:input_type -> tokenvalidation.TokenRequest
	1, // 2: tokenvalidation.TokenValidation.ValidateToken:output_type -> tokenvalidation.TokenResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_token_validation_proto_init() }
func file_token_validation_proto_init() {
	if File_token_validation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_token_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_validation_proto_goTypes,
		DependencyIndexes: file_token_validation_proto_depIdxs,
		MessageInfos:      file_token_validation_proto_msgTypes,
	}.Build()
	File_token_validation_proto = out.File
	file_token_validation_proto_rawDesc = nil
	file_token_validation_proto_goTypes = nil
	file_token_validation_proto_depIdxs = nil
}
