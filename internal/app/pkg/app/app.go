package app

import (
	"WeatherGoDokerApp/internal/app/endpoint"
	"WeatherGoDokerApp/internal/app/service"
	"log"
	"net/http"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	a := &App{}                                      // Создание экземпляра App.
	a.s = service.New()                              // Инициализация сервиса.
	a.e = endpoint.New(a.s)                          // Инициализация конечной точки с переданным сервисом.
	http.HandleFunc("/api/substring", a.e.Substring) // Установка обработчика для пути "/api/substring".
	return a, nil
}

func (a *App) Run() {
	log.Println("HTTP-server started on port - 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
