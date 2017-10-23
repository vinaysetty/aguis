package web

import (
	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/models"
	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateUserRequest updates a user object in the database.
type UpdateUserRequest struct {
	Name      string `json:"name"`
	StudentID string `json:"studentid"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarurl"`
	IsAdmin   *bool  `json:"isadmin"`
}

//
//// GetSelf redirects to GetUser with the current user's id.
//func GetSelf() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		// If type assertions fails, the recover middleware will catch the panic and log a stack trace.
//		user := c.Get("user").(*models.User)
//		return c.Redirect(http.StatusFound, fmt.Sprintf("/api/v1/users/%d", user.ID))
//	}
//}
//
// GetUser returns information about the provided user id.
func GetUser(query *pb.GetRecordRequest, db database.Database) (*pb.User, error) {
	user, err := db.GetUser(query.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, err
	}

	return toProtoUser(user), nil
}

// GetUsers returns all the users in the database.
func GetUsers(db database.Database) (*pb.UsersResponse, error) {

	var results []*pb.User
	users, err := db.GetUsers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "Users not found")
		}
		return nil, err
	}
	for _, u := range users {
		results = append(results, toProtoUser(u))
	}
	return &pb.UsersResponse{Users: results}, nil
}

// UpdateUser promotes a user to an administrator
func UpdateUser(userReq *pb.UpdateUserRequest, db database.Database) (*pb.User, error) {
	user, err := db.GetUser(userReq.User.Id)
	if err != nil {
		return nil, err
	}

	user.Name = userReq.User.Name
	user.Email = userReq.User.Email
	user.AvatarURL = userReq.User.Avatarurl
	user.StudentID = userReq.User.Studentid
	user.IsAdmin = userReq.User.Isadmin
	if err := db.UpdateUser(user); err != nil {
		return nil, err
	}

	return userReq.User, nil
}

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

func toProtoUser(user *models.User) *pb.User {
	pu := &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Studentid: user.StudentID,
		Avatarurl: user.AvatarURL,
		Isadmin:   user.IsAdmin,
	}
	return pu
}
