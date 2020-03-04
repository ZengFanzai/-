// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exp_user1.proto

package foo_bar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type EnumAllowingAlias int32

const (
	EnumAllowingAlias_UNKNOWN EnumAllowingAlias = 0
	EnumAllowingAlias_STARTED EnumAllowingAlias = 1
	EnumAllowingAlias_RUNNING EnumAllowingAlias = 1
)

var EnumAllowingAlias_name = map[int32]string{
	0: "UNKNOWN",
	1: "STARTED",
	// Duplicate value: 1: "RUNNING",
}

var EnumAllowingAlias_value = map[string]int32{
	"UNKNOWN": 0,
	"STARTED": 1,
	"RUNNING": 1,
}

func (x EnumAllowingAlias) String() string {
	return proto.EnumName(EnumAllowingAlias_name, int32(x))
}

func (EnumAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{0}
}

type SearchRequest_Corpus int32

const (
	SearchRequest_UNIVERSAL SearchRequest_Corpus = 0
	SearchRequest_WEB       SearchRequest_Corpus = 1
	SearchRequest_IMAGES    SearchRequest_Corpus = 2
	SearchRequest_LOCAL     SearchRequest_Corpus = 3
	SearchRequest_NEWS      SearchRequest_Corpus = 4
	SearchRequest_PRODUCTS  SearchRequest_Corpus = 5
	SearchRequest_VIDEO     SearchRequest_Corpus = 6
)

var SearchRequest_Corpus_name = map[int32]string{
	0: "UNIVERSAL",
	1: "WEB",
	2: "IMAGES",
	3: "LOCAL",
	4: "NEWS",
	5: "PRODUCTS",
	6: "VIDEO",
}

var SearchRequest_Corpus_value = map[string]int32{
	"UNIVERSAL": 0,
	"WEB":       1,
	"IMAGES":    2,
	"LOCAL":     3,
	"NEWS":      4,
	"PRODUCTS":  5,
	"VIDEO":     6,
}

func (x SearchRequest_Corpus) String() string {
	return proto.EnumName(SearchRequest_Corpus_name, int32(x))
}

func (SearchRequest_Corpus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{3, 0}
}

type ErrorStatus struct {
	Message              string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	Test                 []byte     `protobuf:"bytes,3,opt,name=test,proto3" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ErrorStatus) Reset()         { *m = ErrorStatus{} }
func (m *ErrorStatus) String() string { return proto.CompactTextString(m) }
func (*ErrorStatus) ProtoMessage()    {}
func (*ErrorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{0}
}

func (m *ErrorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorStatus.Unmarshal(m, b)
}
func (m *ErrorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorStatus.Marshal(b, m, deterministic)
}
func (m *ErrorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorStatus.Merge(m, src)
}
func (m *ErrorStatus) XXX_Size() int {
	return xxx_messageInfo_ErrorStatus.Size(m)
}
func (m *ErrorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorStatus proto.InternalMessageInfo

func (m *ErrorStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorStatus) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *ErrorStatus) GetTest() []byte {
	if m != nil {
		return m.Test
	}
	return nil
}

type SampleMessage struct {
	// Types that are valid to be assigned to TestOneof:
	//	*SampleMessage_Name
	//	*SampleMessage_SubMessage
	TestOneof            isSampleMessage_TestOneof `protobuf_oneof:"test_oneof"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SampleMessage) Reset()         { *m = SampleMessage{} }
func (m *SampleMessage) String() string { return proto.CompactTextString(m) }
func (*SampleMessage) ProtoMessage()    {}
func (*SampleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{1}
}

