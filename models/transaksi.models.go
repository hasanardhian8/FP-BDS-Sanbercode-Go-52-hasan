package models

import "time"

type Transaksis struct {
	Id         int       `json:"id" gorm:"primary_key"`
	IdProfil   int       `json:"idProfil" gorm:"ForeignKey:ProfilsRefer;AssociationForeignKey:Id"`
	IdPesan    int       `json:"idPesan"  gorm:"ForeignKey:PemesanansRefer;AssociationForeignKey:Id"`
	Tanggal    time.Time `json:"tanggal"`
	Pembayaran string    `json:"pembayaran"`
	Status     bool      `json:"status"`
}
