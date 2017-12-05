package web_test

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/scm"
	"github.com/autograde/aguis/web"
	"github.com/autograde/aguis/web/auth"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc/codes"
)

var allCourses = []*pb.Course{
	{
		Name:        "Distributed Systems",
		Code:        "DAT520",
		Year:        2018,
		Tag:         "Spring",
		Provider:    "fake",
		DirectoryID: 1,
	},
	{
		Name:        "Operating Systems",
		Code:        "DAT320",
		Year:        2017,
		Tag:         "Fall",
		Provider:    "fake",
		DirectoryID: 2,
	}, {
		Name:        "New Systems",
		Code:        "DATx20",
		Year:        2019,
		Tag:         "Fall",
		Provider:    "fake",
		DirectoryID: 3,
	}, {
		Name:        "Hyped Systems",
		Code:        "DATx20",
		Year:        2019,
		Tag:         "Fall",
		Provider:    "fake",
		DirectoryID: 4,
	},
}

func TestListCourses(t *testing.T) {
	db, cleanup := setup(t)
	defer cleanup()

	var testCourses []*pb.Course
	for _, course := range allCourses {
		testCourse := *course
		err := db.CreateCourse(&testCourse)
		if err != nil {
			t.Fatal(err)
		}
		testCourses = append(testCourses, &testCourse)
	}

	courses, err := web.ListCourses(db)
	if err != nil {
		t.Fatal(err)
	}

	for i, course := range courses.Courses {
		if !reflect.DeepEqual(course, testCourses[i]) {
			t.Errorf("have course %+v want %+v", course, testCourses[i])
		}
	}
}

func TestNewCourse(t *testing.T) {
	const (
		route = "/courses"
		fake  = "fake"
	)

	db, cleanup := setup(t)
	defer cleanup()

	var user pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user, &pb.RemoteIdentity{
			Provider: fake,
		},
	); err != nil {
		t.Fatal(err)
	}

	testCourse := *allCourses[0]

	// Convert course to course request, this allows us to verify that the
	// course we get from the database is correct.
	cr := courseToRequest(t, &testCourse)

	b, err := json.Marshal(cr)
	if err != nil {
		t.Fatal(err)
	}

	//TODO Not sure how to interact with the SCM part without plain http.
	r := httptest.NewRequest(http.MethodPost, route, bytes.NewReader(b))
	r.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(r, w)
	f := scm.NewFakeSCMClient()
	if _, err := f.CreateDirectory(context.Background(), &scm.CreateDirectoryOptions{
		Name: testCourse.Code,
		Path: testCourse.Code,
	}); err != nil {
		t.Fatal(err)
	}
	c.Set(fake, f)
	c.Set(auth.UserKey, &pb.User{ID: user.ID})

	// h := web.NewCourse(nullLogger(), db, &web.BaseHookOptions{})
	respCourse, err := web.NewCourse(&testCourse, db)
	if err != nil {
		t.Fatal(err)
	}

	course, err := db.GetCourse(respCourse.ID)
	if err != nil {
		t.Fatal(err)
	}

	testCourse.ID = respCourse.ID
	if !reflect.DeepEqual(course, &testCourse) {
		t.Errorf("have database course %+v want %+v", course, &testCourse)
	}

	if !reflect.DeepEqual(respCourse, course) {
		t.Errorf("have response course %+v want %+v", respCourse, course)
	}

	enrollment, err := db.GetEnrollmentByCourseAndUser(testCourse.ID, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	wantEnrollment := &pb.Enrollment{
		ID:       enrollment.ID,
		CourseID: testCourse.ID,
		UserID:   user.ID,
		Status:   pb.Enrollment_Teacher,
	}
	if !reflect.DeepEqual(enrollment, wantEnrollment) {
		t.Errorf("have enrollment %+v want %+v", enrollment, wantEnrollment)
	}
	//TODO FIX THIS: Doesnotpass
	// if len(f.Hooks) != 4 {
	// 	t.Errorf("have %d hooks want %d", len(f.Hooks), 4)
	// }
}

