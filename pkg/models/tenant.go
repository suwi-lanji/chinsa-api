package models

type Tenant struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Email      string                 `json:"email"`
	Address    string                 `json:"address"`
	Image      string                 `json:"image"`
	Properties map[string]interface{} `json:"properties"`
}
