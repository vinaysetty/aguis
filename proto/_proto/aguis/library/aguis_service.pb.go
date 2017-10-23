// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aguis/library/aguis_service.proto

/*
Package library is a generated protocol buffer package.

It is generated from these files:
	aguis/library/aguis_service.proto

It has these top-level messages:
	Enrollment
	User
	UsersResponse
	Void
	GetRecordRequest
	UpdateUserRequest
	Assignment
	Submission
	Assignments
	Group
	Course
	Courses
	CoursesWithEnrollmentRequest
*/
package library

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Enrollment struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Courseid uint64 `protobuf:"varint,2,opt,name=courseid" json:"courseid,omitempty"`
	Userid   uint64 `protobuf:"varint,3,opt,name=userid" json:"userid,omitempty"`
	Groupid  uint64 `protobuf:"varint,4,opt,name=groupid" json:"groupid,omitempty"`
	Status   uint32 `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
}

func (m *Enrollment) Reset()                    { *m = Enrollment{} }
func (m *Enrollment) String() string            { return proto.CompactTextString(m) }
func (*Enrollment) ProtoMessage()               {}
func (*Enrollment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Enrollment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Enrollment) GetCourseid() uint64 {
	if m != nil {
		return m.Courseid
	}
	return 0
}

func (m *Enrollment) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *Enrollment) GetGroupid() uint64 {
	if m != nil {
		return m.Groupid
	}
	return 0
}

func (m *Enrollment) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type User struct {
	Id          uint64        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Isadmin     bool          `protobuf:"varint,2,opt,name=isadmin" json:"isadmin,omitempty"`
	Name        string        `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Studentid   string        `protobuf:"bytes,4,opt,name=studentid" json:"studentid,omitempty"`
	Email       string        `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Avatarurl   string        `protobuf:"bytes,6,opt,name=avatarurl" json:"avatarurl,omitempty"`
	Enrollments []*Enrollment `protobuf:"bytes,7,rep,name=enrollments" json:"enrollments,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetIsadmin() bool {
	if m != nil {
		return m.Isadmin
	}
	return false
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetStudentid() string {
	if m != nil {
		return m.Studentid
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAvatarurl() string {
	if m != nil {
		return m.Avatarurl
	}
	return ""
}

func (m *User) GetEnrollments() []*Enrollment {
	if m != nil {
		return m.Enrollments
	}
	return nil
}

type UsersResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *UsersResponse) Reset()                    { *m = UsersResponse{} }
func (m *UsersResponse) String() string            { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()               {}
func (*UsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetRecordRecord can be used for all type of single record request like User, Course etc
type GetRecordRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRecordRequest) Reset()                    { *m = GetRecordRequest{} }
func (m *GetRecordRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRecordRequest) ProtoMessage()               {}
func (*GetRecordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRecordRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateUserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=User" json:"User,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Assignment struct {
	Id          uint64      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Courseid    uint64      `protobuf:"varint,2,opt,name=courseid" json:"courseid,omitempty"`
	Name        string      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Language    string      `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
	Deadline    string      `protobuf:"bytes,5,opt,name=deadline" json:"deadline,omitempty"`
	Autoapprove bool        `protobuf:"varint,6,opt,name=autoapprove" json:"autoapprove,omitempty"`
	Order       uint32      `protobuf:"varint,7,opt,name=order" json:"order,omitempty"`
	Submission  *Submission `protobuf:"bytes,8,opt,name=submission" json:"submission,omitempty"`
}

func (m *Assignment) Reset()                    { *m = Assignment{} }
func (m *Assignment) String() string            { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()               {}
func (*Assignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Assignment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Assignment) GetCourseid() uint64 {
	if m != nil {
		return m.Courseid
	}
	return 0
}

func (m *Assignment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Assignment) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Assignment) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Assignment) GetAutoapprove() bool {
	if m != nil {
		return m.Autoapprove
	}
	return false
}

func (m *Assignment) GetOrder() uint32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Assignment) GetSubmission() *Submission {
	if m != nil {
		return m.Submission
	}
	return nil
}

type Submission struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Assignmentid uint64 `protobuf:"varint,2,opt,name=assignmentid" json:"assignmentid,omitempty"`
	Userid       uint64 `protobuf:"varint,3,opt,name=userid" json:"userid,omitempty"`
	Groupid      uint64 `protobuf:"varint,4,opt,name=groupid" json:"groupid,omitempty"`
	Score        uint32 `protobuf:"varint,5,opt,name=score" json:"score,omitempty"`
	Scoreobjects string `protobuf:"bytes,6,opt,name=scoreobjects" json:"scoreobjects,omitempty"`
	Buildinfo    string `protobuf:"bytes,7,opt,name=buildinfo" json:"buildinfo,omitempty"`
	Commithash   string `protobuf:"bytes,8,opt,name=commithash" json:"commithash,omitempty"`
}

func (m *Submission) Reset()                    { *m = Submission{} }
func (m *Submission) String() string            { return proto.CompactTextString(m) }
func (*Submission) ProtoMessage()               {}
func (*Submission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Submission) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Submission) GetAssignmentid() uint64 {
	if m != nil {
		return m.Assignmentid
	}
	return 0
}

func (m *Submission) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *Submission) GetGroupid() uint64 {
	if m != nil {
		return m.Groupid
	}
	return 0
}

func (m *Submission) GetScore() uint32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Submission) GetScoreobjects() string {
	if m != nil {
		return m.Scoreobjects
	}
	return ""
}

func (m *Submission) GetBuildinfo() string {
	if m != nil {
		return m.Buildinfo
	}
	return ""
}

func (m *Submission) GetCommithash() string {
	if m != nil {
		return m.Commithash
	}
	return ""
}

type Assignments struct {
	Assignments []*Assignment `protobuf:"bytes,1,rep,name=assignments" json:"assignments,omitempty"`
}

func (m *Assignments) Reset()                    { *m = Assignments{} }
func (m *Assignments) String() string            { return proto.CompactTextString(m) }
func (*Assignments) ProtoMessage()               {}
func (*Assignments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Assignments) GetAssignments() []*Assignment {
	if m != nil {
		return m.Assignments
	}
	return nil
}

type Group struct {
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Course struct {
	Id          uint64        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code        string        `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	Year        uint32        `protobuf:"varint,4,opt,name=year" json:"year,omitempty"`
	Tag         string        `protobuf:"bytes,5,opt,name=tag" json:"tag,omitempty"`
	Provider    string        `protobuf:"bytes,6,opt,name=provider" json:"provider,omitempty"`
	Directoryid uint64        `protobuf:"varint,7,opt,name=directoryid" json:"directoryid,omitempty"`
	Enrolled    int32         `protobuf:"varint,8,opt,name=enrolled" json:"enrolled,omitempty"`
	Assignments []*Assignment `protobuf:"bytes,9,rep,name=assignments" json:"assignments,omitempty"`
	Groups      []*Group      `protobuf:"bytes,10,rep,name=groups" json:"groups,omitempty"`
}

func (m *Course) Reset()                    { *m = Course{} }
func (m *Course) String() string            { return proto.CompactTextString(m) }
func (*Course) ProtoMessage()               {}
func (*Course) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Course) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Course) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Course) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Course) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Course) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Course) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Course) GetDirectoryid() uint64 {
	if m != nil {
		return m.Directoryid
	}
	return 0
}

