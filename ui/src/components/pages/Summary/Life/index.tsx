import React, { FC, Fragment, useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import apiClient from "~/utils/api";
import { dateFormat, timeFormat } from "~/utils/format";
import { API_HOST } from "~/utils/constants";
import Title from "~/components/common/Title";

type Time = {
  time: string;
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

const useStyles = makeStyles(() => ({
  cell: {
    minWidth: "120px",
  },
}));

const Life: FC = () => {
  const classes = useStyles();
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
    <Fragment>
      <Title>Life</Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell className={classes.cell}>date</TableCell>
            <TableCell className={classes.cell}>condition</TableCell>
            <TableCell className={classes.cell}>gets up</TableCell>
            <TableCell className={classes.cell}>sleep</TableCell>
            <TableCell className={classes.cell}>lights off</TableCell>
            <TableCell className={classes.cell}>bath</TableCell>
            <TableCell className={classes.cell}>fullness</TableCell>
            <TableCell className={classes.cell}>vitamin pill</TableCell>
            <TableCell className={classes.cell}>weather</TableCell>
            <TableCell className={classes.cell}>working</TableCell>
            <TableCell className={classes.cell}>effort</TableCell>
            <TableCell className={classes.cell}>hobby</TableCell>
            <TableCell className={classes.cell}>walking</TableCell>
            <TableCell className={classes.cell}>running</TableCell>
            <TableCell align="right" className={classes.cell}>
              burpee jump
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {lifes.map((row, i) => (
            <TableRow key={i}>
              <TableCell>
                {!row.date.valid && dateFormat(row.date.time)}
              </TableCell>
              <TableCell>{!row.condition.valid && row.condition.int}</TableCell>
              <TableCell>
                {!row.rising.valid && timeFormat(row.rising.time)}
              </TableCell>
              <TableCell>
                {!row.sleep.valid && `${row.sleep.float} h`}
              </TableCell>
              <TableCell>
                {!row.light_off.valid && row.light_off.bool ? "yes" : "no"}
              </TableCell>
              <TableCell>
                {!row.bath.valid && timeFormat(row.bath.time)}
              </TableCell>
              <TableCell>{!row.fullness.valid && row.fullness.int}</TableCell>
              <TableCell>
                {!row.vitamin.valid && row.vitamin.bool ? "yes" : "no"}
              </TableCell>
              <TableCell>{!row.weather.valid && row.weather.string}</TableCell>
              <TableCell>
                {!row.hunting.valid && `${row.hunting.float} h`}
              </TableCell>
              <TableCell>
                {!row.devotion.valid && `${row.devotion.float} h`}
              </TableCell>
              <TableCell>
                {!row.hobby.valid && `${row.hobby.float} h`}
              </TableCell>
              <TableCell>
                {!row.workout_w.valid && `${row.workout_w.float} h`}
              </TableCell>
              <TableCell>
                {!row.workout_r.valid && `${row.workout_r.float} km`}
              </TableCell>
              <TableCell align="right">
                {!row.workout_b.valid && `${row.workout_b.int} 回`}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Fragment>
  );
};

export default Life;
