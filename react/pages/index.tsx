import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import { Box, Link, Container, Heading, Text } from "@chakra-ui/react";
import axios from "axios";
import { useState } from "react";
interface UserObj {
  id: number;
  username: string;
  email: string;
}
const Home: NextPage = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  async function getData() {
    try {
      const data = await axios.get("http://localhost:8080/v1/user", {
        withCredentials: true,
      });

      if (data.status === 401) {
        console.log({ error: "Not authorized" });
        return { error: "Not authorized" };
      }
      if (data.status === 200) {
        console.log(data.data);
        console.log(data.data.user);
        return data.data;
      }
    } catch (err) {
      console.log("");
      return { error: "Not authorized" };
    }
  }
  const [user, setUser] = useState<UserObj>();
  getData().then((data) => {
    console.log("promise");
    console.log(data.user);
    if (data.user && window !== undefined) {
      setUser(data.user);
      setLoggedIn(true);
    }
  });

  return (
    <>
      <Head>
        <title>Seclusion</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Box>
        {loggedIn ? (
          <Container>
            <Heading>Welcome {user?.username}</Heading>
            <Text>Here is your data:</Text>
            <Text>Username: {user?.username}</Text>
            <Text>Email: {user?.email}</Text>
          </Container>
        ) : (
          <Container>
            <Heading>Welcome to Seclusion</Heading>
            <Text>
              Please
              <Link href="/login"> login </Link>
              or
              <Link href="/register"> register </Link>
              to view your data.
            </Text>
          </Container>
        )}
      </Box>
    </>
  );
};

export default Home;
