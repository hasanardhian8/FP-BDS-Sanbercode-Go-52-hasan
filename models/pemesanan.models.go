package models

type Pemesanans struct {
	Id           int `json:"id" gorm:"primary_key"`
	IdProduk     int `json:"idProduk" gorm:"ForeignKey:ProduksRefer;AssociationForeignKey:Id"`
	JumlahBarang int `json:"jumlahBarang"`
	Total        int `json:"total"`
}
