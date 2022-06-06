package entity

import "gin.go.tpl/db"

type User struct {
	Id            int64  `gorm:"<-:create type:int8;comment:主键"`
	Nickname      string `gorm:"type:varchar(64);comment:昵称"`
	Passwd        string `gorm:"type:varchar(64);comment:密码"`
	Salt          string `gorm:"type:varchar(64);comment:盐"`
	db.TimeModule `gorm:"embedded"`
}
