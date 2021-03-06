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
  sleep: number;
};

type SleepProps = {
  data: Data[];
};

const Sleep: FC<SleepProps> = ({ data }) => {
  return (
    <Fragment>
      <Title>Sleep</Title>
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
          ticks={[...Array(12)].map((_, i) => i + 1)}
          orientation="right"
        />
        <YAxis yAxisId={2} ticks={[...Array(10)].map((_, i) => i + 1)} />
        <Tooltip />
        <Legend />
        <Line yAxisId={1} dataKey="sleep" type="monotone" stroke="#413ea0" />
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

export default Sleep;
