package grpcutil

import (
	"github.com/autograde/aguis/database"
	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/autograde/aguis/web"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type autograderService struct {
	db *database.GormDB
}

func NewAutograderService(db *database.GormDB) *autograderService {
	return &autograderService{
		db: db,
	}
}

func (s *autograderService) GetUser(ctx context.Context, userQuery *pb.GetRecordRequest) (*pb.User, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.GetUser(userQuery, s.db)
}

func (s *autograderService) GetUsers(ctx context.Context, query *pb.Void) (*pb.UsersResponse, error) {
	return web.GetUsers(s.db)
}

func (s *autograderService) UpdateUser(ctx context.Context, u *pb.UpdateUserRequest) (*pb.User, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.UpdateUser(u, s.db)
}

func (s *autograderService) GetCourse(ctx context.Context, query *pb.GetRecordRequest) (*pb.Course, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.GetCourse(query, s.db)
}

func (s *autograderService) GetCoursesWithEnrollment(
	ctx context.Context, request *pb.RecordWithStatusRequest) (*pb.Courses, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.ListCoursesWithEnrollment(request, s.db)
}

func (s *autograderService) GetCourses(ctx context.Context, q *pb.Void) (*pb.Courses, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.ListCourses(s.db)
}

func (s *autograderService) GetAssignments(ctx context.Context, cid *pb.GetRecordRequest) (*pb.Assignments, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.ListAssignments(cid, s.db)

}

func (s *autograderService) GetEnrollmentsByCourse(
	ctx context.Context, req *pb.RecordWithStatusRequest) (*pb.EnrollemntResponse, error) {
	SetGrpcHeaderAndTrailer(ctx)
	return web.GetEnrollmentsByCourse(req, s.db)
}

func SetGrpcHeaderAndTrailer(ctx context.Context) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

}
