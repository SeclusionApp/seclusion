import {
  Button,
  Center,
  Container,
  Heading,
  Input,
  Link,
  Text,
} from "@chakra-ui/react";
import { Formik } from "formik";
import { NextPage } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import React from "react";

const Channel: NextPage = () => {
  const router = useRouter();
  return (
    <Container>
      <Heading>Channel</Heading>
      <Container mt="10">
        <Head>
          <title>Create Channel</title>
        </Head>
        <Formik
          initialValues={{ name: "" }}
          onSubmit={async (values) => {
            console.log(values);

            const res = await fetch("http://localhost:8080/v1/channel", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                name: values.name,
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
                name="channel"
                placeholder="Name"
                value={values.name}
                size="md"
                onChange={handleChange("name")}
                onBlur={handleBlur("name")}
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
                  Create Channel
                </Button>
              </Center>
            </Container>
          )}
        </Formik>
      </Container>
    </Container>
  );
};

export default Channel;
