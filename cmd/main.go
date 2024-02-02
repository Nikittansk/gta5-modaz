package main

import (
	"log"

	gta5modaz "github.com/Nikittansk/gta5-modaz"
	"github.com/Nikittansk/gta5-modaz/pkg/handler"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка при инициализации конфигурации: %s", err)
	}

	handlers := handler.NewHandler()

	srv := new(gta5modaz.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Произошла ошибка при запуске http-сервера: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}