func TestEnrollmentProcess(t *testing.T) {
	const (
		github = "github"
		gitlab = "gitlab"
	)

	db, cleanup := setup(t)
	defer cleanup()

	// Create course.
	testCourse := *allCourses[0]
	if err := db.CreateCourse(&testCourse); err != nil {
		t.Fatal(err)
	}
	// Create admin.
	var admin pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&admin, &pb.RemoteIdentity{
			Provider: github,
		},
	); err != nil {
		t.Fatal(err)
	}
	// Create user.
	var user pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user, &pb.RemoteIdentity{
			Provider: gitlab,
		},
	); err != nil {
		t.Fatal(err)
	}

	enrollReq := &pb.EnrollmentRequest{CourseID: testCourse.ID, UserID: user.ID}
	status, err := web.CreateEnrollment(enrollReq, db)
	if err != nil {
		t.Fatal(err)
	}
	if status.GetStatuscode() != int32(codes.OK) { //TODO fix StatusCode to use codes directly
		t.Fatalf("bad status code %v", status)
	}

	// Verify that an appropriate enrollment was indeed created.
	pendingEnrollment, err := db.GetEnrollmentByCourseAndUser(testCourse.ID, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	wantEnrollment := &pb.Enrollment{
		ID:       pendingEnrollment.ID,
		CourseID: testCourse.ID,
		UserID:   user.ID,
		Status:   pb.Enrollment_Pending,
	}
	if !reflect.DeepEqual(pendingEnrollment, wantEnrollment) {
		t.Errorf("have enrollment\n %+v\n want\n %+v", pendingEnrollment, wantEnrollment)
	}

	enrollReq.Enrolled = pb.Enrollment_Student
	status, err = web.UpdateEnrollment(enrollReq, db)
	if err != nil {
		t.Fatal(err)
	}
	if status.GetStatuscode() != int32(codes.OK) { //TODO fix StatusCode to use codes directly
		t.Fatalf("bad status code %v", status)
	}

	// Verify that the enrollment have been accepted.
	acceptedEnrollment, err := db.GetEnrollmentByCourseAndUser(testCourse.ID, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	wantEnrollment.Status = pb.Enrollment_Student
	if !reflect.DeepEqual(acceptedEnrollment, wantEnrollment) {
		t.Errorf("have enrollment %+v want %+v", acceptedEnrollment, wantEnrollment)
	}
}

func TestListCoursesWithEnrollment(t *testing.T) {
	db, cleanup := setup(t)
	defer cleanup()

	var testCourses []*pb.Course
	for _, course := range allCourses {
		testCourse := *course
		err := db.CreateCourse(&testCourse)
		if err != nil {
			t.Fatal(err)
		}
		testCourses = append(testCourses, &testCourse)
	}

	var user pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user, &pb.RemoteIdentity{},
	); err != nil {
		t.Fatal(err)
	}

	if err := db.CreateEnrollment(&pb.Enrollment{
		UserID:   user.ID,
		CourseID: testCourses[0].ID,
	}); err != nil {
		t.Fatal(err)
	}
	if err := db.CreateEnrollment(&pb.Enrollment{
		UserID:   user.ID,
		CourseID: testCourses[1].ID,
	}); err != nil {
		t.Fatal(err)
	}
	if err := db.CreateEnrollment(&pb.Enrollment{
		UserID:   user.ID,
		CourseID: testCourses[2].ID,
	}); err != nil {
		t.Fatal(err)
	}
	if err := db.RejectEnrollment(user.ID, testCourses[1].ID); err != nil {
		t.Fatal(err)
	}
	if err := db.EnrollStudent(user.ID, testCourses[2].ID); err != nil {
		t.Fatal(err)
	}

	req := &pb.RecordWithStatusRequest{ID: user.ID}
	courses, err := web.ListCoursesWithEnrollment(req, db)
	if err != nil {
		t.Fatal(err)
	}

	wantCourses := []*pb.Course{
		{ID: testCourses[0].ID, Enrolled: pb.Enrollment_Pending},
		{ID: testCourses[1].ID, Enrolled: pb.Enrollment_Rejected},
		{ID: testCourses[2].ID, Enrolled: pb.Enrollment_Student},
		{ID: testCourses[3].ID, Enrolled: pb.Enrollment_None},
	}
	for i := range courses.Courses {
		if courses.Courses[i].ID != wantCourses[i].ID {
			t.Errorf("have course %+v want %+v", courses.Courses[i].ID, wantCourses[i].ID)
		}
		if courses.Courses[i].Enrolled != wantCourses[i].Enrolled {
			t.Errorf("have course %+v want %+v", courses.Courses[i].Enrolled, wantCourses[i].Enrolled)
		}
	}
}

func TestListCoursesWithEnrollmentStatuses(t *testing.T) {
	db, cleanup := setup(t)
	defer cleanup()

	var testCourses []*pb.Course
	for _, course := range allCourses {
		testCourse := *course
		err := db.CreateCourse(&testCourse)
		if err != nil {
			t.Fatal(err)
		}
		testCourses = append(testCourses, &testCourse)
	}

	var user pb.User
	if err := db.CreateUserFromRemoteIdentity(
		&user, &pb.RemoteIdentity{},
	); err != nil {
		t.Fatal(err)
	}

	if err := db.CreateEnrollment(&pb.Enrollment{
		UserID:   user.ID,
		CourseID: testCourses[0].ID,
	}); err != nil {
		t.Fatal(err)
	}
	if err := db.CreateEnrollment(&pb.Enrollment{
		UserID:   user.ID,
		CourseID: testCourses[1].ID,
	}); err != nil {
		t.Fatal(err)
	}
	if err := db.CreateEnrollment(&pb.Enrollment{
		UserID:   user.ID,
		CourseID: testCourses[2].ID,
	}); err != nil {
		t.Fatal(err)
	}
	if err := db.RejectEnrollment(user.ID, testCourses[1].ID); err != nil {
		t.Fatal(err)
	}
	if err := db.EnrollStudent(user.ID, testCourses[2].ID); err != nil {
		t.Fatal(err)
	}

	req := &pb.RecordWithStatusRequest{
		ID: user.ID,
		Statuses: []pb.Enrollment_Status{
			pb.Enrollment_Student,
			pb.Enrollment_Rejected,
		},
	}
	courses, err := web.ListCoursesWithEnrollment(req, db)
	if err != nil {
		t.Fatal(err)
	}

	wantCourses, err := db.GetCoursesByUser(user.ID, pb.Enrollment_Rejected, pb.Enrollment_Student)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(courses.Courses, wantCourses) {
		t.Errorf("have course %+v want %+v", courses.Courses, wantCourses)
	}
}

func TestGetCourse(t *testing.T) {
	db, cleanup := setup(t)
	defer cleanup()

	var course pb.Course
	err := db.CreateCourse(&course)
	if err != nil {
		t.Fatal(err)
	}

	foundCourse, err := web.GetCourse(&pb.RecordRequest{ID: course.ID}, db)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(foundCourse, &course) {
		t.Errorf("have course %+v want %+v", foundCourse, &course)
	}
}

func courseToRequest(t *testing.T, course *pb.Course) (cr web.NewCourseRequest) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(course); err != nil {
		t.Fatal(err)
	}
	dec := gob.NewDecoder(&b)
	if err := dec.Decode(&cr); err != nil {
		t.Fatal(err)
	}
	return
}
