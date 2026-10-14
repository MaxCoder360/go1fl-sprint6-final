package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаём логгер
	logger := log.New(os.Stdout, "Morse_convert: ", log.LstdFlags)

	logger.Println("Server start")

	// Создаём сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}

	logger.Println("Server end")
}
