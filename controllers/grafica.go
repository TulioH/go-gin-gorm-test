package controllers

import (
	"apiGorm/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Proceso      int
	Despachado   int
	NoDespachado int
}

type datos struct {
	Distribuidor string `json:"distribuidor"`
	Fecha        string `json:"fecha"`
}

func PieDistributor(c *gin.Context) {

	var request datos
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Printf("%#v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Printf("%#v", request)

	con, err := models.ConnectToDatabaseAbrageo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result Result

	query := fmt.Sprint(`SELECT count(Case idEstado when 1 then id end) as Proceso,
	count(Case idEstado when 2 then id end) as Despachado,
	count(Case idEstado when 3 then id end) as NoDespachado
	FROM 	gestiondiaria
	WHERE 	gestiondiaria.distribuidor like '`, request.Distribuidor, `' AND gestiondiaria.ingresoFH LIKE '%`, request.Fecha, `%'
	GROUP BY gestiondiaria.distribuidor`)
	// fmt.Print("ESTE ES EL STRING \n", query)
	if res := con.Raw(query).Scan(&result); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}
	fmt.Printf("%#v", result)
	c.JSON(http.StatusOK, result)
}

func GetFichaClienteQuery(c *gin.Context) {

	// var request datos
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	fmt.Printf("%#v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Printf("%#v", request)

	con, err := models.ConnectToDatabaseAbrageo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []models.FichaCliente

	query := `select distinct * from fichacliente`
	// fmt.Print("ESTE ES EL STRING \n", query)
	if res := con.Raw(query).Scan(&result); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}
	// fmt.Printf("%#v", result)
	c.JSON(http.StatusOK, result)
}

func GetFichaClienteModel(c *gin.Context) {
	// fmt.Printf("%#v", request)

	db, err := models.ConnectToDatabaseAbrageo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []models.FichaCliente

	// fmt.Print("ESTE ES EL STRING \n", query)
	if res := db.Distinct().Find(&result); res.Error != nil {
		// log.Fatal(res.Error, result)
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}
	// fmt.Printf("%#v", result)
	c.JSON(http.StatusOK, result)
}
