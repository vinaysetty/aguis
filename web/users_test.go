package web_test

import (
	"reflect"
	"testing"

	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/web"
)

func TestGetSelf(t *testing.T) {
	// const (
	// 	selfURL   = "/user"
	// 	apiPrefix = "/api/v1"
	// )

	// r := httptest.NewRequest(http.MethodGet, selfURL, nil)
	// w := httptest.NewRecorder()
	// e := echo.New()
	// c := e.NewContext(r, w)

	// user := &pb.User{ID: 1}
	// c.Set(auth.UserKey, user)

	// userHandler := web.GetSelf()
	// if err := userHandler(c); err != nil {
	// 	t.Error(err)
	// }

	// userURL := "/users/" + strconv.FormatUint(user.ID, 10)
	// location := w.Header().Get("Location")
	// if location != apiPrefix+userURL {
	// 	t.Errorf("have Location '%v' want '%v'", location, apiPrefix+userURL)
	// }
	// assertCode(t, w.Code, http.StatusFound)
}

func TestGetUser(t *testing.T) {
	const (
		provider    = "github"
		accessToken = "secret"
	)

	db, cleanup := setup(t)
	defer cleanup()

	// Create first user (the admin).
	if err := db.CreateUserFromRemoteIdentity(
		&pb.User{},
		&pb.RemoteIdentity{},
	); err != nil {
		t.Fatal(err)
	}

	var user pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user,
		&pb.RemoteIdentity{
			Provider:    provider,
			AccessToken: accessToken,
		},
	); err != nil {
		t.Fatal(err)
	}

	foundUser, err := web.GetUser(&pb.RecordRequest{ID: user.ID}, db)
	if err != nil {
		t.Fatal(err)
	}

	// Access token should be stripped by web.GetUser().
	user.RemoteIdentities[0].AccessToken = ""
	if !reflect.DeepEqual(foundUser, &user) {
		t.Errorf("have user %+v want %+v", foundUser, &user)
	}
}

func TestGetUsers(t *testing.T) {
	const (
		github = "github"
		gitlab = "gitlab"
	)

	db, cleanup := setup(t)
	defer cleanup()

	var user1 pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user1,
		&pb.RemoteIdentity{
			Provider: github,
		},
	); err != nil {
		t.Fatal(err)
	}
	var user2 pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user2,
		&pb.RemoteIdentity{
			Provider: gitlab,
		},
	); err != nil {
		t.Fatal(err)
	}

	foundUsers, err := web.GetUsers(db)
	if err != nil {
		t.Fatal(err)
	}

	// Remote identities should not be loaded.
	user1.RemoteIdentities = nil
	user2.RemoteIdentities = nil
	// First user should be admin.
	user1.IsAdmin = true
	wantUsers := []*pb.User{&user1, &user2}
	if !reflect.DeepEqual(foundUsers.Users, wantUsers) {
		t.Errorf("have users %+v want %+v", foundUsers.Users, wantUsers)
	}
}

var allUsers = []struct {
	provider string
	remoteID uint64
	secret   string
}{
	{"github", 1, "123"},
	{"github", 2, "456"},
	{"gitlab", 3, "789"},
	{"gitlab", 4, "012"},
	{"bitlab", 5, "345"},
	{"gitlab", 6, "678"},
	{"gitlab", 7, "901"},
	{"gitlab", 8, "234"},
}

func TestGetEnrollmentsByCourse(t *testing.T) {
	db, cleanup := setup(t)
	defer cleanup()

	var users []*pb.User
	for _, u := range allUsers {
		var user pb.User
		if err := db.CreateUserFromRemoteIdentity(&user, &pb.RemoteIdentity{
			Provider:    u.provider,
			RemoteID:    u.remoteID,
			AccessToken: u.secret,
		}); err != nil {
			t.Fatal(err)
		}
		// Remote identities should not be loaded.
		user.RemoteIdentities = nil
		users = append(users, &user)
	}

	for _, course := range allCourses {
		err := db.CreateCourse(course)
		if err != nil {
			t.Fatal(err)
		}
	}

	// users to enroll in course DAT520 Distributed Systems
	wantUsers := users[0 : len(allUsers)-3]

	// users to enroll in course DAT320 Operating Systems
	osUsers := users[3:7]

	for _, user := range wantUsers {
		if err := db.CreateEnrollment(&pb.Enrollment{
			UserID:   user.ID,
			CourseID: allCourses[0].ID,
		}); err != nil {
			t.Fatal(err)
		}
		if err := db.EnrollStudent(user.ID, allCourses[0].ID); err != nil {
			t.Fatal(err)
		}
	}

	for _, user := range osUsers {
		if err := db.CreateEnrollment(&pb.Enrollment{
			UserID:   user.ID,
			CourseID: allCourses[1].ID,
		}); err != nil {
			t.Fatal(err)
		}
		if err := db.EnrollStudent(user.ID, allCourses[1].ID); err != nil {
			t.Fatal(err)
		}
	}

	req := &pb.RecordWithStatusRequest{
		ID: allCourses[0].ID,
	}
	foundEnrollments, err := web.GetEnrollmentsByCourse(req, db)
	if err != nil {
		t.Fatal(err)
	}
	var foundUsers []*pb.User
	for _, e := range foundEnrollments.Enrollments {
		// Remote identities should not be loaded.
		e.User.RemoteIdentities = nil
		foundUsers = append(foundUsers, e.User)
	}

	if !reflect.DeepEqual(foundUsers, wantUsers) {
		t.Errorf("have users %+v want %+v", foundUsers, wantUsers)
	}
}

func TestUpdateUser(t *testing.T) {
	db, cleanup := setup(t)
	defer cleanup()

	var user, admin pb.User
	var remoteIdentity pb.RemoteIdentity
	// First user is always admin
	if err := db.CreateUserFromRemoteIdentity(
		&admin, &remoteIdentity,
	); err != nil {
		t.Fatal(err)
	}
	// Second user shouldn't be admin
	if err := db.CreateUserFromRemoteIdentity(
		&user, &remoteIdentity,
	); err != nil {
		t.Fatal(err)
	}

	// Send empty update request, the user should not be modified.
	noChngUser, err := web.UpdateUser(&pb.User{ID: user.ID}, db)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(&user, noChngUser) {
		t.Errorf("have user %+v want %+v", &user, noChngUser)
	}

	// Send request with IsAdmin set to true, the user should become admin.
	adminUser, err := web.UpdateUser(&pb.User{ID: user.ID, IsAdmin: true}, db)
	if err != nil {
		t.Fatal(err)
	}
	newAdmin, err := db.GetUser(user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !newAdmin.IsAdmin {
		t.Error("expected user to have become admin")
	}
	if !reflect.DeepEqual(newAdmin, adminUser) {
		t.Errorf("have user %+v want %+v", newAdmin, adminUser)
	}

	// Send request with Name.
	_, err = web.UpdateUser(&pb.User{ID: user.ID, Name: "Scrooge McDuck"}, db)
	if err != nil {
		t.Fatal(err)
	}
	withName, err := db.GetUser(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	wantUser := &pb.User{
		ID:               withName.ID,
		Name:             "Scrooge McDuck",
		IsAdmin:          true,
		RemoteIdentities: []*pb.RemoteIdentity{&remoteIdentity},
	}
	if !reflect.DeepEqual(withName, wantUser) {
		t.Errorf("have users %+v want %+v", withName, wantUser)
	}
}
