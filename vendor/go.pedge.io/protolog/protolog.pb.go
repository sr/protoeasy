// Code generated by protoc-gen-go.
// source: protolog.proto
// DO NOT EDIT!

package protolog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
import google_protobuf1 "go.pedge.io/google-protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Level is a logging level.
type Level int32

const (
	Level_LEVEL_NONE  Level = 0
	Level_LEVEL_DEBUG Level = 1
	Level_LEVEL_INFO  Level = 2
	Level_LEVEL_WARN  Level = 3
	Level_LEVEL_ERROR Level = 4
	Level_LEVEL_FATAL Level = 5
	Level_LEVEL_PANIC Level = 6
)

var Level_name = map[int32]string{
	0: "LEVEL_NONE",
	1: "LEVEL_DEBUG",
	2: "LEVEL_INFO",
	3: "LEVEL_WARN",
	4: "LEVEL_ERROR",
	5: "LEVEL_FATAL",
	6: "LEVEL_PANIC",
}
var Level_value = map[string]int32{
	"LEVEL_NONE":  0,
	"LEVEL_DEBUG": 1,
	"LEVEL_INFO":  2,
	"LEVEL_WARN":  3,
	"LEVEL_ERROR": 4,
	"LEVEL_FATAL": 5,
	"LEVEL_PANIC": 6,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}
func (Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// MessageType is the type of protolog message.
type MessageType int32

const (
	MessageType_MESSAGE_TYPE_NONE    MessageType = 0
	MessageType_MESSAGE_TYPE_EVENT   MessageType = 1
	MessageType_MESSAGE_TYPE_CONTEXT MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "MESSAGE_TYPE_NONE",
	1: "MESSAGE_TYPE_EVENT",
	2: "MESSAGE_TYPE_CONTEXT",
}
var MessageType_value = map[string]int32{
	"MESSAGE_TYPE_NONE":    0,
	"MESSAGE_TYPE_EVENT":   1,
	"MESSAGE_TYPE_CONTEXT": 2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Fields is a generic context Message used for
// the Logger functions WithField and WithFields.
type Fields struct {
	Value map[string]string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Fields) Reset()                    { *m = Fields{} }
func (m *Fields) String() string            { return proto.CompactTextString(m) }
func (*Fields) ProtoMessage()               {}
func (*Fields) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Fields) GetValue() map[string]string {
	if m != nil {
		return m.Value
	}
	return nil
}

// Event is a generic event Message used for
// the non-protobuf-specific Logger functions.
type Event struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// WriterOutput is an event Message used for
// the writer Logger functions.
type WriterOutput struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *WriterOutput) Reset()                    { *m = WriterOutput{} }
func (m *WriterOutput) String() string            { return proto.CompactTextString(m) }
func (*WriterOutput) ProtoMessage()               {}
func (*WriterOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Entry is the object serialized for logging.
type Entry struct {
	Id        string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Level     Level                       `protobuf:"varint,2,opt,name=level,enum=protolog.Level" json:"level,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Context   []*Entry_Message            `protobuf:"bytes,4,rep,name=context" json:"context,omitempty"`
	Event     *Entry_Message              `protobuf:"bytes,5,opt,name=event" json:"event,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Entry) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetContext() []*Entry_Message {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Entry) GetEvent() *Entry_Message {
	if m != nil {
		return m.Event
	}
	return nil
}

// Message represents a serialized protobuf Message.
// The name is the name registered with protolog.
type Entry_Message struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Entry_Message) Reset()                    { *m = Entry_Message{} }
func (m *Entry_Message) String() string            { return proto.CompactTextString(m) }
func (*Entry_Message) ProtoMessage()               {}
func (*Entry_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

var E_Event = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50009,
	Name:          "protolog.event",
	Tag:           "varint,50009,opt,name=event",
}

var E_Context = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50010,
	Name:          "protolog.context",
	Tag:           "varint,50010,opt,name=context",
}

