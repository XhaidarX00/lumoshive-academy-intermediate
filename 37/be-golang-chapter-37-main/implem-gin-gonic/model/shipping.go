package model

type ShippingService struct {
	CourierName   string  `json:"courier_name"`
	ServiceType   string  `json:"service_type"`
	BaseCost      float64 `json:"base_cost"`
	CostPerKm     float64 `json:"cost_per_km"`
	MaxWeightKg   float64 `json:"max_weight_kg"`
	EstimatedTime string  `json:"estimated_time"`
}
