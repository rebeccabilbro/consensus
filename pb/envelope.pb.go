// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envelope.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specifies the message type in a peer request and reply
type Type int32

const (
	Type_BEACON       Type = 0
	Type_PREPARE      Type = 1
	Type_ACCEPT       Type = 2
	Type_COMMIT       Type = 3
	Type_TRYPREACCEPT Type = 4
	Type_PREACCEPT    Type = 5
)

var Type_name = map[int32]string{
	0: "BEACON",
	1: "PREPARE",
	2: "ACCEPT",
	3: "COMMIT",
	4: "TRYPREACCEPT",
	5: "PREACCEPT",
}
var Type_value = map[string]int32{
	"BEACON":       0,
	"PREPARE":      1,
	"ACCEPT":       2,
	"COMMIT":       3,
	"TRYPREACCEPT": 4,
	"PREACCEPT":    5,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// A wrapper message that can contain one of the request message types.
type PeerRequest struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,enum=pb.Type" json:"type,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	// only one of these fields can be set, and the field that is set should
	// match the message type described above.
	//
	// Types that are valid to be assigned to Message:
	//	*PeerRequest_Beacon
	//	*PeerRequest_Prepare
	//	*PeerRequest_Accept
	//	*PeerRequest_Commit
	//	*PeerRequest_TryPreAcccept
	//	*PeerRequest_PreAccept
	Message isPeerRequest_Message `protobuf_oneof:"message"`
}

func (m *PeerRequest) Reset()                    { *m = PeerRequest{} }
func (m *PeerRequest) String() string            { return proto.CompactTextString(m) }
func (*PeerRequest) ProtoMessage()               {}
func (*PeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isPeerRequest_Message interface {
	isPeerRequest_Message()
}

type PeerRequest_Beacon struct {
	Beacon *BeaconRequest `protobuf:"bytes,10,opt,name=beacon,oneof"`
}
type PeerRequest_Prepare struct {
	Prepare *PrepareRequest `protobuf:"bytes,11,opt,name=prepare,oneof"`
}
type PeerRequest_Accept struct {
	Accept *AcceptRequest `protobuf:"bytes,12,opt,name=accept,oneof"`
}
type PeerRequest_Commit struct {
	Commit *CommitRequest `protobuf:"bytes,13,opt,name=commit,oneof"`
}
type PeerRequest_TryPreAcccept struct {
	TryPreAcccept *TryPreAcceptRequest `protobuf:"bytes,14,opt,name=tryPreAcccept,oneof"`
}
type PeerRequest_PreAccept struct {
	PreAccept *PreAcceptRequest `protobuf:"bytes,15,opt,name=preAccept,oneof"`
}

func (*PeerRequest_Beacon) isPeerRequest_Message()        {}
func (*PeerRequest_Prepare) isPeerRequest_Message()       {}
func (*PeerRequest_Accept) isPeerRequest_Message()        {}
func (*PeerRequest_Commit) isPeerRequest_Message()        {}
func (*PeerRequest_TryPreAcccept) isPeerRequest_Message() {}
func (*PeerRequest_PreAccept) isPeerRequest_Message()     {}

func (m *PeerRequest) GetMessage() isPeerRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PeerRequest) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_BEACON
}

func (m *PeerRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PeerRequest) GetBeacon() *BeaconRequest {
	if x, ok := m.GetMessage().(*PeerRequest_Beacon); ok {
		return x.Beacon
	}
	return nil
}

func (m *PeerRequest) GetPrepare() *PrepareRequest {
	if x, ok := m.GetMessage().(*PeerRequest_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *PeerRequest) GetAccept() *AcceptRequest {
	if x, ok := m.GetMessage().(*PeerRequest_Accept); ok {
		return x.Accept
	}
	return nil
}

func (m *PeerRequest) GetCommit() *CommitRequest {
	if x, ok := m.GetMessage().(*PeerRequest_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *PeerRequest) GetTryPreAcccept() *TryPreAcceptRequest {
	if x, ok := m.GetMessage().(*PeerRequest_TryPreAcccept); ok {
		return x.TryPreAcccept
	}
	return nil
}

func (m *PeerRequest) GetPreAccept() *PreAcceptRequest {
	if x, ok := m.GetMessage().(*PeerRequest_PreAccept); ok {
		return x.PreAccept
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PeerRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PeerRequest_OneofMarshaler, _PeerRequest_OneofUnmarshaler, _PeerRequest_OneofSizer, []interface{}{
		(*PeerRequest_Beacon)(nil),
		(*PeerRequest_Prepare)(nil),
		(*PeerRequest_Accept)(nil),
		(*PeerRequest_Commit)(nil),
		(*PeerRequest_TryPreAcccept)(nil),
		(*PeerRequest_PreAccept)(nil),
	}
}

func _PeerRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PeerRequest)
	// message
	switch x := m.Message.(type) {
	case *PeerRequest_Beacon:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Beacon); err != nil {
			return err
		}
	case *PeerRequest_Prepare:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *PeerRequest_Accept:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Accept); err != nil {
			return err
		}
	case *PeerRequest_Commit:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *PeerRequest_TryPreAcccept:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TryPreAcccept); err != nil {
			return err
		}
	case *PeerRequest_PreAccept:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PreAccept); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PeerRequest.Message has unexpected type %T", x)
	}
	return nil
}

