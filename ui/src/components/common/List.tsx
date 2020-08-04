import React, { FC, Fragment } from "react";
import { Link } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import DashboardIcon from "@material-ui/icons/Dashboard";
import FilterHdrIcon from "@material-ui/icons/FilterHdr";
import BarChartIcon from "@material-ui/icons/BarChart";
import SubjectIcon from "@material-ui/icons/Subject";

const useStyles = makeStyles((theme) => ({
  link: {
    textDecoration: "none",
    color: theme.palette.text.secondary,
  },
}));

const BasicList: FC = () => {
  const classes = useStyles();
  return (
    <Fragment>
      <Link to="summary" className={classes.link}>
        <ListItem button>
          <ListItemIcon>
            <DashboardIcon />
          </ListItemIcon>
          <ListItemText primary="Summary" />
        </ListItem>
      </Link>
      <Link to="analytics" className={classes.link}>
        <ListItem button>
          <ListItemIcon>
            <BarChartIcon />
          </ListItemIcon>
          <ListItemText primary="Analytics" />
        </ListItem>
      </Link>
      <ListItem button disabled>
        <ListItemIcon>
          <FilterHdrIcon />
        </ListItemIcon>
        <ListItemText primary="Climbing" />
      </ListItem>
      <ListItem button disabled>
        <ListItemIcon>
          <SubjectIcon />
        </ListItemIcon>
        <ListItemText primary="Report" />
      </ListItem>
    </Fragment>
  );
};

export default BasicList;
