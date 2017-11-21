package web

import (
	"time"

	//"github.com/autograde/aguis/ci"
	"github.com/autograde/aguis/yamlparser"

	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/models"
	//"github.com/autograde/aguis/scm"
	//"github.com/jinzhu/gorm"
	//"github.com/sirupsen/logrus"
	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MaxWait is the maximum time a request is allowed to stay open before
// aborting.
const MaxWait = 10 * time.Minute

// NewCourseRequest represents a request for a new course.
type NewCourseRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Year uint   `json:"year"`
	Tag  string `json:"tag"`

	Provider    string `json:"provider"`
	DirectoryID uint64 `json:"directoryid"`
}

func (cr *NewCourseRequest) valid() bool {
	return cr != nil &&
		cr.Name != "" &&
		cr.Code != "" &&
		(cr.Provider == "github" || cr.Provider == "gitlab" || cr.Provider == "fake") &&
		cr.DirectoryID != 0 &&
		cr.Year != 0 &&
		cr.Tag != ""
}

// EnrollUserRequest represent a request for enrolling a user to a course.
type EnrollUserRequest struct {
	Status uint `json:"status"`
}

func (eur *EnrollUserRequest) valid() bool {
	return eur.Status <= models.Teacher
}

// NewGroupRequest represents a new group.
type NewGroupRequest struct {
	Name     string   `json:"name"`
	CourseID uint64   `json:"courseid"`
	UserIDs  []uint64 `json:"userids"`
}

func (grp *NewGroupRequest) valid() bool {
	return grp != nil &&
		grp.Name != "" &&
		len(grp.UserIDs) > 0
}

// UpdateGroupRequest updates group
type UpdateGroupRequest struct {
	Status uint `json:"status"`
}

// ListCourses returns a JSON object containing all the courses in the database.
func ListCourses(db database.Database) (*pb.Courses, error) {
	var results []*pb.Course
	courses, err := db.GetCourses()
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		results = append(results, toProtoCourse(course))
	}
	return &pb.Courses{Courses: results}, nil
}

// ListCoursesWithEnrollment lists all existing courses with the provided users
// enrollment status.
// If status query param is provided, lists only courses of the student filtered by the query param.
func ListCoursesWithEnrollment(request *pb.RecordWithStatusRequest, db database.Database) (*pb.Courses, error) {
	var results []*pb.Course
	id := request.Id
	statuses, err := parseEnrollmentStatus(request.State)
	if err != nil {
		return nil, err
	}

	courses, err := db.GetCoursesByUser(id, statuses...)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		results = append(results, toProtoCourse(course))
	}
	return &pb.Courses{Courses: results}, nil
}

// ListAssignments lists the assignments for the provided course.
func ListAssignments(request *pb.GetRecordRequest, db database.Database) (*pb.Assignments, error) {
	var results []*pb.Assignment
	assignments, err := db.GetAssignmentsByCourse(request.Id)
	if err != nil {
		return nil, err
	}
	for _, asg := range assignments {
		results = append(results, toProtoAssignment(asg))
	}
	return &pb.Assignments{Assignments: results}, nil
}

// Default repository names.
const (
	InfoRepo       = "course-info"
	AssignmentRepo = "assignments"
	TestsRepo      = "tests"
	SolutionsRepo  = "solutions"
)

// BaseHookOptions contains options shared among all webhooks.
type BaseHookOptions struct {
	BaseURL string
	// Secret is used to verify that the event received is legit. GitHub
	// sends back a signature of the payload, while GitLab just sends back
	// the secret. This is all handled by the
	// gopkg.in/go-playground/webhooks.v3 package.
	Secret string
}

