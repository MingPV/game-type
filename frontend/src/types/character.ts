export interface Character {
  character_id: string;
  user_id: string;
  name: string;
  level: number;
  current_exp: number;
  class_id: string;
  inventory_id: string;
  class: {
    name: string;
    description: string;
  };
  status: {
    status_point: number;
    attack_level: number;
    defense_level: number;
    hp_level: number;
    mp_level: number;
    critical_level: number;
    attack: number;
    defense: number;
    hp: number;
    mp: number;
    critical: number;
  };
  equipment_slots: string;
}
