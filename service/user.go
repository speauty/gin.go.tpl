package service

import (
	"gin.go.tpl/db/dao"
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/errors"
)

type UserService struct{}

func (us UserService) Register(_ *lib.Context, userDao *dao.UserDao) errors.Error {
	if err := userDao.UniqueUser(); err != nil {
		return err
	}
	return userDao.CreateUser()
}

func (us UserService) Query(_ *lib.Context, userDao *dao.UserDao) errors.Error {
	return userDao.ReloadById()
}
