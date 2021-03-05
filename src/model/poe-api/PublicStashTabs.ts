import { StashTab } from "./StashTab";

export interface PublicStashTabs {
  next_change_id: string;
  stashes: Array<StashTab>;
}
