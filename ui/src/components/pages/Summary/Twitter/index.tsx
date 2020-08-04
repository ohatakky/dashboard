import React, { FC, Fragment, useState, useEffect } from "react";
import { useTheme } from "@material-ui/core/styles";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Label,
  ResponsiveContainer,
} from "recharts";
import apiClient from "~/utils/api";
import { API_HOST } from "~/utils/constants";
import Title from "~/components/common/Title";

type Tweet = {
  date: string;
  count: number;
};

const Twitter: FC = () => {
  const theme = useTheme();
  const [submissions, setTweets] = useState<Tweet[]>([]);

  useEffect(() => {
    const getTweets = async () => {
      const { response, error } = await apiClient.get<Tweet[]>(
        `${API_HOST}/twitter`
      );
      if (error) return;
      setTweets(response);
    };
    getTweets();
  }, []);

  return (
    <Fragment>
      <Title>Twitter</Title>
      <ResponsiveContainer>
        <BarChart
          data={submissions}
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
              style={{
                textAnchor: "middle",
                fill: theme.palette.text.secondary,
              }}
            >
              tweets
            </Label>
          </YAxis>
          <Bar
            dataKey="count"
            fill={theme.palette.primary.main}
            stroke={theme.palette.primary.main}
          />
        </BarChart>
      </ResponsiveContainer>
    </Fragment>
  );
};

export default Twitter;
