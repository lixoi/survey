// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ClassQuestions int32

const (
	ClassQuestions_UNKNOWN_QUESTIONS_CLASS ClassQuestions = 0
	ClassQuestions_LINUX_QUESTIONS         ClassQuestions = 1
	ClassQuestions_K8S_QUESTIONS           ClassQuestions = 2
	ClassQuestions_NETWORK_QUESTIONS       ClassQuestions = 3
	ClassQuestions_SECURITY_QUESTIONS      ClassQuestions = 4
	ClassQuestions_CONTAINER_QUESTIONS     ClassQuestions = 5
	ClassQuestions_DEVELOPER_UESTIONS      ClassQuestions = 6
)

var ClassQuestions_name = map[int32]string{
	0: "UNKNOWN_QUESTIONS_CLASS",
	1: "LINUX_QUESTIONS",
	2: "K8S_QUESTIONS",
	3: "NETWORK_QUESTIONS",
	4: "SECURITY_QUESTIONS",
	5: "CONTAINER_QUESTIONS",
	6: "DEVELOPER_UESTIONS",
}

var ClassQuestions_value = map[string]int32{
	"UNKNOWN_QUESTIONS_CLASS": 0,
	"LINUX_QUESTIONS":         1,
	"K8S_QUESTIONS":           2,
	"NETWORK_QUESTIONS":       3,
	"SECURITY_QUESTIONS":      4,
	"CONTAINER_QUESTIONS":     5,
	"DEVELOPER_UESTIONS":      6,
}

func (x ClassQuestions) String() string {
	return proto.EnumName(ClassQuestions_name, int32(x))
}

func (ClassQuestions) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

type UserInfoRequest struct {
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExistTo              uint32         `protobuf:"varint,2,opt,name=exist_to,json=existTo,proto3" json:"exist_to,omitempty"`
	BaseQuestion         ClassQuestions `protobuf:"varint,3,opt,name=base_question,json=baseQuestion,proto3,enum=api.ClassQuestions" json:"base_question,omitempty"`
	FirstGuestion        ClassQuestions `protobuf:"varint,4,opt,name=first_guestion,json=firstGuestion,proto3,enum=api.ClassQuestions" json:"first_guestion,omitempty"`
	SecondGuestion       ClassQuestions `protobuf:"varint,5,opt,name=second_guestion,json=secondGuestion,proto3,enum=api.ClassQuestions" json:"second_guestion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(m, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoRequest) GetExistTo() uint32 {
	if m != nil {
		return m.ExistTo
	}
	return 0
}

func (m *UserInfoRequest) GetBaseQuestion() ClassQuestions {
	if m != nil {
		return m.BaseQuestion
	}
	return ClassQuestions_UNKNOWN_QUESTIONS_CLASS
}

func (m *UserInfoRequest) GetFirstGuestion() ClassQuestions {
	if m != nil {
		return m.FirstGuestion
	}
	return ClassQuestions_UNKNOWN_QUESTIONS_CLASS
}

func (m *UserInfoRequest) GetSecondGuestion() ClassQuestions {
	if m != nil {
		return m.SecondGuestion
	}
	return ClassQuestions_UNKNOWN_QUESTIONS_CLASS
}

type AnswerRequest struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Number               uint32   `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerRequest) Reset()         { *m = AnswerRequest{} }
func (m *AnswerRequest) String() string { return proto.CompactTextString(m) }
func (*AnswerRequest) ProtoMessage()    {}
func (*AnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *AnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerRequest.Unmarshal(m, b)
}
func (m *AnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerRequest.Marshal(b, m, deterministic)
}
func (m *AnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerRequest.Merge(m, src)
}
func (m *AnswerRequest) XXX_Size() int {
	return xxx_messageInfo_AnswerRequest.Size(m)
}
func (m *AnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerRequest proto.InternalMessageInfo

func (m *AnswerRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AnswerRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *AnswerRequest) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type UserIdRequest struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (m *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(m, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Survey struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Question             string   `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	Latency              string   `protobuf:"bytes,4,opt,name=latency,proto3" json:"latency,omitempty"`
	Answer               string   `protobuf:"bytes,5,opt,name=answer,proto3" json:"answer,omitempty"`
	Number               uint32   `protobuf:"varint,6,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Survey) Reset()         { *m = Survey{} }
func (m *Survey) String() string { return proto.CompactTextString(m) }
func (*Survey) ProtoMessage()    {}
func (*Survey) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{3}
}

func (m *Survey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Survey.Unmarshal(m, b)
}
func (m *Survey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Survey.Marshal(b, m, deterministic)
}
func (m *Survey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Survey.Merge(m, src)
}
func (m *Survey) XXX_Size() int {
	return xxx_messageInfo_Survey.Size(m)
}
func (m *Survey) XXX_DiscardUnknown() {
	xxx_messageInfo_Survey.DiscardUnknown(m)
}

var xxx_messageInfo_Survey proto.InternalMessageInfo

func (m *Survey) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Survey) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Survey) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Survey) GetLatency() string {
	if m != nil {
		return m.Latency
	}
	return ""
}

