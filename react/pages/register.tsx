import { Box, Button, Container, Input, Link, Text } from "@chakra-ui/react";
import React from "react";
import { Formik } from "formik";
import axios from "axios";
import { useRouter } from "next/router";
interface loginProps {}

const Register: React.FC<loginProps> = ({}) => {
  const router = useRouter();
  return (
    <Container>
      <h1>Register</h1>
      <Formik
        initialValues={{ username: "", password: "", email: "" }}
        onSubmit={async (values) => {
          console.log(values);

          let headersList = {
            Accept: "*/*",
            "Content-Type": "application/json",
          };

          let bodyContent = JSON.stringify({
            email: values.email,
            username: values.username,
            password: values.password,
          });

          let reqOptions = {
            url: "http://localhost:8080/v1/auth/register",
            method: "POST",
            headers: headersList,
            data: bodyContent,
          };

          let response = await axios.request(reqOptions);
          if (response.data.error) {
            console.log(response.data.error);
          }
          if (response.status === 200) {
            console.log(response.data);
            router.push("/");
          }
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
              onError={(err) => console.log(err)}
            />
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
              value={values.password}
              onChange={handleChange("password")}
              onBlur={handleBlur("password")}
            />
            <Button type="submit" onClick={() => handleSubmit()}>
              Submit
            </Button>
            <Text>
              Don't have an account? <Link href="/register">Register</Link>
            </Text>
          </Box>
        )}
      </Formik>
    </Container>
  );
};
export default Register;
