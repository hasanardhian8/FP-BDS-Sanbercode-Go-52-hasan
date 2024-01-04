package models

type Saldos struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Pembayaran string `json:"pembayaran"`
	Nominal    int    `json:"nominal"`
	Total      int    `json:"total"`
}
