package main

import (
	"awesomeProject2/serverHttp"
	"awesomeProject2/serverHttp/pkg/handler"
	"awesomeProject2/serverHttp/pkg/repository"
	"awesomeProject2/serverHttp/pkg/service"
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка инициализации конфига: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLmode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("Ошибка подключения к базе данных: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(serverHttp.Server)

	go func() {
		if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Ошибка запуска сервера: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGQUIT)
	<-quit

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Ошибка остановки сервера: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Fatalf("Ошибка отключения базы данных: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("serverHttp/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
