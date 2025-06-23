package dto

type CreateLevelProgressRequest struct {
	Level       int `json:"level"`
	ExpRequired int `json:"exp_required"`
}
