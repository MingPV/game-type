"use client";

import { useRef, useCallback, useEffect, useState } from "react";
import { Sprite, Container, useTick } from "@pixi/react";
import {
  ANIMATION_SPEED,
  DEFAULT_X_POS,
  DEFAULT_Y_POS,
  MOVE_SPEED,
} from "@/constants/gameConstants";
// import { useCharacterControls } from "./useCharacterControl";
import { Texture } from "pixi.js";
import { calculateNewTarget, handleMovement } from "@/lib/utils/gameUtils";
import { useCharacterAnimation } from "./useCharacterAnimation";
import { Direction } from "@/types/gameWorld";
import { Character as CharacterT } from "@/types/character";
import RestartButton from "../../RestartButton";

interface ICharacterProps {
  texture: Texture;
  hpTexture: Texture;
  onMove: (gridX: number, gridY: number) => void;
  isTypingMode: boolean;
  characterData: CharacterT;
}

export const Character = ({
  texture,
  hpTexture,
  onMove,
  isTypingMode,
  characterData,
}: ICharacterProps) => {
  const character = useRef(characterData);
  const position = useRef({ x: DEFAULT_X_POS, y: DEFAULT_Y_POS });
  const targetPosition = useRef<{ x: number; y: number } | null>(null);
  const currentDirection = useRef<Direction | null>(null);
  //   const { getControlsDirection } = useCharacterControls();
  const isMoving = useRef(false);
  const currentHp = useRef<number>(characterData.status.hp);
  const [isGameOver, setIsGameOver] = useState(false);

  // Animation
  const {
    sprite,
    hpTexture: monsterHpFrameTexture,
    updateSprite,
  } = useCharacterAnimation({
    texture,
    hpTexture,
    character: character.current,
    currentHp: currentHp.current,
    frameWidth: 96, // 64
    frameHeight: 80, // 64
    totalFrames: 8, // 9
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

    // if (checkCanMove(newTarget)) {
    //   targetPosition.current = newTarget;
    // }
    if (true) {
      targetPosition.current = newTarget;
    }
  }, []);

  // Listening game loop
  useTick((delta) => {
    // const direction = getControlsDirection();

    const direction = null;

    if (direction && !isTypingMode) {
      isMoving.current = true;
      setNextTarget(direction);
    } else {
      isMoving.current = false;
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
    updateSprite();
  });

  // Listening to monster attack
  useEffect(() => {
    const onMonsterAttack = () => {
      if (isGameOver) {
        return;
      }
      currentHp.current -= 1;

      if (currentHp.current <= 0) {
        console.log("Game over!");
        setIsGameOver(true);
        // alert("Game over! You have been defeated by the monster.");
      }
    };

    document.addEventListener("monster-attack", onMonsterAttack);
    return () =>
      document.removeEventListener("monster-attack", onMonsterAttack);
  }, [isGameOver]);

  useEffect(() => {
    // call onMove when game started
    onMove(position.current.x, position.current.y);

    // character direction when game started
    updateSprite();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Container>
      {sprite && monsterHpFrameTexture && (
        <>
          <Sprite
            texture={sprite.texture}
            x={position.current.x}
            y={position.current.y}
            scale={0.5}
            anchor={[0, 0.4]}
          />
          <Sprite
            texture={monsterHpFrameTexture}
            x={position.current.x - 8}
            y={position.current.y - 10}
            scale={0.5}
            anchor={[0, 0.4]}
          />
          {isGameOver && (
            <RestartButton
              position={position}
              currentHp={currentHp}
              characterData={characterData}
              setIsGameOver={setIsGameOver}
            />
          )}
        </>
      )}
    </Container>
  );
};
