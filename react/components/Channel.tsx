import { Box, Text } from "@chakra-ui/react";
import React from "react";
import { Messages } from "./Messages";

interface ChannelProps {}

export const Channel: React.FC<ChannelProps> = ({}) => {
  const [data, setData] = React.useState<any>(null);
  const [clicked, setClicked] = React.useState(true);
  React.useEffect(() => {
    (async () => {
      fetch("http://localhost:8080/v1/channels", {
        method: "GET",
        credentials: "include",
      })
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          setData(data);
        });
    })();
  }, []);
  return (
    <Box>
      <Text>Channel</Text>

      {data?.map((channel: { name: any; id: number }) => (
        <>
          <Box
            onClick={() => {
              //setClicked(!clicked);
            }}
          >
            <Box>{channel.name}</Box>

            {clicked ? <Messages id={channel.id} /> : null}
          </Box>
        </>
      ))}
    </Box>
  );
};
