import React, { FC, Fragment, useState, useEffect } from "react";
import clsx from "clsx";
import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Copyright from "~/components/common/Copyright";
import apiClient from "~/utils/api";
import { dateFormat, timeFormat } from "~/utils/format";
import { API_HOST } from "~/utils/constants";
import { Life } from "~/utils/types";
import { Sleep } from "./Charts";

const useStyles = makeStyles((theme) => ({
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: "flex",
    overflow: "auto",
    flexDirection: "column",
  },
  fixedHeight: {
    height: 240,
  },
  lifeHeight: {
    height: 480,
  },
}));

const Analytics: FC = () => {
  const classes = useStyles();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  const [lifes, setLifes] = useState<Life[]>([]);

  useEffect(() => {
    const getLifes = async () => {
      const { response, error } = await apiClient.get<Life[]>(
        `${API_HOST}/life`
      );
      if (error) return;
      setLifes(response);
    };
    getLifes();
  }, []);
  return (
    <Container maxWidth="lg" className={classes.container}>
      <Grid container spacing={3}>
        <Grid item xs={12} md={12} lg={12}>
          <Paper className={fixedHeightPaper}>
            <Sleep />
          </Paper>
        </Grid>
      </Grid>
      <Box pt={4}>
        <Copyright />
      </Box>
    </Container>
  );
};

export default Analytics;
