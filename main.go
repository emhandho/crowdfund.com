package main

import (
	"log"
	// "net/http"

	"crowdfund.com/handler"
	"crowdfund.com/user"
	"github.com/gin-gonic/gin"

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

	userRepository	:= user.NewRepository(db)
	userService		:= user.NewService(userRepository)

	userHandler		:= handler.NewUserHandler(userService) 

	router	:= gin.Default()
	api		:= router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()
}