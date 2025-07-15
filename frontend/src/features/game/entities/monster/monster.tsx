"use client";

import { useRef, useCallback, useEffect } from "react";
import { Sprite, Container, useTick } from "@pixi/react";
import {
  ANIMATION_SPEED,
  MONSTER_MOVE_SPEED,
  PLAYER_MONSTER_DISTANCE,
} from "@/constants/gameConstants";
import { Texture } from "pixi.js";
import {
  calculateNewMonsterTarget,
  checkCanMove,
  handleMovement,
} from "@/lib/utils/gameUtils";
import { useMonsterAnimation } from "./useMonsterAnimation";
import { Direction, IPosition } from "@/types/gameWorld";

interface IMonsterProps {
  texture: Texture;
  characterPosition: { x: number; y: number };
  monsterPosition: { x: number; y: number };
}

export const Monster = ({
  texture,
  characterPosition,
  monsterPosition,
}: IMonsterProps) => {
  const position = useRef({
    x: monsterPosition.x,
    y: monsterPosition.y,
  });
  const targetPosition = useRef<{ x: number; y: number } | null>(null);
  const currentDirection = useRef<Direction | null>(null);
  // const { getControlsDirection } = useMonsterControls();
  const isMoving = useRef(false);
  // Animation
  const { texture: monsterFrameTexture, updateSprite } = useMonsterAnimation({
    texture,
    frameWidth: 64,
    frameHeight: 64,
    totalFrames: 9,
    animationSpeed: ANIMATION_SPEED,
  });

  const onMove = (x: number, y: number) => {
    position.current.x = x;
    position.current.y = y;
  };

  // Function for set target
  const setNextTarget = useCallback(
    (direction: Direction, charcaterPos: IPosition) => {
      if (targetPosition.current) {
        return;
      }
      const { x, y } = position.current;
      currentDirection.current = direction;
      const newTarget = calculateNewMonsterTarget(x, y, charcaterPos);

      // check player&monster distance
      if (
        Math.abs(newTarget.x - charcaterPos.x) < PLAYER_MONSTER_DISTANCE &&
        Math.abs(newTarget.y - charcaterPos.y) < PLAYER_MONSTER_DISTANCE
      ) {
        return;
      }

      console.log(newTarget, charcaterPos);

      if (checkCanMove(newTarget)) {
        targetPosition.current = newTarget;
      }
    },
    []
  );

  // Listening game loop
  useTick((delta) => {
    const direction = "RIGHT";

    // Check direction
    if (direction) {
      setNextTarget(direction, characterPosition);
    }

    // If has target then move
    if (targetPosition.current) {
      isMoving.current = true;

      // Calculate position to move (with different FPS)
      const { position: newPosition } = handleMovement(
        position.current,
        targetPosition.current,
        MONSTER_MOVE_SPEED,
        delta
      );

      // Move
      position.current = newPosition;
      const { x, y } = position.current;
      onMove(x, y);

      // Set target to default
      targetPosition.current = null;
    } else {
      isMoving.current = false;
    }

    // Update monster animation
    updateSprite(currentDirection.current!, isMoving.current);
  });

  useEffect(() => {
    // call onMove when game started
    onMove(position.current.x, position.current.y);

    // monster direction when game started
    updateSprite("DOWN", false);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Container>
      {monsterFrameTexture && (
        <Sprite
          texture={monsterFrameTexture}
          x={position.current.x}
          y={position.current.y}
          scale={0.5}
          anchor={[0, 0.4]}
        />
      )}
    </Container>
  );
};
