import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import { Box, Link, Container, Heading, Text, Button } from "@chakra-ui/react";
import axios from "axios";
import { useEffect, useState } from "react";
import NavBar from "../components/NavBar";
interface Message {
  ID: number;
  Content: string;
  UserID: number;
  ChannelID: number;
  Time: string;
}

interface MessagesObj {
  messages: Message[];
  status: string;
}
const Home: NextPage = () => {
  const [data, setData] = useState<MessagesObj>();
  useEffect(() => {
    (async () => {
      const res = await fetch("http://localhost:8080/v1/messages", {
        method: "GET",
        credentials: "include",
      });
      const data = await res.json();
      console.log(data);
      setData(data);
    })();
  }, []);
  return (
    <>
      <NavBar />

      <Box mt="20">
        <Heading>Home</Heading>
        <Text>Home page</Text>
        {data?.messages.map((message) => (
          <Text>{message.Content}</Text>
        ))}
      </Box>
    </>
  );
};

export default Home;
