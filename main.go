package main

import (
	"base-go/config"
	userDelivery "base-go/user/delivery/http"
	userRepository "base-go/user/repository/mysql"
	userUsecase "base-go/user/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := config.NewDatabase()

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	app := gin.Default()

	_userMysqlRepository := userRepository.NewMysqlUserRepo(db)

	_userUsecase := userUsecase.NewUserUsecase(_userMysqlRepository)

	userDelivery.NewUserHandler(app, _userUsecase)

	app.Run(":8080")
}
