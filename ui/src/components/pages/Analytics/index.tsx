import React, { FC, Fragment, useState, useEffect } from "react";
import clsx from "clsx";
import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Copyright from "~/components/common/Copyright";
import Card from "~/components/common/Card";
import apiClient from "~/utils/api";
import { dateFormat } from "~/utils/format";
import { API_HOST } from "~/utils/constants";
import { Life } from "~/utils/types";
import { Sleep, SleepData, Getup, GetupData } from "./Charts";

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
    height: 360,
  },
}));

const Analytics: FC = () => {
  const classes = useStyles();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  const [getups, setGetups] = useState<GetupData[]>([]);
  const [sleeps, setSleeps] = useState<SleepData[]>([]);

  useEffect(() => {
    const getLifes = async () => {
      const { response, error } = await apiClient.get<Life[]>(
        `${API_HOST}/life`,
      );
      if (error) return;
      let tmpSleeps: SleepData[] = [];
      let tmpGetups: GetupData[] = [];
      response.map((res) => {
        tmpSleeps = !res.sleep.valid
          ? [
            ...tmpSleeps,
            {
              date: dateFormat(res.date.time),
              condition: res.condition.int,
              sleep: res.sleep.float,
            },
          ]
          : tmpSleeps;
        tmpGetups = !res.rising.valid
          ? [
            ...tmpGetups,
            {
              date: dateFormat(res.date.time),
              condition: res.condition.int,
              rising: new Date(res.rising.time).getHours() +
                (new Date(res.rising.time).getMinutes() / 100),
            },
          ]
          : tmpGetups;
      });
      setSleeps(tmpSleeps);
      setGetups(tmpGetups);
    };
    getLifes();
  }, []);

  return (
    <Container maxWidth="lg" className={classes.container}>
      <Grid container spacing={3}>
        <Grid item xs={9} md={8} lg={8}>
          <Paper className={fixedHeightPaper}>
            <Sleep data={sleeps} />
          </Paper>
        </Grid>
        <Grid item xs={3} md={4} lg={4}>
          <Paper className={fixedHeightPaper}>
            <Card
              title="Sleep"
              average={sleeps.reduce((acc, cur) => {
                return acc + cur.sleep;
              }, 0) / sleeps.length}
            />
          </Paper>
        </Grid>
        <Grid item xs={9} md={8} lg={8}>
          <Paper className={fixedHeightPaper}>
            <Getup data={getups} />
          </Paper>
        </Grid>
        <Grid item xs={3} md={4} lg={4}>
          <Paper className={fixedHeightPaper}>
            <Card
              title="Getup"
              average={getups.reduce((acc, cur) => {
                return acc + cur.rising;
              }, 0) / getups.length}
            />
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
