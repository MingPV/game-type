"use client";

import { useRef, useCallback, useEffect, useState } from "react";
import { Sprite, Text, Container, useTick } from "@pixi/react";

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
import { Monster as MonsterT } from "@/types/monster";
import { getRandomWordByDifficulty } from "../../words/words";
import { TextStyle } from "pixi.js";

interface IMonsterProps {
  texture: Texture;
  hpTexture: Texture;
  characterPosition: { x: number; y: number };
  monsterPosition: { x: number; y: number };
  monsterData: MonsterT;
}

export const Monster = ({
  texture,
  hpTexture,
  characterPosition,
  monsterPosition,
  monsterData,
}: IMonsterProps) => {
  const position = useRef({
    x: monsterPosition.x,
    y: monsterPosition.y,
  });
  const targetPosition = useRef<{ x: number; y: number } | null>(null);
  const currentDirection = useRef<Direction | null>(null);
  const [currentWord, setCurrentWord] = useState(
    getRandomWordByDifficulty("easy")
  );
  const isMoving = useRef(false);
  const currentHp = useRef<number>(monsterData.hp);

  // Animation
  const {
    texture: monsterFrameTexture,
    hpTexture: monsterHpFrameTexture,
    updateSprite,
  } = useMonsterAnimation({
    texture,
    hpTexture,
    monster: monsterData,
    currentHp: currentHp.current,
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

  // Listening to player attack
  useEffect(() => {
    const onPlayerAttack = (e: Event) => {
      const customEvent = e as CustomEvent;
      const detail = customEvent.detail;
      const inputWord = detail?.word;
      console.log("Received attack word:", inputWord);

      // calculate range to be hit

      const isInRange = true;
      if (isInRange && inputWord == currentWord.word) {
        currentHp.current -= 3;
        console.log(`Monster got hit! Current HP: ${currentHp.current}`);

        setCurrentWord(getRandomWordByDifficulty("easy"));

        if (currentHp.current <= 0) {
          console.log("Monster is dead!");
          // Respawn monster at spawn point
          position.current.x = 70;
          position.current.y = 30;
          currentHp.current = monsterData.hp;
        }
      }
    };

    document.addEventListener("player-attack", onPlayerAttack);
    return () => document.removeEventListener("player-attack", onPlayerAttack);
  }, [characterPosition, currentWord, monsterData.hp]);

  return (
    <Container>
      {monsterFrameTexture && monsterHpFrameTexture && (
        <>
          <Sprite
            texture={monsterFrameTexture}
            x={position.current.x}
            y={position.current.y}
            scale={0.5}
            anchor={[0, 0.4]}
          />
          <Sprite
            texture={monsterHpFrameTexture}
            x={position.current.x + 3}
            y={position.current.y - 10}
            scale={0.5}
            anchor={[0, 0.4]}
          />
          <Text
            text={currentWord.word}
            x={position.current.x}
            y={position.current.y - 30}
            style={
              new TextStyle({
                fill: "white",
                fontSize: 16,
                fontWeight: "bold",
                align: "center",
                fontFamily: "Arial",
              })
            }
          />
        </>
      )}
    </Container>
  );
};
