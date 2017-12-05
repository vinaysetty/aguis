package web

import (
	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/database"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AutograderService struct
type AutograderService struct {
	db *database.GormDB
}

// NewAutograderService a new grpc service
func NewAutograderService(db *database.GormDB) *AutograderService {
	return &AutograderService{
		db: db,
	}
}

// GetUser returns a user with a given id
func (s *AutograderService) GetUser(ctx context.Context, userQuery *pb.RecordRequest) (*pb.User, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return GetUser(userQuery, s.db)
}

// GetUsers returns all existing users
func (s *AutograderService) GetUsers(ctx context.Context, query *pb.Void) (*pb.Users, error) {
	return GetUsers(s.db)
}

// UpdateUser updates a user record
func (s *AutograderService) UpdateUser(ctx context.Context, u *pb.User) (*pb.User, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return UpdateUser(u, s.db)
}

// GetCourse returns a course record find by ID
func (s *AutograderService) GetCourse(ctx context.Context, query *pb.RecordRequest) (*pb.Course, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return GetCourse(query, s.db)
}

// GetCoursesWithEnrollment returns list of courses with enrollment status
func (s *AutograderService) GetCoursesWithEnrollment(
	ctx context.Context, request *pb.RecordWithStatusRequest) (*pb.Courses, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return ListCoursesWithEnrollment(request, s.db)
}

// GetCourses returns list of all courses
func (s *AutograderService) GetCourses(ctx context.Context, q *pb.Void) (*pb.Courses, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return ListCourses(s.db)
}

// GetAssignments lists the assignments for the provided course
func (s *AutograderService) GetAssignments(ctx context.Context, cid *pb.RecordRequest) (*pb.Assignments, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return ListAssignments(cid, s.db)
}

// GetEnrollmentsByCourse get all enrollments for a course.
func (s *AutograderService) GetEnrollmentsByCourse(
	ctx context.Context, req *pb.RecordWithStatusRequest) (*pb.EnrollmentResponse, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return GetEnrollmentsByCourse(req, s.db)
}

// CreateCourse creates a new course and assigns logged in user to the course
func (s *AutograderService) CreateCourse(ctx context.Context, c *pb.Course) (*pb.Course, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return NewCourse(c, s.db)
}

// UpdateCourse updates a course
func (s *AutograderService) UpdateCourse(ctx context.Context, c *pb.Course) (*pb.Course, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return UpdateCourse(c, s.db)
}

// CreateEnrollment add a user to a course
func (s *AutograderService) CreateEnrollment(ctx context.Context, ucid *pb.EnrollmentRequest) (*pb.StatusCode, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return CreateEnrollment(ucid, s.db)
}

// UpdateEnrollment accept/reject a user to a course
func (s *AutograderService) UpdateEnrollment(
	ctx context.Context, req *pb.EnrollmentRequest) (*pb.StatusCode, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return UpdateEnrollment(req, s.db)
}

// SetGrpcHeaderAndTrailer sets header and treailer to grpc
func SetGrpcHeaderAndTrailer(ctx context.Context) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))
}
