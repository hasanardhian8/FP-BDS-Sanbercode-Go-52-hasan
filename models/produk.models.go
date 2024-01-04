package models

type Produks struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}
