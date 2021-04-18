import * as React from "react";
import { UserContext } from "./UserContext";
import {useContext} from "react";
import {Button} from "antd";

export function  UserDisplayComponent() {
  const userContext = useContext(UserContext);
  return (
    <>
      {userContext.user.name} <Button type="link" onClick={userContext.logout}>Logout</Button>
    </>
  )
}
