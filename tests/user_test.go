package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/suwi-lanji/chinsa-api/pkg/handlers"
	"github.com/suwi-lanji/chinsa-api/pkg/models"
	"github.com/suwi-lanji/chinsa-api/pkg/repositories"
)

func TestRegisterUser(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		return
	}
	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		panic("DATABASE_URL environment variable is not set")
	}
	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Could not connect to database: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	user := models.User{
		Username: "testuser1",
		Password: "testpassword",
		Email:    "test@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Could not marshal user: %v", err)
	}

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.RegisterUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		t.Errorf("Error: %v", rr.Body)
	}
}
