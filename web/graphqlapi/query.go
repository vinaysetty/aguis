package graphqlapi

import (
	"strconv"

	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/models"
	"github.com/autograde/aguis/web/graphqlapi/objects"
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
			"allUsers": &graphql.Field{
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
		},
	})
}
