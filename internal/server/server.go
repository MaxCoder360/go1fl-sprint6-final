package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server — обёртка над http.Server с логгером.
type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

// NewServer создаёт новый сервер с зарегистрированными хендлерами.
func NewServer(logger *log.Logger) *Server {
	// Создаём http-роутер
	mux := http.NewServeMux()
	logger.Println("router is created")

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	logger.Println("handlers are created")

	// Создаём http.Server с настройками из ТЗ
	httpSrv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	logger.Printf("http.Server is configured: addr=%s, readTimeout=%s, writeTimeout=%s, idleTimeout=%s",
		httpSrv.Addr, httpSrv.ReadTimeout, httpSrv.WriteTimeout, httpSrv.IdleTimeout)

	return &Server{
		logger:     logger,
		httpServer: httpSrv,
	}
}

// Start запускает сервер.
func (s *Server) Start() error {
	s.logger.Printf("server is running on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
