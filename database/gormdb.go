package database

import (
	"errors"
	"strings"

	"github.com/autograde/aguis/models"
	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/jinzhu/gorm"
)

// GormDB implements the Database interface.
type GormDB struct {
	conn *gorm.DB
}

// NewGormDB creates a new gorm database using the provided driver.
func NewGormDB(driver, path string, logger GormLogger) (*GormDB, error) {
	conn, err := gorm.Open(driver, path)
	if err != nil {
		return nil, err
	}

	if logger != nil {
		conn.SetLogger(logger)
	}
	conn.LogMode(logger != nil)

	if err := conn.AutoMigrate(
		&pb.User{},
		&pb.RemoteIdentity{},
		&pb.Course{},
		&pb.Enrollment{},
		&pb.Assignment{},
		&pb.Submission{},
		&pb.Group{},
		&models.Repository{},
	).Error; err != nil {
		return nil, err
	}

	return &GormDB{conn}, nil
}

// GetUser implements the Database interface.
func (db *GormDB) GetUser(uid uint64) (*pb.User, error) {
	var user pb.User
	if err := db.conn.Preload("RemoteIdentities").First(&user, uid).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByRemoteIdentity implements the Database interface.
func (db *GormDB) GetUserByRemoteIdentity(provider string, rid uint64, accessToken string) (*pb.User, error) {
	tx := db.conn.Begin()

	// Get the remote identity.
	var remoteIdentity pb.RemoteIdentity
	if err := tx.
		Where(&pb.RemoteIdentity{
			Provider: provider,
			RemoteID: rid,
		}).
		First(&remoteIdentity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update the access token.
	if err := tx.Model(&remoteIdentity).Update("access_token", accessToken).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get the user.
	var user pb.User
	if err := tx.Preload("RemoteIdentities").First(&user, remoteIdentity.UserID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers implements the Database interface.
func (db *GormDB) GetUsers(uids ...uint64) ([]*pb.User, error) {
	m := db.conn
	if len(uids) > 0 {
		m = m.Where(uids)
	}

	var users []*pb.User
	if err := m.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser implements the Database interface
func (db *GormDB) UpdateUser(user *pb.User) error {
	return db.conn.Model(&pb.User{}).Updates(user).Error
}

// SetAdmin implements the Database interface.
func (db *GormDB) SetAdmin(uid uint64) error {
	var user pb.User
	if err := db.conn.First(&user, uid).Error; err != nil {
		return err
	}
	user.IsAdmin = true
	return db.conn.Save(&user).Error
}

// GetRemoteIdentity implements the Database interface.
func (db *GormDB) GetRemoteIdentity(provider string, rid uint64) (*pb.RemoteIdentity, error) {
	var remoteIdentity pb.RemoteIdentity
	if err := db.conn.Model(&pb.RemoteIdentity{}).
		Where(&pb.RemoteIdentity{
			Provider: provider,
			RemoteID: rid,
		}).
		First(&remoteIdentity).Error; err != nil {
		return nil, err
	}
	return &remoteIdentity, nil
}

// CreateUserFromRemoteIdentity implements the Database interface.
func (db *GormDB) CreateUserFromRemoteIdentity(user *pb.User, remoteIdentity *pb.RemoteIdentity) error {
	user.RemoteIdentities = []*pb.RemoteIdentity{remoteIdentity}
	if err := db.conn.Create(&user).Error; err != nil {
		return err
	}
	// The first user defaults to admin user.
	if user.ID == 1 {
		if err := db.SetAdmin(1); err != nil {
			return err
		}
		user.IsAdmin = true
	}
	return nil
}

var (
	// ErrDuplicateIdentity is returned when trying to associate a remote identity
	// with a user account and the identity is already in use.
	ErrDuplicateIdentity = errors.New("remote identity register with another user")
	// ErrDuplicateGroup is returned when trying to create a group with the same
	// name as a previously registered group.
	ErrDuplicateGroup = errors.New("group name already registered")
	// ErrCourseExists is returned when trying to create an association in
	// the database for a DirectoryID that already exists in the database.
	ErrCourseExists = errors.New("course already exists on git provider")
)

// AssociateUserWithRemoteIdentity implements the Database interface.
func (db *GormDB) AssociateUserWithRemoteIdentity(uid uint64, provider string, rid uint64, accessToken string) error {
	var count uint64
	if err := db.conn.
		Model(&pb.RemoteIdentity{}).
		Where(&pb.RemoteIdentity{
			Provider: provider,
			RemoteID: rid,
		}).
		Not(&pb.RemoteIdentity{
			UserID: uid,
		}).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrDuplicateIdentity
	}

	var remoteIdentity pb.RemoteIdentity
	return db.conn.
		Where(pb.RemoteIdentity{Provider: provider, RemoteID: rid, UserID: uid}).
		Assign(pb.RemoteIdentity{AccessToken: accessToken}).
		FirstOrCreate(&remoteIdentity).Error
}

// CreateCourse implements the Database interface.
func (db *GormDB) CreateCourse(course *pb.Course) error {
	var courses uint64
	if err := db.conn.Model(&pb.Course{}).Where(&pb.Course{
		DirectoryID: course.DirectoryID,
	}).Count(&courses).Error; err != nil {
		return err
	}
	if courses > 0 {
		return ErrCourseExists
	}
	return db.conn.Create(course).Error
}

// GetCourses implements the Database interface.
// If one or more course ids are provided, the corresponding courses
// are returned. Otherwise, all courses are returned.
func (db *GormDB) GetCourses(cids ...uint64) ([]*pb.Course, error) {
	m := db.conn
	if len(cids) > 0 {
		m = m.Where(cids)
	}
	var courses []*pb.Course
	if err := m.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

// GetAssignmentsByCourse implements the Database interface
func (db *GormDB) GetAssignmentsByCourse(cid uint64) ([]*pb.Assignment, error) {
	var course pb.Course
	if err := db.conn.Preload("Assignments").First(&course, cid).Error; err != nil {
		return nil, err
	}
	return course.Assignments, nil
}

// CreateSubmission implements the Database interface
// TODO: Also check enrollment to see if the user is
// enrolled in the course the assignment belongs to
func (db *GormDB) CreateSubmission(submission *pb.Submission) error {
	// Primary key must be greater than 0.
	if submission.AssignmentID < 1 {
		return gorm.ErrRecordNotFound
	}

	// Either user or group id must be set.
	var m *gorm.DB
	switch {
	case submission.UserID > 0:
		m = db.conn.First(&pb.User{ID: submission.UserID})
	case submission.GroupID > 0:
		m = db.conn.First(&pb.User{ID: submission.UserID})
	default:
		return gorm.ErrRecordNotFound
	}

	// Check that group exists.
	var group uint64
	if err := m.Count(&group).Error; err != nil {
		return err
	}

	// Checks that the assignment exists.
	var assignment uint64
	if err := db.conn.Model(&pb.Assignment{}).Where(&pb.Assignment{
		ID: submission.AssignmentID,
	}).Count(&assignment).Error; err != nil {
		return err
	}

	if assignment+group != 2 {
		return gorm.ErrRecordNotFound
	}
	return db.conn.Create(submission).Error
}

// GetSubmissionForUser implements the Database interface
func (db *GormDB) GetSubmissionForUser(aid uint64, uid uint64) (*pb.Submission, error) {
	var submission pb.Submission
	if err := db.conn.Where(&pb.Submission{AssignmentID: aid, UserID: uid}).Last(&submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

// GetSubmissions implements the Database interface
func (db *GormDB) GetSubmissions(cid uint64, uid uint64) ([]*pb.Submission, error) {
	var course pb.Course
	if err := db.conn.Preload("Assignments").First(&course, cid).Error; err != nil {
		return nil, err
	}

	var latestSubs []*pb.Submission
	for _, a := range course.Assignments {
		temp, err := db.GetSubmissionForUser(a.ID, uid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return nil, err
		}
		latestSubs = append(latestSubs, temp)
	}
	return latestSubs, nil
}

// CreateAssignment implements the Database interface
func (db *GormDB) CreateAssignment(assignment *pb.Assignment) error {
	// Course id and assignment order must be given.
	if assignment.CourseID < 1 || assignment.Order < 1 {
		return gorm.ErrRecordNotFound
	}

	var course uint64
	if err := db.conn.Model(&pb.Course{}).Where(&pb.Course{
		ID: assignment.CourseID,
	}).Count(&course).Error; err != nil {
		return err
	}
	if course != 1 {
		return gorm.ErrRecordNotFound
	}

	return db.conn.
		Where(pb.Assignment{
			CourseID: assignment.CourseID,
			Order:    assignment.Order,
		}).
		Assign(pb.Assignment{
			Name:        assignment.Name,
			Language:    assignment.Language,
			Deadline:    assignment.Deadline,
			AutoApprove: assignment.AutoApprove,
		}).FirstOrCreate(assignment).Error
}

// CreateEnrollment implements the Database interface.
// This method will overwrite the status field with models.Pending.
func (db *GormDB) CreateEnrollment(enrollment *pb.Enrollment) error {
	var user, course uint64
	if err := db.conn.Model(&pb.User{}).Where(&pb.User{
		ID: enrollment.UserID,
	}).Count(&user).Error; err != nil {
		return err
	}
	if err := db.conn.Model(&pb.Course{}).Where(&pb.Course{
		ID: enrollment.CourseID,
	}).Count(&course).Error; err != nil {
		return err
	}
	if user+course != 2 {
		return gorm.ErrRecordNotFound
	}
	// make sure to mark enrollment status as pending, independent of what is supplied
	enrollment.Status = pb.Enrollment_Pending
	return db.conn.Create(&enrollment).Error
}

// EnrollStudent implements the Database interface.
func (db *GormDB) EnrollStudent(uid, cid uint64) error {
	return db.setEnrollment(uid, cid, pb.Enrollment_Student)
}

// RejectEnrollment implements the Database interface.
func (db *GormDB) RejectEnrollment(uid, cid uint64) error {
	return db.setEnrollment(uid, cid, pb.Enrollment_Rejected)
}

// EnrollTeacher implements the Database interface.
func (db *GormDB) EnrollTeacher(uid, cid uint64) error {
	return db.setEnrollment(uid, cid, pb.Enrollment_Teacher)
}

// GetEnrollmentsByCourse implements the Database interface.
func (db *GormDB) GetEnrollmentsByCourse(cid uint64, statuses ...pb.Enrollment_Status) ([]*pb.Enrollment, error) {
	return db.getEnrollments(&pb.Course{ID: cid}, statuses...)
}

func (db *GormDB) getEnrollments(model interface{}, statuses ...pb.Enrollment_Status) ([]*pb.Enrollment, error) {
	if len(statuses) == 0 {
		statuses = []pb.Enrollment_Status{
			pb.Enrollment_Pending,
			pb.Enrollment_Rejected,
			pb.Enrollment_Student,
			pb.Enrollment_Teacher,
		}
	}
	var enrollments []*pb.Enrollment
	if err := db.conn.Model(model).
		Where("status in (?)", statuses).
		Association("Enrollments").
		Find(&enrollments).Error; err != nil {
		return nil, err
	}

	return enrollments, nil
}

// GetEnrollmentByCourseAndUser return a record of Enrollment
func (db *GormDB) GetEnrollmentByCourseAndUser(cid uint64, uid uint64) (*pb.Enrollment, error) {
	var enrollment pb.Enrollment
	if err := db.conn.
		Where(&pb.Enrollment{
			CourseID: cid,
			UserID:   uid,
		}).
		First(&enrollment).Error; err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (db *GormDB) setEnrollment(uid, cid uint64, status pb.Enrollment_Status) error {
	if status > pb.Enrollment_Teacher {
		panic("invalid status")
	}
	return db.conn.
		Model(&pb.Enrollment{}).
		Where(&pb.Enrollment{CourseID: cid, UserID: uid}).
		Update(&pb.Enrollment{Status: status}).Error
}

// GetCoursesByUser returns all courses (with enrollment status)
// for the given user id.
// If enrollment statuses is provided, the set of courses returned
// is filtered according to these enrollment statuses.
func (db *GormDB) GetCoursesByUser(uid uint64, statuses ...pb.Enrollment_Status) ([]*pb.Course, error) {
	enrollments, err := db.getEnrollments(&pb.User{ID: uid}, statuses...)
	if err != nil {
		return nil, err
	}

	var courseIDs []uint64
	m := make(map[uint64]*pb.Enrollment)
	for _, enrollment := range enrollments {
		m[enrollment.CourseID] = enrollment
		courseIDs = append(courseIDs, enrollment.CourseID)
	}

	if len(statuses) == 0 {
		courseIDs = nil
	} else if len(courseIDs) == 0 {
		// No need to query database since user have no enrolled courses.
		return []*pb.Course{}, nil
	}
	courses, err := db.GetCourses(courseIDs...)
	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		course.Enrolled = pb.Enrollment_None
		if enrollment, ok := m[course.ID]; ok {
			course.Enrolled = enrollment.Status
		}
	}
	return courses, nil
}

// GetCourse implements the Database interface
func (db *GormDB) GetCourse(cid uint64) (*pb.Course, error) {
	var course pb.Course
	if err := db.conn.First(&course, cid).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// GetCourseByDirectoryID implements the Database interface
func (db *GormDB) GetCourseByDirectoryID(did uint64) (*pb.Course, error) {
	var course pb.Course
	if err := db.conn.First(&course, &pb.Course{DirectoryID: did}).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// UpdateCourse implements the Database interface
func (db *GormDB) UpdateCourse(course *pb.Course) error {
	return db.conn.Model(&pb.Course{}).Updates(course).Error
}

// CreateGroup creates a new group and assign users to newly created group
func (db *GormDB) CreateGroup(group *pb.Group) error {
	if group.CourseID == 0 {
		return gorm.ErrRecordNotFound
	}

	tx := db.conn.Begin()
	var course uint64
	if err := db.conn.Model(&pb.Course{}).Where(&pb.Course{
		ID: group.CourseID,
	}).Count(&course).Error; err != nil {
		return err
	}
	if course != 1 {
		return gorm.ErrRecordNotFound
	}

	if err := tx.Model(&pb.Group{}).Create(group).Error; err != nil {
		tx.Rollback()
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			return ErrDuplicateGroup
		}
		return err
	}
	var userids []uint64
	for _, u := range group.Users {
		userids = append(userids, u.ID)
	}
	query := tx.Model(&pb.Enrollment{}).
		Where(&pb.Enrollment{
			CourseID: group.CourseID,
		}).
		Where("user_id IN (?) AND status IN (?)", userids, []pb.Enrollment_Status{
			pb.Enrollment_Student, pb.Enrollment_Teacher}).
		Updates(&pb.Enrollment{
			GroupID: group.ID,
		})

	if query.Error != nil {
		tx.Rollback()
		return query.Error
	}

	if query.RowsAffected != int64(len(userids)) {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	tx.Commit()
	return nil
}

// GetGroup returns a group specified by id return error if does not exits
func (db *GormDB) GetGroup(gid uint64) (*pb.Group, error) {
	var group pb.Group
	if err := db.conn.Preload("Enrollments").First(&group, gid).Error; err != nil {
		return nil, err
	}
	var userIDs []uint64
	for _, enrollment := range group.Enrollments {
		userIDs = append(userIDs, enrollment.UserID)
	}
	if len(userIDs) > 0 {
		users, err := db.GetUsers(userIDs...)
		if err != nil {
			return nil, err
		}
		group.Users = users
	}
	return &group, nil
}

// UpdateGroupStatus updates status field of a group
func (db *GormDB) UpdateGroupStatus(group *pb.Group) error {
	return db.conn.Model(group).Update("status", group.Status).Error
}

// GetGroupsByCourse returns a list of groups
func (db *GormDB) GetGroupsByCourse(cid uint64) ([]*pb.Group, error) {
	var groups []*pb.Group
	if err := db.conn.
		Preload("Enrollments").
		Where(&pb.Group{
			CourseID: cid,
		}).
		Find(&groups).Error; err != nil {
		return nil, err
	}

	for _, group := range groups {
		var userIDs []uint64
		for _, enrollment := range group.Enrollments {
			userIDs = append(userIDs, enrollment.UserID)
		}
		if len(userIDs) > 0 {
			users, err := db.GetUsers(userIDs...)
			if err != nil {
				return nil, err
			}
			group.Users = users
		}

	}
	return groups, nil
}

// DeleteGroup delete a group
func (db *GormDB) DeleteGroup(gid uint64) error {
	group, err := db.GetGroup(gid)
	if err != nil {
		return err
	}

	tx := db.conn.Begin()
	if err := tx.Delete(group).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("UPDATE enrollments SET group_id= ? WHERE group_id= ?", 0, gid).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// CreateRepository implements the Database interface
func (db *GormDB) CreateRepository(repo *models.Repository) error {
	if repo.DirectoryID == 0 || repo.RepositoryID == 0 {
		return gorm.ErrRecordNotFound
	}

	if repo.UserID > 0 {
		err := db.conn.First(&pb.User{}, repo.UserID).Error
		if err != nil {
			return err
		}
	}

	return db.conn.Create(repo).Error
}

// GetRepository imlements the Database interface
func (db *GormDB) GetRepository(rid uint64) (*models.Repository, error) {
	// This uses the repositoryid from the provider to search with, and not
	// the id of the entry in the database
	var repo models.Repository

	if err := db.conn.First(&repo, &models.Repository{RepositoryID: rid}).Error; err != nil {
		return nil, err
	}

	return &repo, nil
}

// UpdateGroup updates a group
func (db *GormDB) UpdateGroup(group *pb.Group) error {
	if group.CourseID == 0 {
		return gorm.ErrRecordNotFound
	}
	tx := db.conn.Begin()
	var course uint64
	if err := db.conn.Model(&pb.Course{}).Where(&pb.Course{
		ID: group.CourseID,
	}).Count(&course).Error; err != nil {
		return err
	}
	if course != 1 {
		return gorm.ErrRecordNotFound
	}
	if err := tx.Model(&pb.Group{}).Updates(group).Error; err != nil {
		tx.Rollback()
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			return ErrDuplicateGroup
		}
		return err
	}
	if err := tx.Exec("UPDATE enrollments SET group_id= ? WHERE group_id= ?", 0, group.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	var userids []uint64
	for _, u := range group.Users {
		userids = append(userids, u.ID)
	}
	query := tx.Model(&pb.Enrollment{}).
		Where(&pb.Enrollment{
			CourseID: group.CourseID,
		}).
		Where("user_id IN (?) AND status IN (?)", userids, []pb.Enrollment_Status{
			pb.Enrollment_Student, pb.Enrollment_Teacher}).
		Updates(&pb.Enrollment{
			GroupID: group.ID,
		})

	if query.Error != nil {
		tx.Rollback()
		return query.Error
	}

	if query.RowsAffected != int64(len(userids)) {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	tx.Commit()
	return nil
}

// Close closes the gorm database.
func (db *GormDB) Close() error {
	return db.conn.Close()
}