func (m *Survey) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *Survey) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SurveyResponse struct {
	StartSurvey          *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_survey,json=startSurvey,proto3" json:"start_survey,omitempty"`
	Mesage               string               `protobuf:"bytes,2,opt,name=mesage,proto3" json:"mesage,omitempty"`
	Qs                   []*Survey            `protobuf:"bytes,3,rep,name=qs,proto3" json:"qs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SurveyResponse) Reset()         { *m = SurveyResponse{} }
func (m *SurveyResponse) String() string { return proto.CompactTextString(m) }
func (*SurveyResponse) ProtoMessage()    {}
func (*SurveyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{4}
}

func (m *SurveyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyResponse.Unmarshal(m, b)
}
func (m *SurveyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyResponse.Marshal(b, m, deterministic)
}
func (m *SurveyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyResponse.Merge(m, src)
}
func (m *SurveyResponse) XXX_Size() int {
	return xxx_messageInfo_SurveyResponse.Size(m)
}
func (m *SurveyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyResponse proto.InternalMessageInfo

func (m *SurveyResponse) GetStartSurvey() *timestamp.Timestamp {
	if m != nil {
		return m.StartSurvey
	}
	return nil
}

func (m *SurveyResponse) GetMesage() string {
	if m != nil {
		return m.Mesage
	}
	return ""
}

func (m *SurveyResponse) GetQs() []*Survey {
	if m != nil {
		return m.Qs
	}
	return nil
}

type QuestionResponse struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Question             string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Number               uint32   `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionResponse) Reset()         { *m = QuestionResponse{} }
func (m *QuestionResponse) String() string { return proto.CompactTextString(m) }
func (*QuestionResponse) ProtoMessage()    {}
func (*QuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{5}
}

func (m *QuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionResponse.Unmarshal(m, b)
}
func (m *QuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionResponse.Marshal(b, m, deterministic)
}
func (m *QuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionResponse.Merge(m, src)
}
func (m *QuestionResponse) XXX_Size() int {
	return xxx_messageInfo_QuestionResponse.Size(m)
}
func (m *QuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionResponse proto.InternalMessageInfo

func (m *QuestionResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *QuestionResponse) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *QuestionResponse) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *QuestionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StatusResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{6}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.ClassQuestions", ClassQuestions_name, ClassQuestions_value)
	proto.RegisterType((*UserInfoRequest)(nil), "api.UserInfoRequest")
	proto.RegisterType((*AnswerRequest)(nil), "api.AnswerRequest")
	proto.RegisterType((*UserIdRequest)(nil), "api.UserIdRequest")
	proto.RegisterType((*Survey)(nil), "api.Survey")
	proto.RegisterType((*SurveyResponse)(nil), "api.SurveyResponse")
	proto.RegisterType((*QuestionResponse)(nil), "api.QuestionResponse")
	proto.RegisterType((*StatusResponse)(nil), "api.StatusResponse")
}