func (m *SampleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleMessage.Unmarshal(m, b)
}
func (m *SampleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleMessage.Marshal(b, m, deterministic)
}
func (m *SampleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleMessage.Merge(m, src)
}
func (m *SampleMessage) XXX_Size() int {
	return xxx_messageInfo_SampleMessage.Size(m)
}
func (m *SampleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SampleMessage proto.InternalMessageInfo

type isSampleMessage_TestOneof interface {
	isSampleMessage_TestOneof()
}

type SampleMessage_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type SampleMessage_SubMessage struct {
	SubMessage *ErrorStatus `protobuf:"bytes,2,opt,name=sub_message,json=subMessage,proto3,oneof"`
}

func (*SampleMessage_Name) isSampleMessage_TestOneof() {}

func (*SampleMessage_SubMessage) isSampleMessage_TestOneof() {}

func (m *SampleMessage) GetTestOneof() isSampleMessage_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (m *SampleMessage) GetName() string {
	if x, ok := m.GetTestOneof().(*SampleMessage_Name); ok {
		return x.Name
	}
	return ""
}

func (m *SampleMessage) GetSubMessage() *ErrorStatus {
	if x, ok := m.GetTestOneof().(*SampleMessage_SubMessage); ok {
		return x.SubMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SampleMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SampleMessage_Name)(nil),
		(*SampleMessage_SubMessage)(nil),
	}
}

type Info struct {
	//field number range [1,2^29-1],不能使用[19000,19999],这些是保留字段号
	//1到15的范围需要一个字节进行编码
	//16打2047需要两个字节，对常出现的字段使用1到15
	Field1               string   `protobuf:"bytes,1,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2               string   `protobuf:"bytes,2,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3               int32    `protobuf:"varint,3,opt,name=field3,proto3" json:"field3,omitempty"`
	Field4               string   `protobuf:"bytes,4,opt,name=field4,proto3" json:"field4,omitempty"`
	Field5               string   `protobuf:"bytes,5,opt,name=field5,proto3" json:"field5,omitempty"`
	Field6               string   `protobuf:"bytes,6,opt,name=field6,proto3" json:"field6,omitempty"`
	Field7               string   `protobuf:"bytes,7,opt,name=field7,proto3" json:"field7,omitempty"`
	Field8               string   `protobuf:"bytes,8,opt,name=field8,proto3" json:"field8,omitempty"`
	Field9               string   `protobuf:"bytes,9,opt,name=field9,proto3" json:"field9,omitempty"`
	Field10              string   `protobuf:"bytes,10,opt,name=field10,proto3" json:"field10,omitempty"`
	Field11              []string `protobuf:"bytes,11,rep,name=field11,proto3" json:"field11,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{2}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *Info) GetField2() string {
	if m != nil {
		return m.Field2
	}
	return ""
}

func (m *Info) GetField3() int32 {
	if m != nil {
		return m.Field3
	}
	return 0
}

func (m *Info) GetField4() string {
	if m != nil {
		return m.Field4
	}
	return ""
}

func (m *Info) GetField5() string {
	if m != nil {
		return m.Field5
	}
	return ""
}

func (m *Info) GetField6() string {
	if m != nil {
		return m.Field6
	}
	return ""
}

func (m *Info) GetField7() string {
	if m != nil {
		return m.Field7
	}
	return ""
}

func (m *Info) GetField8() string {
	if m != nil {
		return m.Field8
	}
	return ""
}

func (m *Info) GetField9() string {
	if m != nil {
		return m.Field9
	}
	return ""
}

func (m *Info) GetField10() string {
	if m != nil {
		return m.Field10
	}
	return ""
}

func (m *Info) GetField11() []string {
	if m != nil {
		return m.Field11
	}
	return nil
}

type SearchRequest struct {
	Query                string               `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32                `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32                `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	Corpus               SearchRequest_Corpus `protobuf:"varint,4,opt,name=corpus,proto3,enum=foo.bar.SearchRequest_Corpus" json:"corpus,omitempty"`
	Enum                 EnumAllowingAlias    `protobuf:"varint,5,opt,name=enum,proto3,enum=foo.bar.EnumAllowingAlias" json:"enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{3}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func (m *SearchRequest) GetCorpus() SearchRequest_Corpus {
	if m != nil {
		return m.Corpus
	}
	return SearchRequest_UNIVERSAL
}

func (m *SearchRequest) GetEnum() EnumAllowingAlias {
	if m != nil {
		return m.Enum
	}
	return EnumAllowingAlias_UNKNOWN
}

type SearchRequest1 struct {
	Src                  SearchRequest_Corpus `protobuf:"varint,1,opt,name=src,proto3,enum=foo.bar.SearchRequest_Corpus" json:"src,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchRequest1) Reset()         { *m = SearchRequest1{} }
func (m *SearchRequest1) String() string { return proto.CompactTextString(m) }
func (*SearchRequest1) ProtoMessage()    {}
func (*SearchRequest1) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{4}
}

