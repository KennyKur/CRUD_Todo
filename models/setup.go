package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	// DB variable for connection DB postgresql
	DB *gorm.DB
)

func ConnectDatabase() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=db_exercise sslmode=disable password=4n4k0nd4")

	if err != nil {

		panic(err)

	}

	// defer db.Close()

	DB = db
}
