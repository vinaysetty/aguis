package graphqlapi

import (
	"errors"
	"strconv"

	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/models"
	gql "github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//Mutation - GraphQL mutation
func Mutation(l logrus.FieldLogger, db database.Database) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEnrollment": &gql.Field{
				Type: gql.Boolean,
				Args: gql.FieldConfigArgument{
					"uid": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
					"cid": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					var enrollment models.Enrollment

					arg1, _ := p.Args["uid"].(string)
					if uid, err := strconv.ParseUint(arg1, 10, 64); err == nil {
						enrollment.UserID = uid
					}
					arg2, _ := p.Args["cid"].(string)
					if cid, err := strconv.ParseUint(arg2, 10, 64); err == nil {
						enrollment.CourseID = cid
					}

					if err := db.CreateEnrollment(&enrollment); err != nil {
						if err == gorm.ErrRecordNotFound {
							return false, errors.New(err.Error()) // make Status type return int, 0,1,2,3 etc
						}
						l.Error(err.Error())
						return false, nil
					}
					return true, nil
				},
			},
		},
	})
}
