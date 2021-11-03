// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: broker.proto

package rest

import (
	activity "github.com/pydio/cells/v4/common/proto/activity"
	log "github.com/pydio/cells/v4/common/proto/log"
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

// Collection of Activities
type ActivitiesCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activities []*activity.Object `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
}

func (x *ActivitiesCollection) Reset() {
	*x = ActivitiesCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivitiesCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitiesCollection) ProtoMessage() {}

func (x *ActivitiesCollection) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivitiesCollection.ProtoReflect.Descriptor instead.
func (*ActivitiesCollection) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{0}
}

func (x *ActivitiesCollection) GetActivities() []*activity.Object {
	if x != nil {
		return x.Activities
	}
	return nil
}

type SubscriptionsCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*activity.Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *SubscriptionsCollection) Reset() {
	*x = SubscriptionsCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionsCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionsCollection) ProtoMessage() {}

func (x *SubscriptionsCollection) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionsCollection.ProtoReflect.Descriptor instead.
func (*SubscriptionsCollection) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionsCollection) GetSubscriptions() []*activity.Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

// Collection of serialized log messages
type LogCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*log.Log `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *LogCollection) Reset() {
	*x = LogCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCollection) ProtoMessage() {}

func (x *LogCollection) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCollection.ProtoReflect.Descriptor instead.
func (*LogCollection) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{2}
}

func (x *LogCollection) GetLines() []*log.Log {
	if x != nil {
		return x.Lines
	}
	return nil
}

// Collection of serialized log messages
type LogMessageCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*log.LogMessage `protobuf:"bytes,1,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (x *LogMessageCollection) Reset() {
	*x = LogMessageCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMessageCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMessageCollection) ProtoMessage() {}

func (x *LogMessageCollection) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMessageCollection.ProtoReflect.Descriptor instead.
func (*LogMessageCollection) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{3}
}

func (x *LogMessageCollection) GetLogs() []*log.LogMessage {
	if x != nil {
		return x.Logs
	}
	return nil
}

// Collection of serialized aggregated result of time range request
// with a cursor to ease navigation implementation
type TimeRangeResultCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*log.TimeRangeResult `protobuf:"bytes,1,rep,name=Results,proto3" json:"Results,omitempty"`
	Links   []*log.TimeRangeCursor `protobuf:"bytes,2,rep,name=Links,proto3" json:"Links,omitempty"`
}

func (x *TimeRangeResultCollection) Reset() {
	*x = TimeRangeResultCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRangeResultCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRangeResultCollection) ProtoMessage() {}

func (x *TimeRangeResultCollection) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRangeResultCollection.ProtoReflect.Descriptor instead.
func (*TimeRangeResultCollection) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{4}
}

func (x *TimeRangeResultCollection) GetResults() []*log.TimeRangeResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *TimeRangeResultCollection) GetLinks() []*log.TimeRangeCursor {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_broker_proto protoreflect.FileDescriptor

var file_broker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x72, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x57, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x14, 0x4c, 0x6f, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x77, 0x0a, 0x19, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79,
	0x64, 0x69, 0x6f, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_broker_proto_rawDescOnce sync.Once
	file_broker_proto_rawDescData = file_broker_proto_rawDesc
)

func file_broker_proto_rawDescGZIP() []byte {
	file_broker_proto_rawDescOnce.Do(func() {
		file_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_proto_rawDescData)
	})
	return file_broker_proto_rawDescData
}

var file_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_broker_proto_goTypes = []interface{}{
	(*ActivitiesCollection)(nil),      // 0: rest.ActivitiesCollection
	(*SubscriptionsCollection)(nil),   // 1: rest.SubscriptionsCollection
	(*LogCollection)(nil),             // 2: rest.LogCollection
	(*LogMessageCollection)(nil),      // 3: rest.LogMessageCollection
	(*TimeRangeResultCollection)(nil), // 4: rest.TimeRangeResultCollection
	(*activity.Object)(nil),           // 5: activity.Object
	(*activity.Subscription)(nil),     // 6: activity.Subscription
	(*log.Log)(nil),                   // 7: log.Log
	(*log.LogMessage)(nil),            // 8: log.LogMessage
	(*log.TimeRangeResult)(nil),       // 9: log.TimeRangeResult
	(*log.TimeRangeCursor)(nil),       // 10: log.TimeRangeCursor
}
var file_broker_proto_depIdxs = []int32{
	5,  // 0: rest.ActivitiesCollection.activities:type_name -> activity.Object
	6,  // 1: rest.SubscriptionsCollection.subscriptions:type_name -> activity.Subscription
	7,  // 2: rest.LogCollection.lines:type_name -> log.Log
	8,  // 3: rest.LogMessageCollection.Logs:type_name -> log.LogMessage
	9,  // 4: rest.TimeRangeResultCollection.Results:type_name -> log.TimeRangeResult
	10, // 5: rest.TimeRangeResultCollection.Links:type_name -> log.TimeRangeCursor
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_broker_proto_init() }
func file_broker_proto_init() {
	if File_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivitiesCollection); i {
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
		file_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionsCollection); i {
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
		file_broker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogCollection); i {
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
		file_broker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMessageCollection); i {
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
		file_broker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRangeResultCollection); i {
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
			RawDescriptor: file_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_broker_proto_goTypes,
		DependencyIndexes: file_broker_proto_depIdxs,
		MessageInfos:      file_broker_proto_msgTypes,
	}.Build()
	File_broker_proto = out.File
	file_broker_proto_rawDesc = nil
	file_broker_proto_goTypes = nil
	file_broker_proto_depIdxs = nil
}