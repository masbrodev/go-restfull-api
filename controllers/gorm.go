package controllers

import "github.com/jinzhu/gorm"

// InDB DB
type InDB struct {
	DB *gorm.DB
}
