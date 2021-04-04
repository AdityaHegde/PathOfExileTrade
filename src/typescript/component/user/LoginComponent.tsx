import React, {useContext, useState} from "react";
import { UserContext } from "./UserContext";
import {Button, Checkbox, Form, Input} from "antd";

export function LoginComponent() {
  const [isSignup, useIsSignup] = useState<boolean>(false);
  const userContext = useContext(UserContext);

  const onFinish = (values: any) => {
    if (isSignup) {
      userContext.signup(values);
    } else {
      userContext.login(values);
    }
  };

  return (
    <Form onFinish={onFinish}>
      <Form.Item name="name" key="name" label="User Name">
        <Input />
      </Form.Item>
      {isSignup ? <Form.Item name="email" key="email" label="Email Address">
        <Input type="email"/>
      </Form.Item> : ""}
      <Form.Item name="password" key="password" label="Password">
        <Input type="password" />
      </Form.Item>
      <Form.Item key="signup" label="Signup">
        <Checkbox defaultChecked={false} onChange={(e) => {
          useIsSignup(e.target.checked)
        }} />
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit">Submit</Button>
      </Form.Item>
    </Form>
  );
}
