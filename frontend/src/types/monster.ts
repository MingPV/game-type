// type Monster struct {
// 	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"monster_id"`
// 	Name          string    `json:"name"`
// 	Description   string    `json:"description"`
// 	Level         int       `json:"level"`
// 	HP            int       `json:"hp"`
// 	Attack        int       `json:"attack"`
// 	Defense       int       `json:"defense"`
// 	ExpReward     int       `json:"exp_reward"`
// 	GoldReward    int       `json:"gold_reward"`
// 	MonsterTypeID uuid.UUID `gorm:"type:uuid" json:"monster_type_id"`

import { MonsterLoot } from "./monsterLoot";
import { MonsterType } from "./monsterType";

// 	MonsterType  MonsterType   `gorm:"foreignKey:MonsterTypeID;references:ID" json:"monster_type"` // this.MonsterTypeID -> MonsterType.ID
// 	MonsterLoots []MonsterLoot `gorm:"foreignKey:MonsterID" json:"monster_loots"`                  // MonsterLoot.MonsterID -> this.ID
// }

export interface Monster {
  monster_id: string;
  name: string;
  description: string;
  level: number;
  hp: number;
  attack: number;
  defense: number;
  exp_reward: number;
  gold_reward: number;
  monster_type_id: string;

  monster_type: MonsterType;

  monster_loots: MonsterLoot[];
}
