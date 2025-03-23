package routes

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/suwi-lanji/chinsa-api/pkg/handlers"
	"github.com/suwi-lanji/chinsa-api/pkg/repositories"
)

func SetupRoutes(db *pgx.Conn) *mux.Router {
	r := mux.NewRouter()

	userRepository := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	return r
}
