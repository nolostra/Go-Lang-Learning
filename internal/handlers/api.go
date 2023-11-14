package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/nolostra/goproject/internal/middleware"
	
)


func Handler(r *chi.Mux){
	// Global Middleware forward slash are removed 
	// account/200/ <- last slash are removed
	r.Use(chimiddle.StripSlashes)


	// http://localhost:8000/login/?username=sai&password=yourpassword
	r.Route("/login",func(router chi.Router){
		router.Get("/",Login)
	})


	r.Route("/account",func(router chi.Router){
		// Middleware For /account route
		// http://localhost:8000/account/coins/?username=sai + token
		router.Use(middleware.Authorisation)

		router.Get("/coins",GetCoinBalance)
		router.Post("/coins",PostCoinBalance)
		router.Put("/coins",ChangeCoinBalance)
		router.Delete("/coins",DelCoinBalance)
	})
}