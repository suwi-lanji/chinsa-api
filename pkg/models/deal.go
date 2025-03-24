package models

import "time"

type Deal struct {
	ID         string                 `json:"id"`
	TenantId   string                 `json:"tenant_id"`
	ListingId  string                 `json:"listing_id"`
	CustomerId string                 `json:"customer_id"`
	Amount     float64                `json:"amount"`
	Stage      string                 `json:"stage"`
	CloseDate  time.Time              `json:"close_date"`
	Properties map[string]interface{} `json:"properties"`
}
