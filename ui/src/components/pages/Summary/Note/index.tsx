import React, { FC, useState, useEffect } from "react";
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

type Post = {
  date: string;
  count: number;
};

const Note: FC = () => {
  const theme = useTheme();
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    const getPosts = async () => {
      const { response, error } = await apiClient.get<Post[]>(
        `${API_HOST}/note`,
      );
      if (error) return;
      setPosts(response);
    };
    getPosts();
  }, []);

  return (
    <React.Fragment>
      <Title>Note</Title>
      <ResponsiveContainer>
        <BarChart
          data={posts}
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

export default Note;
