package main

import (
	Config "book/config"
	Models "book/models"
	Routes "book/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routes.SetupRouter()
	r.Run("localhost:8080")
}
