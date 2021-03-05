import { Helpers } from "../Helpers";

export class RateLimiter {
  // calls per second
  private rate: number;
  private timeDiff: number;

  private previousTime: number;

  constructor(rate: number) {
    this.setLimit(rate);
  }

  public setLimit(rate: number) {
    this.rate = rate;
    this.timeDiff = 1 / rate;
  }

  public async limit() {
    const curTime = Date.now();

    if (curTime - this.previousTime > this.timeDiff) {
      await Helpers.wait(curTime - this.previousTime - this.timeDiff);
    }

    this.previousTime = curTime;
  }
}
