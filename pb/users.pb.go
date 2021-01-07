// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: users.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ActivityType int32

const (
	ActivityType_UNKNOWN ActivityType = 0
	ActivityType_SLEEP   ActivityType = 1
	ActivityType_EAT     ActivityType = 2
	ActivityType_READ    ActivityType = 3
	ActivityType_PLAY    ActivityType = 4
)

// Enum value maps for ActivityType.
var (
	ActivityType_name = map[int32]string{
		0: "UNKNOWN",
		1: "SLEEP",
		2: "EAT",
		3: "READ",
		4: "PLAY",
	}
	ActivityType_value = map[string]int32{
		"UNKNOWN": 0,
		"SLEEP":   1,
		"EAT":     2,
		"READ":    3,
		"PLAY":    4,
	}
)

func (x ActivityType) Enum() *ActivityType {
	p := new(ActivityType)
	*p = x
	return p
}

func (x ActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_users_proto_enumTypes[0].Descriptor()
}

func (ActivityType) Type() protoreflect.EnumType {
	return &file_users_proto_enumTypes[0]
}

func (x ActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityType.Descriptor instead.
func (ActivityType) EnumDescriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_UN_KNOWN Status = 0
	Status_ACTIVE   Status = 1
	Status_DONE     Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UN_KNOWN",
		1: "ACTIVE",
		2: "DONE",
	}
	Status_value = map[string]int32{
		"UN_KNOWN": 0,
		"ACTIVE":   1,
		"DONE":     2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_users_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_users_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNo  string `protobuf:"bytes,3,opt,name=PhoneNo,proto3" json:"PhoneNo,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

type CreateUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *CreateUsersResponse) Reset() {
	*x = CreateUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUsersResponse) ProtoMessage() {}

func (x *CreateUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUsersResponse.ProtoReflect.Descriptor instead.
func (*CreateUsersResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUsersResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day          string               `protobuf:"bytes,6,opt,name=day,proto3" json:"day,omitempty"`
	Activity     ActivityType         `protobuf:"varint,1,opt,name=activity,proto3,enum=gsk967.userActivity.ActivityType" json:"activity,omitempty"`
	TimeDuration uint32               `protobuf:"varint,2,opt,name=time_duration,json=timeDuration,proto3" json:"time_duration,omitempty"`
	Status       Status               `protobuf:"varint,3,opt,name=status,proto3,enum=gsk967.userActivity.Status" json:"status,omitempty"`
	CreatedAt    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *Activity) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Activity) GetActivity() ActivityType {
	if x != nil {
		return x.Activity
	}
	return ActivityType_UNKNOWN
}

func (x *Activity) GetTimeDuration() uint32 {
	if x != nil {
		return x.TimeDuration
	}
	return 0
}

func (x *Activity) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UN_KNOWN
}

func (x *Activity) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Activity) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetUserActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (x *GetUserActivityReq) Reset() {
	*x = GetUserActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserActivityReq) ProtoMessage() {}

func (x *GetUserActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserActivityReq.ProtoReflect.Descriptor instead.
func (*GetUserActivityReq) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserActivityReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type UserActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail       string      `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	DailyActivities []*Activity `protobuf:"bytes,2,rep,name=daily_activities,json=dailyActivities,proto3" json:"daily_activities,omitempty"`
}

func (x *UserActivity) Reset() {
	*x = UserActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserActivity) ProtoMessage() {}

func (x *UserActivity) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserActivity.ProtoReflect.Descriptor instead.
func (*UserActivity) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{4}
}

func (x *UserActivity) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UserActivity) GetDailyActivities() []*Activity {
	if x != nil {
		return x.DailyActivities
	}
	return nil
}

type CreateUserActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail string    `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Activity  *Activity `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *CreateUserActivityReq) Reset() {
	*x = CreateUserActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserActivityReq) ProtoMessage() {}

func (x *CreateUserActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserActivityReq.ProtoReflect.Descriptor instead.
func (*CreateUserActivityReq) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserActivityReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *CreateUserActivityReq) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type CreateUserActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail string    `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Activity  *Activity `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *CreateUserActivityResponse) Reset() {
	*x = CreateUserActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserActivityResponse) ProtoMessage() {}

