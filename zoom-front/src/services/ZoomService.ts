export type ZoomResponse<T> =
  | { success: true; data: T }
  | { success: false; error: string };

export type RGBLChannel = "r" | "g" | "b" | "l";

export type Histogram = {
  r: number[];
  g: number[];
  b: number[];
  l: number[];
};

class ZoomService {
  private constructor() {}

  private static async processImage(
    endpoint: string,
    file: File,
    params: Record<string, string> = {}
  ): Promise<ZoomResponse<File>> {
    try {
      const formData = new FormData();
      formData.append("image", file);

      const queryParams = new URLSearchParams(params).toString();
      const url = `http://localhost:8080${endpoint}?${queryParams}`;

      console.debug("POST", "url");

      const response = await fetch(url, {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("Erro ao processar a imagem.");
      }

      const blob = await response.blob();
      const processedFile = new File([blob], file.name, { type: file.type });

      return { success: true, data: processedFile };
    } catch (error: any) {
      return { success: false, error: error.message };
    }
  }

  private static async uploadImage<T>(
    endpoint: string,
    file: File,
    params: Record<string, string> = {}
  ): Promise<ZoomResponse<T>> {
    try {
      const formData = new FormData();
      formData.append("image", file);

      const queryParams = new URLSearchParams(params).toString();
      const url = `http://localhost:8080${endpoint}?${queryParams}`;

      console.debug("POST", url);

      const response = await fetch(url, {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("Erro ao processar a imagem.");
      }
      const data = await response.json();
      return { success: true, data: data as T };
    } catch (error: any) {
      return { success: false, error: error.message };
    }
  }

  public static async applyNegative(file: File): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/negative", file);
  }

  public static async applySobelX(file: File): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/sobel-x", file);
  }

  public static async applySobelY(file: File): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/sobel-y", file);
  }

  public static async applySobelMag(file: File): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/sobel-mag", file);
  }

  public static async applyGaussian(
    file: File,
    ks: number = 3,
    s: number = 1.0
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/gaussian", file, {
      ks: ks.toString(),
      s: s.toString(),
    });
  }

  public static async applyLaplacian(
    file: File,
    ks: number = 3
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/laplacian", file, {
      ks: ks.toString(),
    });
  }

  public static async applyMean(
    file: File,
    ks: number = 3
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/mean", file, { ks: ks.toString() });
  }

  public static async applyBin(
    file: File,
    t: number = 128
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/bin", file, { t: t.toString() });
  }

  public static async applyGamma(
    file: File,
    g: number = 2.0,
    c: number = 1.0
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/gamma", file, {
      g: g.toString(),
      c: c.toString(),
    });
  }

  public static async applyHighBoost(
    file: File,
    ks: number = 3,
    s: number = 1.0,
    k: number = 1.5
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/high-boost", file, {
      ks: ks.toString(),
      s: s.toString(),
      k: k.toString(),
    });
  }

  public static async applySharpening(
    file: File,
    ks: number = 3
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/sharpening", file, {
      ks: ks.toString(),
    });
  }

  public static async applyFourier(file: File): Promise<ZoomResponse<File>> {
    return this.processImage("/api/effects/fourier", file);
  }

  public static async applyScaleBilinear(
    file: File,
    x: number = 1,
    y: number = 1
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/transform/scale/bilinear", file, {
      x: x.toString(),
      y: y.toString(),
    });
  }

  public static async applyScaleNearestNeighbor(
    file: File,
    x: number = 1,
    y: number = 1
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/transform/scale/nearest-neighbor", file, {
      x: x.toString(),
      y: y.toString(),
    });
  }

  public static async applyRotateBilinear(
    file: File,
    a: number = 0
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/transform/rotate/bilinear", file, {
      a: a.toString(),
    });
  }

  public static async applyRotateNearestNeighbor(
    file: File,
    a: number = 0
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/transform/rotate/nearest-neighbor", file, {
      a: a.toString(),
    });
  }

  public static async getHistogramRGBL(
    file: File
  ): Promise<ZoomResponse<Histogram>> {
    return this.uploadImage("/api/histogram/rgbl", file);
  }

  public static async equalizeHistogram(
    file: File,
    ch: RGBLChannel
  ): Promise<ZoomResponse<File>> {
    return this.processImage("/api/histogram/equalize", file, {
      ch: ch,
    });
  }
}

export default ZoomService;
