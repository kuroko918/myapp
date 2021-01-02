package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Delete(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	FirstOrCreate(interface{}, ...interface{}) *gorm.DB
	Preload(string, ...interface{}) *gorm.DB
	Update(...interface{}) *gorm.DB
}
