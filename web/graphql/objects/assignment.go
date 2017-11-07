package objects

import (
	"github.com/autograde/aguis/models"
	"github.com/graphql-go/graphql"
)

var AssignmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Assignment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.ID, nil
				}
				return nil, nil
			},
		},
		"courseID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.CourseID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.Name, nil
				}
				return nil, nil
			},
		},
		"language": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.Language, nil
				}
				return nil, nil
			},
		},
		"deadline": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.Deadline, nil
				}
				return nil, nil
			},
		},
		"autoApprove": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.AutoApprove, nil
				}
				return nil, nil
			},
		},
		"order": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if assign, ok := p.Source.(*models.Assignment); ok {
					return assign.Order, nil
				}
				return nil, nil
			},
		},
		"submission": &graphql.Field{
			Type: submissionType,
		},
	},
})

var submissionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Submission",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.ID, nil
				}
				return nil, nil
			},
		},
		"assignmentID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.AssignmentID, nil
				}
				return nil, nil
			},
		},
		"userID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.UserID, nil
				}
				return nil, nil
			},
		},
		"groupID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.GroupID, nil
				}
				return nil, nil
			},
		},
		"score": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.Score, nil
				}
				return nil, nil
			},
		},
		"scoreObjects": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.ScoreObjects, nil
				}
				return nil, nil
			},
		},
		"buildInfo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.BuildInfo, nil
				}
				return nil, nil
			},
		},
		"commitHash": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if sub, ok := p.Source.(*models.Submission); ok {
					return sub.CommitHash, nil
				}
				return nil, nil
			},
		},
	},
})
