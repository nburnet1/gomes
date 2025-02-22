// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: namespace.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a node in the namespace
type Node struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentTopic   string                 `protobuf:"bytes,3,opt,name=parent_topic,json=parentTopic,proto3" json:"parent_topic,omitempty"`
	Value         string                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp     string                 `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Node) Reset() {
	*x = Node{}
	mi := &file_namespace_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetParentTopic() string {
	if x != nil {
		return x.ParentTopic
	}
	return ""
}

func (x *Node) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Node) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// Request for creating a node
type CreateNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNodeRequest) Reset() {
	*x = CreateNodeRequest{}
	mi := &file_namespace_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeRequest) ProtoMessage() {}

func (x *CreateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNodeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *CreateNodeRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Request for reading a node
type ReadNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadNodeRequest) Reset() {
	*x = ReadNodeRequest{}
	mi := &file_namespace_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadNodeRequest) ProtoMessage() {}

func (x *ReadNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadNodeRequest.ProtoReflect.Descriptor instead.
func (*ReadNodeRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *ReadNodeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type ReadNodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Node          *Node                  `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadNodeResponse) Reset() {
	*x = ReadNodeResponse{}
	mi := &file_namespace_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadNodeResponse) ProtoMessage() {}

func (x *ReadNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadNodeResponse.ProtoReflect.Descriptor instead.
func (*ReadNodeResponse) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{3}
}

func (x *ReadNodeResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

// Request for updating a node
type UpdateNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNodeRequest) Reset() {
	*x = UpdateNodeRequest{}
	mi := &file_namespace_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeRequest) ProtoMessage() {}

func (x *UpdateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNodeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *UpdateNodeRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Request for deleting a node
type DeleteNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteNodeRequest) Reset() {
	*x = DeleteNodeRequest{}
	mi := &file_namespace_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeRequest) ProtoMessage() {}

func (x *DeleteNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNodeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

// Response for deleting a node
type DeleteNodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteNodeResponse) Reset() {
	*x = DeleteNodeResponse{}
	mi := &file_namespace_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeResponse) ProtoMessage() {}

func (x *DeleteNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeResponse.ProtoReflect.Descriptor instead.
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteNodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request for browsing all nodes
type BrowseNodesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrowseNodesRequest) Reset() {
	*x = BrowseNodesRequest{}
	mi := &file_namespace_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrowseNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowseNodesRequest) ProtoMessage() {}

func (x *BrowseNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowseNodesRequest.ProtoReflect.Descriptor instead.
func (*BrowseNodesRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{7}
}

// Request for browsing root nodes
type BrowseRootNodesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrowseRootNodesRequest) Reset() {
	*x = BrowseRootNodesRequest{}
	mi := &file_namespace_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrowseRootNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowseRootNodesRequest) ProtoMessage() {}

