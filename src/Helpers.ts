export class Helpers {
  public static wait(time: number) {
    return new Promise(resolve => setTimeout(resolve, time));
  }
}
