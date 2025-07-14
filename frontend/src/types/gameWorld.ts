export interface IPosition {
  x: number;
  y: number;
}

export type Direction =
  | "UP"
  | "DOWN"
  | "LEFT"
  | "RIGHT"
  | "UP_RIGHT"
  | "UP_LEFT"
  | "DOWN_RIGHT"
  | "DOWN_LEFT"
  | undefined;

export interface Screen {
  width: number;
  height: number;
}
