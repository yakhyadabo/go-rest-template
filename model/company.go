package model

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	//ID   string `json:"id"`
	Name string `json:"name"`
}
