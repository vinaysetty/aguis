package objects

import (
	"github.com/autograde/aguis/models"
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"isAdmin": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.IsAdmin, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Name, nil
				}
				return nil, nil
			},
		},
		"studentID": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.StudentID, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Email, nil
				}
				return nil, nil
			},
		},
		"avatarURL": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.AvatarURL, nil
				}
				return nil, nil
			},
		},
	},
})

var RemoteIdentityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RemoteIdentity",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.ID, nil
				}
				return nil, nil
			},
		},
		"provider": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.Provider, nil
				}
				return nil, nil
			},
		},
		"remoteID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.RemoteID, nil
				}
				return nil, nil
			},
		},
		"accessToken": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.AccessToken, nil
				}
				return nil, nil
			},
		},
		"userID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.UserID, nil
				}
				return nil, nil
			},
		},
	},
})
