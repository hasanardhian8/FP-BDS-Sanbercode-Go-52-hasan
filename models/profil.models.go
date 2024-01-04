package models

type Profils struct {
	Id         int `json:"id" gorm:"primary_key"`
	IdRegister int `json:"idRegister" gorm:"ForeignKey:RegistersRefer;AssociationForeignKey:Id"`
	IdSaldo    int `json:"idSaldo"  gorm:"ForeignKey:SaldosRefer;AssociationForeignKey:Id"`
}
