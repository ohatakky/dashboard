import React, { FC, Fragment } from "react";
import { Link } from "react-router-dom";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import DashboardIcon from "@material-ui/icons/Dashboard";
import FilterHdrIcon from "@material-ui/icons/FilterHdr";
import BarChartIcon from "@material-ui/icons/BarChart";
import SubjectIcon from "@material-ui/icons/Subject";

const BasicList: FC = () => {
  return (
    <Fragment>
      <Link to="summary">
        <ListItem button>
          <ListItemIcon>
            <DashboardIcon />
          </ListItemIcon>
          <ListItemText primary="Summary" />
        </ListItem>
      </Link>
      <Link to="analytics">
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
