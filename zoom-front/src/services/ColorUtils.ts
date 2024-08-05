export class ColorUtils {
  static RGBToHSV(
    r: number,
    g: number,
    b: number
  ): { h: number; s: number; v: number } {
    r /= 255;
    g /= 255;
    b /= 255;

    const max = Math.max(r, g, b);
    const min = Math.min(r, g, b);
    const delta = max - min;

    let h = 0;
    let s = max === 0 ? 0 : delta / max;
    let v = max;

    if (delta !== 0) {
      if (max === r) {
        h = (g - b) / delta;
      } else if (max === g) {
        h = (b - r) / delta + 2;
      } else {
        h = (r - g) / delta + 4;
      }
      h *= 60;
      if (h < 0) h += 360;
    }

    return { h, s, v };
  }

  static HSVToRGB(
    h: number,
    s: number,
    v: number
  ): { r: number; g: number; b: number } {
    const c = v * s;
    const x = c * (1 - Math.abs(((h / 60) % 2) - 1));
    const m = v - c;

    let r = 0,
      g = 0,
      b = 0;
    if (0 <= h && h < 60) {
      r = c;
      g = x;
      b = 0;
    } else if (60 <= h && h < 120) {
      r = x;
      g = c;
      b = 0;
    } else if (120 <= h && h < 180) {
      r = 0;
      g = c;
      b = x;
    } else if (180 <= h && h < 240) {
      r = 0;
      g = x;
      b = c;
    } else if (240 <= h && h < 300) {
      r = x;
      g = 0;
      b = c;
    } else if (300 <= h && h < 360) {
      r = c;
      g = 0;
      b = x;
    }

    return {
      r: Math.round((r + m) * 255),
      g: Math.round((g + m) * 255),
      b: Math.round((b + m) * 255),
    };
  }

  static toGrayAverage(r: number, g: number, b: number): number {
    return Math.round((r + g + b) / 3);
  }

  static toGrayWeighted(r: number, g: number, b: number): number {
    return Math.round(0.299 * r + 0.587 * g + 0.114 * b);
  }

  static toGrayHSV(h: number, s: number, v: number): number {
    const { r, g, b } = this.HSVToRGB(h, s, v);
    return this.toGrayWeighted(r, g, b);
  }

  static negativeRGB(
    r: number,
    g: number,
    b: number
  ): { nr: number; ng: number; nb: number } {
    return { nr: 255 - r, ng: 255 - g, nb: 255 - b };
  }

  static negativeHSV(
    h: number,
    s: number,
    v: number
  ): { nh: number; ns: number; nv: number } {
    return { nh: (h + 180) % 360, ns: 1 - s, nv: 1 - v };
  }
}
