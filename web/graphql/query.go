package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/hansludvig/graphql-aguis/graphql/objects"
	"github.com/autograde/aguis/models"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"courses": &graphql.Field{
			Type: graphql.NewList(objects.CourseType),
		},
		"course": &graphql.Field{
			Type: objects.CourseType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error){
				if id, ok := p.Args["id"].(int); ok{
					return GetCourse(id), nil
				}
				return nil, nil
			},
		},
	},
})

func (string) GetCourse *models.Course{
	return nil
}

func GetUser(id int) *User {
	if u, ok := UserData[id]; ok {
		return u
	}
	return &User{}
}

// Users list of user
var Users = []*modles.User{
	{
		ID: "2342"
		IsAdmin: false
		Name: "Hans Ludvig"
		StudentID: "220896"
		Email: "hl@kleivdal.no"
		AvatarURL: "vg.no"
		Enrollments: [
			{
				// STOPED
			}
		]
}

// UserData -> User data
var UserData = make(map[int]*models.User)
