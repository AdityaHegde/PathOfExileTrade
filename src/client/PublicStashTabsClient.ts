import { PublicStashTabs } from "@model/poe-api/PublicStashTabs";
import got from "got";
import { InternalLogger } from "../Logger";
import { RateLimiter } from "./RateLimiter";

const PUBLIC_STASH_TABS_API = "http://api.pathofexile.com/public-stash-tabs";

export class PublicStashTabsClient {
  private logger = new InternalLogger("PathOfExileHTTPClient");
  private rateLimiter = new RateLimiter(2);

  private initialId: string;
  private newStashTabsCallback: (newPublicStashTabs: PublicStashTabs) => void;

  constructor(
    initialId: string,
    newStashTabsCallback: (newPublicStashTabs: PublicStashTabs) => void,
  ) {
    this.initialId = initialId;
    this.newStashTabsCallback = newStashTabsCallback;
  }

  public async start() {
    let nextId = this.initialId;
    // eslint-disable-next-line no-constant-condition
    while (true) {
      await this.rateLimiter.limit();

      const publicStashTabs = await this.get(nextId);

      if (publicStashTabs) {
        nextId = publicStashTabs.next_change_id;
        this.newStashTabsCallback(publicStashTabs);
      }
    }
  }

  private async get(id?: string): Promise<PublicStashTabs> {
    try {
      // this.logger.debug(`Making a call using id=${id}`);

      const resp = await got.get(PUBLIC_STASH_TABS_API + (id ? `/?id=${id}` : ""));
      const publicStashTabs: PublicStashTabs = JSON.parse(resp.body);

      // this.logger.debug(`Got response newId=${id} entires=${publicStashTabs.stashes.length}`);

      return publicStashTabs;
    } catch (err) {
      console.log(err);
    }

    return null;
  }
}
