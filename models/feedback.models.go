package models

import "time"

type Feedbacks struct {
	Id          int       `json:"id" gorm:"primary_key"`
	IdTransaksi int       `json:"idTransaksi" gorm:"ForeignKey:TransaksisRefer;AssociationForeignKey:Id"`
	IdProfil    int       `json:"idProfil" gorm:"ForeignKey:ProfilsRefer;AssociationForeignKey:Id"`
	Komen       string    `json:"komen"`
	Rating      int       `json:"rating"`
	Tanggal     time.Time `json:"tanggal"`
}
