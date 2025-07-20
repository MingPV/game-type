"use client";

import { Container, Graphics, Text } from "@pixi/react";
import { TextStyle, Graphics as PixiGraphics } from "pixi.js";

interface RestartButtonProps {
  position: React.MutableRefObject<{ x: number; y: number }>;
  setIsGameOver: (value: boolean) => void;
  currentHp: React.MutableRefObject<number>;
  characterData: { status: { hp: number } };
}

export const RestartButton: React.FC<RestartButtonProps> = ({
  position,
  setIsGameOver,
  currentHp,
  characterData,
}) => {
  const handleRestart = () => {
    setIsGameOver(false);
    currentHp.current = characterData.status.hp;
    document.dispatchEvent(new CustomEvent("restart-game"));
  };

  const drawButton = (g: PixiGraphics) => {
    g.clear();

    const width = 140;
    const height = 40;
    const radius = 10;

    // Outer border (gold)
    g.beginFill(0xffd700); // gold
    g.drawRoundedRect(0, 0, width, height, radius);
    g.endFill();

    // Inner border (dark inner shadow)
    g.beginFill(0x8b4513); // saddle brown (RPG wooden look)
    g.drawRoundedRect(2, 2, width - 4, height - 4, radius - 2);
    g.endFill();

    // Inner layer (light)
    g.beginFill(0xf0e68c); // khaki (fantasy RPG-style highlight)
    g.drawRoundedRect(4, 4, width - 8, height - 8, radius - 4);
    g.endFill();
  };

  return (
    <Container
      x={position.current.x - 66}
      y={position.current.y - 120}
      interactive
      pointertap={handleRestart}
      cursor="pointer"
    >
      <Graphics draw={drawButton} />
      <Text
        text="Play again"
        x={70}
        y={20}
        anchor={0.5}
        style={
          new TextStyle({
            fill: "#333",
            fontSize: 16,
            fontWeight: "bold",
            fontFamily: "Arial",
          })
        }
      />
    </Container>
  );
};

export default RestartButton;
