package model

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"PrimaryKey;type:serial" validate:"required"`
	NameBook  string    `json:"name_book" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
