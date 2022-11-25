import { Box, Text } from "@chakra-ui/react";
import React from "react";

interface ChannelProps {}

export const Channel: React.FC<ChannelProps> = ({}) => {
  const [data, setData] = React.useState<any>(null);

  React.useEffect(() => {
    fetch("http://localhost:8080/v1/messages", {
      method: "GET",
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setData(data);
      });
  }, []);
  return (
    <Box>
      <Text>Channel</Text>
    </Box>
  );
};