func (x *BrowseRootNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowseRootNodesRequest.ProtoReflect.Descriptor instead.
func (*BrowseRootNodesRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{8}
}

// Response containing a list of nodes
type BrowseNodesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nodes         []*Node                `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrowseNodesResponse) Reset() {
	*x = BrowseNodesResponse{}
	mi := &file_namespace_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrowseNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowseNodesResponse) ProtoMessage() {}

func (x *BrowseNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowseNodesResponse.ProtoReflect.Descriptor instead.
func (*BrowseNodesResponse) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{9}
}

func (x *BrowseNodesResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// Request for subscribing to a node's updates
type SubNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubNodeRequest) Reset() {
	*x = SubNodeRequest{}
	mi := &file_namespace_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubNodeRequest) ProtoMessage() {}

func (x *SubNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubNodeRequest.ProtoReflect.Descriptor instead.
func (*SubNodeRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{10}
}

func (x *SubNodeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

// Streamed response for node updates
type NodeUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeUpdate) Reset() {
	*x = NodeUpdate{}
	mi := &file_namespace_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeUpdate) ProtoMessage() {}

func (x *NodeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeUpdate.ProtoReflect.Descriptor instead.
func (*NodeUpdate) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{11}
}

func (x *NodeUpdate) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *NodeUpdate) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *NodeUpdate) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_namespace_proto protoreflect.FileDescriptor

var file_namespace_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x87, 0x01, 0x0a,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x22, 0x37, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x42,
	0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x56, 0x0a, 0x0a, 0x4e,
	0x6f, 0x64, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x32, 0xbe, 0x04, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3b,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_namespace_proto_rawDescOnce sync.Once
	file_namespace_proto_rawDescData []byte
)

func file_namespace_proto_rawDescGZIP() []byte {
	file_namespace_proto_rawDescOnce.Do(func() {
		file_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_namespace_proto_rawDesc), len(file_namespace_proto_rawDesc)))
	})
	return file_namespace_proto_rawDescData
}

var file_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_namespace_proto_goTypes = []any{
	(*Node)(nil),                   // 0: namespace.Node
	(*CreateNodeRequest)(nil),      // 1: namespace.CreateNodeRequest
	(*ReadNodeRequest)(nil),        // 2: namespace.ReadNodeRequest
	(*ReadNodeResponse)(nil),       // 3: namespace.ReadNodeResponse
	(*UpdateNodeRequest)(nil),      // 4: namespace.UpdateNodeRequest
	(*DeleteNodeRequest)(nil),      // 5: namespace.DeleteNodeRequest
	(*DeleteNodeResponse)(nil),     // 6: namespace.DeleteNodeResponse
	(*BrowseNodesRequest)(nil),     // 7: namespace.BrowseNodesRequest
	(*BrowseRootNodesRequest)(nil), // 8: namespace.BrowseRootNodesRequest
	(*BrowseNodesResponse)(nil),    // 9: namespace.BrowseNodesResponse
	(*SubNodeRequest)(nil),         // 10: namespace.SubNodeRequest
	(*NodeUpdate)(nil),             // 11: namespace.NodeUpdate
}
var file_namespace_proto_depIdxs = []int32{
	0,  // 0: namespace.ReadNodeResponse.node:type_name -> namespace.Node
	0,  // 1: namespace.BrowseNodesResponse.nodes:type_name -> namespace.Node
	1,  // 2: namespace.NamespaceService.CreateNode:input_type -> namespace.CreateNodeRequest
	2,  // 3: namespace.NamespaceService.ReadNode:input_type -> namespace.ReadNodeRequest
	4,  // 4: namespace.NamespaceService.UpdateNode:input_type -> namespace.UpdateNodeRequest
	5,  // 5: namespace.NamespaceService.DeleteNode:input_type -> namespace.DeleteNodeRequest
	7,  // 6: namespace.NamespaceService.BrowseNodes:input_type -> namespace.BrowseNodesRequest
	8,  // 7: namespace.NamespaceService.BrowseRootNodes:input_type -> namespace.BrowseRootNodesRequest
	2,  // 8: namespace.NamespaceService.GetChildren:input_type -> namespace.ReadNodeRequest
	10, // 9: namespace.NamespaceService.SubNode:input_type -> namespace.SubNodeRequest
	0,  // 10: namespace.NamespaceService.CreateNode:output_type -> namespace.Node
	0,  // 11: namespace.NamespaceService.ReadNode:output_type -> namespace.Node
	0,  // 12: namespace.NamespaceService.UpdateNode:output_type -> namespace.Node
	6,  // 13: namespace.NamespaceService.DeleteNode:output_type -> namespace.DeleteNodeResponse
	9,  // 14: namespace.NamespaceService.BrowseNodes:output_type -> namespace.BrowseNodesResponse
	9,  // 15: namespace.NamespaceService.BrowseRootNodes:output_type -> namespace.BrowseNodesResponse
	9,  // 16: namespace.NamespaceService.GetChildren:output_type -> namespace.BrowseNodesResponse
	11, // 17: namespace.NamespaceService.SubNode:output_type -> namespace.NodeUpdate
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_namespace_proto_init() }
func file_namespace_proto_init() {
	if File_namespace_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_namespace_proto_rawDesc), len(file_namespace_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_namespace_proto_goTypes,
		DependencyIndexes: file_namespace_proto_depIdxs,
		MessageInfos:      file_namespace_proto_msgTypes,
	}.Build()
	File_namespace_proto = out.File
	file_namespace_proto_goTypes = nil
	file_namespace_proto_depIdxs = nil
}
