"use client";

import { Container, Graphics, Text } from "@pixi/react";
import { TextStyle, TextMetrics } from "pixi.js";
import { useMemo } from "react";

export function CommandBox({
  command,
  characterPosition,
}: {
  command: string;
  characterPosition: { x: number; y: number };
}) {
  // paddingX, paddingY
  const paddingX = 0;
  const paddingY = 0;

  // text style
  const style = useMemo(
    () =>
      new TextStyle({
        fill: "white",
        fontSize: 16,
        fontFamily: "monospace",
        fontWeight: "bold",
        align: "center",
        stroke: "#000000",
        strokeThickness: 2,
        dropShadow: true,
        dropShadowColor: "#000000",
        dropShadowBlur: 4,
        dropShadowAngle: Math.PI / 6,
        dropShadowDistance: 3,
      }),
    []
  );

  // use metrix to calculate word size
  const metrics = TextMetrics.measureText(command, style);
  const boxWidth = metrics.width + paddingX * 2 + 50;
  const boxHeight = metrics.height + paddingY * 2;

  return (
    <Container x={characterPosition.x + 15} y={characterPosition.y - 20}>
      <Graphics
        draw={(g) => {
          g.clear();
          g.lineStyle(2, 0xffffff); // border white 2
          g.beginFill(0x000000, 0.7); // bg-black/70
          g.drawRoundedRect(
            -boxWidth / 2,
            -boxHeight / 2,
            boxWidth,
            boxHeight,
            6
          ); // box size along with the word + padding
          g.endFill();
        }}
      />
      <Text text={command} anchor={0.5} style={style} />
    </Container>
  );
}
