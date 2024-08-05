import React from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";
import { Histogram } from "../services/ZoomService";

type HistogramProps = {
  histogram: Histogram;
};

const HistogramCharts: React.FC<HistogramProps> = ({ histogram }) => {
  const formatData = (data: number[]) =>
    data.map((value, index) => ({ index, value }));

  const dataR = formatData(histogram.r);
  const dataG = formatData(histogram.g);
  const dataB = formatData(histogram.b);
  const dataL = formatData(histogram.l);

  return (
    <div style={{ display: "flex", flexWrap: "wrap", gap: "20px" }}>
      <ResponsiveContainer width="100%" height={200}>
        <BarChart data={dataR}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Bar dataKey="value" fill="#ff0000" />
        </BarChart>
      </ResponsiveContainer>

      <ResponsiveContainer width="100%" height={200}>
        <BarChart data={dataG}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Bar dataKey="value" fill="#00ff00" />
        </BarChart>
      </ResponsiveContainer>

      <ResponsiveContainer width="100%" height={200}>
        <BarChart data={dataB}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Bar dataKey="value" fill="#0000ff" />
        </BarChart>
      </ResponsiveContainer>

      <ResponsiveContainer width="100%" height={200}>
        <BarChart data={dataL}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Bar dataKey="value" fill="#808080" />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default HistogramCharts;