func (m *Course) GetEnrolled() int32 {
	if m != nil {
		return m.Enrolled
	}
	return 0
}

func (m *Course) GetAssignments() []*Assignment {
	if m != nil {
		return m.Assignments
	}
	return nil
}

func (m *Course) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type Courses struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=Courses" json:"Courses,omitempty"`
}

func (m *Courses) Reset()                    { *m = Courses{} }
func (m *Courses) String() string            { return proto.CompactTextString(m) }
func (*Courses) ProtoMessage()               {}
func (*Courses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Courses) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type CoursesWithEnrollmentRequest struct {
	Userid uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	State  string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
}

func (m *CoursesWithEnrollmentRequest) Reset()                    { *m = CoursesWithEnrollmentRequest{} }
func (m *CoursesWithEnrollmentRequest) String() string            { return proto.CompactTextString(m) }
func (*CoursesWithEnrollmentRequest) ProtoMessage()               {}
func (*CoursesWithEnrollmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CoursesWithEnrollmentRequest) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CoursesWithEnrollmentRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Enrollment)(nil), "library.Enrollment")
	proto.RegisterType((*User)(nil), "library.User")
	proto.RegisterType((*UsersResponse)(nil), "library.UsersResponse")
	proto.RegisterType((*Void)(nil), "library.Void")
	proto.RegisterType((*GetRecordRequest)(nil), "library.GetRecordRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "library.UpdateUserRequest")
	proto.RegisterType((*Assignment)(nil), "library.Assignment")
	proto.RegisterType((*Submission)(nil), "library.Submission")
	proto.RegisterType((*Assignments)(nil), "library.Assignments")
	proto.RegisterType((*Group)(nil), "library.Group")
	proto.RegisterType((*Course)(nil), "library.Course")
	proto.RegisterType((*Courses)(nil), "library.Courses")
	proto.RegisterType((*CoursesWithEnrollmentRequest)(nil), "library.CoursesWithEnrollmentRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AutograderService service

type AutograderServiceClient interface {
	GetUser(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UsersResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetCourse(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*Course, error)
	GetCourses(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Courses, error)
	GetCoursesWithEnrollment(ctx context.Context, in *CoursesWithEnrollmentRequest, opts ...grpc.CallOption) (*Courses, error)
	GetAssignments(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*Assignments, error)
}

type autograderServiceClient struct {
	cc *grpc.ClientConn
}

func NewAutograderServiceClient(cc *grpc.ClientConn) AutograderServiceClient {
	return &autograderServiceClient{cc}
}

func (c *autograderServiceClient) GetUser(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) GetUsers(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/library.AutograderService/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) GetCourse(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetCourse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) GetCourses(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetCourses", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) GetCoursesWithEnrollment(ctx context.Context, in *CoursesWithEnrollmentRequest, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetCoursesWithEnrollment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) GetAssignments(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*Assignments, error) {
	out := new(Assignments)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetAssignments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AutograderService service

type AutograderServiceServer interface {
	GetUser(context.Context, *GetRecordRequest) (*User, error)
	GetUsers(context.Context, *Void) (*UsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	GetCourse(context.Context, *GetRecordRequest) (*Course, error)
	GetCourses(context.Context, *Void) (*Courses, error)
	GetCoursesWithEnrollment(context.Context, *CoursesWithEnrollmentRequest) (*Courses, error)
	GetAssignments(context.Context, *GetRecordRequest) (*Assignments, error)
}

func RegisterAutograderServiceServer(s *grpc.Server, srv AutograderServiceServer) {
	s.RegisterService(&_AutograderService_serviceDesc, srv)
}

func _AutograderService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetUser(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetUsers(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetCourse(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetCourses(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_GetCoursesWithEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoursesWithEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetCoursesWithEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetCoursesWithEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetCoursesWithEnrollment(ctx, req.(*CoursesWithEnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_GetAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetAssignments(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutograderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "library.AutograderService",
	HandlerType: (*AutograderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _AutograderService_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _AutograderService_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AutograderService_UpdateUser_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _AutograderService_GetCourse_Handler,
		},
		{
			MethodName: "GetCourses",
			Handler:    _AutograderService_GetCourses_Handler,
		},
		{
			MethodName: "GetCoursesWithEnrollment",
			Handler:    _AutograderService_GetCoursesWithEnrollment_Handler,
		},
		{
			MethodName: "GetAssignments",
			Handler:    _AutograderService_GetAssignments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aguis/library/aguis_service.proto",
}

func init() { proto.RegisterFile("aguis/library/aguis_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xd6, 0xea, 0x6f, 0xb5, 0xa3, 0xca, 0xb5, 0x59, 0xc3, 0xd8, 0x0a, 0x46, 0x21, 0xb3, 0x68,
	0xe1, 0x5e, 0x6c, 0xc0, 0x3f, 0x05, 0x7c, 0x34, 0xdc, 0x42, 0x97, 0x9e, 0x68, 0xb8, 0x39, 0x06,
	0xd4, 0x92, 0x91, 0x19, 0xac, 0x96, 0x0a, 0xc9, 0x35, 0xe0, 0x6b, 0x5e, 0x24, 0xef, 0x92, 0x7b,
	0x5e, 0x24, 0xb7, 0xbc, 0x41, 0x40, 0x2e, 0xc5, 0x5d, 0xfd, 0xc0, 0x48, 0x72, 0x9b, 0x6f, 0x7e,
	0x38, 0x9a, 0x6f, 0xbe, 0x1d, 0xc1, 0x09, 0x9d, 0x97, 0x42, 0x9f, 0xe7, 0x62, 0xa6, 0xa8, 0x7a,
	0x3e, 0x77, 0xe8, 0xb5, 0xe6, 0xea, 0x49, 0x64, 0xfc, 0x6c, 0xa9, 0xa4, 0x91, 0x28, 0xf6, 0x41,
	0xfc, 0x3e, 0x02, 0xf8, 0xb7, 0x50, 0x32, 0xcf, 0x17, 0xbc, 0x30, 0x68, 0x0f, 0xda, 0x82, 0xa5,
	0xd1, 0x24, 0x3a, 0xed, 0x92, 0xb6, 0x60, 0x68, 0x0c, 0x83, 0x4c, 0x96, 0x4a, 0x73, 0xc1, 0xd2,
	0xb6, 0xf3, 0x06, 0x8c, 0x8e, 0xa0, 0x5f, 0x6a, 0xae, 0x04, 0x4b, 0x3b, 0x2e, 0xe2, 0x11, 0x4a,
	0x21, 0x9e, 0x2b, 0x59, 0x2e, 0x05, 0x4b, 0xbb, 0x2e, 0xb0, 0x82, 0xb6, 0x42, 0x1b, 0x6a, 0x4a,
	0x9d, 0xf6, 0x26, 0xd1, 0xe9, 0x88, 0x78, 0x84, 0x3f, 0x45, 0xd0, 0x7d, 0xd0, 0x5c, 0x6d, 0xb5,
	0x4f, 0x21, 0x16, 0x9a, 0xb2, 0x85, 0x28, 0x5c, 0xf7, 0x01, 0x59, 0x41, 0x84, 0xa0, 0x5b, 0xd0,
	0x05, 0x77, 0xad, 0x13, 0xe2, 0x6c, 0x74, 0x0c, 0x89, 0x36, 0x25, 0xe3, 0x85, 0xf1, 0xad, 0x13,
	0x52, 0x3b, 0xd0, 0x21, 0xf4, 0xf8, 0x82, 0x8a, 0xdc, 0xf5, 0x4e, 0x48, 0x05, 0x6c, 0x0d, 0x7d,
	0xa2, 0x86, 0xaa, 0x52, 0xe5, 0x69, 0xbf, 0xaa, 0x09, 0x0e, 0x74, 0x0d, 0x43, 0x1e, 0xc8, 0xd1,
	0x69, 0x3c, 0xe9, 0x9c, 0x0e, 0x2f, 0x7e, 0x39, 0xf3, 0xe4, 0x9d, 0xd5, 0xc4, 0x91, 0x66, 0x1e,
	0xbe, 0x82, 0x91, 0x1d, 0x47, 0x13, 0xae, 0x97, 0xb2, 0xd0, 0x1c, 0xfd, 0x0e, 0x3d, 0xe7, 0x48,
	0x23, 0xf7, 0xc2, 0x28, 0xbc, 0x60, 0xbd, 0xa4, 0x8a, 0xe1, 0x3e, 0x74, 0xff, 0x97, 0x82, 0x61,
	0x0c, 0xfb, 0x53, 0x6e, 0x08, 0xcf, 0xa4, 0x62, 0x84, 0xbf, 0x2b, 0xb9, 0xde, 0xda, 0x0b, 0xfe,
	0x1b, 0x0e, 0x1e, 0x96, 0x8c, 0x1a, 0xee, 0x1e, 0xf0, 0x49, 0x27, 0x15, 0x8b, 0x2e, 0x6d, 0xab,
	0x89, 0x0b, 0xe1, 0x2f, 0x11, 0xc0, 0xad, 0xd6, 0x62, 0x5e, 0x7c, 0xf7, 0xba, 0x77, 0x31, 0x3e,
	0x86, 0x41, 0x4e, 0x8b, 0x79, 0x49, 0xe7, 0xdc, 0x13, 0x1e, 0xb0, 0x8d, 0x31, 0x4e, 0x59, 0x2e,
	0x0a, 0xee, 0x29, 0x0f, 0x18, 0x4d, 0x60, 0x48, 0x4b, 0x23, 0xe9, 0x72, 0xa9, 0xe4, 0x13, 0x77,
	0xbc, 0x0f, 0x48, 0xd3, 0x65, 0xb7, 0x25, 0x15, 0xe3, 0x2a, 0x8d, 0x9d, 0x52, 0x2a, 0x80, 0x2e,
	0x01, 0x74, 0x39, 0x5b, 0x08, 0xad, 0x85, 0x2c, 0xd2, 0x81, 0x9b, 0xb3, 0x5e, 0xc7, 0x7d, 0x08,
	0x91, 0x46, 0x1a, 0xfe, 0x1c, 0x01, 0xd4, 0xa1, 0xad, 0x99, 0x31, 0xfc, 0x44, 0x03, 0x23, 0x61,
	0xee, 0x35, 0xdf, 0x0f, 0x48, 0xfd, 0x10, 0x7a, 0x3a, 0x93, 0x8a, 0x7b, 0xa5, 0x57, 0xc0, 0xf6,
	0x72, 0x86, 0x9c, 0xbd, 0xe5, 0x99, 0xd1, 0x5e, 0x70, 0x6b, 0x3e, 0xab, 0xc8, 0x59, 0x29, 0x72,
	0x26, 0x8a, 0x37, 0xd2, 0x4d, 0x9f, 0x90, 0xda, 0x81, 0x7e, 0x03, 0xc8, 0xe4, 0x62, 0x21, 0xcc,
	0x23, 0xd5, 0x8f, 0x8e, 0x81, 0x84, 0x34, 0x3c, 0xf8, 0x1f, 0x18, 0xd6, 0xfb, 0xd5, 0x56, 0xc0,
	0xf5, 0x20, 0x2b, 0xf9, 0xd5, 0x8c, 0xd5, 0xa9, 0xa4, 0x99, 0x87, 0x63, 0xe8, 0x4d, 0xed, 0x20,
	0xf8, 0x43, 0x1b, 0xfa, 0x77, 0x4e, 0x01, 0x5b, 0xbc, 0xad, 0xf4, 0xd0, 0x6e, 0xe8, 0x01, 0x41,
	0x37, 0x93, 0x2c, 0x68, 0xc4, 0xda, 0xd6, 0xf7, 0xcc, 0xa9, 0x72, 0x04, 0x8d, 0x88, 0xb3, 0xd1,
	0x3e, 0x74, 0x0c, 0x9d, 0x7b, 0x59, 0x58, 0xd3, 0xaa, 0xc5, 0x2e, 0x5e, 0xd8, 0x95, 0x57, 0xac,
	0x04, 0x6c, 0xd5, 0xc2, 0x84, 0xe2, 0x99, 0x91, 0xea, 0x59, 0x30, 0xc7, 0x49, 0x97, 0x34, 0x5d,
	0xb6, 0xba, 0xfa, 0xfe, 0x38, 0x73, 0x9c, 0xf4, 0x48, 0xc0, 0x9b, 0x14, 0x24, 0xdf, 0x46, 0x01,
	0xfa, 0x13, 0xfa, 0x6e, 0x97, 0x3a, 0x05, 0x57, 0xb1, 0x17, 0x2a, 0x1c, 0x33, 0xc4, 0x47, 0xf1,
	0x15, 0xc4, 0x15, 0x41, 0x1a, 0xfd, 0x15, 0x4c, 0x4f, 0xf4, 0xcf, 0xa1, 0xa6, 0xf2, 0x93, 0x55,
	0x1c, 0xff, 0x07, 0xc7, 0xde, 0x7c, 0x25, 0xcc, 0x63, 0xe3, 0x8e, 0xf8, 0x4f, 0xb9, 0x16, 0x5c,
	0xb4, 0x26, 0x38, 0x2b, 0x2b, 0x43, 0xcd, 0x8a, 0xf5, 0x0a, 0x5c, 0x7c, 0xec, 0xc0, 0xc1, 0x6d,
	0x69, 0xe4, 0x5c, 0x51, 0xc6, 0xd5, 0x7d, 0x75, 0xe9, 0xd1, 0x35, 0xc4, 0x53, 0x6e, 0xdc, 0x5d,
	0xfd, 0xb5, 0xfe, 0xf1, 0x1b, 0x97, 0x65, 0xbc, 0x7e, 0x26, 0x70, 0x0b, 0x5d, 0xc2, 0xc0, 0x97,
	0x69, 0x54, 0x07, 0xed, 0x65, 0x1a, 0x1f, 0xad, 0xe5, 0x86, 0xf3, 0x86, 0x5b, 0xe8, 0x06, 0xa0,
	0xbe, 0x47, 0x68, 0x5c, 0xe7, 0x6d, 0x1e, 0xa9, 0xed, 0x7e, 0x37, 0x90, 0x4c, 0xb9, 0xf1, 0x22,
	0x7b, 0xe1, 0x87, 0x6e, 0x92, 0x89, 0x5b, 0xe8, 0x1c, 0x20, 0x94, 0x6e, 0xfd, 0xd8, 0xfd, 0x8d,
	0x7c, 0x8d, 0x5b, 0xe8, 0x01, 0xd2, 0xba, 0x60, 0x9d, 0x79, 0xf4, 0xc7, 0x66, 0xfe, 0xce, 0xcd,
	0xec, 0x7c, 0xf6, 0x0e, 0xf6, 0xa6, 0xdc, 0x34, 0xbf, 0xbb, 0x17, 0xe6, 0x38, 0xdc, 0x21, 0x3d,
	0x8d, 0x5b, 0xb3, 0xbe, 0xfb, 0x67, 0xbe, 0xfc, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x28, 0x6a,
	0xcd, 0xbe, 0x07, 0x00, 0x00,
}