// NewCourse creates a new course and associates it with a directory (organization in github)
// and creates the repositories for the course.
func NewCourse(crs *pb.Course, db database.Database) (*pb.Course, error) {
	newcrs := NewCourseRequest{
		Name:        crs.Name,
		Code:        crs.Code,
		Year:        uint(crs.Year),
		Tag:         crs.Tag,
		Provider:    crs.Provider,
		DirectoryID: crs.Directoryid,
	}
	if !newcrs.valid() {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload")
	}

	// TODO CreateCourse and CreateEnrollment should be combined into a method with transactions.
	course := models.Course{
		Name:        crs.Name,
		Code:        crs.Code,
		Year:        uint(crs.Year),
		Tag:         crs.Tag,
		Provider:    crs.Provider,
		DirectoryID: crs.Directoryid,
	}

	if err := db.CreateCourse(&course); err != nil {
		if err == database.ErrCourseExists {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, err
	}

	// Automatically enroll the teacher creating the course
	// TODO get logged in user in stead of hard coding
	user, err := db.GetUser(1)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "login user not found")
	}
	if err := db.CreateEnrollment(&models.Enrollment{
		UserID:   user.ID,
		CourseID: course.ID,
	}); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if err := db.EnrollTeacher(user.ID, course.ID); err != nil {
		return nil, err
	}
	return toProtoCourse(&course), nil
}

// CreateEnrollment enrolls a user in a course.
func CreateEnrollment(ucid *pb.UserIDCourseID, db database.Database) (*pb.StatusCode, error) {
	userID := ucid.Userid
	courseID := ucid.Courseid

	enrollment := models.Enrollment{
		UserID:   userID,
		CourseID: courseID,
	}
	if err := db.CreateEnrollment(&enrollment); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "record not found")
		}
		return nil, err
	}

	return &pb.StatusCode{Statuscode: int32(codes.OK)}, nil
}

// UpdateEnrollment accepts or rejects a user to enroll in a course.
//func UpdateEnrollment(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		courseID, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return err
//		}
//		userID, err := parseUint(c.Param("uid"))
//		if err != nil {
//			return err
//		}
//
//		var eur EnrollUserRequest
//		if err := c.Bind(&eur); err != nil {
//			return err
//		}
//		if !eur.valid() || userID == 0 || courseID == 0 {
//			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
//		}
//
//		if _, err := db.GetEnrollmentByCourseAndUser(courseID, userID); err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return c.NoContent(http.StatusNotFound)
//			}
//			return err
//		}
//
//		// If type assertions fails, the recover middleware will catch the panic and log a stack trace.
//		user := c.Get("user").(*models.User)
//		// TODO: This check should be performed in AccessControl.
//		if !user.IsAdmin {
//			// Only admin users are allowed to enroll or reject users to a course.
//			// TODO we should also allow users of the 'teachers' team to accept/reject users
//			return c.NoContent(http.StatusUnauthorized)
//		}
//
//		// TODO If the enrollment is accepted, create repositories with webooks.
//		switch eur.Status {
//		case models.Student:
//			err = db.EnrollStudent(userID, courseID)
//		case models.Teacher:
//			err = db.EnrollTeacher(userID, courseID)
//		case models.Rejected:
//			err = db.RejectEnrollment(userID, courseID)
//		}
//		if err != nil {
//			return err
//		}
//		return c.NoContent(http.StatusOK)
//	}
//}

// GetCourse find course by id and return JSON object.
func GetCourse(query *pb.GetRecordRequest, db database.Database) (*pb.Course, error) {
	course, err := db.GetCourse(query.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "Course not found")
		}
		return nil, err

	}

	return toProtoCourse(course), nil
}

