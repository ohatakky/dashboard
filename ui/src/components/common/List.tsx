import React, { FC, Fragment } from "react";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import DashboardIcon from "@material-ui/icons/Dashboard";
import FilterHdrIcon from "@material-ui/icons/FilterHdr";
import BarChartIcon from "@material-ui/icons/BarChart";
import SubjectIcon from "@material-ui/icons/Subject";

export const BasicList: FC = () => (
  <Fragment>
    <ListItem button>
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Summary" />
    </ListItem>
    <ListItem button disabled>
      <ListItemIcon>
        <BarChartIcon />
      </ListItemIcon>
      <ListItemText primary="Analytics" />
    </ListItem>
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

// export const AdvancedList: FC = () => {};
