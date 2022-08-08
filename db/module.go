package db

import "time"

type IdModule struct {
	Id int64 `gorm:"<-:create;type:int8;primaryKey;comment:主键"`
}

// TimeModule time-module
type TimeModule struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;comment:更新时间"`
}

// TimeModuleWithDel time-module with deleted_at
type TimeModuleWithDel struct {
	TimeModule `gorm:"embedded"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null;comment:删除时间"`
}
