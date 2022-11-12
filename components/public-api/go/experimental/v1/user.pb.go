// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/experimental/v1/user.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a UUID of the user
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the username
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// avatar_url is a link to the user avatar
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// created_at is the creation time
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SSHKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a UUID of the SSH key
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the name of the SSH key
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// key is the public SSH key
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// created_at is the creation time
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SSHKey) Reset() {
	*x = SSHKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHKey) ProtoMessage() {}

func (x *SSHKey) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHKey.ProtoReflect.Descriptor instead.
func (*SSHKey) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *SSHKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SSHKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSHKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SSHKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetAuthenticatedUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthenticatedUserRequest) Reset() {
	*x = GetAuthenticatedUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthenticatedUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthenticatedUserRequest) ProtoMessage() {}

func (x *GetAuthenticatedUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthenticatedUserRequest.ProtoReflect.Descriptor instead.
func (*GetAuthenticatedUserRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{2}
}

type GetAuthenticatedUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetAuthenticatedUserResponse) Reset() {
	*x = GetAuthenticatedUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthenticatedUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthenticatedUserResponse) ProtoMessage() {}

func (x *GetAuthenticatedUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthenticatedUserResponse.ProtoReflect.Descriptor instead.
func (*GetAuthenticatedUserResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthenticatedUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ListSSHKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSSHKeysRequest) Reset() {
	*x = ListSSHKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSSHKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSSHKeysRequest) ProtoMessage() {}

func (x *ListSSHKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSSHKeysRequest.ProtoReflect.Descriptor instead.
func (*ListSSHKeysRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{4}
}

type ListSSHKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SSHKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ListSSHKeysResponse) Reset() {
	*x = ListSSHKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSSHKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSSHKeysResponse) ProtoMessage() {}

func (x *ListSSHKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSSHKeysResponse.ProtoReflect.Descriptor instead.
func (*ListSSHKeysResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *ListSSHKeysResponse) GetKeys() []*SSHKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type CreateSSHKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the SSH key name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the public SSH key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateSSHKeyRequest) Reset() {
	*x = CreateSSHKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSSHKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSSHKeyRequest) ProtoMessage() {}

func (x *CreateSSHKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSSHKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateSSHKeyRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *CreateSSHKeyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSSHKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CreateSSHKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *SSHKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateSSHKeyResponse) Reset() {
	*x = CreateSSHKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSSHKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSSHKeyResponse) ProtoMessage() {}

func (x *CreateSSHKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSSHKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateSSHKeyResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *CreateSSHKeyResponse) GetKey() *SSHKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetSSHKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique identifier of the SSH key to retreive.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *GetSSHKeyRequest) Reset() {
	*x = GetSSHKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSSHKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSSHKeyRequest) ProtoMessage() {}

func (x *GetSSHKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSSHKeyRequest.ProtoReflect.Descriptor instead.
func (*GetSSHKeyRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetSSHKeyRequest) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type GetSSHKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *SSHKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetSSHKeyResponse) Reset() {
	*x = GetSSHKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSSHKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSSHKeyResponse) ProtoMessage() {}

func (x *GetSSHKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSSHKeyResponse.ProtoReflect.Descriptor instead.
func (*GetSSHKeyResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetSSHKeyResponse) GetKey() *SSHKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type DeleteSSHKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique identifier of the SSH key to retreive.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *DeleteSSHKeyRequest) Reset() {
	*x = DeleteSSHKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSSHKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSSHKeyRequest) ProtoMessage() {}

func (x *DeleteSSHKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSSHKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteSSHKeyRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteSSHKeyRequest) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type DeleteSSHKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSSHKeyResponse) Reset() {
	*x = DeleteSSHKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSSHKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSSHKeyResponse) ProtoMessage() {}

func (x *DeleteSSHKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSSHKeyResponse.ProtoReflect.Descriptor instead.
func (*DeleteSSHKeyResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_user_proto_rawDescGZIP(), []int{11}
}

var File_gitpod_experimental_v1_user_proto protoreflect.FileDescriptor

var file_gitpod_experimental_v1_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x79, 0x0a, 0x06, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1d,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x69,
	0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53, 0x48,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74,
	0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x22, 0x3b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x48, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x53, 0x48, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x29, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x53,
	0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79,
	0x49, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x53,
	0x48, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xbb, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x83, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53, 0x48,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53,
	0x48, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12,
	0x2b, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79,
	0x12, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x53, 0x48,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x46, 0x5a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_experimental_v1_user_proto_rawDescOnce sync.Once
	file_gitpod_experimental_v1_user_proto_rawDescData = file_gitpod_experimental_v1_user_proto_rawDesc
)

func file_gitpod_experimental_v1_user_proto_rawDescGZIP() []byte {
	file_gitpod_experimental_v1_user_proto_rawDescOnce.Do(func() {
		file_gitpod_experimental_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_experimental_v1_user_proto_rawDescData)
	})
	return file_gitpod_experimental_v1_user_proto_rawDescData
}

