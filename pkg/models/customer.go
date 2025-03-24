package models

type Customer struct {
	ID         string                 `json:"id"`
	TenantId   string                 `json:"tenant_id"`
	Name       string                 `json:"name"`
	Company    string                 `json:"company"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	Status     string                 `json:"status"`
	Industry   string                 `json:"industry"`
	Properties map[string]interface{} `json:"properties"`
}
