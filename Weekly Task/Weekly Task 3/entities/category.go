package entities

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" form:"Name"`
}
