package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/suwi-lanji/chinsa-api/pkg/models"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx,
		`
	INSERT INTO users (username, password, email, created_at)
	VALUES ($1, $2, $3, $4)
	`, user.Name, user.Password, user.Email, user.CreatedAt)
	return err
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := r.db.QueryRow(ctx,
		`
	SELECT id,username, password, email, created_at
	FROM users WHERE username = $1
	`, username).Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
