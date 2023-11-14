package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/nolostra/goproject/internal/middleware"
	
)


func Handler(r *chi.Mux){
	// Global Middleware forware slash are removed 
	// account/200/ <- last slash are removed
	r.Use(chimiddle.StripSlashes)

	r.Route("/account",func(router chi.Router){
		// Middleware For /account route
		router.Use(middleware.Authorisation)

		router.Get("/coins",GetCoinBalance)
		router.Post("/coins",PostCoinBalance)
		router.Delete("/coins",DelCoinBalance)
	})
}