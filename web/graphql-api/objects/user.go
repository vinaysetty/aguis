package objects

import (
	"github.com/autograde/aguis/models"
	. "github.com/graphql-go/graphql"
)

var UserType = NewObject(ObjectConfig{
	Name: "User",
	Fields: Fields{
		"id": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"isAdmin": &Field{
			Type: Boolean,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.IsAdmin, nil
				}
				return nil, nil
			},
		},
		"name": &Field{
			Type: String,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Name, nil
				}
				return nil, nil
			},
		},
		"studentID": &Field{
			Type: String,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.StudentID, nil
				}
				return nil, nil
			},
		},
		"email": &Field{
			Type: String,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Email, nil
				}
				return nil, nil
			},
		},
		"avatarURL": &Field{
			Type: String,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.AvatarURL, nil
				}
				return nil, nil
			},
		},
		"enrollments": &Field{
			Type: NewList(EnrollmentType),
		},
	},
})

var RemoteIdentityType = NewObject(ObjectConfig{
	Name: "RemoteIdentity",
	Fields: Fields{
		"id": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.ID, nil
				}
				return nil, nil
			},
		},
		"provider": &Field{
			Type: String,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.Provider, nil
				}
				return nil, nil
			},
		},
		"remoteID": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.RemoteID, nil
				}
				return nil, nil
			},
		},
		"accessToken": &Field{
			Type: String,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.AccessToken, nil
				}
				return nil, nil
			},
		},
		"userID": &Field{
			Type: Int,
			Resolve: func(p ResolveParams) (interface{}, error) {
				if remoteID, ok := p.Source.(*models.RemoteIdentity); ok {
					return remoteID.UserID, nil
				}
				return nil, nil
			},
		},
	},
})
