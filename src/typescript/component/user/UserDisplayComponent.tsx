import * as React from "React";
import { UserContext } from "./UserContext";
import {useContext} from "react";

export function  UserDisplayComponent() {
  const userContext = useContext(UserContext);
  return (
    <>
      <div>{userContext.user.name}</div>
      <button onClick={userContext.logout}>Logout</button>
    </>
  )
}
