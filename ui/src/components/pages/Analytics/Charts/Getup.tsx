import React, { FC, Fragment } from "react";
import {
  ComposedChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
} from "recharts";
import Title from "~/components/common/Title";

export type Data = {
  date: string;
  condition: number;
  rising: number;
};

type GetupProps = {
  data: Data[];
};

const Getup: FC<GetupProps> = ({ data }) => {
  return (
    <Fragment>
      <Title>Getup</Title>
      <ComposedChart
        width={600}
        height={300}
        data={data}
        margin={{
          top: 20,
          right: 20,
          bottom: 20,
          left: 20,
        }}
      >
        <CartesianGrid stroke="#f5f5f5" />
        <XAxis dataKey="date" />
        <YAxis
          yAxisId={1}
          orientation="right"
          // label={{ value: "get up [h.m]", angle: -90, color: "#413ea0" }}
          ticks={[...Array(24)].map((_, i) => i + 1)}
        />
        <YAxis
          yAxisId={2}
          // label={{ value: "condition [10]", angle: -90, color: "#ff7300" }}
          ticks={[...Array(10)].map((_, i) => i + 1)}
        />
        <Tooltip />
        <Legend />
        <Line
          yAxisId={1}
          dataKey="rising"
          type="monotone"
          stroke="#413ea0"
        />
        <Line
          yAxisId={2}
          dataKey="condition"
          type="monotone"
          stroke="#ff7300"
        />
      </ComposedChart>
    </Fragment>
  );
};

export default Getup;

// tickFormatter={(unixTime) => timeFormat(unixTime)}
