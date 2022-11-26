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

  return (
    <>
      <NavBar />

      <Container mt="20">
        <Heading>Home</Heading>
        <Text>Home page</Text>
      </Container>

      <Channel />
    </>
  );
};

export default Home;