func init() {
	proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784)
}

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0x8d, 0x13, 0x62, 0xc8, 0x04, 0x07, 0x58, 0xbe, 0xd2, 0x70, 0x68, 0xe4, 0x53, 0xc4, 0x21,
	0x91, 0xd2, 0x0b, 0x82, 0x56, 0x55, 0x08, 0x86, 0x5a, 0x20, 0xa7, 0xac, 0x93, 0x42, 0x7b, 0xb1,
	0x36, 0x78, 0x49, 0x2d, 0x25, 0xb6, 0xf1, 0x6e, 0xda, 0x72, 0xac, 0xd4, 0xbf, 0xd1, 0x9f, 0xd1,
	0x1f, 0xd7, 0x5b, 0xe5, 0xb5, 0x9d, 0xd8, 0xa8, 0x16, 0x55, 0x8f, 0x6f, 0x3c, 0x6f, 0x67, 0xde,
	0x9b, 0x19, 0x83, 0x42, 0x7c, 0xa7, 0x43, 0x7c, 0xa7, 0xed, 0x07, 0x1e, 0xf7, 0x50, 0x89, 0xf8,
	0x4e, 0xe3, 0xe5, 0xc4, 0xf3, 0x26, 0x53, 0xda, 0x11, 0xa1, 0xf1, 0xfc, 0xbe, 0xc3, 0x9d, 0x19,
	0x65, 0x9c, 0xcc, 0xfc, 0x28, 0x4b, 0xfd, 0x2d, 0xc1, 0xc6, 0x88, 0xd1, 0x40, 0x77, 0xef, 0x3d,
	0x4c, 0x1f, 0xe6, 0x94, 0x71, 0xb4, 0x0f, 0xab, 0x73, 0x46, 0x03, 0xcb, 0xb1, 0xeb, 0x52, 0x53,
	0x6a, 0xad, 0x60, 0x39, 0x84, 0xba, 0x8d, 0x5e, 0xc0, 0x1a, 0xfd, 0xe6, 0x30, 0x6e, 0x71, 0xaf,
	0x5e, 0x6c, 0x4a, 0x2d, 0x05, 0xaf, 0x0a, 0x3c, 0xf4, 0xd0, 0x11, 0x28, 0x63, 0xc2, 0xa8, 0x25,
	0x5e, 0x70, 0x3c, 0xb7, 0x5e, 0x6a, 0x4a, 0xad, 0x5a, 0x77, 0xbb, 0x1d, 0x36, 0xd4, 0x9f, 0x12,
	0xc6, 0xae, 0xe3, 0x2f, 0x0c, 0xaf, 0x87, 0x99, 0x09, 0x44, 0xc7, 0x50, 0xbb, 0x77, 0x02, 0xc6,
	0xad, 0x49, 0x42, 0x5d, 0xc9, 0xa7, 0x2a, 0x22, 0xf5, 0x22, 0xe1, 0xbe, 0x86, 0x0d, 0x46, 0xef,
	0x3c, 0xd7, 0x5e, 0x92, 0xcb, 0xf9, 0xe4, 0x5a, 0x94, 0x9b, 0xb0, 0xd5, 0x5b, 0x50, 0x7a, 0x2e,
	0xfb, 0x4a, 0x83, 0x67, 0x85, 0xef, 0x81, 0x4c, 0x44, 0xa6, 0x90, 0x5d, 0xc1, 0x31, 0x0a, 0xe3,
	0xee, 0x7c, 0x36, 0xa6, 0x81, 0x90, 0xab, 0xe0, 0x18, 0xa9, 0x2d, 0x50, 0x84, 0xa9, 0xf6, 0x73,
	0x2f, 0xab, 0x3f, 0x25, 0x90, 0xcd, 0x79, 0xf0, 0x85, 0x3e, 0xe6, 0x57, 0xdf, 0x81, 0x32, 0x77,
	0xf8, 0x94, 0xc6, 0xc5, 0x23, 0x80, 0x1a, 0xb0, 0x96, 0x31, 0xbb, 0x82, 0x17, 0x18, 0xd5, 0x61,
	0x75, 0x4a, 0x38, 0x75, 0xef, 0x1e, 0x85, 0x99, 0x15, 0x9c, 0xc0, 0x94, 0x92, 0x72, 0x8e, 0x12,
	0x39, 0xa3, 0xe4, 0x87, 0x04, 0xb5, 0xa8, 0x3f, 0x4c, 0x99, 0xef, 0xb9, 0x8c, 0xa2, 0x37, 0xb0,
	0xce, 0x38, 0x09, 0xb8, 0xc5, 0x44, 0x5c, 0x34, 0x5b, 0xed, 0x36, 0xda, 0xd1, 0xaa, 0xb5, 0x93,
	0x55, 0x6b, 0x0f, 0x93, 0x55, 0xc3, 0x55, 0x91, 0x1f, 0xcb, 0xdc, 0x03, 0x79, 0x46, 0x19, 0x99,
	0x24, 0x72, 0x62, 0x84, 0x0e, 0xa0, 0xf8, 0xc0, 0xea, 0xa5, 0x66, 0xa9, 0x55, 0xed, 0x56, 0xc5,
	0xf8, 0xe2, 0xba, 0xc5, 0x07, 0xa6, 0x3e, 0xc2, 0x66, 0x32, 0xc7, 0x45, 0x1f, 0xb9, 0x7e, 0xa5,
	0x9d, 0x29, 0x3e, 0x71, 0x26, 0x67, 0x62, 0xa1, 0x63, 0x33, 0xca, 0x44, 0x5b, 0xb1, 0x63, 0x31,
	0x54, 0x0f, 0xa1, 0x66, 0x72, 0xc2, 0xe7, 0x6c, 0x51, 0x38, 0x95, 0x2b, 0x65, 0x72, 0x0f, 0x7f,
	0x49, 0x50, 0xcb, 0x2e, 0x1d, 0x3a, 0x80, 0xfd, 0x91, 0x71, 0x69, 0x0c, 0x6e, 0x0c, 0xeb, 0x7a,
	0xa4, 0x99, 0x43, 0x7d, 0x60, 0x98, 0x56, 0xff, 0xaa, 0x67, 0x9a, 0x9b, 0x05, 0xb4, 0x0d, 0x1b,
	0x57, 0xba, 0x31, 0xba, 0x5d, 0x7e, 0xda, 0x94, 0xd0, 0x16, 0x28, 0x97, 0x47, 0x66, 0x2a, 0x54,
	0x44, 0xbb, 0xb0, 0x65, 0x68, 0xc3, 0x9b, 0x01, 0xbe, 0x4c, 0x85, 0x4b, 0x68, 0x0f, 0x90, 0xa9,
	0xf5, 0x47, 0x58, 0x1f, 0x7e, 0x4c, 0xc5, 0x57, 0xd0, 0x3e, 0x6c, 0xf7, 0x07, 0xc6, 0xb0, 0xa7,
	0x1b, 0x1a, 0x4e, 0x7d, 0x28, 0x87, 0x84, 0x33, 0xed, 0x83, 0x76, 0x35, 0x78, 0xaf, 0x61, 0x6b,
	0x11, 0x97, 0xbb, 0xdf, 0x4b, 0x50, 0xd1, 0xfb, 0xef, 0xe2, 0x09, 0x9d, 0xc0, 0x7a, 0xcf, 0xb6,
	0xfb, 0xc4, 0xb5, 0x1d, 0x9b, 0x70, 0x8a, 0x76, 0xc4, 0x34, 0x9e, 0xfc, 0x25, 0x1a, 0xd1, 0x89,
	0x65, 0xad, 0x51, 0x0b, 0xe1, 0x49, 0x9e, 0xd1, 0x29, 0xe5, 0x74, 0xc9, 0x47, 0x4b, 0xbe, 0xfd,
	0x0c, 0xfb, 0x18, 0xaa, 0x66, 0x6a, 0x57, 0xfe, 0xc6, 0xdc, 0x15, 0xb1, 0xa7, 0xdb, 0xa0, 0x16,
	0xd0, 0x11, 0x54, 0x4c, 0xca, 0xa3, 0x8b, 0x8e, 0x99, 0x99, 0xf3, 0xce, 0x67, 0xbe, 0x05, 0x64,
	0x52, 0x7e, 0xee, 0xb8, 0x0e, 0xfb, 0xfc, 0x5f, 0x6d, 0x9f, 0xc2, 0xee, 0x05, 0x8d, 0x9b, 0x3e,
	0xf7, 0x82, 0x7f, 0x7a, 0x23, 0x73, 0x54, 0x6a, 0xe1, 0x74, 0xed, 0x93, 0xdc, 0xee, 0x9c, 0x10,
	0xdf, 0x19, 0xcb, 0xe2, 0x84, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x94, 0x31, 0x89, 0x2d,
	0xd1, 0x05, 0x00, 0x00,
}
