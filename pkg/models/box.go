package models

import "gorm.io/gorm"

type Box struct {
	gorm.Model
	Title       string
	Username    string
	Password    string
	Url         string
	Description string
	Tag         string `gorm:"index"`
	BelongTo    uint   `gorm:"index"`
}
