package models

import (
	"member/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Registers struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Nama     string `json:"nama"`
	Email    string `json:"Email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null;"`
	Role     string `json:"role"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Registers{}

	err = db.Model(Registers{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	uid := uint(u.Id)
	token, err := token.GenerateToken(uid)

	if err != nil {
		return "", err
	}

	return token, nil

}
