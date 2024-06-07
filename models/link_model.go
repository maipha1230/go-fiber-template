package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Title  string `json:"tile"`
	Url    string `json:"url"`
	Type   string `json:"type"`
	UserID uint   `json:"userId"`
}

type LinkBodyRequest struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Type  string `json:"type"`
}
