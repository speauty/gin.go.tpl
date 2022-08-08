package dao

import (
	"gin.go.tpl/kernel/code"
	"gin.go.tpl/kernel/errors"
	"gin.go.tpl/util"
	"gorm.io/gorm"
)

type GormScopeFn func(*gorm.DB) *gorm.DB
type GormScopeFnSlice []func(*gorm.DB) *gorm.DB

type IDao interface {
	newDao() interface{}
	getEntity(isSlice bool) interface{}
	toEntity() interface{}
	fromEntity(data interface{})
	getModel(tx *gorm.DB, loadNewEntity bool) *gorm.DB
}

type BaseDao struct{}

func (d *BaseDao) DaoCreateNoScope(tx *gorm.DB, stdDao IDao) errors.IError {
	return d.GenericCreate(tx, stdDao, nil)
}

func (d *BaseDao) DaoListNormal(stdDao IDao, whereScope GormScopeFn, orderScope GormScopeFn, pageScope GormScopeFn, result *[]interface{}, count *int64) errors.IError {
	*count = 0
	whereScope = d.fillScope(whereScope)
	if err := d.GenericCount(stdDao, whereScope, count); err != nil {
		return err
	}
	if *count > 0 {
		scopes := d.genScopes(whereScope, orderScope, pageScope)
		tmpEntities := stdDao.getEntity(true)
		if err := d.GenericFind(stdDao, &tmpEntities, scopes); err != nil {
			return err
		}
		tmpEntitiesSwap := util.ToSlice(tmpEntities)
		for _, tmpEntity := range tmpEntitiesSwap {
			tmpDao := stdDao.newDao().(IDao)
			tmpDao.fromEntity(tmpEntity)
			*result = append(*result, tmpDao)
		}
	}
	return nil
}

func (d *BaseDao) GenericCount(stdDao IDao, whereScope GormScopeFn, count *int64) errors.IError {
	if err := stdDao.getModel(nil, true).Scopes(d.fillScope(whereScope)).Count(count).Error; err != nil {
		return errors.Logic().NewFromCode(code.StdDbQuery, err)
	}
	return nil
}

func (d *BaseDao) GenericFindOne(stdDao IDao, result interface{}, scopes GormScopeFnSlice) errors.IError {
	if err := stdDao.getModel(nil, true).Scopes(scopes...).Take(result).Error; err != nil {
		return errors.Logic().NewFromCode(code.StdDbQuery, err)
	}
	return nil
}

func (d *BaseDao) GenericFind(stdDao IDao, result interface{}, scopes GormScopeFnSlice) errors.IError {
	if err := stdDao.getModel(nil, true).Scopes(scopes...).Find(result).Error; err != nil {
		return errors.Logic().NewFromCode(code.StdDbQuery, err)
	}
	return nil
}

func (d *BaseDao) GenericCreate(tx *gorm.DB, stdDao IDao, scope GormScopeFn) errors.IError {
	res := stdDao.getModel(tx, true).Scopes(d.fillScope(scope)).Create(stdDao.toEntity())
	if res.Error != nil {
		return errors.Logic().NewFromCode(code.StdDb, res.Error)
	}
	if res.RowsAffected == 0 {
		return errors.Logic().NewFromCode(code.StdDb, nil)
	}
	return nil
}

func (d *BaseDao) GenericUpdate(tx *gorm.DB, stdDao IDao, scope GormScopeFn) errors.IError {
	if err := stdDao.getModel(tx, true).Scopes(d.fillScope(scope)).Updates(stdDao.toEntity()).Error; err != nil {
		return errors.Logic().NewFromCode(code.StdDbUpdate, err)
	}
	return nil
}

func (d *BaseDao) GenericDelete(tx *gorm.DB, stdDao IDao, scope GormScopeFn) errors.IError {
	if err := stdDao.getModel(tx, true).Scopes(d.fillScope(scope)).Delete(stdDao.toEntity()).Error; err != nil {
		return errors.Logic().NewFromCode(code.StdDbDelete, err)
	}
	return nil
}

func (d *BaseDao) fillScope(fn GormScopeFn) GormScopeFn {
	if fn == nil {
		fn = func(query *gorm.DB) *gorm.DB {
			return query
		}
	}
	return fn
}

func (d *BaseDao) genScopes(fns ...GormScopeFn) GormScopeFnSlice {
	var scopes GormScopeFnSlice
	for _, fn := range fns {
		if fn != nil {
			scopes = append(scopes, fn)
		}
	}
	return scopes
}
