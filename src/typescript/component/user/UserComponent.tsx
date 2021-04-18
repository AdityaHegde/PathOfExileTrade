import * as React from "react";
import {UserDisplayComponent} from "./UserDisplayComponent";
import { UserContext } from "./UserContext";
import {useContext} from "react";

export function UserComponent() {
  const userContext = useContext(UserContext);

  return (
    <div >{userContext.user ? <UserDisplayComponent /> : <a href="/login">Login</a>}</div>
  );
}