// RefreshCourse refreshes the information to a course
//func RefreshCourse(logger logrus.FieldLogger, db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		id, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return err
//		}
//
//		user := c.Get("user").(*models.User)
//
//		course, err := db.GetCourse(id)
//		if err != nil {
//			return err
//		}
//		if c.Get(course.Provider) == nil {
//			return echo.NewHTTPError(http.StatusBadRequest, "provider "+course.Provider+" not registered")
//		}
//
//		remoteID, err := getRemoteIDFor(user, course.Provider)
//		if err != nil {
//			return err
//		}
//
//		s := c.Get(course.Provider).(scm.SCM)
//
//		ctx, cancel := context.WithTimeout(c.Request().Context(), MaxWait)
//		defer cancel()
//
//		directory, err := s.GetDirectory(ctx, course.DirectoryID)
//		if err != nil {
//			return err
//		}
//
//		path, err := s.CreateCloneURL(ctx, &scm.CreateClonePathOptions{
//			Directory:  directory.Path,
//			Repository: "tests.git",
//			UserToken:  remoteID.AccessToken,
//		})
//		if err != nil {
//			return err
//		}
//
//		runner := ci.Local{}
//
//		// This does not work that well on Windows because the path should be
//		// /mnt/c/Users/{user}/AppData/Local/Temp
//		// cloneDirectory := filepath.Join(os.TempDir(), "agclonepath")
//		cloneDirectory := "agclonepath"
//
//		// Clone all tests from tests repositry
//		_, err = runner.Run(c.Request().Context(), &ci.Job{
//			Commands: []string{
//				"mkdir " + cloneDirectory,
//				"cd " + cloneDirectory,
//				"git clone " + path,
//			},
//		})
//		if err != nil {
//			runner.Run(c.Request().Context(), &ci.Job{
//				Commands: []string{
//					"yes | rm -r " + cloneDirectory,
//				},
//			})
//			return err
//		}
//
//		// Parse assignments in the test directory
//		assignments, err := yamlparser.Parse("agclonepath/tests")
//		if err != nil {
//			return err
//		}
//
//		// Cleanup downloaded
//		runner.Run(c.Request().Context(), &ci.Job{
//			Commands: []string{
//				"yes | rm -r " + cloneDirectory,
//			},
//		})
//
//		for _, v := range assignments {
//			assignment, err := createAssignment(&v, course)
//			if err != nil {
//				return err
//			}
//			if err := db.CreateAssignment(assignment); err != nil {
//				return err
//			}
//		}
//
//		return c.JSONPretty(http.StatusOK, assignments, "\t")
//	}
//}

//func getRemoteIDFor(user *models.User, provider string) (*models.RemoteIdentity, error) {
//	var remoteID *models.RemoteIdentity
//	for _, v := range user.RemoteIdentities {
//		if v.Provider == provider {
//			remoteID = v
//			break
//		}
//	}
//	if remoteID == nil {
//		return nil, echo.ErrNotFound
//	}
//	return remoteID, nil
//}

func createAssignment(request *yamlparser.NewAssignmentRequest, course *models.Course) (*models.Assignment, error) {
	date, err := time.Parse("02-01-2006 15:04", request.Deadline)
	if err != nil {
		return nil, err
	}
	return &models.Assignment{
		AutoApprove: request.AutoApprove,
		CourseID:    course.ID,
		Deadline:    date,
		Language:    request.Language,
		Name:        request.Name,
		Order:       request.AssignmentID,
	}, nil
}

// GetSubmission returns a single submission for a assignment and a user
//func GetSubmission(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		assignmentID, err := parseUint(c.Param("aid"))
//		if err != nil {
//			return err
//		}
//
//		user := c.Get("user").(*models.User)
//
//		submission, err := db.GetSubmissionForUser(assignmentID, user.ID)
//		if err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return c.NoContent(http.StatusNotFound)
//			}
//			return err
//		}
//
//		return c.JSONPretty(http.StatusOK, submission, "\t")
//	}
//}

// ListSubmissions returns all the latests submissions for a user to a course
//func ListSubmissions(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		courseID, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return err
//		}
//
//		// Check if a user is provided, else used logged in user
//		userID, err := parseUint(c.Param("uid"))
//		if err != nil {
//			userID = c.Get("user").(*models.User).ID
//		}
//
//		submission, err := db.GetSubmissions(courseID, userID)
//		if err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return c.NoContent(http.StatusNotFound)
//			}
//			return err
//		}
//
//		return c.JSONPretty(http.StatusOK, submission, "\t")
//	}
//}

// UpdateCourse updates an existing course
func UpdateCourse(crs *pb.Course, db database.Database) (*pb.Course, error) {
	course, err := db.GetCourse(crs.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "Course not found")
		}
		return nil, err
	}

	newcrs := NewCourseRequest{
		Name:        crs.Name,
		Code:        crs.Code,
		Year:        uint(crs.Year),
		Tag:         crs.Tag,
		Provider:    crs.Provider,
		DirectoryID: crs.Directoryid,
	}
	if !newcrs.valid() {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payload")
	}

	course.Name = crs.Name
	course.Code = crs.Code
	course.Year = uint(crs.Year)
	course.Tag = crs.Tag
	course.Provider = crs.Provider
	course.DirectoryID = crs.Directoryid
	if err := db.UpdateCourse(course); err != nil {
		return nil, err
	}
	return toProtoCourse(course), nil
}

