package main

import (
	"apiGorm/controllers"
	"apiGorm/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// mes, err := models.ConnectToDatabaseEpicor()
	// if err != nil {
	// 	panic(err.Error())
	// }
	mess, err := models.ConnectToDatabaseAbrageo()
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(mes)
	fmt.Println(mess)
	fmt.Println(controllers.Message())
	router := gin.Default()
	router.GET("/", controllers.Login)
	router.POST("/grafica", controllers.PieDistributor)
	router.GET("/fichaClienteModel", controllers.GetFichaClienteModel)
	router.GET("/fichaClienteQuery", controllers.GetFichaClienteQuery)
	router.Run("localhost:8080")
}
