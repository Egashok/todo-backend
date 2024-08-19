package main

import (
	"github/Egashok/todo-backend"
	"github/Egashok/todo-backend/pkg/handler"
	"github/Egashok/todo-backend/pkg/repository"
	"github/Egashok/todo-backend/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err :=initConfig(); err != nil {
		log.Fatal("Confiig error: ", err.Error())
	}
	if err:=godotenv.Load(); err !=nil{
		log.Fatal("Env error", err.Error())
	}
	db,err:= repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMODE:  viper.GetString("db.sslmode"),

	})
	
	if err != nil {
		log.Fatalf(err.Error())
	}
	repos :=repository.NewRepository(db)
	services:=service.NewService(repos)
	handlers :=handler.NewHandler(services)
	server := new(todo.Server)
	if err:=server.Run(viper.GetString("server.port"),handlers.InitRoutes()); err !=nil {
		log.Fatal(err.Error())
	}  
}
func initConfig() error{
viper.SetConfigName("config")
viper.SetConfigType("yaml")
viper.AddConfigPath("configs")
viper.AddConfigPath("../configs") 
return viper.ReadInConfig()
}