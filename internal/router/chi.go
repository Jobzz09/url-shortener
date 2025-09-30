package router

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux{
	router := chi.NewRouter()
	router.ServeHTTP()
	return router
}