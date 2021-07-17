package models

import (
	"some/api/pkg/api"
	"some/api/pkg/api/code"
	"some/api/pkg/db"

	"github.com/jinzhu/gorm"
)

// gorm 文档：https://gorm.io/zh_CN/docs/models.html
type IModel interface {
	GetTableName() string
	SetTableName(tableName string)
}

type BaseModel struct {
	TableName string
}

func (m *BaseModel) GetTableName() string {
	return m.TableName
}

func (m *BaseModel) SetTableName(tableName string) {
	m.TableName = tableName
}

func (m *BaseModel) Db() *gorm.DB {
	return db.Mysql.Table(m.TableName)
}

// 对数据库的err做统计处理
func (m *BaseModel) ProcessError(err error) *api.Error {
	if err == nil {
		return nil
	}

	return api.NewError(code.ErrorDatabase, err.Error())
}
