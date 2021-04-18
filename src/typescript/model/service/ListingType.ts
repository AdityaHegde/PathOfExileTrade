export class ListingType {
  public name: string;
  public active: string;

  constructor(listingTypeJson) {
    this.name = listingTypeJson.id;
    this.active = listingTypeJson.attributes.active;
  }
}
