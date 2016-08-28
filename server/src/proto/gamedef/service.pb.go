// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package gamedef is a generated protocol buffer package.

It is generated from these files:
	service.proto
	game.proto
	login.proto
	model.proto
	behavior.proto
	router.proto
	char.proto
	tool.proto

It has these top-level messages:
	ServiceDefine
	ServiceFile
	LocalFile
	VerifyGameREQ
	VerifyGameACK
	SimpleCharInfo
	CharListREQ
	CharListACK
	CreateCharREQ
	CreateCharACK
	EnterGameREQ
	EnterGameACK
	LoginREQ
	ServerInfo
	LoginACK
	LoginModel
	ModelACK
	ModelValue
	ModelSyncACK
	BehaviorDefine
	BehaviorFile
	UpstreamACK
	CloseClientACK
	DownstreamACK
	RegisterRouterBackendACK
	CharModel
	TableCodeOption
*/
package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 一个服务的配置类型
type ServiceDefine struct {
	MainConfig bool `protobuf:"varint,10,opt,name=MainConfig,json=mainConfig" json:"MainConfig,omitempty"`
	// Peer配置
	IP        string `protobuf:"bytes,11,opt,name=IP,json=iP" json:"IP,omitempty"`
	Port      uint32 `protobuf:"varint,12,opt,name=Port,json=port" json:"Port,omitempty"`
	PeerName  string `protobuf:"bytes,13,opt,name=PeerName,json=peerName" json:"PeerName,omitempty"`
	PeerIndex int32  `protobuf:"varint,14,opt,name=PeerIndex,json=peerIndex" json:"PeerIndex,omitempty"`
	Enable    bool   `protobuf:"varint,15,opt,name=Enable,json=enable" json:"Enable,omitempty"`
	// 基础服务
	Version     string `protobuf:"bytes,201,opt,name=Version,json=version" json:"Version,omitempty"`
	Name        string `protobuf:"bytes,205,opt,name=Name,json=name" json:"Name,omitempty"`
	DisplayName string `protobuf:"bytes,206,opt,name=DisplayName,json=displayName" json:"DisplayName,omitempty"`
	DSN         string `protobuf:"bytes,210,opt,name=DSN,json=dSN" json:"DSN,omitempty"`
	DBConnCount int32  `protobuf:"varint,211,opt,name=DBConnCount,json=dBConnCount" json:"DBConnCount,omitempty"`
}

