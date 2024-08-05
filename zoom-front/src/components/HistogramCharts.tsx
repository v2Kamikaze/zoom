import React from "react";
import {
  LineChart,
  Line,
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
      <ResponsiveContainer width="25%" height={150}>
        <LineChart data={dataR}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line type="monotone" dataKey="value" stroke="#ff0000" />
        </LineChart>
      </ResponsiveContainer>

      <ResponsiveContainer width="25%" height={150}>
        <LineChart data={dataG}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line type="monotone" dataKey="value" stroke="#00ff00" />
        </LineChart>
      </ResponsiveContainer>

      <ResponsiveContainer width="25%" height={150}>
        <LineChart data={dataB}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line type="monotone" dataKey="value" stroke="#0000ff" />
        </LineChart>
      </ResponsiveContainer>

      <ResponsiveContainer width="25%" height={150}>
        <LineChart data={dataL}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="index" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line type="monotone" dataKey="value" stroke="#808080" />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};

export default HistogramCharts;
