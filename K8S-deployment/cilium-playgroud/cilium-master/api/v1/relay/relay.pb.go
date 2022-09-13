// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relay/relay.proto

package relay

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeState int32

const (
	// UNKNOWN_NODE_STATE indicates that the state of this node is unknown.
	NodeState_UNKNOWN_NODE_STATE NodeState = 0
	// NODE_CONNECTED indicates that we have established a connection
	// to this node. The client can expect to observe flows from this node.
	NodeState_NODE_CONNECTED NodeState = 1
	// NODE_UNAVAILABLE indicates that the connection to this
	// node is currently unavailable. The client can expect to not see any
	// flows from this node until either the connection is re-established or
	// the node is gone.
	NodeState_NODE_UNAVAILABLE NodeState = 2
	// NODE_GONE indicates that a node has been removed from the
	// cluster. No reconnection attempts will be made.
	NodeState_NODE_GONE NodeState = 3
	// NODE_ERROR indicates that a node has reported an error while processing
	// the request. No reconnection attempts will be made.
	NodeState_NODE_ERROR NodeState = 4
)

var NodeState_name = map[int32]string{
	0: "UNKNOWN_NODE_STATE",
	1: "NODE_CONNECTED",
	2: "NODE_UNAVAILABLE",
	3: "NODE_GONE",
	4: "NODE_ERROR",
}

var NodeState_value = map[string]int32{
	"UNKNOWN_NODE_STATE": 0,
	"NODE_CONNECTED":     1,
	"NODE_UNAVAILABLE":   2,
	"NODE_GONE":          3,
	"NODE_ERROR":         4,
}

func (x NodeState) String() string {
	return proto.EnumName(NodeState_name, int32(x))
}

func (NodeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6da3b5c0d1535b3, []int{0}
}

// NodeStatusEvent is a message sent by hubble-relay to inform clients about
// the state of a particular node.
type NodeStatusEvent struct {
	// state_change contains the new node state
	StateChange NodeState `protobuf:"varint,1,opt,name=state_change,json=stateChange,proto3,enum=relay.NodeState" json:"state_change,omitempty"`
	// node_names is the list of nodes for which the above state changes applies
	NodeNames []string `protobuf:"bytes,2,rep,name=node_names,json=nodeNames,proto3" json:"node_names,omitempty"`
	// message is an optional message attached to the state change (e.g. an
	// error message). The message applies to all nodes in node_names.
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStatusEvent) Reset()         { *m = NodeStatusEvent{} }
func (m *NodeStatusEvent) String() string { return proto.CompactTextString(m) }
func (*NodeStatusEvent) ProtoMessage()    {}
func (*NodeStatusEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6da3b5c0d1535b3, []int{0}
}

func (m *NodeStatusEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStatusEvent.Unmarshal(m, b)
}
func (m *NodeStatusEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStatusEvent.Marshal(b, m, deterministic)
}
func (m *NodeStatusEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStatusEvent.Merge(m, src)
}
func (m *NodeStatusEvent) XXX_Size() int {
	return xxx_messageInfo_NodeStatusEvent.Size(m)
}
func (m *NodeStatusEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStatusEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStatusEvent proto.InternalMessageInfo

func (m *NodeStatusEvent) GetStateChange() NodeState {
	if m != nil {
		return m.StateChange
	}
	return NodeState_UNKNOWN_NODE_STATE
}

func (m *NodeStatusEvent) GetNodeNames() []string {
	if m != nil {
		return m.NodeNames
	}
	return nil
}

func (m *NodeStatusEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("relay.NodeState", NodeState_name, NodeState_value)
	proto.RegisterType((*NodeStatusEvent)(nil), "relay.NodeStatusEvent")
}

func init() { proto.RegisterFile("relay/relay.proto", fileDescriptor_b6da3b5c0d1535b3) }

var fileDescriptor_b6da3b5c0d1535b3 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x40, 0x4d, 0xe3, 0x07, 0x3b, 0x6a, 0x5c, 0x07, 0x91, 0x5c, 0x84, 0xe0, 0x29, 0x78, 0xa8,
	0x60, 0x7f, 0x41, 0x4c, 0x17, 0x11, 0xcb, 0x2c, 0x6c, 0x53, 0x3d, 0x86, 0xd5, 0x0e, 0xf5, 0xd0,
	0x26, 0xd2, 0x5d, 0x05, 0x4f, 0xfe, 0x75, 0xe9, 0x48, 0xbc, 0x0c, 0xbc, 0xc7, 0x9b, 0x81, 0x81,
	0xf3, 0x2d, 0xaf, 0xfd, 0xf7, 0xad, 0xcc, 0xf1, 0xc7, 0xb6, 0x8f, 0x3d, 0x1e, 0x08, 0x5c, 0xff,
	0xc0, 0x19, 0xf5, 0x4b, 0x9e, 0x47, 0x1f, 0x3f, 0x83, 0xf9, 0xe2, 0x2e, 0xe2, 0x04, 0x4e, 0x42,
	0xf4, 0x91, 0xdb, 0xb7, 0x77, 0xdf, 0xad, 0x38, 0x4f, 0x8a, 0xa4, 0xcc, 0xee, 0xf4, 0xf8, 0x6f,
	0x7b, 0xa8, 0xd9, 0x1d, 0x4b, 0x55, 0x4b, 0x84, 0x57, 0x00, 0x5d, 0xbf, 0xe4, 0xb6, 0xf3, 0x1b,
	0x0e, 0xf9, 0xa8, 0x48, 0x4b, 0xe5, 0xd4, 0xce, 0xd0, 0x4e, 0x60, 0x0e, 0x47, 0x1b, 0x0e, 0xc1,
	0xaf, 0x38, 0x4f, 0x8b, 0xa4, 0x54, 0x6e, 0xc0, 0x9b, 0x35, 0xa8, 0xff, 0x93, 0x78, 0x09, 0xb8,
	0xa0, 0x27, 0xb2, 0x2f, 0xd4, 0x92, 0x9d, 0x9a, 0x76, 0xde, 0x54, 0x8d, 0xd1, 0x7b, 0x88, 0x90,
	0x09, 0xd7, 0x96, 0xc8, 0xd4, 0x8d, 0x99, 0xea, 0x04, 0x2f, 0x40, 0x8b, 0x5b, 0x50, 0xf5, 0x5c,
	0x3d, 0xce, 0xaa, 0xfb, 0x99, 0xd1, 0x23, 0x3c, 0x05, 0x25, 0xf6, 0xc1, 0x92, 0xd1, 0x29, 0x66,
	0x00, 0x82, 0xc6, 0x39, 0xeb, 0xf4, 0xfe, 0xeb, 0xa1, 0x3c, 0x3f, 0xf9, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x81, 0xcf, 0x9e, 0x11, 0x01, 0x00, 0x00,
}
