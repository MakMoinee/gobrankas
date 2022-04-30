package localhttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	Router *chi.Mux
	port   string
}

func NewService(port string) *Service {
	svc := Service{}
	svc.Router = chi.NewRouter()
	svc.port = port
	return &svc
}

func (svc *Service) Start() error {
	server := http.Server{
		Addr:    svc.port,
		Handler: svc.Router,
	}
	err := server.ListenAndServe()
	return err
}
