"use client";

import { Graphics, Text } from "@pixi/react";
import { TextStyle } from "pixi.js";

export const TextWithBackground = ({
  text,
  x,
  y,
}: {
  text: string;
  x: number;
  y: number;
}) => {
  const padding = 6;
  const fontSize = 16;

  const textStyle = new TextStyle({
    fill: "white",
    fontSize,
    fontWeight: "bold",
    align: "center",
    fontFamily: "Arial",
  });

  const textWidth = text.length * (fontSize * 0.6); // approx width per character
  const textHeight = fontSize + 4;

  return (
    <>
      <Graphics
        draw={(g) => {
          g.clear();
          g.beginFill(0x000000, 0.8); // black background, 60% opacity
          g.drawRoundedRect(
            x - textWidth / 2 - padding,
            y - textHeight / 2 - padding,
            textWidth + padding * 2,
            textHeight + padding * 2,
            4
          );
          g.endFill();
        }}
      />
      <Text text={text} x={x} y={y} anchor={0.5} style={textStyle} />
    </>
  );
};
