import React, { FC } from "react";
import Typography from "@material-ui/core/Typography";
import Link from "@material-ui/core/Link";

const Copyright: FC = () => {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link
        color="inherit"
        href="https://www.notion.so/ohatakky/4e859353a38643ab9dbda1f257d22909"
        target="_blank"
      >
        ohatakky
      </Link>
      {" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
};

export default Copyright;
