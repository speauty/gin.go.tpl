package entity

import (
	"gin.go.tpl/db"
	"gin.go.tpl/util"
	"gorm.io/gorm"
)

type User struct {
	Id                   int64  `gorm:"<-:create;type:int8;primaryKey;comment:主键"`
	Nickname             string `gorm:"type:varchar(64);not null;comment:昵称"`
	Passwd               string `gorm:"type:varchar(64);not null;comment:密码"`
	Salt                 string `gorm:"type:varchar(64);not null;comment:盐"`
	db.TimeModuleWithDel `gorm:"embedded"`
}

func (u *User) AfterFind(_ *gorm.DB) (err error) {
	u.CreatedAt = util.LocalDateTime(u.CreatedAt)
	u.UpdatedAt = util.LocalDateTime(u.UpdatedAt)
	return
}
