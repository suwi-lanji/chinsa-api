package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/suwi-lanji/chinsa-api/pkg/auth"
	"github.com/suwi-lanji/chinsa-api/pkg/models"
	"github.com/suwi-lanji/chinsa-api/pkg/repositories"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepo}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with a username, password, and email
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User information"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	if err := h.userRepository.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	user, err := h.userRepository.GetUserByUsername(r.Context(), creds.Username)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		return
	}

	if !auth.CheckPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
