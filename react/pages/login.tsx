import {
  Box,
  Button,
  Center,
  Container,
  Input,
  Link,
  Text,
} from "@chakra-ui/react";
import React from "react";
import { Formik } from "formik";
import axios from "axios";
import { useRouter } from "next/router";
import Head from "next/head";
interface loginProps {}

const Login: React.FC<loginProps> = ({}) => {
  const router = useRouter();
  return (
    <Container mt="20">
      <Head>
        <title>Login</title>
      </Head>
      <Formik
        initialValues={{ email: "", password: "" }}
        onSubmit={async (values) => {
          console.log(values);

          const res = await fetch("http://localhost:8080/v1/auth/login", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              email: values.email,
              password: values.password,
            }),
            credentials: "include",
          });

          const data = await res.json();
          console.log(data);

          router.push("/");
        }}
      >
        {({ values, handleChange, handleBlur, handleSubmit }) => (
          <Container>
            <Input
              name="email"
              placeholder="email"
              value={values.email}
              size="md"
              onChange={handleChange("email")}
              onBlur={handleBlur("email")}
            />
            <Input
              name="password"
              placeholder="password"
              type={"password"}
              size="md"
              pt={5}
              pb={5}
              value={values.password}
              onChange={handleChange("password")}
              onBlur={handleBlur("password")}
            />
            <Center pt={5} pb={5}>
              <Button
                type="submit"
                pt={2}
                size="md"
                variant="solid"
                bg={"blue.500"}
                onClick={() => handleSubmit()}
              >
                Submit
              </Button>
            </Center>
            <Text>
              Don't have an account? <Link href="/register">Register</Link>
            </Text>
          </Container>
        )}
      </Formik>
    </Container>
  );
};
export default Login;
