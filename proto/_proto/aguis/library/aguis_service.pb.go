// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aguis/library/aguis_service.proto

/*
Package library is a generated protocol buffer package.

It is generated from these files:
	aguis/library/aguis_service.proto

It has these top-level messages:
	Enrollment
	EnrollemntResponse
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
	RecordWithStatusRequest
	UserIDCourseID
	StatusCode
	UpdateEnrollmentRequest
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
	Id       uint64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Courseid uint64  `protobuf:"varint,2,opt,name=courseid" json:"courseid,omitempty"`
	Userid   uint64  `protobuf:"varint,3,opt,name=userid" json:"userid,omitempty"`
	Groupid  uint64  `protobuf:"varint,4,opt,name=groupid" json:"groupid,omitempty"`
	Status   uint32  `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	User     *User   `protobuf:"bytes,6,opt,name=user" json:"user,omitempty"`
	Course   *Course `protobuf:"bytes,7,opt,name=course" json:"course,omitempty"`
	Group    *Group  `protobuf:"bytes,8,opt,name=group" json:"group,omitempty"`
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

func (m *Enrollment) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Enrollment) GetCourse() *Course {
	if m != nil {
		return m.Course
	}
	return nil
}

func (m *Enrollment) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type EnrollemntResponse struct {
	Enrollments []*Enrollment `protobuf:"bytes,1,rep,name=enrollments" json:"enrollments,omitempty"`
}

func (m *EnrollemntResponse) Reset()                    { *m = EnrollemntResponse{} }
func (m *EnrollemntResponse) String() string            { return proto.CompactTextString(m) }
func (*EnrollemntResponse) ProtoMessage()               {}
func (*EnrollemntResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EnrollemntResponse) GetEnrollments() []*Enrollment {
	if m != nil {
		return m.Enrollments
	}
	return nil
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
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*UsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GetRecordRecord can be used for all type of single record request like User, Course etc
type GetRecordRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRecordRequest) Reset()                    { *m = GetRecordRequest{} }
func (m *GetRecordRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRecordRequest) ProtoMessage()               {}
func (*GetRecordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*Assignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*Submission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*Assignments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
func (*Course) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
func (*Courses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Courses) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type RecordWithStatusRequest struct {
	Id    uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
}

func (m *RecordWithStatusRequest) Reset()                    { *m = RecordWithStatusRequest{} }
func (m *RecordWithStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*RecordWithStatusRequest) ProtoMessage()               {}
func (*RecordWithStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RecordWithStatusRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RecordWithStatusRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type UserIDCourseID struct {
	Userid   uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Courseid uint64 `protobuf:"varint,2,opt,name=courseid" json:"courseid,omitempty"`
}

func (m *UserIDCourseID) Reset()                    { *m = UserIDCourseID{} }
func (m *UserIDCourseID) String() string            { return proto.CompactTextString(m) }
func (*UserIDCourseID) ProtoMessage()               {}
func (*UserIDCourseID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UserIDCourseID) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *UserIDCourseID) GetCourseid() uint64 {
	if m != nil {
		return m.Courseid
	}
	return 0
}

type StatusCode struct {
	Statuscode int32 `protobuf:"varint,1,opt,name=statuscode" json:"statuscode,omitempty"`
}

func (m *StatusCode) Reset()                    { *m = StatusCode{} }
func (m *StatusCode) String() string            { return proto.CompactTextString(m) }
func (*StatusCode) ProtoMessage()               {}
func (*StatusCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *StatusCode) GetStatuscode() int32 {
	if m != nil {
		return m.Statuscode
	}
	return 0
}

type UpdateEnrollmentRequest struct {
	Userid   uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Courseid uint64 `protobuf:"varint,2,opt,name=courseid" json:"courseid,omitempty"`
	Status   int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
}

func (m *UpdateEnrollmentRequest) Reset()                    { *m = UpdateEnrollmentRequest{} }
func (m *UpdateEnrollmentRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEnrollmentRequest) ProtoMessage()               {}
func (*UpdateEnrollmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *UpdateEnrollmentRequest) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *UpdateEnrollmentRequest) GetCourseid() uint64 {
	if m != nil {
		return m.Courseid
	}
	return 0
}

func (m *UpdateEnrollmentRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Enrollment)(nil), "library.Enrollment")
	proto.RegisterType((*EnrollemntResponse)(nil), "library.EnrollemntResponse")
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
	proto.RegisterType((*RecordWithStatusRequest)(nil), "library.RecordWithStatusRequest")
	proto.RegisterType((*UserIDCourseID)(nil), "library.UserIDCourseID")
	proto.RegisterType((*StatusCode)(nil), "library.StatusCode")
	proto.RegisterType((*UpdateEnrollmentRequest)(nil), "library.UpdateEnrollmentRequest")
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
	CreateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error)
	GetCourse(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*Course, error)
	UpdateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error)
	GetCourses(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Courses, error)
	GetCoursesWithEnrollment(ctx context.Context, in *RecordWithStatusRequest, opts ...grpc.CallOption) (*Courses, error)
	GetAssignments(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*Assignments, error)
	GetEnrollmentsByCourse(ctx context.Context, in *RecordWithStatusRequest, opts ...grpc.CallOption) (*EnrollemntResponse, error)
	CreateEnrollment(ctx context.Context, in *UserIDCourseID, opts ...grpc.CallOption) (*StatusCode, error)
	UpdateEnrollment(ctx context.Context, in *UpdateEnrollmentRequest, opts ...grpc.CallOption) (*StatusCode, error)
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

func (c *autograderServiceClient) CreateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := grpc.Invoke(ctx, "/library.AutograderService/CreateCourse", in, out, c.cc, opts...)
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

func (c *autograderServiceClient) UpdateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := grpc.Invoke(ctx, "/library.AutograderService/UpdateCourse", in, out, c.cc, opts...)
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

func (c *autograderServiceClient) GetCoursesWithEnrollment(ctx context.Context, in *RecordWithStatusRequest, opts ...grpc.CallOption) (*Courses, error) {
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

func (c *autograderServiceClient) GetEnrollmentsByCourse(ctx context.Context, in *RecordWithStatusRequest, opts ...grpc.CallOption) (*EnrollemntResponse, error) {
	out := new(EnrollemntResponse)
	err := grpc.Invoke(ctx, "/library.AutograderService/GetEnrollmentsByCourse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) CreateEnrollment(ctx context.Context, in *UserIDCourseID, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := grpc.Invoke(ctx, "/library.AutograderService/CreateEnrollment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autograderServiceClient) UpdateEnrollment(ctx context.Context, in *UpdateEnrollmentRequest, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := grpc.Invoke(ctx, "/library.AutograderService/UpdateEnrollment", in, out, c.cc, opts...)
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
	CreateCourse(context.Context, *Course) (*Course, error)
	GetCourse(context.Context, *GetRecordRequest) (*Course, error)
	UpdateCourse(context.Context, *Course) (*Course, error)
	GetCourses(context.Context, *Void) (*Courses, error)
	GetCoursesWithEnrollment(context.Context, *RecordWithStatusRequest) (*Courses, error)
	GetAssignments(context.Context, *GetRecordRequest) (*Assignments, error)
	GetEnrollmentsByCourse(context.Context, *RecordWithStatusRequest) (*EnrollemntResponse, error)
	CreateEnrollment(context.Context, *UserIDCourseID) (*StatusCode, error)
	UpdateEnrollment(context.Context, *UpdateEnrollmentRequest) (*StatusCode, error)
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

func _AutograderService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).CreateCourse(ctx, req.(*Course))
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

func _AutograderService_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/UpdateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).UpdateCourse(ctx, req.(*Course))
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
	in := new(RecordWithStatusRequest)
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
		return srv.(AutograderServiceServer).GetCoursesWithEnrollment(ctx, req.(*RecordWithStatusRequest))
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

func _AutograderService_GetEnrollmentsByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordWithStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).GetEnrollmentsByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/GetEnrollmentsByCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).GetEnrollmentsByCourse(ctx, req.(*RecordWithStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_CreateEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDCourseID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).CreateEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/CreateEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).CreateEnrollment(ctx, req.(*UserIDCourseID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutograderService_UpdateEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutograderServiceServer).UpdateEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.AutograderService/UpdateEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutograderServiceServer).UpdateEnrollment(ctx, req.(*UpdateEnrollmentRequest))
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
			MethodName: "CreateCourse",
			Handler:    _AutograderService_CreateCourse_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _AutograderService_GetCourse_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _AutograderService_UpdateCourse_Handler,
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
		{
			MethodName: "GetEnrollmentsByCourse",
			Handler:    _AutograderService_GetEnrollmentsByCourse_Handler,
		},
		{
			MethodName: "CreateEnrollment",
			Handler:    _AutograderService_CreateEnrollment_Handler,
		},
		{
			MethodName: "UpdateEnrollment",
			Handler:    _AutograderService_UpdateEnrollment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aguis/library/aguis_service.proto",
}

func init() { proto.RegisterFile("aguis/library/aguis_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 978 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0xce, 0xa4, 0xf9, 0x9b, 0x93, 0xa6, 0x64, 0x4d, 0xd5, 0x0e, 0x01, 0xad, 0x52, 0x83, 0xa0,
	0x48, 0xa8, 0x95, 0xd2, 0x5d, 0x24, 0xae, 0xd0, 0x6e, 0x8b, 0xa2, 0xd5, 0xde, 0x20, 0x57, 0x0b,
	0xe2, 0x0a, 0x39, 0xb1, 0x49, 0x8d, 0x92, 0x71, 0xb0, 0x3d, 0x95, 0xfa, 0x18, 0x3c, 0x01, 0x4f,
	0xc4, 0x8b, 0x70, 0x87, 0xc4, 0x03, 0x20, 0xff, 0xc4, 0x33, 0x99, 0xb4, 0xdd, 0x9f, 0x9b, 0x91,
	0xcf, 0x8f, 0x7d, 0x7c, 0xbe, 0xf3, 0x9d, 0xe3, 0x81, 0x13, 0xba, 0x28, 0x84, 0x3e, 0x5f, 0x8a,
	0x99, 0xa2, 0xea, 0xee, 0xdc, 0x49, 0xbf, 0x6a, 0xae, 0x6e, 0xc5, 0x9c, 0x9f, 0xad, 0x95, 0x34,
	0x12, 0x75, 0x83, 0x11, 0xff, 0x97, 0x00, 0xfc, 0x90, 0x2b, 0xb9, 0x5c, 0xae, 0x78, 0x6e, 0xd0,
	0x01, 0x34, 0x05, 0xcb, 0x92, 0x71, 0x72, 0xda, 0x22, 0x4d, 0xc1, 0xd0, 0x08, 0x7a, 0x73, 0x59,
	0x28, 0xcd, 0x05, 0xcb, 0x9a, 0x4e, 0x1b, 0x65, 0x74, 0x04, 0x9d, 0x42, 0x73, 0x25, 0x58, 0xb6,
	0xe7, 0x2c, 0x41, 0x42, 0x19, 0x74, 0x17, 0x4a, 0x16, 0x6b, 0xc1, 0xb2, 0x96, 0x33, 0x6c, 0x44,
	0xbb, 0x43, 0x1b, 0x6a, 0x0a, 0x9d, 0xb5, 0xc7, 0xc9, 0xe9, 0x80, 0x04, 0x09, 0x9d, 0x40, 0xcb,
	0xee, 0xcd, 0x3a, 0xe3, 0xe4, 0xb4, 0x3f, 0x19, 0x9c, 0x85, 0xcb, 0x9d, 0xbd, 0xd1, 0x5c, 0x11,
	0x67, 0x42, 0x5f, 0x41, 0xc7, 0x07, 0xce, 0xba, 0xce, 0xe9, 0xa3, 0xe8, 0x74, 0xe9, 0xd4, 0x24,
	0x98, 0xd1, 0x17, 0xd0, 0x76, 0xe1, 0xb2, 0x9e, 0xf3, 0x3b, 0x88, 0x7e, 0x53, 0xab, 0x25, 0xde,
	0x88, 0x5f, 0x03, 0xf2, 0x59, 0xf3, 0x55, 0x6e, 0x08, 0xd7, 0x6b, 0x99, 0x6b, 0x8e, 0x9e, 0x43,
	0x9f, 0x47, 0x2c, 0x74, 0x96, 0x8c, 0xf7, 0x4e, 0xfb, 0x93, 0x8f, 0xe3, 0x09, 0x25, 0x4e, 0xa4,
	0xea, 0x87, 0xff, 0x4e, 0xa0, 0x65, 0xaf, 0xba, 0x83, 0x5e, 0x06, 0x5d, 0xa1, 0x29, 0x5b, 0x89,
	0xdc, 0x81, 0xd7, 0x23, 0x1b, 0x11, 0x21, 0x68, 0xe5, 0x74, 0xc5, 0x1d, 0x72, 0x29, 0x71, 0x6b,
	0xf4, 0x19, 0xa4, 0xda, 0x14, 0x8c, 0xe7, 0x26, 0x20, 0x97, 0x92, 0x52, 0x81, 0x0e, 0xa1, 0xcd,
	0x57, 0x54, 0x2c, 0x1d, 0x74, 0x29, 0xf1, 0x82, 0xdd, 0x43, 0x6f, 0xa9, 0xa1, 0xaa, 0x50, 0x4b,
	0x07, 0x5f, 0x4a, 0x4a, 0x45, 0x3d, 0x9f, 0xee, 0x3b, 0xe6, 0xf3, 0x0c, 0x06, 0x36, 0x1d, 0x1d,
	0x71, 0xf9, 0x1c, 0xda, 0x4e, 0x11, 0x10, 0xa9, 0x15, 0xc8, 0xdb, 0x70, 0x07, 0x5a, 0x3f, 0x49,
	0xc1, 0x30, 0x86, 0xe1, 0x94, 0x1b, 0xc2, 0xe7, 0x52, 0x31, 0xc2, 0xff, 0x28, 0xb8, 0xde, 0xa1,
	0x15, 0xfe, 0x16, 0x9e, 0xbc, 0x59, 0x33, 0x6a, 0xb8, 0x3b, 0x20, 0x38, 0x9d, 0x78, 0x14, 0x9d,
	0xdb, 0x2e, 0x0b, 0xec, 0x17, 0xff, 0x9b, 0x00, 0xbc, 0xd0, 0x5a, 0x2c, 0xf2, 0xf7, 0x66, 0xeb,
	0x7d, 0x88, 0x8f, 0xa0, 0xb7, 0xa4, 0xf9, 0xa2, 0xa0, 0x0b, 0x1e, 0x00, 0x8f, 0xb2, 0xb5, 0x31,
	0x4e, 0xd9, 0x52, 0xe4, 0x3c, 0x40, 0x1e, 0x65, 0x34, 0x86, 0x3e, 0x2d, 0x8c, 0xa4, 0xeb, 0xb5,
	0x92, 0xb7, 0xdc, 0xe1, 0xde, 0x23, 0x55, 0x95, 0xad, 0x96, 0x54, 0x8c, 0x2b, 0xc7, 0xd6, 0x01,
	0xf1, 0x02, 0xba, 0x00, 0xd0, 0xc5, 0x6c, 0x25, 0xb4, 0x16, 0x32, 0x0f, 0x04, 0x2d, 0xcb, 0x71,
	0x1d, 0x4d, 0xa4, 0xe2, 0x86, 0xff, 0x49, 0x00, 0x4a, 0xd3, 0x4e, 0xce, 0x18, 0xf6, 0x69, 0x44,
	0x24, 0xe6, 0xbd, 0xa5, 0xfb, 0x80, 0x4e, 0x3d, 0x84, 0xb6, 0x9e, 0x4b, 0xc5, 0x43, 0xa3, 0x7a,
	0xc1, 0xc6, 0x72, 0x0b, 0x39, 0xfb, 0x9d, 0xcf, 0x8d, 0x0e, 0x84, 0xdb, 0xd2, 0x59, 0x46, 0xce,
	0x0a, 0xb1, 0x64, 0x22, 0xff, 0x4d, 0xba, 0xec, 0x53, 0x52, 0x2a, 0xd0, 0x53, 0x80, 0xb9, 0x5c,
	0xad, 0x84, 0xb9, 0xa1, 0xfa, 0xc6, 0x21, 0x90, 0x92, 0x8a, 0x06, 0x5f, 0x41, 0xbf, 0xac, 0xaf,
	0xb6, 0x04, 0x2e, 0x13, 0xd9, 0x6d, 0xc8, 0xd2, 0x95, 0x54, 0xfd, 0x70, 0x17, 0xda, 0xae, 0xdb,
	0xf1, 0x5f, 0x4d, 0xe8, 0xf8, 0xf9, 0xb0, 0x83, 0xdb, 0x86, 0x0f, 0xcd, 0x0a, 0x1f, 0x10, 0xb4,
	0xe6, 0x92, 0x45, 0x8e, 0xd8, 0xb5, 0xd5, 0xdd, 0x71, 0xaa, 0x1c, 0x40, 0x03, 0xe2, 0xd6, 0x68,
	0x08, 0x7b, 0x86, 0x2e, 0x02, 0x2d, 0xec, 0xd2, 0xb2, 0xc5, 0x16, 0x5e, 0xb0, 0x30, 0xc5, 0x52,
	0x12, 0x65, 0xcb, 0x16, 0x26, 0x14, 0x9f, 0x1b, 0xa9, 0xee, 0x04, 0x73, 0x98, 0xb4, 0x48, 0x55,
	0x65, 0x77, 0xfb, 0xfe, 0xe3, 0xcc, 0x61, 0xd2, 0x26, 0x51, 0xae, 0x43, 0x90, 0xbe, 0x1b, 0x04,
	0xe8, 0x4b, 0xe8, 0xb8, 0x5a, 0xea, 0x0c, 0xdc, 0x8e, 0xfa, 0x1c, 0x0c, 0x56, 0xfc, 0x0c, 0xba,
	0x1e, 0x20, 0x8d, 0xbe, 0x8e, 0xcb, 0x00, 0xf4, 0xce, 0x8c, 0xdd, 0xd8, 0xf1, 0xf7, 0x70, 0xec,
	0x1b, 0xfc, 0x67, 0x61, 0x6e, 0xae, 0xdd, 0x10, 0x7f, 0xa0, 0xd5, 0x1d, 0x93, 0x0c, 0x35, 0x1b,
	0xa0, 0xbd, 0x80, 0xaf, 0xe0, 0xc0, 0x36, 0xf4, 0xab, 0x2b, 0x7f, 0xe2, 0xab, 0xab, 0x0a, 0x47,
	0x93, 0x2d, 0x8e, 0x3e, 0xd2, 0xd3, 0xf8, 0x1b, 0x00, 0x1f, 0xfc, 0xd2, 0x56, 0xea, 0x29, 0x80,
	0x7f, 0x4f, 0x5c, 0x0d, 0x13, 0x87, 0x63, 0x45, 0x83, 0x39, 0x1c, 0xfb, 0xa1, 0x53, 0x99, 0x7b,
	0xe1, 0xd2, 0x1f, 0x10, 0xbc, 0xf2, 0x98, 0xed, 0xb9, 0x50, 0x41, 0x9a, 0xfc, 0xd9, 0x81, 0x27,
	0x2f, 0x0a, 0x23, 0x17, 0x8a, 0x32, 0xae, 0xae, 0xfd, 0xb3, 0x8b, 0x9e, 0x43, 0x77, 0xca, 0x8d,
	0x7b, 0x25, 0x3e, 0x29, 0x4b, 0x51, 0x9b, 0x93, 0xa3, 0xed, 0xa1, 0x87, 0x1b, 0xe8, 0x02, 0x7a,
	0x61, 0x9b, 0x46, 0xa5, 0xd1, 0xce, 0xd9, 0xd1, 0xd1, 0x96, 0x6f, 0x1c, 0xd6, 0xb8, 0x81, 0xbe,
	0x03, 0x28, 0xa7, 0x2b, 0x1a, 0x95, 0x7e, 0xf5, 0x91, 0xbb, 0x1b, 0x6f, 0x02, 0xfb, 0x97, 0x8a,
	0x53, 0xc3, 0x43, 0xd7, 0xd4, 0x29, 0x30, 0xaa, 0x2b, 0x5c, 0xb8, 0x74, 0xca, 0x4d, 0xd8, 0xf0,
	0x48, 0x72, 0xf7, 0x6c, 0x9d, 0xc0, 0xbe, 0xbf, 0xd4, 0x7b, 0x84, 0x3b, 0x07, 0x88, 0xe1, 0x76,
	0x40, 0x19, 0xd6, 0xfc, 0x35, 0x6e, 0xa0, 0x1f, 0x21, 0x2b, 0x37, 0x58, 0xc2, 0x56, 0xfe, 0x77,
	0xc6, 0xd1, 0xff, 0x01, 0x3e, 0xdf, 0x7b, 0xe2, 0x25, 0x1c, 0x4c, 0xb9, 0xa9, 0x0e, 0xaa, 0x47,
	0xd2, 0x3e, 0xbc, 0xa7, 0x57, 0xed, 0x21, 0xbf, 0xc0, 0xd1, 0x94, 0x9b, 0xf2, 0x2e, 0xfa, 0xe5,
	0x5d, 0x40, 0xe1, 0xed, 0x97, 0xfa, 0xb4, 0xf6, 0x86, 0x57, 0xff, 0x62, 0x70, 0x03, 0xbd, 0x84,
	0xa1, 0xaf, 0x62, 0x25, 0xd3, 0xe3, 0xad, 0x52, 0x97, 0x8d, 0x37, 0xaa, 0x3c, 0x40, 0xb1, 0x97,
	0x70, 0x03, 0xbd, 0x86, 0x61, 0xbd, 0x5b, 0x2a, 0x17, 0x7b, 0xa0, 0x91, 0x1e, 0x38, 0x6c, 0xd6,
	0x71, 0x7f, 0x9d, 0x17, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xbf, 0x07, 0xac, 0x9a, 0x0a,
	0x00, 0x00,
}
