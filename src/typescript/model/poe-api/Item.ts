export interface Item {
  id: string;
  name: string;
  category: any;
  typeLine: string;
  descrText: string;
  
  craftedMods: Array<string>;
  enchantMods: Array<string>;
  explicitMods: Array<string>;
  implicitMods: Array<string>;
  properties: Array<any>;
  additionalProperties: Array<any>;

  verified: boolean;
  x: number;
  y: number;
  icon: string;
  league: string;
  note: string;
}
