package app

import (
	"log"
	"net/http"

	"github.com/ivanov-nikolay/calculator/internal/api"
	"github.com/ivanov-nikolay/calculator/internal/config"
)

// Application структура, содержащая конфигурационные параметры приложения
type Application struct {
	config *config.Config
}

// NewApplication создает новый экземпляр структуры Application
func NewApplication() *Application {
	return &Application{
		config: config.LoadConfig(),
	}
}

// RunApplication запускает сервер приложежния
func (a *Application) RunApplication() {
	mux := http.NewServeMux()
	calculation := http.HandlerFunc(api.GetCalculation)
	mux.Handle("/api/v1/calculate", api.LoggingMiddleWare(calculation))

	log.Printf("server starting on port %s", a.config.ServerPort)
	if err := http.ListenAndServe(a.config.ServerPort, mux); err != nil {
		log.Fatalf("application exit with error: %v", err)
	}
}
