import {UserRole} from "@model/UserRole";

export class User {
  public name: string;
  public email: string;
  public roles: Array<UserRole>;

  constructor(userJson) {
    this.name = userJson.id;
    this.email = userJson.attributes.email;
    this.roles = userJson.attributes.roles;
  }
}
