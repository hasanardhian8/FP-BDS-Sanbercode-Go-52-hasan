package models

import "time"

type Saldos struct {
	Id         int       `json:"id" gorm:"primary_key"`
	Pembayaran string    `json:"pembayaran"`
	Nominal    int       `json:"nominal"`
	Total      int       `json:"total"`
	Tanggal    time.Time `json:"tanggal"`
}
