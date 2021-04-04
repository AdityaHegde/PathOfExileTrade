export class User {
  public name: string;
  public email: string;
  public roles: Array<string>;

  constructor(userJson) {
    this.name = userJson.id;
    this.email = userJson.attributes.email;
    this.roles = userJson.attributes.roles;
  }
}
