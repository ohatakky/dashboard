import React, { FC, useState, useEffect } from "react";
import clsx from "clsx";
import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import apiClient from "~/utils/api";
import { API_HOST } from "~/utils/constants";
import Title from "~/components/common/Title";
import Copyright from "~/components/common/Copyright";
import Deposit from "./Deposit";
import Order from "./Order";
import DailyCountChart, { Data } from "./DateCountChart";

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
}));

type Submission = {
  Date: string;
  Count: number;
};

const Summary: FC = () => {
  const classes = useStyles();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  const [submissions, setSubmissions] = useState<Data[]>([]);

  useEffect(() => {
    const getSubmissions = async () => {
      const { response, error } = await apiClient.get<Submission[]>(
        `${API_HOST}/atcoder`,
      );
      if (error) return;
      const data = response.map((a) => {
        return {
          date: a.Date,
          count: a.Count,
        };
      });
      setSubmissions(data);
    };
    getSubmissions();
  }, []);

  return (
    <Container maxWidth="lg" className={classes.container}>
      <Grid container spacing={3}>
        <Grid item xs={12} md={12} lg={12}>
          <Paper className={fixedHeightPaper}>
            <DailyCountChart title="Atcoder" data={submissions} />
          </Paper>
        </Grid>
        <Grid item xs={12} md={12} lg={12}>
          <Paper className={classes.paper}>
            <Title>Github</Title>
            <img src="https://grass-graph.moshimo.works/images/ohatakky.png" />
          </Paper>
        </Grid>
        <Grid item xs={12} md={4} lg={3}>
          <Paper className={fixedHeightPaper}>
            <Deposit />
          </Paper>
        </Grid>
        <Grid item xs={12}>
          <Paper className={classes.paper}>
            <Order />
          </Paper>
        </Grid>
      </Grid>
      <Box pt={4}>
        <Copyright />
      </Box>
    </Container>
  );
};

export default Summary;
