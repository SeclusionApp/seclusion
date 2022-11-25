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
import NavBar from "../components/NavBar";
import { dateToHowLong, unixToDate } from "../utils/time";
interface Message {
  id: number;
  content: string;
  user_id: number;
  channel_id: number;
  time: string;
}

interface User {
  id: number;
  username: string;
  email: string;
}

interface MessagesObj {
  messages: Message[];
  status: string;
}
const Home: NextPage = () => {
  const [data, setData] = useState<MessagesObj>();
  const [user, setUser] = useState<User>();
  const textInputRef = useRef<typeof Input>();
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

      <Container mt="20">
        <Heading>Home</Heading>
        <Text>Home page</Text>
        {data?.messages.map((message) => (
          <Box
            onClick={async () => {
              const d = await fetch(
                "http://localhost:8080/v1/users/" + message.user_id,
                {
                  method: "GET",
                  credentials: "include",
                }
              );
              console.log(d);
              const json = await d.json();
              console.log("json", json);
              setUser(json.user);
            }}
          >
            <Text>{message.content}</Text>
            <Text>{dateToHowLong(unixToDate(message.time))}</Text>
            <Text>{user?.username}</Text>
          </Box>
        ))}
      </Container>
      <Container>
        <Formik
          initialValues={{ content: "" }}
          onSubmit={async (values) => {
            const meRes = await fetch("http://localhost:8080/v1/user", {
              method: "GET",
              credentials: "include",
            });
            const me = await meRes.json();

            const cusotmValues = {
              ...values,
              channel_id: "1",
              user_id: String(me.user.id),
            };
            const res = await fetch("http://localhost:8080/v1/message", {
              method: "POST",
              credentials: "include",
              body: JSON.stringify(cusotmValues),
              headers: {
                "Content-Type": "application/json",
              },
            });
            //@ts-ignore
            textInputRef?.current?.clear();

            const data = await res.json();
            console.log(data);
          }}
        >
          {({ values, handleChange, handleSubmit }) => (
            <>
              <Input
                name="message"
                placeholder="Type a message"
                w="50%"
                //@ts-ignore
                ref={textInputRef}
                value={values.content}
                onChange={handleChange}
                onSubmitCapture={() => handleSubmit()}
                onSubmit={() => handleSubmit()}
                type="text"
              />
              <Button pl={25} onClick={() => handleSubmit()}>
                Submit
              </Button>
            </>
          )}
        </Formik>
      </Container>
    </>
  );
};

export default Home;
