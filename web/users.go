package web

import (
	"github.com/autograde/aguis/database"
	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//TODO remove, only used by users_test.go; can just use pb.User
// UpdateUserRequest updates a user object in the database.
type UpdateUserRequest struct {
	Name      string `json:"name"`
	StudentID string `json:"studentid"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarurl"`
	IsAdmin   *bool  `json:"isadmin"`
}

//TODO How can we get the current user with gRPC???
//// GetSelf redirects to GetUser with the current user's id.
//func GetSelf() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		// If type assertions fails, the recover middleware will catch the panic and log a stack trace.
//		user := c.Get("user").(*models.User)
//		return c.Redirect(http.StatusFound, fmt.Sprintf("/api/v1/users/%d", user.ID))
//	}
//}

// GetUser returns information about the provided user id.
func GetUser(query *pb.RecordRequest, db database.Database) (*pb.User, error) {
	user, err := db.GetUser(query.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, err
	}
	return user, nil
}

// GetUsers returns all the users in the database.
func GetUsers(db database.Database) (*pb.Users, error) {
	users, err := db.GetUsers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "Users not found")
		}
		return nil, err
	}
	return &pb.Users{Users: users}, nil
}

// UpdateUser promotes a user to an administrator or makes other changes to the user database entry.
func UpdateUser(userReq *pb.User, db database.Database) (*pb.User, error) {
	user, err := db.GetUser(userReq.ID)
	if err != nil {
		return nil, err
	}

	user.Name = userReq.Name
	user.Email = userReq.Email
	user.AvatarURL = userReq.AvatarURL
	user.StudentID = userReq.StudentID
	user.IsAdmin = userReq.IsAdmin
	if err := db.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

//TODO Move this to groups.go
// GetGroupByUserAndCourse returns a single group of a user for a course
// func GetGroupByUserAndCourse(db database.Database) (*pb.Group, error) {
// }

//
//// GetGroupByUserAndCourse returns a single group of a user for a course
//func GetGroupByUserAndCourse(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		uid, err := parseUint(c.Param("uid"))
//		if err != nil {
//			return err
//		}
//		cid, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return nil
//		}
//		enrollment, err := db.GetEnrollmentByCourseAndUser(cid, uid)
//		if err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return c.NoContent(http.StatusNotFound)
//			}
//			return err
//		}
//		if enrollment.GroupID > 0 {
//			group, err := db.GetGroup(enrollment.GroupID)
//			if err != nil {
//				return nil
//			}
//			return c.JSONPretty(http.StatusFound, group, "\t")
//		}
//		return c.NoContent(http.StatusNotFound)
//	}
//}
