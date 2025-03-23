package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/suwi-lanji/chinsa-api/docs"
	"github.com/suwi-lanji/chinsa-api/pkg/database"
	"github.com/suwi-lanji/chinsa-api/pkg/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Chinsa API
// @version 1.0
// @description This is a sample REST API for managing users and tasks.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@mygoapi.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env file")
	}
	connString := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	if connString == "" {
		panic("DATABASE_URL environment variable is not set")
	}
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	if err := database.InitDB(connString); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	r := routes.SetupRoutes(database.DB)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
