package models

type Registers struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Nama     string `json:"nama"`
	Email    string `json:"Email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
