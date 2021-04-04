import * as React from "React";
import { Link } from "react-router-dom";
import {UserDisplayComponent} from "./UserDisplayComponent";
import { UserContext } from "./UserContext";
import {useContext} from "react";

export function UserComponent() {
  const userContext = useContext(UserContext);

  return (
    <>{userContext.user ? <UserDisplayComponent /> : <Link to="/login">Login</Link>}</>
  );
}
