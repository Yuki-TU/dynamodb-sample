package controller

import (
	"github.com/Yuki-TU/dynamodb-sample/gen"
	"github.com/Yuki-TU/dynamodb-sample/repository"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) CreateUser(ctx *gin.Context, params gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {
	users, err := c.repo.GetUsers(ctx)
	if err != nil {
		ctx.Error(err)
		return gen.CreateUser500JSONResponse{
			N500ErrorJSONResponse: gen.N500ErrorJSONResponse{
				Message: "internal server error",
			},
		}, err
	}
	var userID int64 = 1
	if len(users) > 0 {
		userID = users[len(users)-1].UserId + 1
	}
	err = c.repo.CreateUser(ctx, repository.User{
		UserId:        userID,
		FirstName:     params.Body.FirstName,
		LastName:      params.Body.FamilyName,
		FirstNameKana: params.Body.FirstNameKana,
		LastNameKana:  params.Body.FamilyNameKana,
	})
	if err != nil {
		ctx.Error(err)
		return gen.CreateUser201JSONResponse{}, err
	}

	return gen.CreateUser201JSONResponse{
		UserId: userID,
	}, nil
}
