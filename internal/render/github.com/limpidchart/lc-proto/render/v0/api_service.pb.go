// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api_service.proto

package render

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

// ChartStatus contains available chart statuses.
type ChartStatus int32

const (
	ChartStatus_UNSPECIFIED_STATUS ChartStatus = 0
	ChartStatus_CREATED            ChartStatus = 1
	ChartStatus_ERROR              ChartStatus = 2
)

// Enum value maps for ChartStatus.
var (
	ChartStatus_name = map[int32]string{
		0: "UNSPECIFIED_STATUS",
		1: "CREATED",
		2: "ERROR",
	}
	ChartStatus_value = map[string]int32{
		"UNSPECIFIED_STATUS": 0,
		"CREATED":            1,
		"ERROR":              2,
	}
)

func (x ChartStatus) Enum() *ChartStatus {
	p := new(ChartStatus)
	*p = x
	return p
}

func (x ChartStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_service_proto_enumTypes[0].Descriptor()
}

func (ChartStatus) Type() protoreflect.EnumType {
	return &file_api_service_proto_enumTypes[0]
}

func (x ChartStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartStatus.Descriptor instead.
func (ChartStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{0}
}

// CreateChartRequest represents chart creation request.
type CreateChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chart title.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Configured chart sizes.
	Sizes *ChartSizes `protobuf:"bytes,2,opt,name=sizes,proto3" json:"sizes,omitempty"`
	// Configured chart margins.
	Margins *ChartMargins `protobuf:"bytes,3,opt,name=margins,proto3" json:"margins,omitempty"`
	// Configured chart axes.
	Axes *ChartAxes `protobuf:"bytes,4,opt,name=axes,proto3" json:"axes,omitempty"`
	// Configured chart views.
	Views []*ChartView `protobuf:"bytes,5,rep,name=views,proto3" json:"views,omitempty"`
}

func (x *CreateChartRequest) Reset() {
	*x = CreateChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChartRequest) ProtoMessage() {}

func (x *CreateChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChartRequest.ProtoReflect.Descriptor instead.
func (*CreateChartRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChartRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateChartRequest) GetSizes() *ChartSizes {
	if x != nil {
		return x.Sizes
	}
	return nil
}

func (x *CreateChartRequest) GetMargins() *ChartMargins {
	if x != nil {
		return x.Margins
	}
	return nil
}

func (x *CreateChartRequest) GetAxes() *ChartAxes {
	if x != nil {
		return x.Axes
	}
	return nil
}

func (x *CreateChartRequest) GetViews() []*ChartView {
	if x != nil {
		return x.Views
	}
	return nil
}

// GetChartRequest represents chart get request.
type GetChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the chart.
	ChartId string `protobuf:"bytes,1,opt,name=chart_id,json=chartId,proto3" json:"chart_id,omitempty"`
}

func (x *GetChartRequest) Reset() {
	*x = GetChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChartRequest) ProtoMessage() {}

func (x *GetChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChartRequest.ProtoReflect.Descriptor instead.
func (*GetChartRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetChartRequest) GetChartId() string {
	if x != nil {
		return x.ChartId
	}
	return ""
}

// ChartReply represents chart reply from create or get requests.
type ChartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the request.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// ID of the chart.
	ChartId string `protobuf:"bytes,2,opt,name=chart_id,json=chartId,proto3" json:"chart_id,omitempty"`
	// One of the available chart statuses.
	ChartStatus ChartStatus `protobuf:"varint,3,opt,name=chart_status,json=chartStatus,proto3,enum=render.ChartStatus" json:"chart_status,omitempty"`
	// Chart creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// When the chart will become unavailable.
	// It depends of `ttl` value that was provided in chart create request.
	// It equals to `created_at` value if `ttl` was set zero.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// Chart raw bytes representation.
	ChartData []byte `protobuf:"bytes,6,opt,name=chart_data,json=chartData,proto3" json:"chart_data,omitempty"`
}

