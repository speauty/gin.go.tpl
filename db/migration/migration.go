package migration

import (
	"gin.go.tpl/db"
)

type Migration struct {
	models []*interface{}
}

func (m *Migration) AddModel(model *interface{}) {
	if m.models == nil {
		m.models = make([]*interface{}, 1)
	}
	m.models = append(m.models, model)
}

func (m Migration) Exec() {
	err := db.DB{}.Get().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(m.models[:])
	if err != nil {
		panic(err)
	}
}
