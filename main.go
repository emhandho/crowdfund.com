package main

import (
	// "fmt"
	"log"
	// "net/http"

	"crowdfunding.com/user"
	model "crowdfunding.com/user"
	// "github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/startup_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	user := model.User{
		Name: "Nurman Ridho",
	}

	userRepository.Save(user)
}