// GetEnrollmentsByCourse get all enrollments for a course.
func GetEnrollmentsByCourse(request *pb.RecordWithStatusRequest, db database.Database) (*pb.EnrollemntResponse, error) {
	statuses, err := parseEnrollmentStatus(request.State)
	if err != nil {
		return nil, err
	}

	enrollments, err := db.GetEnrollmentsByCourse(request.Id, statuses...)
	if err != nil {
		return nil, err
	}

	for _, enrollment := range enrollments {
		enrollment.User, err = db.GetUser(enrollment.UserID)
		if err != nil {
			return nil, err
		}
	}

	var results []*pb.Enrollment
	for _, enroll := range enrollments {
		results = append(results, toProtoEnrollment(enroll))
	}
	return &pb.EnrollemntResponse{Enrollments: results}, nil
}

// NewGroup creates a new group under a course
//func NewGroup(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cid, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return err
//		}
//
//		if _, err := db.GetCourse(cid); err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return echo.NewHTTPError(http.StatusNotFound, "course not found")
//			}
//			return err
//		}
//
//		var grp NewGroupRequest
//		if err := c.Bind(&grp); err != nil {
//			return err
//		}
//		if !grp.valid() {
//			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
//		}
//
//		users, err := db.GetUsers(grp.UserIDs...)
//		if err != nil {
//			return err
//		}
//		// check if provided user ids are valid
//		if len(users) != len(grp.UserIDs) {
//			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
//		}
//		// if logged in user is student, he must need to be member of the group
//		user := c.Get("user").(*models.User)
//		enrollment, err := db.GetEnrollmentByCourseAndUser(cid, user.ID)
//		if err != nil {
//			return err
//		}
//		if enrollment.Status == models.Student {
//			found := false
//			for _, id := range grp.UserIDs {
//				if user.ID == id {
//					found = true
//					break
//				}
//			}
//			if !found {
//				return echo.NewHTTPError(http.StatusBadRequest,
//					"you must need to be a member of the group")
//			}
//		}
//		// only enrolled user i.e accepted to the course can join a group
//		// prevent group override if a student is already in a group in this course
//		for _, user := range users {
//			enrollment, err := db.GetEnrollmentByCourseAndUser(cid, user.ID)
//			switch {
//			case err == gorm.ErrRecordNotFound:
//				return echo.NewHTTPError(http.StatusNotFound, "user is not enrolled to this course")
//			case err != nil:
//				return err
//			case enrollment.GroupID > 0:
//				return echo.NewHTTPError(http.StatusBadRequest, "user is already in another group")
//			case enrollment.Status < models.Student:
//				return echo.NewHTTPError(http.StatusBadRequest, "user is not yet accepted to this course")
//			}
//		}
//
//		group := models.Group{
//			Name:     grp.Name,
//			CourseID: cid,
//			Users:    users,
//		}
//		// CreateGroup creates a new group and update group_id in enrollment table
//		if err := db.CreateGroup(&group); err != nil {
//			if err == database.ErrDuplicateGroup {
//				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//			}
//			return err
//		}
//
//		return c.JSONPretty(http.StatusCreated, &group, "\t")
//	}
//}

