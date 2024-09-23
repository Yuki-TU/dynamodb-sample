package controller

import "github.com/Yuki-TU/dynamodb-sample/repository"

type Controllers struct {
	repo *repository.Client
}

func NewController(repo *repository.Client) *Controllers {
	return &Controllers{
		repo: repo,
	}
}
