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

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(TrackTasksNew.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
