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

type Submission = {
  Date: string;
  Count: number;
};

const Atcoder: FC = () => {
  const theme = useTheme();
  const [submissions, setSubmissions] = useState<Submission[]>([]);

  useEffect(() => {
    const getSubmissions = async () => {
      const { response, error } = await apiClient.get<Submission[]>(
        `${API_HOST}/atcoder`,
      );
      if (error) return;
      setSubmissions(response);
    };
    getSubmissions();
  }, []);

  return (
    <React.Fragment>
      <Title>Atcoder</Title>
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
          <XAxis dataKey="Date" stroke={theme.palette.text.secondary} />
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
            dataKey="Count"
            fill={theme.palette.primary.main}
            stroke={theme.palette.primary.main}
          />
        </BarChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
};

export default Atcoder;
