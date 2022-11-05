package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Blog struct {
	ID         uuid.UUID
	Title      string `json:"title" form:"Title"`
	Date       time.Time
	Image      string `json:"image" form:"Image"`
	Content    string `json:"content" form:"Content"`
	Author     string `json:"author" form:"Author"`
	UserID     int    `json:"user_id" form:"User_id"`
	User       User
	CategoryID int `json:"category_id" form:"Category_id"`
	Category   Category
}

func (blog *Blog) BeforeCreate(scope *gorm.Scope) (err error) {
	blogID := uuid.New()
	scope.SetColumn("ID", blogID)
	return
}
