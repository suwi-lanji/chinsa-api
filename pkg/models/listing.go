package models

type Listing struct {
	ID                string                 `json:"id"`
	TenantId          string                 `json:"tenant_id"`
	Name              string                 `json:"name"`
	Type              string                 `json:"type"`
	Purpose           string                 `json:"purpose"`
	Status            string                 `json:"status"`
	Address           string                 `json:"address"`
	City              string                 `json:"city"`
	Region            string                 `json:"region"`
	Zip               string                 `json:"zip"`
	Community         string                 `json:"community"`
	Coordinates       map[string]interface{} `json:"coordinates"`
	Price             float64                `json:"price"`
	Currency          string                 `json:"currency"`
	PriceType         string                 `json:"price_type"`
	Taxes             float64                `json:"taxes"`
	Fees              float64                `json:"fees"`
	Size              float64                `json:"size"`
	Bedrooms          int                    `json:"bedrooms"`
	Bathrooms         int                    `json:"bathrooms"`
	Floor             float64                `json:"floor"`
	YearBuilt         string                 `json:"year_built"`
	Parking           int                    `json:"parking"`
	FurnishingStatus  string                 `json:"furnishing_status"`
	IndoorFeatures    map[string]interface{} `json:"indoor_features"`
	OutdoorFeatures   map[string]interface{} `json:"outdoor_features"`
	CommunityFeatures map[string]interface{} `json:"community_features"`
	Utilities         map[string]interface{} `json:"utilities"`
	AgentName         string                 `json:"agent_name"`
	AgentContact      string                 `json:"agent_contact"`
	AgencyName        string                 `json:"agency_name"`
	Description       string                 `json:"description"`
	NearbyLandmarks   map[string]interface{} `json:"nearby_landmarks"`
	Availability      string                 `json:"availability"`
	Condition         string                 `json:"condition"`
	Tags              map[string]interface{} `json:"tags"`
	Images            map[string]interface{} `json:"images"`
	Videos            map[string]interface{} `json:"videos"`
	Properties        map[string]interface{} `json:"properties"`
}
