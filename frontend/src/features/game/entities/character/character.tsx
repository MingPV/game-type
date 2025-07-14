"use client";

import { useRef, useCallback, useEffect } from "react";
import { Sprite, Container, useTick } from "@pixi/react";
import {
  ANIMATION_SPEED,
  DEFAULT_X_POS,
  DEFAULT_Y_POS,
  MOVE_SPEED,
} from "@/constants/gameWorld";
import { useCharacterControls } from "./useCharacterControl";
import { Texture } from "pixi.js";
import {
  calculateNewTarget,
  checkCanMove,
  handleMovement,
} from "@/lib/utils/gameUtils";
import { useCharacterAnimation } from "./useCharacterAnimation";
import { Direction } from "@/types/gameWorld";

interface ICharacterProps {
  texture: Texture;
  onMove: (gridX: number, gridY: number) => void;
}

export const Character = ({ texture, onMove }: ICharacterProps) => {
  const position = useRef({ x: DEFAULT_X_POS, y: DEFAULT_Y_POS });
  const targetPosition = useRef<{ x: number; y: number } | null>(null);
  const currentDirection = useRef<Direction | null>(null);
  const { getControlsDirection } = useCharacterControls();
  const isMoving = useRef(false);

  // Animation
  const { sprite, updateSprite } = useCharacterAnimation({
    texture,
    frameWidth: 64,
    frameHeight: 64,
    totalFrames: 9,
    animationSpeed: ANIMATION_SPEED,
  });

  // Handle Move
  useEffect(() => {
    onMove(position.current.x, position.current.y);
  }, [onMove]);

  // Function for set target
  const setNextTarget = useCallback((direction: Direction) => {
    if (targetPosition.current) return;
    const { x, y } = position.current;
    currentDirection.current = direction;
    const newTarget = calculateNewTarget(x, y, direction);

    console.log(newTarget);

    if (checkCanMove(newTarget)) {
      targetPosition.current = newTarget;
    }
  }, []);

  // Listening game loop
  useTick((delta) => {
    const direction = getControlsDirection();

    // Check direction if no direction just return
    if (direction) {
      isMoving.current = true;
      setNextTarget(direction);
    } else {
      isMoving.current = false;
      return;
    }

    // If has target then move
    if (targetPosition.current) {
      // Calculate position to move (with different FPS)
      const { position: newPosition } = handleMovement(
        position.current,
        targetPosition.current,
        MOVE_SPEED,
        delta
      );

      // Move
      position.current = newPosition;
      const { x, y } = position.current;
      onMove(x, y);

      // Set target to default
      targetPosition.current = null;
    }

    // Update character animation
    updateSprite(currentDirection.current!, isMoving.current);
  });

  useEffect(() => {
    // call onMove when game started
    onMove(position.current.x, position.current.y);

    // character direction when game started
    updateSprite("DOWN", false);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Container>
      {sprite && (
        <Sprite
          texture={sprite.texture}
          x={position.current.x}
          y={position.current.y}
          scale={0.5}
          anchor={[0, 0.4]}
        />
      )}
    </Container>
  );
};
