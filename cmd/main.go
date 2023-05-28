package main

import (
	"apiInterview/pkg/handler"
	"apiInterview/pkg/repo"
	"apiInterview/pkg/server"
	"apiInterview/pkg/service"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("error opening .log file:", err)
	}
	defer file.Close()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetOutput(file)
	repo := repo.NewMongoDB()
	repo.Connect("27017")
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srvr := &server.Server{}
	if err := srvr.Run("8080", handlers.Init()); err != nil {
		logrus.Error(err)
		fmt.Print(err.Error())
	}

}
