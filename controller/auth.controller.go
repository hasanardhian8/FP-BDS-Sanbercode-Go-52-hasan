package controller

import (
	"fmt"
	"member/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inLo loginInput

	if err := c.ShouldBindJSON(&inLo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Registers{}

	u.Email = inLo.Email
	u.Password = inLo.Password

	token, err := models.LoginCheck(u.Email, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	user := map[string]string{
		"email": u.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}
