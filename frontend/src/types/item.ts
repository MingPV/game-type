import { ItemLevelStat } from "./itemLevelStat";
import { ItemType } from "./itemType";
import { Rarity } from "./rarity";

export interface Item {
  item_id: string;
  name: string;
  description: string;
  item_type_id: string;
  rarity_id: string;
  required_level: number;
  max_stack: number;

  item_type: ItemType;
  rarity: Rarity;
  item_level_stat: ItemLevelStat;
}
