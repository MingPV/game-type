"use client";

import { useRef, PropsWithChildren } from "react";
import { Container, useTick } from "@pixi/react";
import { Graphics as PIXIGraphics } from "pixi.js";
import { ZOOM } from "@/constants/gameConstants";

interface ICameraProps {
  characterPosition: { x: number; y: number };
  canvasSize: { width: number; height: number };
}

const lerp = (start: number, end: number) => {
  return start + (end - start) * 0.03;
};

export const Camera = ({
  characterPosition,
  canvasSize,
  children,
}: PropsWithChildren<ICameraProps>) => {
  const containerRef = useRef<PIXIGraphics>(null);

  const cameraPosition = useRef<{ x: number; y: number }>({
    x: canvasSize.width / 2,
    y: canvasSize.height / 2,
  });

  // Listening to game loop
  useTick(() => {
    if (containerRef.current) {
      const targetX = canvasSize.width / 2 - characterPosition.x * ZOOM;
      const targetY = canvasSize.height / 2 - characterPosition.y * ZOOM;

      cameraPosition.current.x = lerp(cameraPosition.current.x, targetX);
      cameraPosition.current.y = lerp(cameraPosition.current.y, targetY);

      containerRef.current.x = cameraPosition.current.x;
      containerRef.current.y = cameraPosition.current.y;
    }
  });

  return (
    <Container ref={containerRef} scale={ZOOM}>
      {children}
    </Container>
  );
};
