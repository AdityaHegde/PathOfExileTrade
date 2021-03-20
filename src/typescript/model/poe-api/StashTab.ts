import { Item } from "./Item";

export enum StashTabType {
  PremiumStash = "PremiumStash",
}

export interface StashTab {
  accountName: string;
  lastCharacterName: string;
  id: string;
  stash: string;
  stashType: StashTabType;
  items: Array<Item>;
  public: boolean;
  league: string;
}
