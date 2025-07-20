"use client";

import { useRef, useCallback, useEffect, useState } from "react";
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
import { Monster as MonsterT } from "@/types/monster";
import { getRandomWordByDifficulty } from "../../words/words";
import { TextWithBackground } from "../textBox/textWithBackground";
import { Character as CharacterT } from "@/types/character";
import {
  DEFAULT_MONSTER_X_POS_CENTER,
  DEFAULT_MONSTER_Y_POS_CENTER,
  DEFAULT_MONSTER_X_POS_LEFT,
  DEFAULT_MONSTER_Y_POS_LEFT,
  DEFAULT_MONSTER_X_POS_RIGHT,
  DEFAULT_MONSTER_Y_POS_RIGHT,
} from "@/constants/gameConstants";

interface IMonsterProps {
  texture: Texture;
  hpTexture: Texture;
  characterPosition: { x: number; y: number };
  monsterPosition: { x: number; y: number };
  monsterData: MonsterT;
  monsterPositions: boolean[][];
  addPoint: (addPoint: number) => void;
}

export const Monster = ({
  texture,
  hpTexture,
  characterPosition,
  monsterPosition,
  monsterData,
  monsterPositions,
  addPoint,
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
    frameWidth: 80,
    frameHeight: 71,
    totalFrames: 4,
    animationSpeed: ANIMATION_SPEED,
  });

  const onMove = (x: number, y: number) => {
    // clear previous position
    // monsterPositions[Math.floor(position.current.x)][
    //   Math.floor(position.current.y)
    // ] = false;
    // change monster position
    position.current.x = x;
    position.current.y = y;
    // set new position
    // monsterPositions[Math.floor(x)][Math.floor(y)] = true;
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
        document.dispatchEvent(new CustomEvent("monster-attack"));
        return;
      }

      // const floorCurrentX = Math.floor(x);
      // const floorCurrentY = Math.floor(x);
      // const floorNextX = Math.floor(newTarget.x);
      // const floorNextY = Math.floor(newTarget.y);

      if (checkCanMove(monsterPositions, position.current, newTarget)) {
        targetPosition.current = newTarget;
      }
    },
    [monsterPositions]
  );

  // Listening game loop
  useTick((delta) => {
    const direction = "RIGHT";

    // Check direction
    if (direction) {
      const adjustCharacterPosition = {
        x: characterPosition.x - 12, // Adjusted for monster position
        y: characterPosition.y,
      };
      setNextTarget(direction, adjustCharacterPosition);
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
      console.log(detail);
      const inputWord = detail?.word;
      const characterData = detail?.character as CharacterT;
      console.log("Received attack word:", inputWord);

      // calculate range to be hit

      const isInRange = true;
      if (isInRange && inputWord == currentWord.word) {
        currentHp.current -= characterData.status.attack;
        console.log(`Monster got hit! Current HP: ${currentHp.current}`);

        setCurrentWord(getRandomWordByDifficulty("easy"));

        if (currentHp.current <= 0) {
          console.log("Monster is dead!");

          addPoint(12);

          // Store the real respawn position
          const spawnPositions = [
            {
              x: DEFAULT_MONSTER_X_POS_CENTER,
              y: DEFAULT_MONSTER_Y_POS_CENTER,
            },
            { x: DEFAULT_MONSTER_X_POS_LEFT, y: DEFAULT_MONSTER_Y_POS_LEFT },
            { x: DEFAULT_MONSTER_X_POS_RIGHT, y: DEFAULT_MONSTER_Y_POS_RIGHT },
          ];
          const randomIndex = Math.floor(Math.random() * spawnPositions.length);
          const respawnPos = spawnPositions[randomIndex];

          // Hide monster at (0,0)
          position.current.x = 0;
          position.current.y = 0;

          setTimeout(() => {
            position.current.x = respawnPos.x;
            position.current.y = respawnPos.y;
            currentHp.current = monsterData.hp;
          }, 2000);
        }
      }
    };

    document.addEventListener("player-attack", onPlayerAttack);
    return () => document.removeEventListener("player-attack", onPlayerAttack);
  }, [characterPosition, currentWord, monsterData.hp, addPoint]);

  // Listening to restart game
  useEffect(() => {
    const onRestart = () => {
      const spawnPositions = [
        {
          x: DEFAULT_MONSTER_X_POS_CENTER,
          y: DEFAULT_MONSTER_Y_POS_CENTER,
        },
        { x: DEFAULT_MONSTER_X_POS_LEFT, y: DEFAULT_MONSTER_Y_POS_LEFT },
        { x: DEFAULT_MONSTER_X_POS_RIGHT, y: DEFAULT_MONSTER_Y_POS_RIGHT },
      ];
      const randomIndex = Math.floor(Math.random() * spawnPositions.length);
      const respawnPos = spawnPositions[randomIndex];

      // Hide monster at (0,0)
      position.current.x = 0;
      position.current.y = 0;

      setTimeout(() => {
        position.current.x = respawnPos.x;
        position.current.y = respawnPos.y;
        currentHp.current = monsterData.hp;
      }, 2000);
    };

    document.addEventListener("restart-game", onRestart);
    return () => document.removeEventListener("restart-game", onRestart);
  });

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
          <TextWithBackground
            text={currentWord.word}
            x={position.current.x + 15}
            y={position.current.y - 30}
          />
        </>
      )}
    </Container>
  );
};
