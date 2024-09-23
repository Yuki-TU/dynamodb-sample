package controller

import (
	"github.com/Yuki-TU/dynamodb-sample/gen"
	"github.com/Yuki-TU/dynamodb-sample/repository"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) UpdateUserByID(ctx *gin.Context, params gen.UpdateUserByIDRequestObject) (gen.UpdateUserByIDResponseObject, error) {
	err := c.repo.UpdateUser(ctx, repository.User{
		UserId:        int64(params.UserId),
		FirstName:     params.Body.FirstName,
		LastName:      params.Body.FamilyName,
		FirstNameKana: params.Body.FirstNameKana,
		LastNameKana:  params.Body.FamilyNameKana,
	})
	if err != nil {
		ctx.Error(err)
		return gen.UpdateUserByID500JSONResponse{}, err
	}

	return gen.UpdateUserByID200JSONResponse{
		Status: "ok",
	}, nil
}
