package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Item struct {
	Id          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Price       string `json:"price" form:"price"`
	Category    string `json:"category" form:"category"`
}
