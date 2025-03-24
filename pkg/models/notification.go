package models

type Notification struct {
	ID         string                 `json:"id"`
	TenantId   string                 `json:"tenant_id"`
	Data       map[string]interface{} `json:"data"`
	Properties map[string]interface{} `json:"properties"`
	Type       string                 `json:"type"`
}
