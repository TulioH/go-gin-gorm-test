package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// type credentials struct {
// 	User     string
// 	Password string
// 	Port     string
// 	Database string
// }

// var DbEpicor credentials = &credentials{"SA", "AD17abra88$", "test10db", "evaluation"}

func ConnectToDatabaseEpicor() (*gorm.DB, error) {
	user := "SA"
	password := "AD17abra88$"
	port := "test10db"
	database := "evaluation"
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, password, port, database)
	Db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "error connecting to database")
	}
	// data, err := db.DB()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = data.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// log.Println(data)
	// log.Println("se conecto a la base de dato")

	return Db, err
}

func ConnectToDatabaseAbrageo() (*gorm.DB, error) {
	user := "root"
	password := "Abracol2014*"
	port := "10.1.0.184:3306"
	database := "abrageo"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, port, database)
	Db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		PrepareStmt: true,
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err, "error connecting to database")
	}
	// data, err := db.DB()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = data.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// log.Println(data)
	// log.Println("se conecto a la base de datos")

	return Db, err
}
