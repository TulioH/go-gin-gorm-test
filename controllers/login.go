package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test de esta cosa...")
}

func Message() string {
	return "entro al paquete"
}
