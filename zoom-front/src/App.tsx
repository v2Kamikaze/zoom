import {
  Button,
  ConfigProvider,
  Drawer,
  Flex,
  Form,
  Image,
  InputNumber,
  Layout,
  Menu,
  Switch,
  theme,
} from "antd";
import React from "react";
import { UploadFile } from "./components/UploadFile";
import ZoomService, { ZoomResponse } from "./services/ZoomService";
import {
  UploadOutlined,
  SwapOutlined,
  BgColorsOutlined,
  ArrowsAltOutlined,
  RedoOutlined,
} from "@ant-design/icons";
import { useNotification } from "./hooks/useNotification";

const { Sider, Content, Header } = Layout;

const Zoom: React.FC = () => {
  // Theme
  const { defaultAlgorithm, darkAlgorithm } = theme;
  const [isDarkMode, setIsDarkMode] = React.useState(false);
  const toggleTheme = () => setIsDarkMode((prev) => !prev);
  const themeType = isDarkMode ? "dark" : "light";

  // Notification
  const { notificationContext, openNotification } = useNotification();

  const [siderOpened, setSiderOpened] = React.useState(false);
  const toggleSider = () => setSiderOpened(!siderOpened);

  const [drawerOpened, setDrawerOpened] = React.useState(false);
  const openDrawer = () => setDrawerOpened(true);
  const closeDrawer = () => setDrawerOpened(false);

  // Data States
  const [file, setFile] = React.useState<File>(new File([], ""));
  const [processedFile, setProcessedFile] = React.useState<File>(file);

  const [preview, setPreview] = React.useState("");
  const [processedPreview, setProcessedPreview] = React.useState("");

  // Queries
  const [angle, setAngle] = React.useState(0);
  const handleAngleChange = (n: number | null) => setAngle(n ?? 0);

  const handleFileUploadChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files === null || e.target.files.length === 0) {
      openNotification("warning", "Escolha uma imagem.");
      return;
    }
    setFile(e.target.files[0]);
  };

  console.log(angle);

  // Mount events
  React.useEffect(() => {
    if (file === null) {
      return;
    }

    const reader = new FileReader();

    reader.onloadend = () => {
      setPreview(reader.result?.toString()!!);
    };

    reader.readAsDataURL(file);
  }, [file]);

  const applyEffect = async (effect: string) => {
    let processed: ZoomResponse<File>;

    switch (effect) {
      case "negative":
        processed = await ZoomService.applyNegative(file);
        break;
      case "sobel-x":
        processed = await ZoomService.applySobelX(file);
        break;
      case "sobel-y":
        processed = await ZoomService.applySobelY(file);
        break;
      case "sobel-mag":
        processed = await ZoomService.applySobelMag(file);
        break;
      case "gaussian":
        processed = await ZoomService.applyGaussian(file);
        break;
      case "laplacian":
        processed = await ZoomService.applyLaplacian(file);
        break;
      case "mean":
        processed = await ZoomService.applyMean(file);
        break;
      case "bin":
        processed = await ZoomService.applyBin(file);
        break;
      case "gamma":
        processed = await ZoomService.applyGamma(file);
        break;
      case "high-boost":
        processed = await ZoomService.applyHighBoost(file);
        break;
      case "sharpening":
        processed = await ZoomService.applySharpening(file);
        break;
      case "fourier":
        processed = await ZoomService.applyFourier(file);
        break;
      case "scale-bilinear":
        processed = await ZoomService.applyScaleBilinear(file);
        break;
      case "scale-nearest-neighbor":
        processed = await ZoomService.applyScaleNearestNeighbor(file);
        break;
      case "rotate-bilinear":
        processed = await ZoomService.applyRotateBilinear(file, angle);
        break;
      case "rotate-nearest-neighbor":
        processed = await ZoomService.applyRotateNearestNeighbor(file, angle);
        break;

      default:
        openNotification("error", "Opção inválida");
        return;
    }

    if (processed.success) {
      setProcessedFile(processed.data);
      const reader = new FileReader();
      reader.onloadend = () => {
        setProcessedPreview(reader.result?.toString()!!);
      };
      reader.readAsDataURL(processed.data);
    } else {
      openNotification("error", processed.error);
    }
  };

  const invert = () => {
    setFile(processedFile);
    setProcessedFile(file);
    setPreview(processedPreview);
    setProcessedPreview(preview);
  };

  const menuItems = [
    {
      key: "upload",
      icon: <UploadOutlined />,
      label: <UploadFile onChange={handleFileUploadChange} />,
    },
    {
      key: "negative",
      icon: <BgColorsOutlined />,
      label: "Negativo",
      onClick: () => applyEffect("negative"),
    },
    {
      key: "sobel-x",
      icon: <BgColorsOutlined />,
      label: "Sobel X",
      onClick: () => applyEffect("sobel-x"),
    },
    {
      key: "sobel-y",
      icon: <BgColorsOutlined />,
      label: "Sobel Y",
      onClick: () => applyEffect("sobel-y"),
    },
    {
      key: "sobel-mag",
      icon: <BgColorsOutlined />,
      label: "Sobel Mag",
      onClick: () => applyEffect("sobel-mag"),
    },
    {
      key: "gaussian",
      icon: <BgColorsOutlined />,
      label: "Gaussian",
      onClick: () => applyEffect("gaussian"),
    },
    {
      key: "laplacian",
      icon: <BgColorsOutlined />,
      label: "Laplacian",
      onClick: () => applyEffect("laplacian"),
    },
    {
      key: "mean",
      icon: <BgColorsOutlined />,
      label: "Mean",
      onClick: () => applyEffect("mean"),
    },
    {
      key: "bin",
      icon: <BgColorsOutlined />,
      label: "Bin",
      onClick: () => applyEffect("bin"),
    },
    {
      key: "gamma",
      icon: <BgColorsOutlined />,
      label: "Gamma",
      onClick: () => applyEffect("gamma"),
    },
    {
      key: "high-boost",
      icon: <BgColorsOutlined />,
      label: "High Boost",
      onClick: () => applyEffect("high-boost"),
    },
    {
      key: "sharpening",
      icon: <BgColorsOutlined />,
      label: "Sharpening",
      onClick: () => applyEffect("sharpening"),
    },
    {
      key: "fourier",
      icon: <BgColorsOutlined />,
      label: "Fourier",
      onClick: () => applyEffect("fourier"),
    },
    {
      key: "scale-bilinear",
      icon: <ArrowsAltOutlined />,
      label: "Scale Bilinear",
      onClick: () => applyEffect("scale-bilinear"),
    },
    {
      key: "scale-nearest-neighbor",
      icon: <ArrowsAltOutlined />,
      label: "Scale Nearest Neighbor",
      onClick: () => applyEffect("scale-nearest-neighbor"),
    },
    {
      key: "rotate-bilinear",
      icon: <RedoOutlined />,
      label: "Rotate Bilinear",
      onClick: () => applyEffect("rotate-bilinear"),
    },
    {
      key: "rotate-nearest-neighbor",
      icon: <RedoOutlined />,
      label: "Rotate Nearest Neighbor",
      onClick: () => applyEffect("rotate-nearest-neighbor"),
    },
  ];

  return (
    <ConfigProvider
      theme={{
        algorithm: isDarkMode ? darkAlgorithm : defaultAlgorithm,
      }}
    >
      {notificationContext}
      <Layout
        style={{ minHeight: "100vh", maxHeight: "100vh", minWidth: "100vw" }}
      >
        <Header className="flex items-center justify-between">
          <span className="text-white text-lg font-bold">Zoom</span>
          <Flex align="center" justify="center" gap="middle">
            <Button type="primary" onClick={openDrawer}>
              Parâmetros
            </Button>
            <Switch
              value={isDarkMode}
              checkedChildren="Tema claro"
              unCheckedChildren="Tema escuro"
              onChange={toggleTheme}
            />
          </Flex>
        </Header>
        <Drawer open={drawerOpened} onClose={closeDrawer}>
          <Flex vertical gap="middle">
            <Flex gap="middle">
              <Flex gap="middle">
                <Form.Item label="Tamanho do Kernel" name="ks">
                  <InputNumber value={angle} onChange={handleAngleChange} />
                </Form.Item>
              </Flex>
            </Flex>
            <Flex gap="middle">
              <Form.Item label="Ângulo" name="a">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>

              <Form.Item label="Escala em X" name="x">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>

              <Form.Item label="Escala em Y" name="y">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>
            </Flex>

            <Flex gap="middle">
              <Form.Item label="Sigma" name="s">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>

              <Form.Item label="Fator de boost" name="k">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>
            </Flex>

            <Flex gap="middle">
              <Form.Item label="Gamma" name="g">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>

              <Form.Item label="Constante de correção" name="c">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>
            </Flex>

            <Flex gap="middle">
              <Form.Item label="Limiar de binarização" name="t">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>
            </Flex>

            <Flex gap="middle">
              <Form.Item label="Sigma" name="s">
                <InputNumber value={angle} onChange={handleAngleChange} />
              </Form.Item>
            </Flex>
          </Flex>
        </Drawer>
        <Layout>
          <Sider
            width={300}
            collapsed={siderOpened}
            collapsible
            onCollapse={toggleSider}
            breakpoint="lg"
            theme={themeType}
          >
            <Menu theme={themeType} mode="inline" items={menuItems} />
          </Sider>
          <Content>
            <Flex
              align="center"
              justify="center"
              gap="large"
              className="h-full"
            >
              <Image
                src={preview}
                fallback="https://placehold.co/600x400"
                className="object-cover w-full"
              />
              <Button
                size="large"
                type="primary"
                onClick={invert}
                icon={<SwapOutlined />}
              />
              <Image
                src={processedPreview}
                fallback="https://placehold.co/600x400"
                className="object-cover w-full"
              />
            </Flex>
          </Content>
        </Layout>
      </Layout>
    </ConfigProvider>
  );
};

export default Zoom;