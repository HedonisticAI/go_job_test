package models

import "gorm.io/gorm"

type Human struct {
	gorm.Model  `json:"-"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronimyc  string `json:"patronimyc ,omitempty"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}
