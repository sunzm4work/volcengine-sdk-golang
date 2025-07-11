// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.4
// source: volcengine/vod/business/vod_edit.proto

package business

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubmitDirectEditTaskAsyncResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId string `protobuf:"bytes,1,opt,name=ReqId,proto3" json:"ReqId,omitempty"` // 视频编辑执行Id
}

func (x *SubmitDirectEditTaskAsyncResult) Reset() {
	*x = SubmitDirectEditTaskAsyncResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDirectEditTaskAsyncResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDirectEditTaskAsyncResult) ProtoMessage() {}

func (x *SubmitDirectEditTaskAsyncResult) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDirectEditTaskAsyncResult.ProtoReflect.Descriptor instead.
func (*SubmitDirectEditTaskAsyncResult) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_edit_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitDirectEditTaskAsyncResult) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

type SubmitDirectEditTaskSyncResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	ReqId  string `protobuf:"bytes,2,opt,name=ReqId,proto3" json:"ReqId,omitempty"`
	Output string `protobuf:"bytes,3,opt,name=Output,proto3" json:"Output,omitempty"`
}

func (x *SubmitDirectEditTaskSyncResult) Reset() {
	*x = SubmitDirectEditTaskSyncResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDirectEditTaskSyncResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDirectEditTaskSyncResult) ProtoMessage() {}

func (x *SubmitDirectEditTaskSyncResult) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDirectEditTaskSyncResult.ProtoReflect.Descriptor instead.
func (*SubmitDirectEditTaskSyncResult) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_edit_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitDirectEditTaskSyncResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitDirectEditTaskSyncResult) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *SubmitDirectEditTaskSyncResult) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type GetDirectEditProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *GetDirectEditProgress) Reset() {
	*x = GetDirectEditProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectEditProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectEditProgress) ProtoMessage() {}

func (x *GetDirectEditProgress) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectEditProgress.ProtoReflect.Descriptor instead.
func (*GetDirectEditProgress) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_edit_proto_rawDescGZIP(), []int{2}
}

func (x *GetDirectEditProgress) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type GetDirectEditResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId        string `protobuf:"bytes,1,opt,name=ReqId,proto3" json:"ReqId,omitempty"`               // 视频编辑执行Id
	EditParam    []byte `protobuf:"bytes,2,opt,name=EditParam,proto3" json:"EditParam,omitempty"`       // 视频编辑参数
	Priority     int32  `protobuf:"varint,3,opt,name=Priority,proto3" json:"Priority,omitempty"`        // 优先级
	CallbackUri  string `protobuf:"bytes,4,opt,name=CallbackUri,proto3" json:"CallbackUri,omitempty"`   // 回调地址
	CallbackArgs string `protobuf:"bytes,5,opt,name=CallbackArgs,proto3" json:"CallbackArgs,omitempty"` // 回调参数
	Status       string `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`             // 编辑任务状态
	OutputVid    string `protobuf:"bytes,7,opt,name=OutputVid,proto3" json:"OutputVid,omitempty"`       // 产物vid
	Message      string `protobuf:"bytes,8,opt,name=Message,proto3" json:"Message,omitempty"`           // 错误信息
}

func (x *GetDirectEditResult) Reset() {
	*x = GetDirectEditResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectEditResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectEditResult) ProtoMessage() {}

func (x *GetDirectEditResult) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectEditResult.ProtoReflect.Descriptor instead.
func (*GetDirectEditResult) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_edit_proto_rawDescGZIP(), []int{3}
}

func (x *GetDirectEditResult) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *GetDirectEditResult) GetEditParam() []byte {
	if x != nil {
		return x.EditParam
	}
	return nil
}

func (x *GetDirectEditResult) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *GetDirectEditResult) GetCallbackUri() string {
	if x != nil {
		return x.CallbackUri
	}
	return ""
}

func (x *GetDirectEditResult) GetCallbackArgs() string {
	if x != nil {
		return x.CallbackArgs
	}
	return ""
}

func (x *GetDirectEditResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetDirectEditResult) GetOutputVid() string {
	if x != nil {
		return x.OutputVid
	}
	return ""
}

