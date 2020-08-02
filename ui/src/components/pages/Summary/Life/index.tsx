import React, { FC, useState, useEffect } from "react";
import Link from "@material-ui/core/Link";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import apiClient from "~/utils/api";
import { API_HOST } from "~/utils/constants";
import Title from "~/components/common/Title";

type Time = {
  time: Date;
  valid: boolean;
};

type String = {
  string: string;
  valid: boolean;
};

type Int = {
  int: number;
  valid: boolean;
};

type Float = {
  float: number;
  valid: boolean;
};

type Bool = {
  bool: boolean;
  valid: boolean;
};

type Life = {
  date: Time;
  condition: Int;
  rising: Time;
  sleep: Float;
  light_off: Bool;
  bath: Time;
  fullness: Int;
  vitamin: Bool;
  weather: String;
  hunting: Float;
  devotion: Float;
  hobby: Float;
  workout_w: Float;
  workout_r: Float;
  workout_b: Int;
};

const useStyles = makeStyles((theme) => ({
  seeMore: {
    marginTop: theme.spacing(3),
  },
}));

const Life: FC = () => {
  const classes = useStyles();
  const [lifes, setLifes] = useState<Life[]>([]);

  useEffect(() => {
    const getLifes = async () => {
      const { response, error } = await apiClient.get<Life[]>(
        `${API_HOST}/life`,
      );
      if (error) return;
      setLifes(response);
    };
    getLifes();
  }, []);
  return (
    <React.Fragment>
      <Title>Life</Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell>Date</TableCell>
            <TableCell>Condition</TableCell>
            <TableCell>Rising</TableCell>
            <TableCell>Sleep</TableCell>
            <TableCell>Light Off</TableCell>
            <TableCell>Bath</TableCell>
            <TableCell>Fullnes</TableCell>
            <TableCell>Vitamin</TableCell>
            <TableCell>Weather</TableCell>
            <TableCell>Hunting</TableCell>
            <TableCell>Devotion</TableCell>
            <TableCell>Hobby</TableCell>
            <TableCell>WorkoutW</TableCell>
            <TableCell>WorkoutR</TableCell>
            <TableCell align="right">WorkoutB</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {lifes.map((row, i) => (
            <TableRow key={i}>
              <TableCell>{row.date.time}</TableCell>
              <TableCell>{row.condition.int}</TableCell>
              <TableCell>{row.rising.time}</TableCell>
              <TableCell>{row.sleep.float}</TableCell>
              <TableCell>{row.light_off.bool}</TableCell>
              <TableCell>{row.bath.time}</TableCell>
              <TableCell>{row.fullness.int}</TableCell>
              <TableCell>{row.vitamin.bool}</TableCell>
              <TableCell>{row.weather.string}</TableCell>
              <TableCell>{`${row.hunting.float} h`}</TableCell>
              <TableCell>{`${row.devotion.float} h`}</TableCell>
              <TableCell>{`${row.hobby.float} h`}</TableCell>
              <TableCell>{`${row.workout_w.float} h`}</TableCell>
              <TableCell>{`${row.workout_r.float} km`}</TableCell>
              <TableCell align="right">
                {`${row.workout_b.int} times`}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <div className={classes.seeMore}>
        <Link color="primary" href="#" onClick={(e) => e.preventDefault()}>
          See Charts
        </Link>
      </div>
    </React.Fragment>
  );
};

export default Life;
