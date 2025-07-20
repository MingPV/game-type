"use client";

import { useRef, useState } from "react";
import { Rectangle, Sprite, Texture } from "pixi.js";
import { Character as CharacterT } from "@/types/character";

interface UseSpriteAnimationProps {
  texture: Texture;
  hpTexture: Texture;
  character: CharacterT;
  currentHp: number;
  frameWidth: number;
  frameHeight: number;
  totalFrames: number;
  animationSpeed: number;
}

export const useCharacterAnimation = ({
  texture,
  hpTexture,
  character,
  currentHp,
  frameWidth,
  frameHeight,
  totalFrames,
  animationSpeed,
}: UseSpriteAnimationProps) => {
  const spriteRef = useRef<Sprite | null>(null);
  const frameRef = useRef(0);
  const [currrentHPTexture, setCurrentHPTexture] = useState<Texture | null>(
    null
  );
  const elapsedTimeRef = useRef(0);

  // Get row index for image
  //   const getRowByDirection = (direction: Direction | null) => {
  //     switch (direction) {
  //       case "UP":
  //         return 8;
  //       case "LEFT":
  //         return 9;
  //       case "DOWN":
  //         return 10;
  //       case "RIGHT":
  //         return 11;
  //       default:
  //         return 10;
  //     }
  //   };

  // Update image
  const updateSprite = () => {
    // const row = getRowByDirection(direction);
    const row = 0;
    let column = 0;

    // if (isMoving) {
    elapsedTimeRef.current += animationSpeed;

    if (elapsedTimeRef.current >= 1) {
      elapsedTimeRef.current = 0;
      frameRef.current = (frameRef.current + 1) % totalFrames;
    }

    column = frameRef.current;
    // }

    const frameTexture = new Texture(
      texture.baseTexture,
      new Rectangle(
        column * frameWidth + 38,
        row * frameHeight,
        20,
        frameHeight
      )
    );

    if (!spriteRef.current) {
      const newSprite = new Sprite(frameTexture);
      newSprite.width = 20;
      newSprite.height = 34;
      spriteRef.current = newSprite;
    } else {
      spriteRef.current.texture = frameTexture;
    }

    const frameTextureHP = new Texture(
      hpTexture.baseTexture,
      new Rectangle(0, 0, 50 * (currentHp / character.status.hp), 5)
    );

    setCurrentHPTexture(frameTextureHP);
  };

  return {
    sprite: spriteRef.current,
    hpTexture: currrentHPTexture,
    updateSprite,
  };
};
