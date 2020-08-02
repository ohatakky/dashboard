import React, { FC, useState, useEffect } from "react";
import clsx from "clsx";
import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Title from "~/components/common/Title";
import Copyright from "~/components/common/Copyright";
import Atcoder from "./Atcoder";
import Note from "./Note";

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

const Summary: FC = () => {
  const classes = useStyles();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);

  return (
    <Container maxWidth="lg" className={classes.container}>
      <Grid container spacing={3}>
        <Grid item xs={12} md={12} lg={12}>
          <Paper className={fixedHeightPaper}>
            <Atcoder />
          </Paper>
        </Grid>
        <Grid item xs={12} md={12} lg={12}>
          <Paper className={fixedHeightPaper}>
            <Note />
          </Paper>
        </Grid>
        <Grid item xs={12} md={10} lg={7}>
          <Paper className={classes.paper}>
            <Title>Github</Title>
            <img src="https://grass-graph.moshimo.works/images/ohatakky.png" />
          </Paper>
        </Grid>
        <Grid item xs={12} md={8} lg={5}>
          <Paper className={fixedHeightPaper}>
            <Title>Github</Title>
            <a href="https://github.com/ohatakky" target="_blank">
              <img
                src="https://github-readme-stats.vercel.app/api?username=ohatakky&count_private=true&hide=contribs,prs,issues,stars"
              />
            </a>
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
