package main

import (
	"github.com/TrackTasks"
	"github.com/TrackTasks/pkg/handler"
	"github.com/TrackTasks/pkg/repository"
	"github.com/TrackTasks/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		SSlMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(TrackTasksNew.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
