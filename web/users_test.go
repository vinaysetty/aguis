package web_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/web"
	"github.com/autograde/aguis/web/auth"
	"github.com/labstack/echo"
)

func TestGetSelf(t *testing.T) {
	const (
		selfURL   = "/user"
		apiPrefix = "/api/v1"
	)

	r := httptest.NewRequest(http.MethodGet, selfURL, nil)
	w := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(r, w)

	user := &pb.User{ID: 1}
	c.Set(auth.UserKey, user)

	userHandler := web.GetSelf()
	if err := userHandler(c); err != nil {
		t.Error(err)
	}

	userURL := "/users/" + strconv.FormatUint(user.ID, 10)
	location := w.Header().Get("Location")
	if location != apiPrefix+userURL {
		t.Errorf("have Location '%v' want '%v'", location, apiPrefix+userURL)
	}
	assertCode(t, w.Code, http.StatusFound)
}

func TestGetUser(t *testing.T) {
	const (
		route       = "/users/:uid"
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

	e := echo.New()
	router := echo.NewRouter(e)

	// Add the route to handler.
	router.Add(http.MethodGet, route, web.GetUser(db))

	requestURL := "/users/" + strconv.FormatUint(user.ID, 10)
	r := httptest.NewRequest(http.MethodGet, requestURL, nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	// Prepare context with user request.
	router.Find(http.MethodGet, requestURL, c)

	// Invoke the prepared handler.
	if err := c.Handler()(c); err != nil {
		t.Error(err)
	}

	var foundUser *pb.User
	if err := json.Unmarshal(w.Body.Bytes(), &foundUser); err != nil {
		t.Fatal(err)
	}

	// Access token should be stripped.
	user.RemoteIdentities[0].AccessToken = ""
	if !reflect.DeepEqual(foundUser, &user) {
		t.Errorf("have user %+v want %+v", foundUser, &user)
	}
	assertCode(t, w.Code, http.StatusFound)
}

func TestGetUsers(t *testing.T) {
	const (
		route = "/users"

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

	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, route, nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	h := web.GetUsers(db)
	if err := h(c); err != nil {
		t.Error(err)
	}

	var foundUsers []*pb.User
	if err := json.Unmarshal(w.Body.Bytes(), &foundUsers); err != nil {
		t.Fatal(err)
	}

	// Remote identities should not be loaded.
	user1.RemoteIdentities = nil
	user2.RemoteIdentities = nil
	// First user should be admin.
	user1.IsAdmin = true
	wantUsers := []*pb.User{&user1, &user2}
	if !reflect.DeepEqual(foundUsers, wantUsers) {
		t.Errorf("have users %+v want %+v", foundUsers, wantUsers)
	}

	assertCode(t, w.Code, http.StatusFound)
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
	const route = "/courses/:cid/users"

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

	e := echo.New()
	router := echo.NewRouter(e)

	// Add the route to handler.
	router.Add(http.MethodGet, route, web.GetEnrollmentsByCourse(db))
	requestURL := "/courses/" + strconv.FormatUint(allCourses[0].ID, 10) + "/users"
	r := httptest.NewRequest(http.MethodGet, requestURL, nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	// Prepare context with user request.
	router.Find(http.MethodGet, requestURL, c)

	// Invoke the prepared handler.
	if err := c.Handler()(c); err != nil {
		t.Error(err)
	}

	var foundEnrollments []*pb.Enrollment
	if err := json.Unmarshal(w.Body.Bytes(), &foundEnrollments); err != nil {
		t.Fatal(err)
	}
	var foundUsers []*pb.User
	for _, e := range foundEnrollments {
		// Remote identities should not be loaded.
		e.User.RemoteIdentities = nil
		foundUsers = append(foundUsers, e.User)
	}

	if !reflect.DeepEqual(foundUsers, wantUsers) {
		t.Errorf("have users %+v want %+v", foundUsers, wantUsers)
	}

	assertCode(t, w.Code, http.StatusOK)
}

func TestPatchUser(t *testing.T) {
	const route = "/users/:uid"

	db, cleanup := setup(t)
	defer cleanup()

	var user pb.User
	var remoteIdentity pb.RemoteIdentity
	if err := db.CreateUserFromRemoteIdentity(
		&user, &remoteIdentity,
	); err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	router := echo.NewRouter(e)

	// Add the route to handler.
	router.Add(http.MethodPatch, route, web.PatchUser(db))

	// Send empty request, the user should not be modified.
	emptyJSON, err := json.Marshal(&web.UpdateUserRequest{})
	if err != nil {
		t.Fatal(err)
	}
	requestBody := bytes.NewReader(emptyJSON)

	requestURL := "/users/" + strconv.FormatUint(user.ID, 10)
	r := httptest.NewRequest(http.MethodPatch, requestURL, requestBody)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	// Prepare context with user request.
	router.Find(http.MethodPatch, requestURL, c)

	// Invoke the prepared handler.
	if err := c.Handler()(c); err != nil {
		t.Error(err)
	}
	assertCode(t, w.Code, http.StatusNotModified)

	tmp := true
	// Send request with IsAdmin set to true, the user should become admin.
	trueJSON, err := json.Marshal(&web.UpdateUserRequest{
		IsAdmin: &tmp,
	})
	if err != nil {
		t.Fatal(err)
	}
	requestBody.Reset(trueJSON)

	r = httptest.NewRequest(http.MethodPatch, requestURL, requestBody)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w = httptest.NewRecorder()
	c.Reset(r, w)
	// Prepare context with user request.
	router.Find(http.MethodPatch, requestURL, c)

	// Invoke the prepared handler.
	if err := c.Handler()(c); err != nil {
		t.Error(err)
	}
	assertCode(t, w.Code, http.StatusOK)

	admin, err := db.GetUser(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !admin.IsAdmin {
		t.Error("expected user to have become admin")
	}

	// Send request with Name.
	nameChangeJSON, err := json.Marshal(&web.UpdateUserRequest{
		Name: "Scrooge McDuck",
	})
	if err != nil {
		t.Fatal(err)
	}
	requestBody.Reset(nameChangeJSON)

	r = httptest.NewRequest(http.MethodPatch, requestURL, requestBody)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w = httptest.NewRecorder()
	c.Reset(r, w)
	// Prepare context with user request.
	router.Find(http.MethodPatch, requestURL, c)

	// Invoke the prepared handler.
	if err := c.Handler()(c); err != nil {
		t.Error(err)
	}
	assertCode(t, w.Code, http.StatusOK)

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