var file_gitpod_experimental_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_gitpod_experimental_v1_user_proto_goTypes = []interface{}{
	(*User)(nil),                         // 0: gitpod.experimental.v1.User
	(*SSHKey)(nil),                       // 1: gitpod.experimental.v1.SSHKey
	(*GetAuthenticatedUserRequest)(nil),  // 2: gitpod.experimental.v1.GetAuthenticatedUserRequest
	(*GetAuthenticatedUserResponse)(nil), // 3: gitpod.experimental.v1.GetAuthenticatedUserResponse
	(*ListSSHKeysRequest)(nil),           // 4: gitpod.experimental.v1.ListSSHKeysRequest
	(*ListSSHKeysResponse)(nil),          // 5: gitpod.experimental.v1.ListSSHKeysResponse
	(*CreateSSHKeyRequest)(nil),          // 6: gitpod.experimental.v1.CreateSSHKeyRequest
	(*CreateSSHKeyResponse)(nil),         // 7: gitpod.experimental.v1.CreateSSHKeyResponse
	(*GetSSHKeyRequest)(nil),             // 8: gitpod.experimental.v1.GetSSHKeyRequest
	(*GetSSHKeyResponse)(nil),            // 9: gitpod.experimental.v1.GetSSHKeyResponse
	(*DeleteSSHKeyRequest)(nil),          // 10: gitpod.experimental.v1.DeleteSSHKeyRequest
	(*DeleteSSHKeyResponse)(nil),         // 11: gitpod.experimental.v1.DeleteSSHKeyResponse
	(*timestamppb.Timestamp)(nil),        // 12: google.protobuf.Timestamp
}
var file_gitpod_experimental_v1_user_proto_depIdxs = []int32{
	12, // 0: gitpod.experimental.v1.User.created_at:type_name -> google.protobuf.Timestamp
	12, // 1: gitpod.experimental.v1.SSHKey.created_at:type_name -> google.protobuf.Timestamp
	0,  // 2: gitpod.experimental.v1.GetAuthenticatedUserResponse.user:type_name -> gitpod.experimental.v1.User
	1,  // 3: gitpod.experimental.v1.ListSSHKeysResponse.keys:type_name -> gitpod.experimental.v1.SSHKey
	1,  // 4: gitpod.experimental.v1.CreateSSHKeyResponse.key:type_name -> gitpod.experimental.v1.SSHKey
	1,  // 5: gitpod.experimental.v1.GetSSHKeyResponse.key:type_name -> gitpod.experimental.v1.SSHKey
	2,  // 6: gitpod.experimental.v1.UserService.GetAuthenticatedUser:input_type -> gitpod.experimental.v1.GetAuthenticatedUserRequest
	4,  // 7: gitpod.experimental.v1.UserService.ListSSHKeys:input_type -> gitpod.experimental.v1.ListSSHKeysRequest
	6,  // 8: gitpod.experimental.v1.UserService.CreateSSHKey:input_type -> gitpod.experimental.v1.CreateSSHKeyRequest
	8,  // 9: gitpod.experimental.v1.UserService.GetSSHKey:input_type -> gitpod.experimental.v1.GetSSHKeyRequest
	10, // 10: gitpod.experimental.v1.UserService.DeleteSSHKey:input_type -> gitpod.experimental.v1.DeleteSSHKeyRequest
	3,  // 11: gitpod.experimental.v1.UserService.GetAuthenticatedUser:output_type -> gitpod.experimental.v1.GetAuthenticatedUserResponse
	5,  // 12: gitpod.experimental.v1.UserService.ListSSHKeys:output_type -> gitpod.experimental.v1.ListSSHKeysResponse
	7,  // 13: gitpod.experimental.v1.UserService.CreateSSHKey:output_type -> gitpod.experimental.v1.CreateSSHKeyResponse
	9,  // 14: gitpod.experimental.v1.UserService.GetSSHKey:output_type -> gitpod.experimental.v1.GetSSHKeyResponse
	11, // 15: gitpod.experimental.v1.UserService.DeleteSSHKey:output_type -> gitpod.experimental.v1.DeleteSSHKeyResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_gitpod_experimental_v1_user_proto_init() }
func file_gitpod_experimental_v1_user_proto_init() {
	if File_gitpod_experimental_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_experimental_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHKey); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthenticatedUserRequest); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthenticatedUserResponse); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSSHKeysRequest); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSSHKeysResponse); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSSHKeyRequest); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSSHKeyResponse); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSSHKeyRequest); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSSHKeyResponse); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSSHKeyRequest); i {
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
		file_gitpod_experimental_v1_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSSHKeyResponse); i {
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
			RawDescriptor: file_gitpod_experimental_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_experimental_v1_user_proto_goTypes,
		DependencyIndexes: file_gitpod_experimental_v1_user_proto_depIdxs,
		MessageInfos:      file_gitpod_experimental_v1_user_proto_msgTypes,
	}.Build()
	File_gitpod_experimental_v1_user_proto = out.File
	file_gitpod_experimental_v1_user_proto_rawDesc = nil
	file_gitpod_experimental_v1_user_proto_goTypes = nil
	file_gitpod_experimental_v1_user_proto_depIdxs = nil
}