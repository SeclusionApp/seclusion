import {
  Box,
  Button,
  Container,
  Flex,
  Heading,
  Input,
  Text,
} from "@chakra-ui/react";
import { Formik } from "formik";
import type { NextPage } from "next";
import { useEffect, useRef, useState } from "react";
import { Channel } from "../components/Channel";
import NavBar from "../components/NavBar";
import { dateToHowLong, unixToDate } from "../utils/time";

const Home: NextPage = () => {
  const textInputRef = useRef<typeof Input>();
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    (async () => {
      const res = await fetch("http://localhost:8080/v1/user", {
        method: "GET",
        credentials: "include",
      });
      const data = await res.json();
      console.log(data);
      if (data.message === "Invalid Token") {
        setLoggedIn(false);
      }
      if (data.user) {
        setLoggedIn(true);
      }
    })();
  }, []);

  return (
    <>
      <NavBar />

      <Container mt="20">
        <Heading>Home</Heading>
        <Text>Home page</Text>
      </Container>
      {loggedIn ? <Channel /> : null}
    </>
  );
};

export default Home;
