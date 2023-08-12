package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id" gorm:"primarykey"`
	Nama      string         `json:"nama"`
	Email     string         `json:"email"`
	Password  string         `json:"-" gorm:"coloumn:password"`
	Adress    string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `jsom:"deleted_at" gorm:"index"`
}
