import * as React from "react";
import * as ReactDOM from "react-dom";

import { Home } from "@component/Home";

const div = document.createElement("div");
document.body.appendChild(div)

ReactDOM.render(
  <Home />,
  div,
);