// UpdateGroup update a group
//func UpdateGroup(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cid, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return err
//		}
//
//		if _, err := db.GetCourse(cid); err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return echo.NewHTTPError(http.StatusNotFound, "course not found")
//			}
//			return err
//		}
//
//		gid, err := parseUint(c.Param("gid"))
//		if err != nil {
//			return err
//		}
//		oldgrp, err := db.GetGroup(gid)
//		if err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return echo.NewHTTPError(http.StatusNotFound, "group not found")
//			}
//			return err
//		}
//
//		user := c.Get("user").(*models.User)
//		enrollment, err := db.GetEnrollmentByCourseAndUser(cid, user.ID)
//		if err != nil {
//			return err
//		}
//		if enrollment.Status != models.Teacher {
//			return echo.NewHTTPError(http.StatusForbidden, "only teacher can update a group")
//		}
//
//		var grp NewGroupRequest
//		if err := c.Bind(&grp); err != nil {
//			return err
//		}
//		if !grp.valid() {
//			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
//		}
//		users, err := db.GetUsers(grp.UserIDs...)
//		if err != nil {
//			return err
//		}
//		// check if provided user ids are valid
//		if len(users) != len(grp.UserIDs) {
//			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
//		}
//
//		// only enrolled user i.e accepted to the course can join a group
//		// prevent group override if a student is already in a group in this course
//		for _, user := range users {
//			enrollment, err := db.GetEnrollmentByCourseAndUser(cid, user.ID)
//			switch {
//			case err == gorm.ErrRecordNotFound:
//				return echo.NewHTTPError(http.StatusNotFound, "user is not enrolled to this course")
//			case err != nil:
//				return err
//			case enrollment.GroupID > 0 && enrollment.GroupID != oldgrp.ID:
//				return echo.NewHTTPError(http.StatusBadRequest, "user is already in another group")
//			case enrollment.Status < models.Student:
//				return echo.NewHTTPError(http.StatusBadRequest, "user is not yet accepted to this course")
//			}
//		}
//
//		if err := db.UpdateGroup(&models.Group{
//			ID:       oldgrp.ID,
//			Name:     grp.Name,
//			CourseID: cid,
//			Users:    users,
//		}); err != nil {
//			if err == database.ErrDuplicateGroup {
//				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//			}
//			return err
//		}
//
//		return c.NoContent(http.StatusOK)
//	}
//}

// GetGroups returns all groups under a course
//func GetGroups(db database.Database) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cid, err := parseUint(c.Param("cid"))
//		if err != nil {
//			return err
//		}
//		if _, err := db.GetCourse(cid); err != nil {
//			if err == gorm.ErrRecordNotFound {
//				return echo.NewHTTPError(http.StatusNotFound, "course not found")
//			}
//			return err
//		}
//		groups, err := db.GetGroupsByCourse(cid)
//		if err != nil {
//			return err
//		}
//		return c.JSONPretty(http.StatusOK, groups, "\t")
//	}
//}

// convert a model Course to proto Course
func toProtoCourse(course *models.Course) *pb.Course {
	pc := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Code:        course.Code,
		Provider:    course.Provider,
		Tag:         course.Tag,
		Year:        uint32(course.Year),
		Directoryid: course.DirectoryID,
		Enrolled:    int32(course.Enrolled),
	}
	return pc
}

// convert a model Assignment to proto Assignment
func toProtoAssignment(asg *models.Assignment) *pb.Assignment {
	pa := &pb.Assignment{
		Id:          asg.ID,
		Name:        asg.Name,
		Courseid:    asg.CourseID,
		Language:    asg.Language,
		Autoapprove: asg.AutoApprove,
		Deadline:    asg.Deadline.Format("2006-01-02 15:04:05"),
		Order:       uint32(asg.Order),
		Submission:  toProtoSubmission(asg.Submission),
	}
	return pa
}

func toProtoSubmission(submission *models.Submission) *pb.Submission {
	ps := &pb.Submission{
		Id:           submission.ID,
		Userid:       submission.UserID,
		Assignmentid: submission.AssignmentID,
		Groupid:      submission.GroupID,
		Buildinfo:    submission.BuildInfo,
		Score:        uint32(submission.Score),
		Scoreobjects: submission.ScoreObjects,
		Commithash:   submission.CommitHash,
	}
	return ps
}

func toProtoEnrollment(enrollment *models.Enrollment) *pb.Enrollment {
	pe := &pb.Enrollment{
		Id:       enrollment.ID,
		Userid:   enrollment.UserID,
		Courseid: enrollment.CourseID,
		Groupid:  enrollment.GroupID,
		User:     toProtoUser(enrollment.User),
		Status:   uint32(enrollment.Status),
	}
	return pe
}
