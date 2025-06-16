package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/brunnossanttos/libri/internal/config"
	"github.com/brunnossanttos/libri/internal/db"
	"github.com/brunnossanttos/libri/internal/handler"
	"github.com/brunnossanttos/libri/internal/repository"
	"github.com/brunnossanttos/libri/internal/service"
)

func main() {
	config.LoadConfig()

	if err := db.Connect(); err != nil {
		log.Fatalf("DB connect error: %v", err)
	}

	userRepo := repository.NewUserRepository(db.DB)
	userSvc  := service.NewUserService(userRepo)
	userH    := handler.NewUserHandler(userSvc)
	bookSvc := service.NewBookService()
	bookH   := handler.NewBookHandler(bookSvc)

	r := gin.Default()
	r.POST("/users", userH.CreateUser)
	r.GET("/", func(c *gin.Context) {
		log.Println("hello from go")
		c.JSON(http.StatusOK, gin.H{"message": "hello from go"})
	})
	r.GET("/books", bookH.SearchBooks)
	r.Run(":" + viper.GetString("APP_PORT"))
}
