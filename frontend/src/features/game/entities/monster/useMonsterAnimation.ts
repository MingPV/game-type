import { useState, useRef } from "react";
import { Rectangle, Texture } from "pixi.js";
import { Direction } from "@/types/gameWorld";
import { Monster as MonsterT } from "@/types/monster";

interface UseSpriteAnimationProps {
  texture: Texture;
  hpTexture: Texture;
  monster: MonsterT;
  currentHp: number;
  frameWidth: number;
  frameHeight: number;
  totalFrames: number;
  animationSpeed: number;
}

export const useMonsterAnimation = ({
  texture,
  hpTexture,
  monster,
  currentHp,
  frameWidth,
  frameHeight,
  totalFrames,
  animationSpeed,
}: UseSpriteAnimationProps) => {
  const [currentTexture, setCurrentTexture] = useState<Texture | null>(null);
  const [currrentHPTexture, setCurrentHPTexture] = useState<Texture | null>(
    null
  );
  const frameRef = useRef(0);
  const elapsedTimeRef = useRef(0);

  const getRowByDirection = (direction: Direction | null) => {
    switch (direction) {
      case "UP":
        return 8;
      case "LEFT":
        return 9;
      case "DOWN":
        return 10;
      case "RIGHT":
        return 11;
      default:
        return 10;
    }
  };

  const updateSprite = (direction: Direction | null, isMoving: boolean) => {
    const row = getRowByDirection(direction);
    let column = 0;

    if (isMoving) {
      elapsedTimeRef.current += animationSpeed;

      if (elapsedTimeRef.current >= 1) {
        elapsedTimeRef.current = 0;
        frameRef.current = (frameRef.current + 1) % totalFrames;
      }

      column = frameRef.current;
    }

    const frameTexture = new Texture(
      texture.baseTexture,
      new Rectangle(
        column * frameWidth,
        row * frameHeight,
        frameWidth,
        frameHeight
      )
    );

    const frameTextureHP = new Texture(
      hpTexture.baseTexture,
      new Rectangle(0, 0, 50 * (currentHp / monster.hp), 5)
    );

    // ✅ tell React to re-render
    setCurrentTexture(frameTexture);
    setCurrentHPTexture(frameTextureHP);
  };

  return {
    texture: currentTexture,
    hpTexture: currrentHPTexture,
    updateSprite,
  };
};
