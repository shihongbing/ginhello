package service

import (
	"ginhello/model"
	"ginhello/repository"
	"github.com/gin-gonic/gin"
)

var userRepository = repository.UserRepository{}

type UserService struct {
	repository.UserRepository
}

func (*UserService) Get(ctx *gin.Context,id string,user *model.User) (int64, error) {
	return userRepository.Get(ctx, id, user)
}

