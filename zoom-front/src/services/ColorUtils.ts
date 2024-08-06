export class ColorUtils {
  static RGBToHSV(
    r: number,
    g: number,
    b: number
  ): { h: number; s: number; v: number } {
    const rf = r / 255;
    const gf = g / 255;
    const bf = b / 255;

    const max = Math.max(rf, gf, bf);
    const min = Math.min(rf, gf, bf);
    const delta = max - min;

    let h = 0;
    let s = max === 0 ? 0 : delta / max;
    let v = max;

    if (delta !== 0) {
      if (max === rf) {
        h = 60 * (((gf - bf) / delta) % 6);
      } else if (max === gf) {
        h = 60 * ((bf - rf) / delta + 2);
      } else {
        h = 60 * ((rf - gf) / delta + 4);
      }
    }
    if (h < 0) h += 360;

    return { h, s, v };
  }

  static HSVToRGB(
    h: number,
    s: number,
    v: number
  ): { r: number; g: number; b: number } {
    const c = v * s; // Chroma
    const x = c * (1 - Math.abs(((h / 60) % 2) - 1)); // Second largest component
    const m = v - c;

    let r1 = 0,
      g1 = 0,
      b1 = 0;

    if (0 <= h && h < 60) {
      r1 = c;
      g1 = x;
    } else if (60 <= h && h < 120) {
      r1 = x;
      g1 = c;
    } else if (120 <= h && h < 180) {
      g1 = c;
      b1 = x;
    } else if (180 <= h && h < 240) {
      g1 = x;
      b1 = c;
    } else if (240 <= h && h < 300) {
      r1 = x;
      b1 = c;
    } else if (300 <= h && h < 360) {
      r1 = c;
      b1 = x;
    }

    const r = Math.round((r1 + m) * 255);
    const g = Math.round((g1 + m) * 255);
    const b = Math.round((b1 + m) * 255);

    return {
      r: Math.min(Math.max(r, 0), 255),
      g: Math.min(Math.max(g, 0), 255),
      b: Math.min(Math.max(b, 0), 255),
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
    const { r, g, b } = this.HSVToRGB(h, s, v);
    const { nr, ng, nb } = this.negativeRGB(r, g, b);
    const n = this.RGBToHSV(nr, ng, nb);
    return {
      nh: n.h,
      ns: n.s,
      nv: n.v,
    };
  }
}
