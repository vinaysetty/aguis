package graphqlapi

import (
	"strconv"

	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/graphqlapi/objects"
	"github.com/autograde/aguis/models"
	"github.com/graphql-go/graphql"
)

//Query - GraphQL query structure
func Query(db database.Database) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: objects.UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if id, ok := p.Args["id"].(string); ok {
						i, _ := strconv.ParseUint(id, 10, 64)
						user, err := db.GetUser(i)
						if err != nil {
							return err, nil
						}
						return user, nil
					}
					return nil, nil
				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(objects.UserType),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if first, ok := p.Args["first"].(int); ok {
						users, err := db.GetUsers()
						if err != nil {
							return err, nil
						}
						if first != 0 {
							var u []*models.User
							if first > len(users) {
								first = len(users)
							}
							for i := 0; i < first; i++ {
								u = append(u, users[i])
							}
							return u, nil
						}
						return users, nil
					}
					return nil, nil
				},
			},
			"course": &graphql.Field{
				Type: objects.CourseType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if id, ok := p.Args["id"].(string); ok {
						i, _ := strconv.ParseUint(id, 10, 64)
						user, err := db.GetCourse(i)
						if err != nil {
							return nil, err
						}
						return user, nil
					}
					return nil, nil
				},
			},
			"courses": &graphql.Field{
				Type: graphql.NewList(objects.CourseType),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if first, ok := p.Args["first"].(int); ok {
						users, err := db.GetCourses()
						if err != nil {
							return err, nil
						}
						if first != 0 {
							var u []*models.Course
							if first > len(users) {
								first = len(users)
							}
							for i := 0; i < first; i++ {
								u = append(u, users[i])
							}
							return u, nil
						}
						return users, nil
					}
					return nil, nil
				},
			},
			"assigments": &graphql.Field{
				Type: graphql.NewList(objects.AssignmentType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if id, ok := p.Args["id"].(string); ok {
						i, _ := strconv.ParseUint(id, 10, 64)
						assigments, err := db.GetAssignmentsByCourse(i)
						if err != nil {
							return nil, err
						}
						return assigments, nil
					}
					return nil, nil
				},
			},
		},
	})
}
