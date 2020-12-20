package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kuroko918/myapp/cmd/app/interfaces/database"
)

func NewSqlHandler() database.SqlHandler {
	conn, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/myapp?parseTime=true")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

type SqlHandler struct {
	Conn *gorm.DB
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Delete(value interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Delete(value, where...)
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.FirstOrCreate(out, where...)
}

func (handler *SqlHandler) Preload(column string, conditions ...interface{}) *gorm.DB {
	return handler.Conn.Preload(column, conditions...)
}

func (handler *SqlHandler) Update(attrs ...interface{}) *gorm.DB {
	return handler.Conn.Update(attrs...)
}
