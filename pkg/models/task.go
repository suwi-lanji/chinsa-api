package models

import "time"

type Task struct {
	ID          string                 `json:"id"`
	TenantId    string                 `json:"tenant_id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	DueDate     time.Time              `json:"due_date"`
	Priority    string                 `json:"priority"`
	Status      string                 `json:"status"`
	Properties  map[string]interface{} `json:"properties"`
}
