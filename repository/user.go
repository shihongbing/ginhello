package repository

import (
	"ginhello/core"
	"ginhello/model"
	"github.com/gin-gonic/gin"
)

type UserRepository struct {

}

func (*UserRepository) Get(ctx *gin.Context,id string,user *model.User) (affectRows int64, err error) {
	DB := core.DB(ctx)
	result := DB.Where(model.User{Id: id},"id").Find(user)
	return result.RowsAffected, result.Error
}

func (*UserRepository) FindByAccount(ctx *gin.Context,account string,user *model.User) (affectRows int64, err error) {
	DB := core.DB(ctx)
	result := DB.Where(model.User{Account: account},"account").Find(user)
	return result.RowsAffected, result.Error
}


func List(ctx *gin.Context)  {

}
