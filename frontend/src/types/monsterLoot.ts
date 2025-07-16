import { Item } from "./item";

export interface MonsterLoot {
  id: string;
  item_id: string;
  quantity_min: number;
  quantity_max: number;

  item: Item;
}
