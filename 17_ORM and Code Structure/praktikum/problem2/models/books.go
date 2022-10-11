package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"Title"`
	Publisher string `json:"publisher" form:"Publisher"`
	Year      int    `json:"year" form:"Year"`
}
