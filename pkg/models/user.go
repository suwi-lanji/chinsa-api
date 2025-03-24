package models

import (
	"time"
)

type User struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Email           string                 `json:"email"`
	EmailVerifiedAt time.Time              `json:"email_verified_at"`
	Image           string                 `json:"image"`
	Password        string                 `json:"-"`
	TenantId        string                 `json:"tenant_id"`
	Properties      map[string]interface{} `json:"properties"`
	CreatedAt       time.Time              `json:"created_at"`
}
