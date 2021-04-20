package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//A user of the system
type Usuario struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

func GetUsuario(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Usuario{
		Login: "administrador",
	}})
}
