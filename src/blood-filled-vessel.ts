import { Item } from "@model/poe-api/Item";
import { PublicStashTabs } from "@model/poe-api/PublicStashTabs";
import { StashTab } from "@model/poe-api/StashTab";
import { PublicStashTabsClient } from "./client/PublicStashTabsClient";

const Interest = new Set([
  "Kitava, The Destroyer",

  "Nassar, Lion of the Seas",
  "Ambrius, Legion Slayer",
  "Varhesh, Shimmering Aberration",

  "Stone of the Currents",
  "Gisale, Thought Thief",
  "Shock and Horror",

  "Penitentiary Incarcerator",
  "The Brittle Emperor",
  "Stalker of the Endless Dunes",

  "Gorulis, Will-Thief",
  "Armala, the Widow",
  "Visceris",
  "Queen of the Great Tangle",
]);

function checkBloodFilledVessel(stash: StashTab, item: Item) {
  const monstors = item.properties[0];
  if (monstors.values[0][1] === 10) {
    if (Interest.has(monstors.values[0][0])) {
      console.log(`${stash.accountName}(${stash.lastCharacterName}) : ${item.note} ` +
        `(stash tab "${stash.stash}"; position: left ${item.x}, top ${item.y}) ` +
        `(${monstors.values[0][0]})`);
    }
  }
}

const stashTabsClient = new PublicStashTabsClient(
  "1090659386-1098552584-1058944583-1187446950-1139423730",
  (newPublicStashTabs: PublicStashTabs) => {
    newPublicStashTabs.stashes.forEach((stash) => {
      if (stash.public && stash.league === "Ritual") {
        stash.items.forEach((item) => {
          if (item.typeLine === "Blood-filled Vessel") {
            checkBloodFilledVessel(stash, item);
          }
        });
      }
    });
  },
);
stashTabsClient.start();
