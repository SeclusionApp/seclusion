import {
  Box,
  Button,
  Container,
  Flex,
  Heading,
  Icon,
  Link,
  Menu,
  MenuButton,
  MenuGroup,
  MenuItem,
  MenuList,
  Text,
} from "@chakra-ui/react";
import { AddIcon, ChevronDownIcon } from "@chakra-ui/icons";

import type { NextPage } from "next";
import Head from "next/head";
import { FaUser, FaUserPlus, FaUserSlash } from "react-icons/fa";

import router from "next/router";
import { useEffect, useState } from "react";
import DarkMenuItem from "./DarkMenuItem";
interface UserObj {
  id: number;
  username: string;
  email: string;
}
const NavBar: NextPage = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  async function getData() {
    try {
      const res = await fetch("http://localhost:8080/v1/user", {
        method: "GET",
        credentials: "include",
      });
      const data = await res.json();

      console.log("Headers: ", data);

      if (res.status === 401) {
        console.log({ error: "Not authorized" });
        return { error: "Not authorized" };
      }
      if (res.status === 200) {
        console.log(data);
        console.log(data.status);
        return data;
      }
    } catch (err) {
      return { error: "Not authorized" };
    }
  }
  const [user, setUser] = useState<UserObj>();
  useEffect(() => {
    getData().then((data) => {
      if (data.error) {
        setLoggedIn(false);
      } else {
        console.log("Data: ", data);
        setLoggedIn(true);

        setUser(data.user);
      }
    });
  }, []);

  let body = null;

  if (loggedIn) {
    body = (
      <Flex>
        <Box>
          <Menu closeOnSelect={false}>
            <MenuButton as={Button} rightIcon={<ChevronDownIcon />}>
              Welcome: {user?.username}
            </MenuButton>
            <MenuList>
              <MenuGroup title="View Options">
                <DarkMenuItem />
              </MenuGroup>
              <MenuGroup title="Logout">
                <MenuItem
                  icon={<Icon as={FaUserSlash} />}
                  onClick={async () => {
                    console.log();
                    fetch("http://localhost:8080/v1/auth/logout", {
                      method: "POST",
                      credentials: "include",
                    }).then((res) => {
                      if (res.status === 200) {
                        setLoggedIn(false);
                      }
                    });
                  }}
                >
                  Logout
                </MenuItem>
              </MenuGroup>
            </MenuList>
          </Menu>
        </Box>
      </Flex>
    );
  } else {
    body = (
      <Container>
        <Menu closeOnSelect={false}>
          <MenuButton as={Button} rightIcon={<ChevronDownIcon />}>
            Login / Register
          </MenuButton>
          <MenuList>
            <MenuItem
              icon={<Icon as={FaUser} />}
              onClick={() => router.push("/login")}
            >
              Login
            </MenuItem>
            <MenuItem
              icon={<Icon as={FaUserPlus} />}
              onClick={() => router.push("/register")}
            >
              Register
            </MenuItem>
            <MenuGroup title="View Options"></MenuGroup>
            <DarkMenuItem />
          </MenuList>
        </Menu>
      </Container>
    );
  }

  return (
    <>
      <Head>
        <title>Seclusion</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Flex
        p={4}
        ml={"auto"}
        position="sticky"
        top={0}
        zIndex={1}
        align="center"
      >
        <Heading>Seclusion</Heading>
        <Box ml={"auto"}>{body}</Box>
      </Flex>
    </>
  );
};

export default NavBar;
