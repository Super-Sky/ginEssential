package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID uuid.UUID         `json:"id" gorm:"type:char(36);primary_key"`
	UserID uint          `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category *Category
	Title string         `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg string       `json:"head_img"`
	Content string       `json:"content" gorm:"type:text;nol null"`
	CreatedAt TimeNormal `gorm:"column:created_at;default:null" json:"created_at"`
	UpdateAt TimeNormal  `gorm:"column:updated_at;default:null" json:"updated_at"`
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}