func (x *GetDirectEditResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CancelDirectEditTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"` // 取消信息
}

func (x *CancelDirectEditTask) Reset() {
	*x = CancelDirectEditTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelDirectEditTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelDirectEditTask) ProtoMessage() {}

func (x *CancelDirectEditTask) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_edit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelDirectEditTask.ProtoReflect.Descriptor instead.
func (*CancelDirectEditTask) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_edit_proto_rawDescGZIP(), []int{4}
}

func (x *CancelDirectEditTask) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_volcengine_vod_business_vod_edit_proto protoreflect.FileDescriptor

var file_volcengine_vod_business_vod_edit_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x64,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x5f, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x1f, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x45, 0x64, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x52,
	0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x71, 0x49,
	0x64, 0x22, 0x62, 0x0a, 0x1e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x45, 0x64, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x71, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xfb, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x52, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52,
	0x65, 0x71, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x64, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x45, 0x64, 0x69, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x69,
	0x12, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x56, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x56, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x45, 0x64, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xc8, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x6f, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x42, 0x07, 0x56, 0x6f, 0x64, 0x45, 0x64, 0x69, 0x74, 0x50, 0x01, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xca, 0x02, 0x20, 0x56, 0x6f, 0x6c, 0x63, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x23, 0x56, 0x6f,
	0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volcengine_vod_business_vod_edit_proto_rawDescOnce sync.Once
	file_volcengine_vod_business_vod_edit_proto_rawDescData = file_volcengine_vod_business_vod_edit_proto_rawDesc
)

func file_volcengine_vod_business_vod_edit_proto_rawDescGZIP() []byte {
	file_volcengine_vod_business_vod_edit_proto_rawDescOnce.Do(func() {
		file_volcengine_vod_business_vod_edit_proto_rawDescData = protoimpl.X.CompressGZIP(file_volcengine_vod_business_vod_edit_proto_rawDescData)
	})
	return file_volcengine_vod_business_vod_edit_proto_rawDescData
}

var file_volcengine_vod_business_vod_edit_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_volcengine_vod_business_vod_edit_proto_goTypes = []interface{}{
	(*SubmitDirectEditTaskAsyncResult)(nil), // 0: Volcengine.Vod.Models.Business.SubmitDirectEditTaskAsyncResult
	(*SubmitDirectEditTaskSyncResult)(nil),  // 1: Volcengine.Vod.Models.Business.SubmitDirectEditTaskSyncResult
	(*GetDirectEditProgress)(nil),           // 2: Volcengine.Vod.Models.Business.GetDirectEditProgress
	(*GetDirectEditResult)(nil),             // 3: Volcengine.Vod.Models.Business.GetDirectEditResult
	(*CancelDirectEditTask)(nil),            // 4: Volcengine.Vod.Models.Business.CancelDirectEditTask
}
var file_volcengine_vod_business_vod_edit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_volcengine_vod_business_vod_edit_proto_init() }
func file_volcengine_vod_business_vod_edit_proto_init() {
	if File_volcengine_vod_business_vod_edit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_volcengine_vod_business_vod_edit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDirectEditTaskAsyncResult); i {
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
		file_volcengine_vod_business_vod_edit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDirectEditTaskSyncResult); i {
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
		file_volcengine_vod_business_vod_edit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectEditProgress); i {
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
		file_volcengine_vod_business_vod_edit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectEditResult); i {
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
		file_volcengine_vod_business_vod_edit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelDirectEditTask); i {
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
			RawDescriptor: file_volcengine_vod_business_vod_edit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_volcengine_vod_business_vod_edit_proto_goTypes,
		DependencyIndexes: file_volcengine_vod_business_vod_edit_proto_depIdxs,
		MessageInfos:      file_volcengine_vod_business_vod_edit_proto_msgTypes,
	}.Build()
	File_volcengine_vod_business_vod_edit_proto = out.File
	file_volcengine_vod_business_vod_edit_proto_rawDesc = nil
	file_volcengine_vod_business_vod_edit_proto_goTypes = nil
	file_volcengine_vod_business_vod_edit_proto_depIdxs = nil
}
