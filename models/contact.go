package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
