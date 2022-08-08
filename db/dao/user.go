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
	BaseDao
	Id        int64      `json:"id"`
	Nickname  string     `json:"nickname"`
	Passwd    string     `json:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (d *UserDao) newDao() interface{} { return &UserDao{} }

func (d *UserDao) getModel(tx *gorm.DB, loadNewEntity bool) *gorm.DB {
	if tx == nil {
		tx = db.DB{}.Get()
	}
	model := d.toEntity()
	if loadNewEntity {
		model = d.getEntity(false)
	}
	return tx.Model(model)
}

func (d *UserDao) getEntity(isSlice bool) interface{} {
	if isSlice {
		return []*entity.User{}
	} else {
		return &entity.User{}
	}
}

func (d *UserDao) toEntity() interface{} {
	newPasswd, salt := util.GenStrEncodedAndSalt(&d.Passwd)
	return &entity.User{Nickname: d.Nickname, Passwd: newPasswd, Salt: salt}
}

func (d *UserDao) fromEntity(data interface{}) {
	d.Id = data.(*UserDao).Id
	d.Nickname = data.(*UserDao).Nickname
	d.CreatedAt = data.(*UserDao).CreatedAt
	d.UpdatedAt = data.(*UserDao).UpdatedAt
}

func (d *UserDao) ReloadById() errors.IError {
	user := &entity.User{}
	if d.Id == 0 {
	}
	if err := d.getModel(nil, true).Where("id = ?", d.Id).First(user).Error; err != nil {
	}
	d.fromEntity(user)
	return nil
}

func (d *UserDao) Unique() errors.IError {
	count := int64(0)
	d.getModel(nil, true).Where("nickname = ?", d.Nickname).Count(&count)
	if count > 0 {
		return errors.Logic().NewFromCode(code.StdDbUnique, nil)
	}
	return nil
}