func init() {
	proto.RegisterType((*Fields)(nil), "protolog.Fields")
	proto.RegisterType((*Event)(nil), "protolog.Event")
	proto.RegisterType((*WriterOutput)(nil), "protolog.WriterOutput")
	proto.RegisterType((*Entry)(nil), "protolog.Entry")
	proto.RegisterType((*Entry_Message)(nil), "protolog.Entry.Message")
	proto.RegisterEnum("protolog.Level", Level_name, Level_value)
	proto.RegisterEnum("protolog.MessageType", MessageType_name, MessageType_value)
	proto.RegisterExtension(E_Event)
	proto.RegisterExtension(E_Context)
}

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x26, 0x69, 0xd2, 0x6e, 0x6f, 0x4a, 0x17, 0x2c, 0x3e, 0xa2, 0x20, 0xb1, 0x69, 0x48, 0x53,
	0x55, 0x41, 0x86, 0xca, 0x05, 0xe5, 0x16, 0x86, 0x3b, 0x4d, 0xea, 0x92, 0xa9, 0x0b, 0x1d, 0x9c,
	0xa6, 0x6e, 0x35, 0x25, 0x22, 0x4d, 0xa2, 0xc4, 0xa9, 0xa8, 0xb8, 0xf3, 0x8f, 0xf8, 0x0f, 0x1c,
	0xb9, 0x82, 0xf8, 0x31, 0x38, 0x76, 0x52, 0x6b, 0x7c, 0x88, 0x9b, 0xdf, 0xd7, 0xcf, 0x97, 0x1f,
	0x43, 0x2f, 0xcb, 0x53, 0x9a, 0xc6, 0xe9, 0xc2, 0xe1, 0x07, 0xb4, 0xd5, 0xcc, 0xf6, 0xde, 0x22,
	0x4d, 0x17, 0x31, 0x39, 0xe4, 0x8b, 0xab, 0xf2, 0xdd, 0xe1, 0x9c, 0x14, 0xd7, 0x79, 0x94, 0xd1,
	0x34, 0x17, 0x58, 0x7b, 0xf7, 0x77, 0x04, 0x8d, 0x96, 0xa4, 0xa0, 0xb3, 0x65, 0x26, 0x00, 0xfb,
	0xef, 0xa1, 0x3d, 0x8a, 0x48, 0x3c, 0x2f, 0xd0, 0x00, 0xf4, 0xd5, 0x2c, 0x2e, 0x89, 0xa5, 0xec,
	0xb5, 0xfa, 0xc6, 0xf0, 0xa1, 0xb3, 0xb1, 0x15, 0x00, 0x67, 0x5a, 0xdd, 0xe2, 0x84, 0xe6, 0x6b,
	0xfb, 0x09, 0x80, 0x9c, 0x90, 0x01, 0xad, 0x0f, 0x64, 0xcd, 0x78, 0x4a, 0x7f, 0x1b, 0xdd, 0x6e,
	0x64, 0xd4, 0x6a, 0x74, 0xd5, 0x17, 0x8a, 0xab, 0x7d, 0xfb, 0x62, 0x29, 0xfb, 0x8f, 0x40, 0xc7,
	0x2b, 0x92, 0x50, 0xb4, 0x03, 0x1d, 0x96, 0xa1, 0x98, 0x2d, 0x88, 0xa0, 0xb8, 0xda, 0xd7, 0xea,
	0xfe, 0x31, 0x74, 0x2f, 0xf2, 0x88, 0x92, 0x3c, 0x28, 0x69, 0x56, 0x52, 0x29, 0x54, 0x81, 0xba,
	0x35, 0xe8, 0xa7, 0xc2, 0x54, 0xb8, 0x29, 0x80, 0x1a, 0xcd, 0x6b, 0x4f, 0x26, 0x1d, 0x93, 0x15,
	0x89, 0xb9, 0x67, 0x6f, 0xb8, 0x23, 0xa3, 0x8f, 0xab, 0x35, 0x7a, 0x0a, 0xdb, 0x9b, 0x77, 0x5b,
	0x2d, 0x86, 0x31, 0x86, 0xb6, 0x23, 0x9a, 0x71, 0x9a, 0x66, 0x9c, 0xb0, 0x41, 0xa0, 0x3e, 0x74,
	0xae, 0xd3, 0x84, 0x92, 0x8f, 0xd4, 0xd2, 0x78, 0x17, 0x0f, 0xa4, 0x20, 0x37, 0x77, 0x4e, 0x45,
	0x7e, 0x74, 0x00, 0x3a, 0xa9, 0xde, 0x64, 0xe9, 0x5c, 0xf4, 0x5f, 0x38, 0xfb, 0x00, 0x3a, 0x0d,
	0xa5, 0x0b, 0x5a, 0x32, 0x5b, 0x92, 0xbf, 0xb5, 0xd5, 0x1d, 0x7c, 0x02, 0x5d, 0x24, 0xee, 0x01,
	0x8c, 0xf1, 0x14, 0x8f, 0x2f, 0xfd, 0xc0, 0xc7, 0xe6, 0x2d, 0xd6, 0x99, 0x21, 0xe6, 0x57, 0xf8,
	0xe5, 0xeb, 0x63, 0x53, 0x91, 0x80, 0x13, 0x7f, 0x14, 0x98, 0xaa, 0x9c, 0x2f, 0xbc, 0x89, 0x6f,
	0xb6, 0x24, 0x01, 0x4f, 0x26, 0xc1, 0xc4, 0xd4, 0xe4, 0x62, 0xe4, 0x85, 0xde, 0xd8, 0xd4, 0xe5,
	0xe2, 0xcc, 0xf3, 0x4f, 0x8e, 0xcc, 0xf6, 0x60, 0x0a, 0x46, 0x1d, 0x32, 0x5c, 0x67, 0x04, 0xdd,
	0x83, 0x3b, 0xa7, 0xf8, 0xfc, 0xdc, 0x3b, 0xc6, 0x97, 0xe1, 0xdb, 0x33, 0xdc, 0x24, 0xb9, 0x0f,
	0xe8, 0xc6, 0x9a, 0x49, 0xf8, 0x21, 0x0b, 0x64, 0xc1, 0xdd, 0x1b, 0xfb, 0xa3, 0xc0, 0x0f, 0xf1,
	0x9b, 0xd0, 0x54, 0xdd, 0x67, 0x75, 0x49, 0x68, 0xf7, 0x8f, 0xce, 0x6b, 0xbf, 0x20, 0xa3, 0x51,
	0x9a, 0x14, 0xd6, 0xf7, 0xcf, 0xd5, 0xe7, 0x6c, 0xb9, 0xc3, 0xcd, 0x07, 0xfc, 0x9f, 0xf3, 0x43,
	0x70, 0xae, 0xda, 0x1c, 0xf0, 0xfc, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0xce, 0x71, 0x5d,
	0x2e, 0x03, 0x00, 0x00,
}