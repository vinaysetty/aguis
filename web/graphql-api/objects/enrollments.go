package objects

import (
	"github.com/autograde/aguis/models"
	"github.com/graphql-go/graphql"
)

var EnrollmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Enrollment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.ID, nil
				}
				return nil, nil
			},
		},
		"courseID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.CourseID, nil
				}
				return nil, nil
			},
		},
		"userID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.UserID, nil
				}
				return nil, nil
			},
		},
		"groupID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.GroupID, nil
				}
				return nil, nil
			},
		},
		"status": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if enroll, ok := p.Source.(*models.Enrollment); ok {
					return enroll.Status, nil
				}
				return nil, nil
			},
		},
	},
})
