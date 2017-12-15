package objects

import (
	"github.com/autograde/aguis/models"
	"github.com/graphql-go/graphql"
)

var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Group",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if group, ok := p.Source.(*models.Group); ok {
					return group.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if group, ok := p.Source.(*models.Group); ok {
					return group.Name, nil
				}
				return nil, nil
			},
		},
		"status": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if group, ok := p.Source.(*models.Group); ok {
					return group.Status, nil
				}
				return nil, nil
			},
		},
		"courseID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if group, ok := p.Source.(*models.Group); ok {
					return group.CourseID, nil
				}
				return nil, nil
			},
		},
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
		},
		"entollments": &graphql.Field{
			Type: graphql.NewList(EnrollmentType),
		},
	},
})
