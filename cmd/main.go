package main

import (
	"log"

	gta5modaz "github.com/Nikittansk/gta5-modaz"
	"github.com/Nikittansk/gta5-modaz/pkg/handler"
	"github.com/Nikittansk/gta5-modaz/pkg/repository"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка при инициализации конфигурации: %s", err)
	}

	_, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %s", err)
	}

	handlers := new(handler.Handler)

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