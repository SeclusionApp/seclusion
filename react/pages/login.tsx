import { Box, Button, Container, Input } from "@chakra-ui/react";
import React from "react";
import { Formik } from "formik";
import axios from "axios";
interface loginProps {}

const Login: React.FC<loginProps> = ({}) => {
  return (
    <Container>
      <h1>Login</h1>
      <Formik
        initialValues={{ username: "", password: "" }}
        onSubmit={async (values) => {
          console.log(values);

          const res = await axios.post("http://localhost:8080/v1/auth/login", {
            values,
          });

          console.log(res);
        }}
      >
        {({ values, handleChange, handleBlur, handleSubmit }) => (
          <Box>
            <Input
              name="username"
              placeholder="username"
              value={values.username}
              size="md"
              onChange={handleChange("username")}
              onBlur={handleBlur("username")}
            />
            <Input
              name="password"
              placeholder="password"
              type={"password"}
              size="md"
              value={values.password}
              onChange={handleChange("password")}
              onBlur={handleBlur("password")}
            />
            <Button type="submit" onClick={() => handleSubmit()}>
              Submit
            </Button>
          </Box>
        )}
      </Formik>
    </Container>
  );
};
export default Login;