func (m *ServiceDefine) Reset()                    { *m = ServiceDefine{} }
func (m *ServiceDefine) String() string            { return proto.CompactTextString(m) }
func (*ServiceDefine) ProtoMessage()               {}
func (*ServiceDefine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 服务配置, 所有服务通用
type ServiceFile struct {
	Service []*ServiceDefine `protobuf:"bytes,1,rep,name=Service,json=service" json:"Service,omitempty"`
}

func (m *ServiceFile) Reset()                    { *m = ServiceFile{} }
func (m *ServiceFile) String() string            { return proto.CompactTextString(m) }
func (*ServiceFile) ProtoMessage()               {}
func (*ServiceFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceFile) GetService() []*ServiceDefine {
	if m != nil {
		return m.Service
	}
	return nil
}

// 对ServiceDefine配置的选择
type LocalFile struct {
	ServiceConfig string `protobuf:"bytes,1,opt,name=ServiceConfig,json=serviceConfig" json:"ServiceConfig,omitempty"`
}

func (m *LocalFile) Reset()                    { *m = LocalFile{} }
func (m *LocalFile) String() string            { return proto.CompactTextString(m) }
func (*LocalFile) ProtoMessage()               {}
func (*LocalFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ServiceDefine)(nil), "gamedef.ServiceDefine")
	proto.RegisterType((*ServiceFile)(nil), "gamedef.ServiceFile")
	proto.RegisterType((*LocalFile)(nil), "gamedef.LocalFile")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0x56, 0xd7, 0xf5, 0xc4, 0x4e, 0x8c, 0x30, 0xa2, 0x88, 0xcc, 0xe1, 0xc5, 0xae,
	0x86, 0x7f, 0x1e, 0x40, 0xd0, 0x29, 0x0c, 0xb4, 0x94, 0x16, 0xbc, 0xef, 0xd6, 0xd3, 0x11, 0xe8,
	0x92, 0xd2, 0xd6, 0xa1, 0x2f, 0xe5, 0x7b, 0x78, 0xa1, 0x17, 0xfa, 0x42, 0x66, 0x49, 0xb6, 0xe1,
	0x5d, 0xbf, 0xdf, 0xd7, 0xe4, 0x3b, 0xf9, 0x0e, 0x04, 0x35, 0x56, 0x2b, 0x3e, 0xc7, 0x71, 0x59,
	0xc9, 0x46, 0x52, 0x6f, 0x91, 0x2e, 0x31, 0xc3, 0x7c, 0xf8, 0xd1, 0x82, 0x20, 0x31, 0xd6, 0x04,
	0x73, 0x2e, 0x90, 0x9e, 0x01, 0x3c, 0xa7, 0x5c, 0xdc, 0x4b, 0x91, 0xf3, 0x05, 0x83, 0x81, 0x33,
	0xea, 0xc6, 0xb0, 0xdc, 0x12, 0xda, 0x83, 0xd6, 0x34, 0x62, 0x44, 0x71, 0x3f, 0x6e, 0xf1, 0x88,
	0x52, 0x70, 0x23, 0x59, 0x35, 0x6c, 0x5f, 0x91, 0x20, 0x76, 0x4b, 0xf5, 0x4d, 0x4f, 0xa0, 0x1b,
	0x21, 0x56, 0xa1, 0x0a, 0x61, 0x81, 0xfe, 0xb3, 0x5b, 0x5a, 0x4d, 0x4f, 0xc1, 0x5f, 0x7b, 0x53,
	0x91, 0xe1, 0x1b, 0xeb, 0x29, 0x73, 0x2f, 0xf6, 0xcb, 0x0d, 0xa0, 0x7d, 0xe8, 0x3c, 0x88, 0x74,
	0x56, 0x20, 0x3b, 0xd0, 0xc9, 0x1d, 0xd4, 0x8a, 0x1e, 0x83, 0xf7, 0x82, 0x55, 0xcd, 0xa5, 0x60,
	0x9f, 0x8e, 0xbe, 0xd1, 0x5b, 0x19, 0x4d, 0x8f, 0xc0, 0xd5, 0x41, 0x5f, 0x86, 0xbb, 0x62, 0x9d,
	0x72, 0x0e, 0x64, 0xc2, 0xeb, 0xb2, 0x48, 0xdf, 0xb5, 0xf7, 0x6d, 0x3c, 0x92, 0xed, 0x18, 0x3d,
	0x84, 0xf6, 0x24, 0x09, 0xd9, 0x8f, 0xb1, 0xda, 0x59, 0x12, 0xea, 0x53, 0x77, 0xea, 0x9d, 0xea,
	0xad, 0xaf, 0xa2, 0x61, 0xbf, 0x8e, 0x1e, 0x8f, 0x64, 0x3b, 0x36, 0xbc, 0x05, 0x62, 0xfb, 0x7a,
	0xe4, 0x6a, 0xae, 0x4b, 0xf0, 0xac, 0x64, 0xce, 0xa0, 0x3d, 0x22, 0xd7, 0xfd, 0xb1, 0xad, 0x76,
	0xfc, 0xaf, 0xd6, 0xd8, 0xb3, 0x0b, 0x18, 0x5e, 0x81, 0xff, 0x24, 0xe7, 0x69, 0xa1, 0x8f, 0x5f,
	0x6c, 0xdb, 0xb7, 0x7d, 0x9b, 0x61, 0x36, 0xdb, 0x32, 0x70, 0xd6, 0xd1, 0x4b, 0xbb, 0xf9, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xe2, 0xb4, 0xb5, 0xeb, 0xc5, 0x01, 0x00, 0x00,
}
