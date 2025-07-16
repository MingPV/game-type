export interface ItemLevelStat {
  item_id: string;
  bonus_stat: Map<string, BonusStat>;
}

type BonusStat = {
  bonus_attack: number;
  bonus_defense: number;
  bonus_hp: number;
  bonus_mp: number;
  bonus_critical: number;
};
