"use client";

import { useRef } from "react";
import { TILE_SIZE } from "@/constants/gameWorld";
import { Rectangle, Sprite, Texture } from "pixi.js";
import { Direction } from "@/types/gameWorld";

interface UseSpriteAnimationProps {
  texture: Texture;
  frameWidth: number;
  frameHeight: number;
  totalFrames: number;
  animationSpeed: number;
}

export const useCharacterAnimation = ({
  texture,
  frameWidth,
  frameHeight,
  totalFrames,
  animationSpeed,
}: UseSpriteAnimationProps) => {
  const spriteRef = useRef<Sprite | null>(null);
  const frameRef = useRef(0);
  const elapsedTimeRef = useRef(0);

  // Get row index for image
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

    if (!spriteRef.current) {
      const newSprite = new Sprite(frameTexture);
      newSprite.width = TILE_SIZE;
      newSprite.height = TILE_SIZE;
      spriteRef.current = newSprite;
    } else {
      spriteRef.current.texture = frameTexture;
    }
  };

  return { sprite: spriteRef.current, updateSprite };
};
