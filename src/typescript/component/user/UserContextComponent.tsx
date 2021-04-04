import * as React from "react";
import {UserContext, useUserContext} from "./UserContext";

export function UserContextComponent(props) {
  const userContextType = useUserContext();

  console.log("UserContextComponent", userContextType);

  return (
    <UserContext.Provider value={userContextType}>
      {props.children}
    </UserContext.Provider>
  )
}
