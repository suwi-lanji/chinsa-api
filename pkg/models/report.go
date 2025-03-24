package models

type Report struct {
	ID          string                 `json:"id"`
	TenantId    string                 `json:"tenant_id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Data        map[string]interface{} `json:"data"`
	Properties  map[string]interface{} `json:"properties"`
}