func (m *SearchRequest1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest1.Unmarshal(m, b)
}
func (m *SearchRequest1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest1.Marshal(b, m, deterministic)
}
func (m *SearchRequest1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest1.Merge(m, src)
}
func (m *SearchRequest1) XXX_Size() int {
	return xxx_messageInfo_SearchRequest1.Size(m)
}
func (m *SearchRequest1) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest1.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest1 proto.InternalMessageInfo

func (m *SearchRequest1) GetSrc() SearchRequest_Corpus {
	if m != nil {
		return m.Src
	}
	return SearchRequest_UNIVERSAL
}

type MapUse struct {
	Projects             map[string]EnumAllowingAlias `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=foo.bar.EnumAllowingAlias"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MapUse) Reset()         { *m = MapUse{} }
func (m *MapUse) String() string { return proto.CompactTextString(m) }
func (*MapUse) ProtoMessage()    {}
func (*MapUse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8dc0843561a3a97, []int{5}
}

func (m *MapUse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapUse.Unmarshal(m, b)
}
func (m *MapUse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapUse.Marshal(b, m, deterministic)
}
func (m *MapUse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapUse.Merge(m, src)
}
func (m *MapUse) XXX_Size() int {
	return xxx_messageInfo_MapUse.Size(m)
}
func (m *MapUse) XXX_DiscardUnknown() {
	xxx_messageInfo_MapUse.DiscardUnknown(m)
}

var xxx_messageInfo_MapUse proto.InternalMessageInfo

func (m *MapUse) GetProjects() map[string]EnumAllowingAlias {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterEnum("foo.bar.EnumAllowingAlias", EnumAllowingAlias_name, EnumAllowingAlias_value)
	proto.RegisterEnum("foo.bar.SearchRequest_Corpus", SearchRequest_Corpus_name, SearchRequest_Corpus_value)
	proto.RegisterType((*ErrorStatus)(nil), "foo.bar.ErrorStatus")
	proto.RegisterType((*SampleMessage)(nil), "foo.bar.SampleMessage")
	proto.RegisterType((*Info)(nil), "foo.bar.Info")
	proto.RegisterType((*SearchRequest)(nil), "foo.bar.SearchRequest")
	proto.RegisterType((*SearchRequest1)(nil), "foo.bar.SearchRequest1")
	proto.RegisterType((*MapUse)(nil), "foo.bar.MapUse")
	proto.RegisterMapType((map[string]EnumAllowingAlias)(nil), "foo.bar.MapUse.ProjectsEntry")
}

func init() {
	proto.RegisterFile("exp_user1.proto", fileDescriptor_b8dc0843561a3a97)
}

var fileDescriptor_b8dc0843561a3a97 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcb, 0x6e, 0xda, 0x4c,
	0x14, 0xc7, 0xf1, 0x15, 0x38, 0x84, 0xc4, 0xdf, 0x28, 0xfa, 0xe4, 0x22, 0x55, 0x45, 0x5e, 0x54,
	0x51, 0x17, 0x4e, 0x20, 0xf7, 0x2e, 0x2a, 0x39, 0x89, 0x95, 0xa0, 0x26, 0x06, 0x8d, 0x43, 0x90,
	0xba, 0x41, 0x86, 0x0c, 0x94, 0xc6, 0xd8, 0xce, 0x8c, 0xdd, 0x96, 0xd7, 0xe8, 0xba, 0x2f, 0xd1,
	0x37, 0xac, 0x3c, 0x36, 0xc6, 0x51, 0xa5, 0x76, 0xe7, 0xf3, 0x3b, 0x7f, 0x9d, 0xcb, 0xff, 0x68,
	0x0c, 0x3b, 0xe4, 0x7b, 0x34, 0x4e, 0x18, 0xa1, 0x1d, 0x33, 0xa2, 0x61, 0x1c, 0xa2, 0xea, 0x2c,
	0x0c, 0xcd, 0x89, 0x47, 0x5b, 0xaf, 0xe6, 0x61, 0x38, 0xf7, 0xc9, 0x3e, 0xc7, 0x93, 0x64, 0xb6,
	0xef, 0x05, 0xab, 0x4c, 0x63, 0x3c, 0x41, 0xc3, 0xa6, 0x34, 0xa4, 0x6e, 0xec, 0xc5, 0x09, 0x43,
	0x3a, 0x54, 0x97, 0x84, 0x31, 0x6f, 0x4e, 0x74, 0xa1, 0x2d, 0xec, 0xd5, 0xf1, 0x3a, 0x44, 0x26,
	0x54, 0x1f, 0x49, 0xec, 0x2d, 0x7c, 0xa6, 0x8b, 0x6d, 0x69, 0xaf, 0xd1, 0xdd, 0x35, 0xb3, 0xaa,
	0xe6, 0xba, 0xaa, 0x69, 0x05, 0x2b, 0xbc, 0x16, 0x21, 0x04, 0x72, 0x4c, 0x58, 0xac, 0x4b, 0x6d,
	0x61, 0x6f, 0x0b, 0xf3, 0x6f, 0xc3, 0x87, 0xa6, 0xeb, 0x2d, 0x23, 0x9f, 0xdc, 0xe5, 0x45, 0x77,
	0x41, 0x0e, 0xbc, 0x65, 0xde, 0xeb, 0xa6, 0x82, 0x79, 0x84, 0x4e, 0xa1, 0xc1, 0x92, 0xc9, 0x78,
	0x3d, 0x88, 0xd8, 0x16, 0x78, 0xbb, 0x7c, 0x1b, 0xb3, 0x34, 0xef, 0x4d, 0x05, 0x03, 0x4b, 0x26,
	0x79, 0xb9, 0x8b, 0x2d, 0x80, 0xb4, 0xcf, 0x38, 0x0c, 0x48, 0x38, 0x33, 0x7e, 0x88, 0x20, 0xf7,
	0x82, 0x59, 0x88, 0xfe, 0x07, 0x75, 0xb6, 0x20, 0xfe, 0x63, 0x27, 0xdf, 0x29, 0x8f, 0x0a, 0xde,
	0xe5, 0x2d, 0xd6, 0xbc, 0x5b, 0xf0, 0x43, 0x3e, 0xbc, 0x92, 0xf3, 0xc3, 0x82, 0x1f, 0xe9, 0x72,
	0x49, 0x7f, 0x54, 0xf0, 0x63, 0x5d, 0x29, 0xf1, 0xe3, 0x82, 0x9f, 0xe8, 0x6a, 0x89, 0x9f, 0x14,
	0xfc, 0x54, 0xaf, 0x96, 0xf8, 0x69, 0xc1, 0xcf, 0xf4, 0x5a, 0x89, 0x9f, 0x15, 0xfc, 0x5c, 0xaf,
	0x97, 0xf8, 0x79, 0x7a, 0xac, 0x6c, 0x93, 0x03, 0x1d, 0xb2, 0x63, 0xe5, 0xe1, 0x26, 0xd3, 0xd1,
	0x1b, 0x6d, 0x69, 0x93, 0xe9, 0x18, 0xbf, 0x44, 0x68, 0xba, 0xc4, 0xa3, 0xd3, 0xcf, 0x98, 0x3c,
	0x27, 0x84, 0xc5, 0x68, 0x17, 0x94, 0xe7, 0x84, 0xd0, 0x55, 0x6e, 0x4e, 0x16, 0xa0, 0x37, 0xd0,
	0x88, 0xbc, 0x39, 0x19, 0x07, 0xc9, 0x72, 0x42, 0x28, 0x37, 0x48, 0xc1, 0x90, 0x22, 0x87, 0x13,
	0xf4, 0x16, 0x76, 0x28, 0x61, 0x89, 0x1f, 0x8f, 0x23, 0x42, 0xc7, 0x69, 0x22, 0x77, 0xab, 0x99,
	0xe1, 0x01, 0xa1, 0x83, 0xf4, 0xc4, 0xc7, 0xa0, 0x4e, 0x43, 0x1a, 0x25, 0x8c, 0x9b, 0xb6, 0xdd,
	0x7d, 0x5d, 0xdc, 0xf1, 0xc5, 0x18, 0xe6, 0x25, 0x17, 0xe1, 0x5c, 0x8c, 0x4c, 0x90, 0x49, 0x90,
	0x2c, 0xb9, 0xa3, 0xdb, 0xdd, 0xd6, 0xe6, 0xf8, 0x41, 0xb2, 0xb4, 0x7c, 0x3f, 0xfc, 0xb6, 0x08,
	0xe6, 0x96, 0xbf, 0xf0, 0x18, 0xe6, 0x3a, 0xe3, 0x13, 0xa8, 0x59, 0x05, 0xd4, 0x84, 0xfa, 0xd0,
	0xe9, 0x3d, 0xd8, 0xd8, 0xb5, 0x6e, 0xb5, 0x0a, 0xaa, 0x82, 0x34, 0xb2, 0x2f, 0x34, 0x01, 0x01,
	0xa8, 0xbd, 0x3b, 0xeb, 0xda, 0x76, 0x35, 0x11, 0xd5, 0x41, 0xb9, 0xed, 0x5f, 0x5a, 0xb7, 0x9a,
	0x84, 0x6a, 0x20, 0x3b, 0xf6, 0xc8, 0xd5, 0x64, 0xb4, 0x05, 0xb5, 0x01, 0xee, 0x5f, 0x0d, 0x2f,
	0xef, 0x5d, 0x4d, 0x49, 0x25, 0x0f, 0xbd, 0x2b, 0xbb, 0xaf, 0xa9, 0x86, 0x05, 0xdb, 0x2f, 0x66,
	0xed, 0xa0, 0x7d, 0x90, 0x18, 0x9d, 0x72, 0xc7, 0xfe, 0xb9, 0x51, 0xaa, 0x34, 0x7e, 0x0a, 0xa0,
	0xde, 0x79, 0xd1, 0x90, 0x11, 0x74, 0x0e, 0xb5, 0x88, 0x86, 0x5f, 0xc8, 0x34, 0x66, 0xba, 0xc0,
	0x5f, 0xd2, 0xa6, 0x40, 0x26, 0x31, 0x07, 0x79, 0xde, 0x0e, 0x62, 0xba, 0xc2, 0x85, 0xbc, 0x35,
	0x82, 0xe6, 0x8b, 0x14, 0xd2, 0x40, 0x7a, 0x22, 0xeb, 0xcb, 0xa5, 0x9f, 0xe8, 0x00, 0x94, 0xaf,
	0x9e, 0x9f, 0x64, 0xaf, 0xe6, 0xef, 0xc6, 0x65, 0xc2, 0xf7, 0xe2, 0x99, 0xf0, 0xee, 0x03, 0xfc,
	0xf7, 0x47, 0x1e, 0x35, 0xa0, 0x3a, 0x74, 0x3e, 0x3a, 0xfd, 0x91, 0xa3, 0x55, 0xd2, 0xc0, 0xbd,
	0xb7, 0xf0, 0xbd, 0x7d, 0xa5, 0x09, 0x69, 0x80, 0x87, 0x8e, 0xd3, 0x73, 0xae, 0x35, 0xa1, 0x25,
	0x6a, 0xc2, 0x44, 0xe5, 0xff, 0x80, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xd1, 0xa3,
	0x1b, 0x83, 0x04, 0x00, 0x00,
}