package app

import (
	"challange-8_2/config"
	"challange-8_2/repository"
	"challange-8_2/route"
	"challange-8_2/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApp() {
	repo := repository.NewRepo(config.GORM.DB)
	service := service.NewService(repo)
	// server := handler.NewHttpServer(service)

	route.RegisterApi(router, service)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
