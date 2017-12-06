package objects

import (
	"github.com/autograde/aguis/models"
	. "github.com/graphql-go/graphql"
)

func init() {
	EnrollmentType.AddFieldConfig("user", &Field{Type: UserType})
}

var EnrollmentType = NewObject(ObjectConfig{
	Name: "Enrollment",
	Fields: Fields{
		"id": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.ID, nil
				}
				return nil, nil
			},
		},
		"courseID": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.CourseID, nil
				}
				return nil, nil
			},
		},
		"userID": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.UserID, nil
				}
				return nil, nil
			},
		},
		"groupID": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.GroupID, nil
				}
				return nil, nil
			},
		},
		"status": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.Status, nil
				}
				return nil, nil
			},
		},
	},
})
