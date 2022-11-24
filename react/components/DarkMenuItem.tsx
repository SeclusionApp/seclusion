import {
  MenuItem,
  Icon,
  Switch,
  ColorMode,
  useColorMode,
} from "@chakra-ui/react";
import React from "react";
import { FaMoon } from "react-icons/fa";

function DarkMenuItem() {
  const { colorMode, toggleColorMode } = useColorMode();
  let light: any = null;
  if (colorMode === "light") {
    light = false;
  } else {
    light = true;
  }
  return (
    <MenuItem onClick={toggleColorMode} icon={<Icon as={FaMoon} />}>
      Dark Mode
      <Switch
        ml={2}
        isChecked={light}
        isReadOnly={true}
        onClick={toggleColorMode}
      >
        {" "}
      </Switch>
    </MenuItem>
  );
}

export default DarkMenuItem;
