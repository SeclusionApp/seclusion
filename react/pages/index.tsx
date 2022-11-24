import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import { Box, Link, Container, Heading, Text, Button } from "@chakra-ui/react";
import axios from "axios";
import { useEffect, useState } from "react";
import NavBar from "../components/NavBar";
interface UserObj {
  id: number;
  username: string;
  email: string;
}
const Home: NextPage = () => {
  return (
    <>
      <NavBar />

      <Box mt="20">
        <Heading>Home</Heading>
        <Text>Home page</Text>
      </Box>
    </>
  );
};

export default Home;
