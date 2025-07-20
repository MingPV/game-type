"use client";
import { useEffect, useState, useCallback } from "react";
import { Stage } from "@pixi/react";
import MainContainer from "./MainContainer";
import { Screen } from "@/types/gameWorld";
import { calculateCanvasSize } from "@/lib/utils/gameUtils";

export const GameCanvas = () => {
  const [canvasSize, setCanvasSize] = useState<Screen>({ width: 0, height: 0 });

  // Update canvas
  const updateCanvasSize = useCallback(() => {
    setCanvasSize(calculateCanvasSize());
  }, []);

  useEffect(() => {
    setCanvasSize(calculateCanvasSize());

    window.addEventListener("resize", updateCanvasSize);
    return () => window.removeEventListener("resize", updateCanvasSize);
  }, [updateCanvasSize]);

  return (
    <Stage width={canvasSize.width} height={canvasSize.height}>
      <MainContainer canvasSize={canvasSize} />
    </Stage>
  );
};