func (x *ChartReply) Reset() {
	*x = ChartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartReply) ProtoMessage() {}

func (x *ChartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartReply.ProtoReflect.Descriptor instead.
func (*ChartReply) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChartReply) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ChartReply) GetChartId() string {
	if x != nil {
		return x.ChartId
	}
	return ""
}

func (x *ChartReply) GetChartStatus() ChartStatus {
	if x != nil {
		return x.ChartStatus
	}
	return ChartStatus_UNSPECIFIED_STATUS
}

func (x *ChartReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ChartReply) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *ChartReply) GetChartData() []byte {
	if x != nil {
		return x.ChartData
	}
	return nil
}

var File_api_service_proto protoreflect.FileDescriptor

var file_api_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x0b, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x73, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x07,
	0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4d, 0x61, 0x72, 0x67,
	0x69, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x04,
	0x61, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x41, 0x78, 0x65, 0x73, 0x52, 0x04, 0x61,
	0x78, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x2c, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x74, 0x49, 0x64, 0x22, 0x93, 0x02, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x2a, 0x3d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32,
	0x86, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x72, 0x74, 0x41, 0x50, 0x49, 0x12, 0x3f, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6d, 0x70, 0x69, 0x64, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x2f, 0x6c, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x30, 0x3b, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_service_proto_rawDescOnce sync.Once
	file_api_service_proto_rawDescData = file_api_service_proto_rawDesc
)

func file_api_service_proto_rawDescGZIP() []byte {
	file_api_service_proto_rawDescOnce.Do(func() {
		file_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_service_proto_rawDescData)
	})
	return file_api_service_proto_rawDescData
}

var file_api_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_service_proto_goTypes = []interface{}{
	(ChartStatus)(0),              // 0: render.ChartStatus
	(*CreateChartRequest)(nil),    // 1: render.CreateChartRequest
	(*GetChartRequest)(nil),       // 2: render.GetChartRequest
	(*ChartReply)(nil),            // 3: render.ChartReply
	(*ChartSizes)(nil),            // 4: render.ChartSizes
	(*ChartMargins)(nil),          // 5: render.ChartMargins
	(*ChartAxes)(nil),             // 6: render.ChartAxes
	(*ChartView)(nil),             // 7: render.ChartView
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_api_service_proto_depIdxs = []int32{
	4, // 0: render.CreateChartRequest.sizes:type_name -> render.ChartSizes
	5, // 1: render.CreateChartRequest.margins:type_name -> render.ChartMargins
	6, // 2: render.CreateChartRequest.axes:type_name -> render.ChartAxes
	7, // 3: render.CreateChartRequest.views:type_name -> render.ChartView
	0, // 4: render.ChartReply.chart_status:type_name -> render.ChartStatus
	8, // 5: render.ChartReply.created_at:type_name -> google.protobuf.Timestamp
	8, // 6: render.ChartReply.deleted_at:type_name -> google.protobuf.Timestamp
	1, // 7: render.ChartAPI.CreateChart:input_type -> render.CreateChartRequest
	2, // 8: render.ChartAPI.GetChart:input_type -> render.GetChartRequest
	3, // 9: render.ChartAPI.CreateChart:output_type -> render.ChartReply
	3, // 10: render.ChartAPI.GetChart:output_type -> render.ChartReply
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_service_proto_init() }
func file_api_service_proto_init() {
	if File_api_service_proto != nil {
		return
	}
	file_chart_proto_init()
	file_view_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChartRequest); i {
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
		file_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChartRequest); i {
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
		file_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartReply); i {
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
			RawDescriptor: file_api_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_proto_goTypes,
		DependencyIndexes: file_api_service_proto_depIdxs,
		EnumInfos:         file_api_service_proto_enumTypes,
		MessageInfos:      file_api_service_proto_msgTypes,
	}.Build()
	File_api_service_proto = out.File
	file_api_service_proto_rawDesc = nil
	file_api_service_proto_goTypes = nil
	file_api_service_proto_depIdxs = nil
}
