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
    str: number;
    agi: number;
    int: number;
    dex: number;
    vit: number;
    luk: number;
  };
  equipment_slots: string;
}
