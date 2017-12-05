package web

import (
	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/database"
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
	// Remove access token for user
	for _, remoteID := range user.GetRemoteIdentities() {
		remoteID.AccessToken = ""
	}
	return user, nil
}

// GetUsers returns all the users in the database.
func GetUsers(db database.Database) (*pb.Users, error) {
	// This call does not preload the remote identities,
	// and therefore we do not need to remove the access token.
	users, err := db.GetUsers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "No users found")
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
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.StudentID != "" {
		user.StudentID = userReq.StudentID
	}
	if userReq.Email != "" {
		user.Email = userReq.Email
	}
	if userReq.AvatarURL != "" {
		user.AvatarURL = userReq.AvatarURL
	}
	if userReq.IsAdmin {
		user.IsAdmin = userReq.IsAdmin
	}
	if err := db.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}
