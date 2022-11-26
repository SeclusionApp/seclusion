import {
  Box,
  Button,
  Container,
  Flex,
  Input,
  Stack,
  Text,
} from "@chakra-ui/react";
import { Formik } from "formik";
import React, { useEffect, useState } from "react";
import { MessagesObj, User } from "../types";
import { dateToHowLong, unixToDate } from "../utils/time";

interface MessagesProps {
  id: number;
}

export const Messages: React.FC<MessagesProps> = ({ id }) => {
  const [data, setData] = useState<MessagesObj>();
  const [user, setUser] = useState<User>();
  useEffect(() => {
    (async () => {
      const res = await fetch(`http://localhost:8080/v1/messages/${id}`, {
        method: "GET",
        credentials: "include",
      });

      const data = await res.json();
      console.log(data);
      setData(data);
    })();
  }, []);
  return (
    <Flex>
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
          <Flex>
            <Text>{message.content}</Text>
            <Text>{dateToHowLong(unixToDate(message.time))}</Text>
            <Text>{user?.username}</Text>
          </Flex>
        </Box>
      ))}
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
              channel_id: String(id),
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
            console.log(res);
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
                value={values.content}
                onChange={handleChange("content")}
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
    </Flex>
  );
};
