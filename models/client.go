package models

import "gorm.io/gorm"

type My_client struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"` 
	Slug string `json:"slug" gorm:"not null"`
	IsProject string `json:"is_poject" gorm:"not null"` 
	SelfCapture string `json:"self_capture" gorm:"not null"`
	ClientPrefix string `json:"client_prefix" gorm:"not null"`
	ClientLogo string `json:"client_logo" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
}