func (x *CreateUserActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserActivityResponse.ProtoReflect.Descriptor instead.
func (*CreateUserActivityResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserActivityResponse) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *CreateUserActivityResponse) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type UpdateActivityStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail string       `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Day       string       `protobuf:"bytes,2,opt,name=day,proto3" json:"day,omitempty"`
	Activity  ActivityType `protobuf:"varint,3,opt,name=activity,proto3,enum=gsk967.userActivity.ActivityType" json:"activity,omitempty"`
	Status    Status       `protobuf:"varint,4,opt,name=status,proto3,enum=gsk967.userActivity.Status" json:"status,omitempty"`
}

func (x *UpdateActivityStatusReq) Reset() {
	*x = UpdateActivityStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateActivityStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActivityStatusReq) ProtoMessage() {}

func (x *UpdateActivityStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActivityStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateActivityStatusReq) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateActivityStatusReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UpdateActivityStatusReq) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *UpdateActivityStatusReq) GetActivity() ActivityType {
	if x != nil {
		return x.Activity
	}
	return ActivityType_UNKNOWN
}

func (x *UpdateActivityStatusReq) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UN_KNOWN
}

type UpdateActivityStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail string    `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Activity  *Activity `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *UpdateActivityStatusResponse) Reset() {
	*x = UpdateActivityStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateActivityStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActivityStatusResponse) ProtoMessage() {}

func (x *UpdateActivityStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActivityStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateActivityStatusResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateActivityStatusResponse) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UpdateActivityStatusResponse) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67,
	0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x56, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x22, 0x50, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x08,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67,
	0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x33, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x77,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x48, 0x0a,
	0x10, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x39, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x76, 0x0a, 0x1a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x73, 0x6b, 0x39,
	0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x22, 0xbe, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12,
	0x3d, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x33,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x78, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2a, 0x43, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4c,
	0x45, 0x45, 0x50, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x41, 0x54, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4c, 0x41, 0x59,
	0x10, 0x04, 0x2a, 0x2c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08,
	0x55, 0x4e, 0x5f, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02,
	0x32, 0xc7, 0x04, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x2e, 0x67,
	0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x28, 0x2e, 0x67, 0x73,
	0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x27, 0x2e, 0x67, 0x73, 0x6b, 0x39,
	0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x21, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9d, 0x01, 0x0a, 0x19, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x9e, 0x01, 0x0a, 0x1c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x2e, 0x67, 0x73, 0x6b,
	0x39, 0x36, 0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x67, 0x73, 0x6b, 0x39, 0x36,
	0x37, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_users_proto_goTypes = []interface{}{
	(ActivityType)(0),                    // 0: gsk967.userActivity.ActivityType
	(Status)(0),                          // 1: gsk967.userActivity.Status
	(*UserInfo)(nil),                     // 2: gsk967.userActivity.UserInfo
	(*CreateUsersResponse)(nil),          // 3: gsk967.userActivity.CreateUsersResponse
	(*Activity)(nil),                     // 4: gsk967.userActivity.Activity
	(*GetUserActivityReq)(nil),           // 5: gsk967.userActivity.GetUserActivityReq
	(*UserActivity)(nil),                 // 6: gsk967.userActivity.UserActivity
	(*CreateUserActivityReq)(nil),        // 7: gsk967.userActivity.CreateUserActivityReq
	(*CreateUserActivityResponse)(nil),   // 8: gsk967.userActivity.CreateUserActivityResponse
	(*UpdateActivityStatusReq)(nil),      // 9: gsk967.userActivity.UpdateActivityStatusReq
	(*UpdateActivityStatusResponse)(nil), // 10: gsk967.userActivity.UpdateActivityStatusResponse
	(*timestamp.Timestamp)(nil),          // 11: google.protobuf.Timestamp
}
var file_users_proto_depIdxs = []int32{
	2,  // 0: gsk967.userActivity.CreateUsersResponse.userInfo:type_name -> gsk967.userActivity.UserInfo
	0,  // 1: gsk967.userActivity.Activity.activity:type_name -> gsk967.userActivity.ActivityType
	1,  // 2: gsk967.userActivity.Activity.status:type_name -> gsk967.userActivity.Status
	11, // 3: gsk967.userActivity.Activity.created_at:type_name -> google.protobuf.Timestamp
	11, // 4: gsk967.userActivity.Activity.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 5: gsk967.userActivity.UserActivity.daily_activities:type_name -> gsk967.userActivity.Activity
	4,  // 6: gsk967.userActivity.CreateUserActivityReq.activity:type_name -> gsk967.userActivity.Activity
	4,  // 7: gsk967.userActivity.CreateUserActivityResponse.activity:type_name -> gsk967.userActivity.Activity
	0,  // 8: gsk967.userActivity.UpdateActivityStatusReq.activity:type_name -> gsk967.userActivity.ActivityType
	1,  // 9: gsk967.userActivity.UpdateActivityStatusReq.status:type_name -> gsk967.userActivity.Status
	4,  // 10: gsk967.userActivity.UpdateActivityStatusResponse.activity:type_name -> gsk967.userActivity.Activity
	2,  // 11: gsk967.userActivity.Users.CreateUsersReq:input_type -> gsk967.userActivity.UserInfo
	5,  // 12: gsk967.userActivity.Users.GetUserActivityServiceReq:input_type -> gsk967.userActivity.GetUserActivityReq
	7,  // 13: gsk967.userActivity.Users.AddUserActivityServiceReq:input_type -> gsk967.userActivity.CreateUserActivityReq
	9,  // 14: gsk967.userActivity.Users.UpdateUserActivityServiceReq:input_type -> gsk967.userActivity.UpdateActivityStatusReq
	3,  // 15: gsk967.userActivity.Users.CreateUsersReq:output_type -> gsk967.userActivity.CreateUsersResponse
	6,  // 16: gsk967.userActivity.Users.GetUserActivityServiceReq:output_type -> gsk967.userActivity.UserActivity
	8,  // 17: gsk967.userActivity.Users.AddUserActivityServiceReq:output_type -> gsk967.userActivity.CreateUserActivityResponse
	10, // 18: gsk967.userActivity.Users.UpdateUserActivityServiceReq:output_type -> gsk967.userActivity.UpdateActivityStatusResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUsersResponse); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserActivityReq); i {
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
		file_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserActivity); i {
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
		file_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserActivityReq); i {
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
		file_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserActivityResponse); i {
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
		file_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateActivityStatusReq); i {
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
		file_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateActivityStatusResponse); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		EnumInfos:         file_users_proto_enumTypes,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}
