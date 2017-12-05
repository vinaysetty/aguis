package web

import (
	"net/http"

	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// PatchGroup updates status of a group
func PatchGroup(db database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := parseUint(c.Param("gid"))
		if err != nil {
			return err
		}
		oldgrp, err := db.GetGroup(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "group not found")
			}
			return err
		}
		// UpdateGroupRequest updates group
		var ngrp struct {
			Status pb.Enrollment_Status `json:"status"`
		}
		if err := c.Bind(&ngrp); err != nil { //TODO should be replaced with some grpc args
			return err
		}
		if ngrp.Status > pb.Enrollment_Teacher {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
		}

		user := c.Get("user").(*models.User)
		// TODO: This check should be performed in AccessControl.
		if !user.IsAdmin {
			// Ony Admin i.e Teacher can update status of a group
			return c.NoContent(http.StatusForbidden)
		}

		if err := db.UpdateGroupStatus(&pb.Group{
			ID:     oldgrp.ID,
			Status: ngrp.Status,
		}); err != nil {
			return err
		}
		return c.NoContent(http.StatusOK)
	}
}

// GetGroup returns a group
func GetGroup(db database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		gid, err := parseUint(c.Param("gid"))
		if err != nil {
			return err
		}
		group, err := db.GetGroup(gid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "group not found")
			}
			return err
		}
		return c.JSONPretty(http.StatusOK, group, "\t")
	}
}

// DeleteGroup deletes a pending or rejected group
func DeleteGroup(db database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		gid, err := parseUint(c.Param("gid"))
		if err != nil {
			return err
		}
		group, err := db.GetGroup(gid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "group not found")
			}
			return err
		}
		if group.Status > pb.Enrollment_Rejected {
			return echo.NewHTTPError(http.StatusForbidden, "accepted group cannot be deleted")
		}
		if err := db.DeleteGroup(gid); err != nil {
			return nil
		}
		return c.NoContent(http.StatusOK)
	}
}

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
