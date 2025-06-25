package dto

type CreateRarityRequest struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"drop_rate"`
}
