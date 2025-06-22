package dto

type CreateRarityRequest struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"drop_rate"`
}

// type Rarity struct {
// 	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"rarity_id"`
// 	Name     string    `json:"name"`
// 	DropRate float64   `json:"drop_rate"`

// 	Items []Item `gorm:"foreignKey:RarityID;references:ID" json:"items"` // Item.RarityID -> this.ID
// }
