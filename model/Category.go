package model

type Category struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt TimeNormal `gorm:"column:created_at;default:null" json:"created_at"`
	UpdateAt TimeNormal `gorm:"column:updated_at;default:null" json:"updated_at"`
}
