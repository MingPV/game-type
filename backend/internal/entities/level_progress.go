package entities

type LevelProgress struct {
	Level       int `gorm:"primaryKey" json:"level"`
	ExpRequired int `json:"exp_required"`
}
