package config

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
	
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DATABASE_URL", "postgresql://postgres:1999@localhost:5432/libridb?sslmode=disable")
	viper.SetDefault("GOOGLE_BOOKS_API_KEY", "")
	viper.AutomaticEnv()
}
