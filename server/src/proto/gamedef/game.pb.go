// Code generated by protoc-gen-go.
// source: game.proto
// DO NOT EDIT!

package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VerifyGameResult int32

const (
	VerifyGameResult_VerifyOK       VerifyGameResult = 0
	VerifyGameResult_DataException  VerifyGameResult = 1
	VerifyGameResult_AccountBlocked VerifyGameResult = 2
)

var VerifyGameResult_name = map[int32]string{
	0: "VerifyOK",
	1: "DataException",
	2: "AccountBlocked",
}
var VerifyGameResult_value = map[string]int32{
	"VerifyOK":       0,
	"DataException":  1,
	"AccountBlocked": 2,
}

func (x VerifyGameResult) String() string {
	return proto.EnumName(VerifyGameResult_name, int32(x))
}
func (VerifyGameResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// 创建角色结果
type CreateCharResult int32

const (
	CreateCharResult_CreateCharOK    CreateCharResult = 0
	CreateCharResult_CharNotExist    CreateCharResult = 1
	CreateCharResult_CreateCharError CreateCharResult = 2
)

var CreateCharResult_name = map[int32]string{
	0: "CreateCharOK",
	1: "CharNotExist",
	2: "CreateCharError",
}
var CreateCharResult_value = map[string]int32{
	"CreateCharOK":    0,
	"CharNotExist":    1,
	"CreateCharError": 2,
}

func (x CreateCharResult) String() string {
	return proto.EnumName(CreateCharResult_name, int32(x))
}
func (CreateCharResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type EnterGameResult int32

const (
	EnterGameResult_EnterGameOK  EnterGameResult = 0
	EnterGameResult_GameUserFull EnterGameResult = 1
	EnterGameResult_CharBlocked  EnterGameResult = 2
)

var EnterGameResult_name = map[int32]string{
	0: "EnterGameOK",
	1: "GameUserFull",
	2: "CharBlocked",
}
var EnterGameResult_value = map[string]int32{
	"EnterGameOK":  0,
	"GameUserFull": 1,
	"CharBlocked":  2,
}

func (x EnterGameResult) String() string {
	return proto.EnumName(EnterGameResult_name, int32(x))
}
func (EnterGameResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// client -> game
type VerifyGameREQ struct {
	Token string `protobuf:"bytes,1,opt,name=Token,json=token" json:"Token,omitempty"`
}

func (m *VerifyGameREQ) Reset()                    { *m = VerifyGameREQ{} }
func (m *VerifyGameREQ) String() string            { return proto.CompactTextString(m) }
func (*VerifyGameREQ) ProtoMessage()               {}
func (*VerifyGameREQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// game -> client
type VerifyGameACK struct {
	Result VerifyGameResult `protobuf:"varint,1,opt,name=Result,json=result,enum=gamedef.VerifyGameResult" json:"Result,omitempty"`
}

func (m *VerifyGameACK) Reset()                    { *m = VerifyGameACK{} }
func (m *VerifyGameACK) String() string            { return proto.CompactTextString(m) }
func (*VerifyGameACK) ProtoMessage()               {}
func (*VerifyGameACK) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// 初始角色信息
type SimpleCharInfo struct {
	ModelID      int64  `protobuf:"varint,1,opt,name=ModelID,json=modelID" json:"ModelID,omitempty"`
	CharName     string `protobuf:"bytes,2,opt,name=CharName,json=charName" json:"CharName,omitempty"`
	LastLoginUTC int64  `protobuf:"varint,3,opt,name=LastLoginUTC,json=lastLoginUTC" json:"LastLoginUTC,omitempty"`
}

func (m *SimpleCharInfo) Reset()                    { *m = SimpleCharInfo{} }
func (m *SimpleCharInfo) String() string            { return proto.CompactTextString(m) }
func (*SimpleCharInfo) ProtoMessage()               {}
func (*SimpleCharInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// 请求角色列表
// client -> game
type CharListREQ struct {
}

func (m *CharListREQ) Reset()                    { *m = CharListREQ{} }
func (m *CharListREQ) String() string            { return proto.CompactTextString(m) }
func (*CharListREQ) ProtoMessage()               {}
func (*CharListREQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// game -> client
type CharListACK struct {
}

func (m *CharListACK) Reset()                    { *m = CharListACK{} }
func (m *CharListACK) String() string            { return proto.CompactTextString(m) }
func (*CharListACK) ProtoMessage()               {}
func (*CharListACK) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// 创建角色
// client -> game
type CreateCharREQ struct {
	CandidateID int32  `protobuf:"varint,1,opt,name=CandidateID,json=candidateID" json:"CandidateID,omitempty"`
	CharName    string `protobuf:"bytes,2,opt,name=CharName,json=charName" json:"CharName,omitempty"`
}

func (m *CreateCharREQ) Reset()                    { *m = CreateCharREQ{} }
func (m *CreateCharREQ) String() string            { return proto.CompactTextString(m) }
func (*CreateCharREQ) ProtoMessage()               {}
func (*CreateCharREQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// 创建角色结果
// game -> client
type CreateCharACK struct {
	CandidateID int32            `protobuf:"varint,1,opt,name=CandidateID,json=candidateID" json:"CandidateID,omitempty"`
	Result      CreateCharResult `protobuf:"varint,2,opt,name=Result,json=result,enum=gamedef.CreateCharResult" json:"Result,omitempty"`
}

func (m *CreateCharACK) Reset()                    { *m = CreateCharACK{} }
func (m *CreateCharACK) String() string            { return proto.CompactTextString(m) }
func (*CreateCharACK) ProtoMessage()               {}
func (*CreateCharACK) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// 进入游戏
// client -> game
type EnterGameREQ struct {
	CharName string `protobuf:"bytes,1,opt,name=CharName,json=charName" json:"CharName,omitempty"`
}

func (m *EnterGameREQ) Reset()                    { *m = EnterGameREQ{} }
func (m *EnterGameREQ) String() string            { return proto.CompactTextString(m) }
func (*EnterGameREQ) ProtoMessage()               {}
func (*EnterGameREQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

// 进入游戏信息
// game -> client
type EnterGameACK struct {
	Result EnterGameResult `protobuf:"varint,1,opt,name=Result,json=result,enum=gamedef.EnterGameResult" json:"Result,omitempty"`
}

func (m *EnterGameACK) Reset()                    { *m = EnterGameACK{} }
func (m *EnterGameACK) String() string            { return proto.CompactTextString(m) }
func (*EnterGameACK) ProtoMessage()               {}
func (*EnterGameACK) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func init() {
	proto.RegisterType((*VerifyGameREQ)(nil), "gamedef.VerifyGameREQ")
	proto.RegisterType((*VerifyGameACK)(nil), "gamedef.VerifyGameACK")
	proto.RegisterType((*SimpleCharInfo)(nil), "gamedef.SimpleCharInfo")
	proto.RegisterType((*CharListREQ)(nil), "gamedef.CharListREQ")
	proto.RegisterType((*CharListACK)(nil), "gamedef.CharListACK")
	proto.RegisterType((*CreateCharREQ)(nil), "gamedef.CreateCharREQ")
	proto.RegisterType((*CreateCharACK)(nil), "gamedef.CreateCharACK")
	proto.RegisterType((*EnterGameREQ)(nil), "gamedef.EnterGameREQ")
	proto.RegisterType((*EnterGameACK)(nil), "gamedef.EnterGameACK")
	proto.RegisterEnum("gamedef.VerifyGameResult", VerifyGameResult_name, VerifyGameResult_value)
	proto.RegisterEnum("gamedef.CreateCharResult", CreateCharResult_name, CreateCharResult_value)
	proto.RegisterEnum("gamedef.EnterGameResult", EnterGameResult_name, EnterGameResult_value)
}

func init() { proto.RegisterFile("game.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0x5d, 0x4f, 0xfa, 0x30,
	0x14, 0xc6, 0xff, 0x83, 0xf0, 0xf2, 0x3f, 0x6c, 0x6c, 0x56, 0x2f, 0xa6, 0x57, 0x64, 0x89, 0x89,
	0xe1, 0x82, 0xf8, 0xf2, 0x05, 0x84, 0x31, 0x09, 0x01, 0x34, 0x4e, 0xf0, 0x7e, 0x6e, 0x05, 0x27,
	0x63, 0x25, 0x5d, 0x49, 0xf0, 0xdb, 0xdb, 0x96, 0x97, 0x95, 0x5d, 0xe8, 0x15, 0x3d, 0x0f, 0x4f,
	0x9f, 0xfe, 0xce, 0x39, 0x19, 0xc0, 0x22, 0x58, 0xe1, 0xce, 0x9a, 0x12, 0x46, 0x50, 0x4d, 0x9c,
	0x23, 0x3c, 0x77, 0xae, 0xc1, 0x78, 0xc7, 0x34, 0x9e, 0x7f, 0x0f, 0xb8, 0xe0, 0x7b, 0xaf, 0xe8,
	0x02, 0x2a, 0x53, 0xb2, 0xc4, 0xa9, 0xad, 0xb5, 0xb4, 0x9b, 0xff, 0x7e, 0x85, 0x89, 0xc2, 0xe9,
	0xa9, 0xb6, 0xae, 0x3b, 0x42, 0x77, 0x50, 0xf5, 0x71, 0xb6, 0x49, 0x98, 0xf4, 0x35, 0xef, 0x2f,
	0x3b, 0xfb, 0xc4, 0x8e, 0x12, 0x27, 0x0d, 0x7e, 0x95, 0xca, 0x5f, 0xe7, 0x0b, 0x9a, 0x6f, 0xf1,
	0x6a, 0x9d, 0x60, 0xf7, 0x33, 0xa0, 0xc3, 0x74, 0x4e, 0x90, 0x0d, 0xb5, 0x09, 0x89, 0x70, 0x32,
	0xec, 0xcb, 0x94, 0xb2, 0x5f, 0x5b, 0xed, 0x4a, 0x74, 0x05, 0x75, 0xe1, 0x7a, 0xe6, 0x29, 0x76,
	0x49, 0x82, 0xd4, 0xc3, 0x7d, 0x8d, 0x1c, 0xd0, 0xc7, 0x41, 0xc6, 0xc6, 0x64, 0x11, 0xa7, 0xb3,
	0xa9, 0x6b, 0x97, 0xe5, 0x55, 0x3d, 0x51, 0x34, 0xc7, 0x80, 0x86, 0xb8, 0x3f, 0x8e, 0x33, 0xc6,
	0x9b, 0x52, 0x4b, 0x0e, 0xef, 0x4c, 0xc0, 0x70, 0x29, 0x0e, 0x98, 0x24, 0x11, 0x4d, 0xb7, 0xf8,
	0xff, 0x41, 0x1a, 0xc5, 0x11, 0xd7, 0xf6, 0x30, 0x15, 0xbf, 0x11, 0xe6, 0xd2, 0x6f, 0x40, 0x4e,
	0xa4, 0xc6, 0x89, 0xe1, 0xfc, 0x1d, 0x97, 0x8f, 0xaf, 0x54, 0x18, 0x9f, 0x02, 0x76, 0x3a, 0xbe,
	0x36, 0xe8, 0x5e, 0xca, 0x30, 0x3d, 0x2c, 0x4a, 0x25, 0xd2, 0x0a, 0x44, 0x8f, 0x8a, 0x57, 0x00,
	0xdd, 0x16, 0xb6, 0x65, 0x1f, 0x9f, 0xcb, 0x23, 0x4f, 0x5e, 0x6b, 0x0f, 0xc0, 0x2a, 0x2e, 0x12,
	0xe9, 0x50, 0xdf, 0x69, 0x2f, 0x23, 0xeb, 0x1f, 0x3a, 0x03, 0xa3, 0x1f, 0xb0, 0xc0, 0xdb, 0x86,
	0x78, 0xcd, 0x62, 0x92, 0x5a, 0x1a, 0x42, 0xd0, 0xec, 0x86, 0x21, 0xd9, 0xa4, 0xac, 0x97, 0x90,
	0x70, 0x89, 0x23, 0xab, 0xd4, 0x1e, 0x81, 0x55, 0x6c, 0x09, 0x59, 0xa0, 0xe7, 0x9a, 0x0c, 0x13,
	0x8a, 0x80, 0x27, 0xcc, 0xdb, 0xf2, 0x25, 0xf1, 0xac, 0x73, 0x30, 0x73, 0x8f, 0x47, 0x29, 0xa1,
	0x3c, 0xcc, 0x03, 0xb3, 0x00, 0x8c, 0x4c, 0x68, 0x1c, 0xa5, 0x43, 0x94, 0x38, 0xcf, 0x32, 0x4c,
	0x9f, 0x36, 0x49, 0xc2, 0xa3, 0xcc, 0xdd, 0xf6, 0x8f, 0x4c, 0x1f, 0x55, 0xf9, 0x11, 0x3c, 0xfc,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x91, 0x9d, 0xae, 0x94, 0x12, 0x03, 0x00, 0x00,
}
