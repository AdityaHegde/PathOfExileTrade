import * as React from "react";
import {
  MemoryRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import { Layout } from "antd";
import {UserComponent} from "@component/user/UserComponent";
import {UserContextComponent} from "@component/user/UserContextComponent";
import {LoginComponent} from "@component/user/LoginComponent";
import {Dashboard} from "@component/Dashboard";

const { Header, Content } = Layout;

export function Home() {
  return (
    <Router>
      <UserContextComponent>
        <Layout className="layout">
          <Header>
            <UserComponent />
          </Header>
          <Content style={{ minHeight: "1000px", margin: "25px" }}>
            <Switch>
              <Route key="/" path="/" exact>
                HOME
              </Route>
              <Route key="/dash" path="/dash" exact>
                <Dashboard />
              </Route>
              <Route key="/login" path="/login" exact>
                <LoginComponent />
              </Route>
            </Switch>
          </Content>
        </Layout>
      </UserContextComponent>
    </Router>
  )
}
