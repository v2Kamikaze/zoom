import React, { useState } from "react";
import { Input, Button, Card, Col, Row, Typography, Radio } from "antd";
import { ColorUtils } from "../services/ColorUtils";

const { Title, Text } = Typography;

const ColorConverter: React.FC = () => {
  const [r, setR] = useState<number>(0);
  const [g, setG] = useState<number>(0);
  const [b, setB] = useState<number>(0);
  const [h, setH] = useState<number>(0);
  const [s, setS] = useState<number>(0);
  const [v, setV] = useState<number>(0);
  const [conversionType, setConversionType] = useState<string>("RGBtoHSV");
  const [result, setResult] = useState<any>(null);

  const handleConvert = () => {
    let result;
    if (conversionType === "RGBtoHSV") {
      result = ColorUtils.RGBToHSV(r, g, b);
      result = {
        ...result,
        grayAverage: ColorUtils.toGrayAverage(r, g, b),
        grayWeighted: ColorUtils.toGrayWeighted(r, g, b),
        negative: ColorUtils.negativeRGB(r, g, b),
      };
    } else {
      result = ColorUtils.HSVToRGB(h, s, v);
      result = {
        ...result,
        gray: ColorUtils.toGrayHSV(h, s, v),
        negative: ColorUtils.negativeHSV(h, s, v),
      };
    }
    setResult(result);
  };

  return (
    <div style={{ padding: "20px" }}>
      <Title level={3}>Color Converter</Title>
      <Radio.Group
        value={conversionType}
        onChange={(e) => setConversionType(e.target.value)}
        style={{ marginBottom: "20px" }}
      >
        <Radio.Button value="RGBtoHSV">RGB to HSV</Radio.Button>
        <Radio.Button value="HSVtoRGB">HSV to RGB</Radio.Button>
      </Radio.Group>

      {conversionType === "RGBtoHSV" ? (
        <Row gutter={16}>
          <Col span={8}>
            <Input
              type="number"
              value={r}
              onChange={(e) => setR(parseInt(e.target.value, 10))}
              placeholder="Red (0-255)"
              min={0}
              max={255}
            />
          </Col>
          <Col span={8}>
            <Input
              type="number"
              value={g}
              onChange={(e) => setG(parseInt(e.target.value, 10))}
              placeholder="Green (0-255)"
              min={0}
              max={255}
            />
          </Col>
          <Col span={8}>
            <Input
              type="number"
              value={b}
              onChange={(e) => setB(parseInt(e.target.value, 10))}
              placeholder="Blue (0-255)"
              min={0}
              max={255}
            />
          </Col>
        </Row>
      ) : (
        <Row gutter={16}>
          <Col span={8}>
            <Input
              type="number"
              value={h}
              onChange={(e) => setH(parseFloat(e.target.value))}
              placeholder="Hue (0-360)"
              min={0}
              max={360}
            />
          </Col>
          <Col span={8}>
            <Input
              type="number"
              value={s}
              onChange={(e) => setS(parseFloat(e.target.value))}
              placeholder="Saturation (0-1)"
              min={0}
              max={1}
              step={0.01}
            />
          </Col>
          <Col span={8}>
            <Input
              type="number"
              value={v}
              onChange={(e) => setV(parseFloat(e.target.value))}
              placeholder="Value (0-1)"
              min={0}
              max={1}
              step={0.01}
            />
          </Col>
        </Row>
      )}

      <Button
        type="primary"
        onClick={handleConvert}
        style={{ marginTop: "20px" }}
      >
        Convert
      </Button>

      <div style={{ marginTop: "20px" }}>
        <Card title="Results">
          {result && (
            <div>
              {conversionType === "RGBtoHSV" ? (
                <Row gutter={16}>
                  <Col span={6}>
                    <Title level={4}>HSV</Title>
                    <Text>H: {result.h?.toFixed(2)}</Text>
                  </Col>
                  <Col span={6}>
                    <Text>S: {result.s?.toFixed(2)}</Text>
                  </Col>
                  <Col span={6}>
                    <Text>V: {result.v?.toFixed(2)}</Text>
                  </Col>
                  <Col span={6}>
                    <Title level={4}>Gray (Average)</Title>
                    <Text>Gray: {result.grayAverage}</Text>
                  </Col>
                  <Col span={6}>
                    <Title level={4}>Gray (Weighted)</Title>
                    <Text>Gray: {result.grayWeighted}</Text>
                  </Col>
                  <Col span={6}>
                    <Title level={4}>Negative RGB</Title>
                    <Text>R: {result.negative?.nr}</Text>
                    <br />
                    <Text>G: {result.negative?.ng}</Text>
                    <br />
                    <Text>B: {result.negative?.nb}</Text>
                  </Col>
                </Row>
              ) : (
                <Row gutter={16}>
                  <Col span={6}>
                    <Title level={4}>RGB</Title>
                    <Text>R: {result.r}</Text>
                  </Col>
                  <Col span={6}>
                    <Text>G: {result.g}</Text>
                  </Col>
                  <Col span={6}>
                    <Text>B: {result.b}</Text>
                  </Col>
                  <Col span={6}>
                    <Title level={4}>Gray (HSV)</Title>
                    <Text>Gray: {result.gray}</Text>
                  </Col>
                  <Col span={6}>
                    <Title level={4}>Negative HSV</Title>
                    <Text>H: {result.nh?.toFixed(2)}</Text>
                    <br />
                    <Text>S: {result.ns?.toFixed(2)}</Text>
                    <br />
                    <Text>V: {result.nv?.toFixed(2)}</Text>
                  </Col>
                </Row>
              )}
            </div>
          )}
        </Card>
      </div>
    </div>
  );
};

export default ColorConverter;