func _PeerRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PeerRequest)
	switch tag {
	case 10: // message.beacon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BeaconRequest)
		err := b.DecodeMessage(msg)
		m.Message = &PeerRequest_Beacon{msg}
		return true, err
	case 11: // message.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrepareRequest)
		err := b.DecodeMessage(msg)
		m.Message = &PeerRequest_Prepare{msg}
		return true, err
	case 12: // message.accept
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AcceptRequest)
		err := b.DecodeMessage(msg)
		m.Message = &PeerRequest_Accept{msg}
		return true, err
	case 13: // message.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CommitRequest)
		err := b.DecodeMessage(msg)
		m.Message = &PeerRequest_Commit{msg}
		return true, err
	case 14: // message.tryPreAcccept
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TryPreAcceptRequest)
		err := b.DecodeMessage(msg)
		m.Message = &PeerRequest_TryPreAcccept{msg}
		return true, err
	case 15: // message.preAccept
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PreAcceptRequest)
		err := b.DecodeMessage(msg)
		m.Message = &PeerRequest_PreAccept{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PeerRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PeerRequest)
	// message
	switch x := m.Message.(type) {
	case *PeerRequest_Beacon:
		s := proto.Size(x.Beacon)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerRequest_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerRequest_Accept:
		s := proto.Size(x.Accept)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerRequest_Commit:
		s := proto.Size(x.Commit)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerRequest_TryPreAcccept:
		s := proto.Size(x.TryPreAcccept)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerRequest_PreAccept:
		s := proto.Size(x.PreAccept)
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A wrapper message that can contain one of the reply message types.
type PeerReply struct {
	Type    Type   `protobuf:"varint,1,opt,name=type,enum=pb.Type" json:"type,omitempty"`
	Sender  string `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Success bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	// only one of these fields can be set, and the field that is set should
	// match the message type described above.
	//
	// Types that are valid to be assigned to Message:
	//	*PeerReply_Beacon
	//	*PeerReply_Prepare
	//	*PeerReply_Accept
	//	*PeerReply_TryPreAccept
	//	*PeerReply_PreAccept
	Message isPeerReply_Message `protobuf_oneof:"message"`
}

func (m *PeerReply) Reset()                    { *m = PeerReply{} }
func (m *PeerReply) String() string            { return proto.CompactTextString(m) }
func (*PeerReply) ProtoMessage()               {}
func (*PeerReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type isPeerReply_Message interface {
	isPeerReply_Message()
}

type PeerReply_Beacon struct {
	Beacon *BeaconReply `protobuf:"bytes,10,opt,name=beacon,oneof"`
}
type PeerReply_Prepare struct {
	Prepare *PrepareReply `protobuf:"bytes,11,opt,name=prepare,oneof"`
}
type PeerReply_Accept struct {
	Accept *AcceptReply `protobuf:"bytes,12,opt,name=accept,oneof"`
}
type PeerReply_TryPreAccept struct {
	TryPreAccept *TryPreAcceptReply `protobuf:"bytes,14,opt,name=tryPreAccept,oneof"`
}
type PeerReply_PreAccept struct {
	PreAccept *PreAcceptReply `protobuf:"bytes,15,opt,name=preAccept,oneof"`
}

func (*PeerReply_Beacon) isPeerReply_Message()       {}
func (*PeerReply_Prepare) isPeerReply_Message()      {}
func (*PeerReply_Accept) isPeerReply_Message()       {}
func (*PeerReply_TryPreAccept) isPeerReply_Message() {}
func (*PeerReply_PreAccept) isPeerReply_Message()    {}

func (m *PeerReply) GetMessage() isPeerReply_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PeerReply) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_BEACON
}

func (m *PeerReply) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PeerReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PeerReply) GetBeacon() *BeaconReply {
	if x, ok := m.GetMessage().(*PeerReply_Beacon); ok {
		return x.Beacon
	}
	return nil
}

func (m *PeerReply) GetPrepare() *PrepareReply {
	if x, ok := m.GetMessage().(*PeerReply_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *PeerReply) GetAccept() *AcceptReply {
	if x, ok := m.GetMessage().(*PeerReply_Accept); ok {
		return x.Accept
	}
	return nil
}

func (m *PeerReply) GetTryPreAccept() *TryPreAcceptReply {
	if x, ok := m.GetMessage().(*PeerReply_TryPreAccept); ok {
		return x.TryPreAccept
	}
	return nil
}

func (m *PeerReply) GetPreAccept() *PreAcceptReply {
	if x, ok := m.GetMessage().(*PeerReply_PreAccept); ok {
		return x.PreAccept
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PeerReply) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PeerReply_OneofMarshaler, _PeerReply_OneofUnmarshaler, _PeerReply_OneofSizer, []interface{}{
		(*PeerReply_Beacon)(nil),
		(*PeerReply_Prepare)(nil),
		(*PeerReply_Accept)(nil),
		(*PeerReply_TryPreAccept)(nil),
		(*PeerReply_PreAccept)(nil),
	}
}

func _PeerReply_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PeerReply)
	// message
	switch x := m.Message.(type) {
	case *PeerReply_Beacon:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Beacon); err != nil {
			return err
		}
	case *PeerReply_Prepare:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *PeerReply_Accept:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Accept); err != nil {
			return err
		}
	case *PeerReply_TryPreAccept:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TryPreAccept); err != nil {
			return err
		}
	case *PeerReply_PreAccept:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PreAccept); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PeerReply.Message has unexpected type %T", x)
	}
	return nil
}

func _PeerReply_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PeerReply)
	switch tag {
	case 10: // message.beacon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BeaconReply)
		err := b.DecodeMessage(msg)
		m.Message = &PeerReply_Beacon{msg}
		return true, err
	case 11: // message.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrepareReply)
		err := b.DecodeMessage(msg)
		m.Message = &PeerReply_Prepare{msg}
		return true, err
	case 12: // message.accept
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AcceptReply)
		err := b.DecodeMessage(msg)
		m.Message = &PeerReply_Accept{msg}
		return true, err
	case 14: // message.tryPreAccept
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TryPreAcceptReply)
		err := b.DecodeMessage(msg)
		m.Message = &PeerReply_TryPreAccept{msg}
		return true, err
	case 15: // message.preAccept
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PreAcceptReply)
		err := b.DecodeMessage(msg)
		m.Message = &PeerReply_PreAccept{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PeerReply_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PeerReply)
	// message
	switch x := m.Message.(type) {
	case *PeerReply_Beacon:
		s := proto.Size(x.Beacon)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerReply_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerReply_Accept:
		s := proto.Size(x.Accept)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerReply_TryPreAccept:
		s := proto.Size(x.TryPreAccept)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PeerReply_PreAccept:
		s := proto.Size(x.PreAccept)
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*PeerRequest)(nil), "pb.PeerRequest")
	proto.RegisterType((*PeerReply)(nil), "pb.PeerReply")
	proto.RegisterEnum("pb.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x8e, 0x94, 0x40,
	0x14, 0x85, 0x07, 0xa6, 0x85, 0xe1, 0x42, 0x77, 0x63, 0xc5, 0x9f, 0xca, 0xc4, 0x05, 0x99, 0x15,
	0xa3, 0x86, 0x45, 0xeb, 0xce, 0x85, 0xa1, 0x09, 0x89, 0x2e, 0xc6, 0x21, 0x15, 0x16, 0xba, 0x04,
	0xbc, 0x31, 0x26, 0xdd, 0x4d, 0x49, 0x31, 0x46, 0x9e, 0xd5, 0x57, 0x71, 0x61, 0xa8, 0x2a, 0xe4,
	0x67, 0x5c, 0xcd, 0x8e, 0x7b, 0xea, 0x3b, 0xb7, 0x52, 0xe7, 0x04, 0xd8, 0xe0, 0xe9, 0x27, 0x1e,
	0x6a, 0x8e, 0x11, 0x6f, 0xea, 0xb6, 0x26, 0x26, 0x2f, 0x2f, 0xbd, 0x12, 0x8b, 0xaa, 0x3e, 0x29,
	0xe5, 0xd2, 0x43, 0x5e, 0xfc, 0xaa, 0x85, 0x9a, 0xae, 0xfe, 0x98, 0xe0, 0x66, 0x88, 0x0d, 0xc3,
	0x1f, 0x77, 0x28, 0x5a, 0xf2, 0x02, 0x56, 0x6d, 0xc7, 0x91, 0x1a, 0x81, 0x11, 0x6e, 0x76, 0x17,
	0x11, 0x2f, 0xa3, 0xbc, 0xe3, 0xc8, 0xa4, 0x4a, 0x9e, 0x81, 0x25, 0xf0, 0xf4, 0x15, 0x1b, 0x6a,
	0x06, 0x46, 0xe8, 0x30, 0x3d, 0x91, 0x57, 0x60, 0xa9, 0x3b, 0x28, 0x04, 0x46, 0xe8, 0xee, 0x1e,
	0xf7, 0xbe, 0xbd, 0x54, 0xf4, 0xe2, 0x0f, 0x67, 0x4c, 0x23, 0x24, 0x02, 0x9b, 0x37, 0xc8, 0x8b,
	0x06, 0xa9, 0x2b, 0x69, 0xd2, 0xd3, 0x99, 0x92, 0x46, 0x7c, 0x80, 0xfa, 0xe5, 0x45, 0x55, 0x21,
	0x6f, 0xa9, 0x37, 0x2e, 0x8f, 0xa5, 0x32, 0x59, 0xae, 0x90, 0x1e, 0xae, 0xea, 0xe3, 0xf1, 0x7b,
	0x4b, 0xd7, 0x23, 0x9c, 0x48, 0x65, 0x02, 0x2b, 0x84, 0xbc, 0x87, 0x75, 0xdb, 0x74, 0x59, 0x83,
	0x71, 0xa5, 0x2e, 0xd8, 0x48, 0xcf, 0x73, 0xf9, 0xea, 0xe1, 0x60, 0x76, 0xcd, 0x9c, 0x27, 0x6f,
	0xc1, 0xe1, 0x03, 0x44, 0xb7, 0xd2, 0xfc, 0x44, 0x3f, 0x66, 0xe9, 0x1c, 0xc1, 0xbd, 0x03, 0xf6,
	0x11, 0x85, 0x28, 0xbe, 0xe1, 0xd5, 0x6f, 0x13, 0x1c, 0x15, 0x3f, 0x3f, 0x74, 0x0f, 0x0c, 0x9f,
	0x82, 0x2d, 0xee, 0xaa, 0x0a, 0x85, 0xa0, 0xe7, 0x81, 0x11, 0x5e, 0xb0, 0x61, 0x24, 0xd7, 0x8b,
	0x5a, 0xb6, 0xd3, 0x5a, 0xf8, 0xa1, 0x9b, 0x94, 0xf2, 0x7a, 0x59, 0x8a, 0x3f, 0x2b, 0x45, 0xc1,
	0xff, 0x2a, 0xb9, 0x5e, 0x54, 0xb2, 0x9d, 0x56, 0xa2, 0x17, 0xeb, 0x42, 0xde, 0x81, 0xd7, 0x4e,
	0xa2, 0xd4, 0x11, 0x3f, 0xbd, 0x1f, 0xb1, 0xb2, 0xcd, 0x60, 0xb2, 0xbb, 0x9f, 0x2f, 0x59, 0xe4,
	0xab, 0x6c, 0xff, 0x4d, 0xf7, 0xe5, 0x67, 0x58, 0xf5, 0xf9, 0x11, 0x00, 0x6b, 0x9f, 0xc6, 0xc9,
	0xed, 0x27, 0xff, 0x8c, 0xb8, 0x60, 0x67, 0x2c, 0xcd, 0x62, 0x96, 0xfa, 0x46, 0x7f, 0x10, 0x27,
	0x49, 0x9a, 0xe5, 0xbe, 0xd9, 0x7f, 0x27, 0xb7, 0x37, 0x37, 0x1f, 0x73, 0xff, 0x9c, 0xf8, 0xe0,
	0xe5, 0xec, 0x4b, 0xc6, 0x52, 0x7d, 0xba, 0x22, 0x6b, 0x70, 0xc6, 0xf1, 0x51, 0x69, 0xc9, 0xbf,
	0xe7, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x86, 0x0d, 0xee, 0x6f, 0x03, 0x00, 0x00,
}
