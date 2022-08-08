package dao

import (
	"gin.go.tpl/db"
	"gin.go.tpl/db/entity"
	"gin.go.tpl/kernel/code"
	"gin.go.tpl/kernel/errors"
	"gin.go.tpl/util"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	Id        int64     `json:"id"`
	Nickname  string    `json:"nickname"`
	Passwd    string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ud *UserDao) NewModel() *gorm.DB {
	return db.DB{}.Get().Model(&entity.User{})
}

func (ud *UserDao) toEntity() *entity.User {
	newPasswd, salt := util.GenStrEncodedAndSalt(&ud.Passwd)
	return &entity.User{Nickname: ud.Nickname, Passwd: newPasswd, Salt: salt}
}

func (ud *UserDao) fromEntity(user *entity.User) {
	ud.Id = user.Id
	ud.Nickname = user.Nickname
	ud.CreatedAt = user.CreatedAt
	ud.UpdatedAt = user.UpdatedAt
}

func (ud *UserDao) ReloadById() errors.IError {
	user := &entity.User{}
	if ud.Id == 0 {
	}
	if err := ud.NewModel().Where("id = ?", ud.Id).First(user).Error; err != nil {
	}
	ud.fromEntity(user)
	return nil
}

func (ud *UserDao) CreateUser() errors.IError {
	if err := (db.DB{}).Create(ud.toEntity()); err != nil {
	}
	return nil
}

func (ud *UserDao) Unique() errors.IError {
	count := int64(0)
	ud.NewModel().Where("nickname = ?", ud.Nickname).Count(&count)
	if count > 0 {
		return errors.Logic().NewFromCode(code.StdDbUnique, nil)
	}
	return nil
}
