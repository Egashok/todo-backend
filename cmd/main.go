package main

import (
	"github/Egashok/todo-backend"
	"log"
)

func main() {
	server := new(todo.Server)
	if err:=server.Run("8000"); err !=nil{
		log.Fatal(err.Error())
	}  
}