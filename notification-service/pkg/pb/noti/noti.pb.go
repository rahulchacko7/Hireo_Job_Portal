// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: pkg/pb/noti/noti.proto

package noti

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

type GetAllNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetAllNotificationsRequest) Reset() {
	*x = GetAllNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotificationsRequest) ProtoMessage() {}

func (x *GetAllNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetAllNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllNotificationsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification []*AllMessage `protobuf:"bytes,1,rep,name=notification,proto3" json:"notification,omitempty"`
}

func (x *GetAllNotificationsResponse) Reset() {
	*x = GetAllNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotificationsResponse) ProtoMessage() {}

func (x *GetAllNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetAllNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllNotificationsResponse) GetNotification() []*AllMessage {
	if x != nil {
		return x.Notification
	}
	return nil
}

type MarkAllAsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MarkAllAsReadRequest) Reset() {
	*x = MarkAllAsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAllAsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAllAsReadRequest) ProtoMessage() {}

func (x *MarkAllAsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAllAsReadRequest.ProtoReflect.Descriptor instead.
func (*MarkAllAsReadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{2}
}

func (x *MarkAllAsReadRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type MarkAllAsReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MarkAllAsReadResponse) Reset() {
	*x = MarkAllAsReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAllAsReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAllAsReadResponse) ProtoMessage() {}

func (x *MarkAllAsReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAllAsReadResponse.ProtoReflect.Descriptor instead.
func (*MarkAllAsReadResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{3}
}

func (x *MarkAllAsReadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadNotificationRequest) Reset() {
	*x = ReadNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadNotificationRequest) ProtoMessage() {}

func (x *ReadNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadNotificationRequest.ProtoReflect.Descriptor instead.
func (*ReadNotificationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{4}
}

func (x *ReadNotificationRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReadNotificationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ReadNotificationResponse) Reset() {
	*x = ReadNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadNotificationResponse) ProtoMessage() {}

func (x *ReadNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadNotificationResponse.ProtoReflect.Descriptor instead.
func (*ReadNotificationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{5}
}

func (x *ReadNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetNotificationRequest) Reset() {
	*x = GetNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationRequest) ProtoMessage() {}

func (x *GetNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotificationRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetNotificationRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetNotificationRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Id       int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Time     string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	PostId   int64  `protobuf:"varint,6,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{7}
}

func (x *Message) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Message) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Message) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type AllMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Id       int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Time     string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	PostId   int64  `protobuf:"varint,6,opt,name=postId,proto3" json:"postId,omitempty"`
	Read     bool   `protobuf:"varint,7,opt,name=read,proto3" json:"read,omitempty"`
}

func (x *AllMessage) Reset() {
	*x = AllMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllMessage) ProtoMessage() {}

func (x *AllMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllMessage.ProtoReflect.Descriptor instead.
func (*AllMessage) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{8}
}

func (x *AllMessage) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AllMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AllMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AllMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AllMessage) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *AllMessage) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *AllMessage) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

type GetNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification []*Message `protobuf:"bytes,1,rep,name=notification,proto3" json:"notification,omitempty"`
}

func (x *GetNotificationResponse) Reset() {
	*x = GetNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_noti_noti_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationResponse) ProtoMessage() {}

func (x *GetNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_noti_noti_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_noti_noti_proto_rawDescGZIP(), []int{9}
}

func (x *GetNotificationResponse) GetNotification() []*Message {
	if x != nil {
		return x.Notification
	}
	return nil
}

var File_pkg_pb_noti_noti_proto protoreflect.FileDescriptor

var file_pkg_pb_noti_noti_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x14, 0x4d, 0x61, 0x72,
	0x6b, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x4d, 0x61, 0x72,
	0x6b, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x17,
	0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x34, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a,
	0x0a, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x22, 0x54, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xa4, 0x03, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x10, 0x52, 0x65, 0x61,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_noti_noti_proto_rawDescOnce sync.Once
	file_pkg_pb_noti_noti_proto_rawDescData = file_pkg_pb_noti_noti_proto_rawDesc
)

func file_pkg_pb_noti_noti_proto_rawDescGZIP() []byte {
	file_pkg_pb_noti_noti_proto_rawDescOnce.Do(func() {
		file_pkg_pb_noti_noti_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_noti_noti_proto_rawDescData)
	})
	return file_pkg_pb_noti_noti_proto_rawDescData
}

var file_pkg_pb_noti_noti_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_pb_noti_noti_proto_goTypes = []interface{}{
	(*GetAllNotificationsRequest)(nil),  // 0: notification.GetAllNotificationsRequest
	(*GetAllNotificationsResponse)(nil), // 1: notification.GetAllNotificationsResponse
	(*MarkAllAsReadRequest)(nil),        // 2: notification.MarkAllAsReadRequest
	(*MarkAllAsReadResponse)(nil),       // 3: notification.MarkAllAsReadResponse
	(*ReadNotificationRequest)(nil),     // 4: notification.ReadNotificationRequest
	(*ReadNotificationResponse)(nil),    // 5: notification.ReadNotificationResponse
	(*GetNotificationRequest)(nil),      // 6: notification.GetNotificationRequest
	(*Message)(nil),                     // 7: notification.Message
	(*AllMessage)(nil),                  // 8: notification.AllMessage
	(*GetNotificationResponse)(nil),     // 9: notification.GetNotificationResponse
}
var file_pkg_pb_noti_noti_proto_depIdxs = []int32{
	8, // 0: notification.GetAllNotificationsResponse.notification:type_name -> notification.AllMessage
	7, // 1: notification.GetNotificationResponse.notification:type_name -> notification.Message
	6, // 2: notification.NotificationService.GetNotification:input_type -> notification.GetNotificationRequest
	4, // 3: notification.NotificationService.ReadNotification:input_type -> notification.ReadNotificationRequest
	2, // 4: notification.NotificationService.MarkAllAsRead:input_type -> notification.MarkAllAsReadRequest
	0, // 5: notification.NotificationService.GetAllNotifications:input_type -> notification.GetAllNotificationsRequest
	9, // 6: notification.NotificationService.GetNotification:output_type -> notification.GetNotificationResponse
	5, // 7: notification.NotificationService.ReadNotification:output_type -> notification.ReadNotificationResponse
	3, // 8: notification.NotificationService.MarkAllAsRead:output_type -> notification.MarkAllAsReadResponse
	1, // 9: notification.NotificationService.GetAllNotifications:output_type -> notification.GetAllNotificationsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_pb_noti_noti_proto_init() }
func file_pkg_pb_noti_noti_proto_init() {
	if File_pkg_pb_noti_noti_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_noti_noti_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNotificationsRequest); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNotificationsResponse); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAllAsReadRequest); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAllAsReadResponse); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadNotificationRequest); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadNotificationResponse); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationRequest); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllMessage); i {
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
		file_pkg_pb_noti_noti_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationResponse); i {
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
			RawDescriptor: file_pkg_pb_noti_noti_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_noti_noti_proto_goTypes,
		DependencyIndexes: file_pkg_pb_noti_noti_proto_depIdxs,
		MessageInfos:      file_pkg_pb_noti_noti_proto_msgTypes,
	}.Build()
	File_pkg_pb_noti_noti_proto = out.File
	file_pkg_pb_noti_noti_proto_rawDesc = nil
	file_pkg_pb_noti_noti_proto_goTypes = nil
	file_pkg_pb_noti_noti_proto_depIdxs = nil
}
