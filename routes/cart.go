package routes

import (
	"backend-API/handlers"
	"backend-API/pkg/middleware"
	"backend-API/pkg/mysql"
	"backend-API/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cartsall", middleware.Auth(h.FindCarts)).Methods("GET")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.GetCart)).Methods("GET")
	r.HandleFunc("/carts", middleware.Auth(h.FindCartId)).Methods("GET") //use
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST") //use
	r.HandleFunc("/cart/{id}", h.UpdateCartTransaction).Methods("PATCH") //use
	// r.HandleFunc("/carts/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}
