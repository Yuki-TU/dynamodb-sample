package controller

import (
	"errors"

	"github.com/Yuki-TU/dynamodb-sample/gen"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) GetUserByID(ctx *gin.Context, params gen.GetUserByIDRequestObject) (gen.GetUserByIDResponseObject, error) {
	res, err := c.repo.GetUserByID(ctx, int64(params.UserId))
	if err != nil {
		ctx.Error(err)
		return gen.GetUserByID500JSONResponse{
			N500ErrorJSONResponse: gen.N500ErrorJSONResponse{
				Message: "internal server error",
			},
		}, err
	}
	if res.FirstName == "" {
		ctx.Error(errors.New("user not found"))
		return gen.GetUserByID404JSONResponse{
			N404ErrorJSONResponse: gen.N404ErrorJSONResponse{
				Message: "user not found",
			},
		}, nil
	}

	return gen.GetUserByID200JSONResponse{
		UserId:         float32(res.UserId),
		FamilyName:     res.LastName,
		FamilyNameKana: res.LastNameKana,
		FirstName:      res.FirstName,
		FirstNameKana:  res.FirstNameKana,
	}, nil
}
