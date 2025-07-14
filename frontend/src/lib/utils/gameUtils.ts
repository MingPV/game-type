import { Direction, IPosition } from "@/types/gameWorld";

export const calculateCanvasSize = () => {
  const width = window.innerWidth;
  const height = window.innerHeight;
  return { width, height };
};

export const calculateNewTarget = (
  x: number,
  y: number,
  direction: Direction
): IPosition => {
  return {
    x:
      x +
      (direction === "UP_LEFT" || direction === "DOWN_LEFT"
        ? -0.7071
        : direction === "UP_RIGHT" || direction === "DOWN_RIGHT"
        ? 0.7071
        : direction === "LEFT"
        ? -1
        : direction === "RIGHT"
        ? 1
        : 0),
    y:
      y +
      (direction === "UP_LEFT" || direction === "UP_RIGHT"
        ? -0.7071
        : direction === "DOWN_LEFT" || direction === "DOWN_RIGHT"
        ? 0.7071
        : direction === "DOWN"
        ? 1
        : direction === "UP"
        ? -1
        : 0),
  };
};

export const checkCanMove = (target: IPosition) => {
  if (!target) {
    return false;
  }
  return true;
};

export const moveTowards = (
  current: number,
  target: number,
  maxStep: number
) => {
  return (
    current +
    Math.sign(target - current) * Math.min(Math.abs(target - current), maxStep)
  );
};

export const continueMovement = (
  currentPosition: IPosition,
  targetPosition: IPosition,
  step: number
): IPosition => {
  return {
    x: moveTowards(currentPosition.x, targetPosition.x, step),
    y: moveTowards(currentPosition.y, targetPosition.y, step),
  };
};

export const handleMovement = (
  currentPosition: IPosition,
  targetPosition: IPosition,
  moveSpeed: number,
  delta: number
) => {
  const step = moveSpeed * delta;

  return {
    position: continueMovement(currentPosition, targetPosition, step),
  };
};
