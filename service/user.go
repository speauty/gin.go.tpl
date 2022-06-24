package service

import (
	"gin.go.tpl/db/dao"
	"gin.go.tpl/kernel/errors"
	"github.com/gin-gonic/gin"
)

type UserService struct{}

func (us UserService) Register(_ *gin.Context, userDao *dao.UserDao) errors.Error {
	if err := userDao.UniqueUser(); err != nil {
		return err
	}
	return userDao.CreateUser()
}

func (us UserService) Query(_ *gin.Context, userDao *dao.UserDao) errors.Error {
	return userDao.ReloadById()
}
