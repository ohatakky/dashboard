import React, { FC, Fragment } from "react";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import DashboardIcon from "@material-ui/icons/Dashboard";
import FilterHdrIcon from "@material-ui/icons/FilterHdr";
import FitnessCenterIcon from "@material-ui/icons/FitnessCenter";
import BarChartIcon from "@material-ui/icons/BarChart";

export const BasicList: FC = () => (
  <Fragment>
    <ListItem button>
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Summary" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <FitnessCenterIcon />
      </ListItemIcon>
      <ListItemText primary="Progress" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <FilterHdrIcon />
      </ListItemIcon>
      <ListItemText primary="Climbing" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <BarChartIcon />
      </ListItemIcon>
      <ListItemText primary="Records" />
    </ListItem>
  </Fragment>
);

// export const AdvancedList: FC = () => {};
