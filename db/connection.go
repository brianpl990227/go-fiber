package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connection =  "host=localhost user=postgres password=9902 dbname=golang port=5432 sslmode=disable"

var Context *gorm.DB

func DbConnection(){
	var err error
	
	Context,err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Conectado a base de datos PostgresSQL")
	}



}