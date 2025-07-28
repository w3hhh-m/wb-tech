package main

import (
	"fmt"
	"os"

	_ "wb-tech-l0/docs/swagger"
	"wb-tech-l0/internal/app"
)

// @title WB Tech L0 Orders API
// @version 1.0
// @description Сервис для хранения и получения заказов.
// @contact.url https://github.com/w3hhh-m/wb-tech/tree/main/l0
// @contact.email w3hhh.m@gmail.com
// @host localhost:8080
// @BasePath /
// @schemes http

// main is the entry point of the application
func main() {
	if err := app.Start(); err != nil {
		fmt.Printf("Application error: %s\n", err)
		os.Exit(1)
	}
}
