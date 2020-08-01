import React, { FC } from "react";
import { useTheme } from "@material-ui/core/styles";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Label,
  ResponsiveContainer,
} from "recharts";
import Title from "~/components/common/Title";

export type Data = {
  date: string;
  count: number;
};

type DailyCountChartProps = {
  title: string;
  data: Data[];
};

const DailyCountChart: FC<DailyCountChartProps> = ({ title, data }) => {
  const theme = useTheme();

  return (
    <React.Fragment>
      <Title>{title}</Title>
      <ResponsiveContainer>
        <BarChart
          data={data}
          margin={{
            top: 16,
            right: 16,
            bottom: 0,
            left: 24,
          }}
        >
          <XAxis dataKey="date" stroke={theme.palette.text.secondary} />
          <YAxis stroke={theme.palette.text.secondary}>
            <Label
              angle={270}
              position="left"
              style={{ textAnchor: "middle", fill: theme.palette.text.primary }}
            >
              Count
            </Label>
          </YAxis>
          <Bar
            dataKey="count"
            fill={theme.palette.primary.main}
            stroke={theme.palette.primary.main}
          />
        </BarChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
};

export default DailyCountChart;
