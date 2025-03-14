// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: api.proto

package buf

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

// Socket
type SocketMessageType int32

const (
	SocketMessageType_SET_NAME    SocketMessageType = 0
	SocketMessageType_CREATE_ROOM SocketMessageType = 1
	SocketMessageType_JOIN_ROOM   SocketMessageType = 2
	SocketMessageType_LEAVE_ROOM  SocketMessageType = 3
	SocketMessageType_ERROR       SocketMessageType = 4
)

// Enum value maps for SocketMessageType.
var (
	SocketMessageType_name = map[int32]string{
		0: "SET_NAME",
		1: "CREATE_ROOM",
		2: "JOIN_ROOM",
		3: "LEAVE_ROOM",
		4: "ERROR",
	}
	SocketMessageType_value = map[string]int32{
		"SET_NAME":    0,
		"CREATE_ROOM": 1,
		"JOIN_ROOM":   2,
		"LEAVE_ROOM":  3,
		"ERROR":       4,
	}
)

func (x SocketMessageType) Enum() *SocketMessageType {
	p := new(SocketMessageType)
	*p = x
	return p
}

func (x SocketMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (SocketMessageType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x SocketMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketMessageType.Descriptor instead.
func (SocketMessageType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if x != nil {
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
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Users) Reset() {
	*x = Users{}
	mi := &file_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type WSMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Type  SocketMessageType      `protobuf:"varint,1,opt,name=type,proto3,enum=buf.SocketMessageType" json:"type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//
	//	*WSMessage_SetName
	//	*WSMessage_CreateRoom
	//	*WSMessage_JoinRoom
	//	*WSMessage_Error
	Payload       isWSMessage_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WSMessage) Reset() {
	*x = WSMessage{}
	mi := &file_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSMessage) ProtoMessage() {}

func (x *WSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSMessage.ProtoReflect.Descriptor instead.
func (*WSMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *WSMessage) GetType() SocketMessageType {
	if x != nil {
		return x.Type
	}
	return SocketMessageType_SET_NAME
}

func (x *WSMessage) GetPayload() isWSMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *WSMessage) GetSetName() *SetNameRequest {
	if x != nil {
		if x, ok := x.Payload.(*WSMessage_SetName); ok {
			return x.SetName
		}
	}
	return nil
}

func (x *WSMessage) GetCreateRoom() *CreateRoomRequest {
	if x != nil {
		if x, ok := x.Payload.(*WSMessage_CreateRoom); ok {
			return x.CreateRoom
		}
	}
	return nil
}

func (x *WSMessage) GetJoinRoom() *JoinRoomRequest {
	if x != nil {
		if x, ok := x.Payload.(*WSMessage_JoinRoom); ok {
			return x.JoinRoom
		}
	}
	return nil
}

func (x *WSMessage) GetError() *Error {
	if x != nil {
		if x, ok := x.Payload.(*WSMessage_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isWSMessage_Payload interface {
	isWSMessage_Payload()
}

type WSMessage_SetName struct {
	SetName *SetNameRequest `protobuf:"bytes,2,opt,name=set_name,json=setName,proto3,oneof"`
}

type WSMessage_CreateRoom struct {
	CreateRoom *CreateRoomRequest `protobuf:"bytes,3,opt,name=create_room,json=createRoom,proto3,oneof"`
}

type WSMessage_JoinRoom struct {
	JoinRoom *JoinRoomRequest `protobuf:"bytes,4,opt,name=join_room,json=joinRoom,proto3,oneof"`
}

type WSMessage_Error struct {
	Error *Error `protobuf:"bytes,5,opt,name=error,proto3,oneof"`
}

func (*WSMessage_SetName) isWSMessage_Payload() {}

func (*WSMessage_CreateRoom) isWSMessage_Payload() {}

func (*WSMessage_JoinRoom) isWSMessage_Payload() {}

func (*WSMessage_Error) isWSMessage_Payload() {}

type SetNameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNameRequest) Reset() {
	*x = SetNameRequest{}
	mi := &file_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNameRequest) ProtoMessage() {}

func (x *SetNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNameRequest.ProtoReflect.Descriptor instead.
func (*SetNameRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *SetNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	mi := &file_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *JoinRoomRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Room struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Room          *Room                  `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	mi := &file_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *CreateRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

// Also API
type GetRooms struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rooms         []*Room                `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRooms) Reset() {
	*x = GetRooms{}
	mi := &file_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRooms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRooms) ProtoMessage() {}

func (x *GetRooms) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRooms.ProtoReflect.Descriptor instead.
func (*GetRooms) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetRooms) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x75, 0x66,
	0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x28, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x09, 0x57, 0x53,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x73, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x12, 0x33, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x6f, 0x69,
	0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x2b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x2a, 0x5c, 0x0a, 0x11, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x54,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x4f, 0x49, 0x4e,
	0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x41, 0x56, 0x45,
	0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_goTypes = []any{
	(SocketMessageType)(0),     // 0: buf.SocketMessageType
	(*Error)(nil),              // 1: buf.Error
	(*User)(nil),               // 2: buf.User
	(*Users)(nil),              // 3: buf.Users
	(*WSMessage)(nil),          // 4: buf.WSMessage
	(*SetNameRequest)(nil),     // 5: buf.SetNameRequest
	(*CreateRoomRequest)(nil),  // 6: buf.CreateRoomRequest
	(*JoinRoomRequest)(nil),    // 7: buf.JoinRoomRequest
	(*Room)(nil),               // 8: buf.Room
	(*CreateRoomResponse)(nil), // 9: buf.CreateRoomResponse
	(*GetRooms)(nil),           // 10: buf.GetRooms
}
var file_api_proto_depIdxs = []int32{
	2, // 0: buf.Users.users:type_name -> buf.User
	0, // 1: buf.WSMessage.type:type_name -> buf.SocketMessageType
	5, // 2: buf.WSMessage.set_name:type_name -> buf.SetNameRequest
	6, // 3: buf.WSMessage.create_room:type_name -> buf.CreateRoomRequest
	7, // 4: buf.WSMessage.join_room:type_name -> buf.JoinRoomRequest
	1, // 5: buf.WSMessage.error:type_name -> buf.Error
	8, // 6: buf.CreateRoomResponse.room:type_name -> buf.Room
	8, // 7: buf.GetRooms.rooms:type_name -> buf.Room
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_api_proto_msgTypes[3].OneofWrappers = []any{
		(*WSMessage_SetName)(nil),
		(*WSMessage_CreateRoom)(nil),
		(*WSMessage_JoinRoom)(nil),
		(*WSMessage_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
