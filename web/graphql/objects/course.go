package objects

import (
	"github.com/autograde/aguis/models"
	"github.com/graphql-go/graphql"
)

var CourseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Course",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.Name, nil
				}
				return nil, nil
			},
		},
		"code": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.Code, nil
				}
				return nil, nil
			},
		},
		"year": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.Year, nil
				}
				return nil, nil
			},
		},
		"tag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.Tag, nil
				}
				return nil, nil
			},
		},
		"provider": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.Provider, nil
				}
				return nil, nil
			},
		},
		"directoryID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.DirectoryID, nil
				}
				return nil, nil
			},
		},
		"enrollments": &graphql.Field{
			Type: graphql.NewList(EnrollmentType),
		},
		"enrolled": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if course, ok := p.Source.(*models.Course); ok {
					return course.Enrolled, nil
				}
				return nil, nil
			},
		},
		"assignments": &graphql.Field{
			Type: graphql.NewList(AssignmentType),
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(GroupType),
		},
	},
})
