"use client";

import {
  useState,
  useMemo,
  PropsWithChildren,
  useCallback,
  useRef,
  useEffect,
} from "react";
import { TextStyle, Texture } from "pixi.js";
import { Container, Sprite, Text } from "@pixi/react";
import { Character } from "./entities/character/character";
import { Map1 } from "./maps/map1";
import { Camera } from "./Camera";
import backgroundAsset from "@/gameAssets/black.jpg";
// import characterAsset from "@/gameAssets/hero.png";
import characterAsset from "@/gameAssets/swordman/idle_up.png";
import demonBatAsset from "@/gameAssets/demonBat/idle.png";
import redHpAsset from "@/gameAssets/redHp50.png";
import hpAsset from "@/gameAssets/hp50.png";
import { Monster as MonsterEntity } from "./entities/monster/monster";
import { Monster } from "@/types/monster";
import { CommandBox } from "./entities/textBox/inputBox";
import { useTyping } from "./useTyping";
import { Character as CharacterT } from "@/types/character";
import {
  DEFAULT_MONSTER_X_POS_LEFT,
  DEFAULT_MONSTER_Y_POS_LEFT,
  DEFAULT_MONSTER_X_POS_RIGHT,
  DEFAULT_MONSTER_Y_POS_RIGHT,
  DEFAULT_MONSTER_X_POS_CENTER,
  DEFAULT_MONSTER_Y_POS_CENTER,
  DEFAULT_X_POS,
  DEFAULT_Y_POS,
} from "@/constants/gameConstants";

interface IMainContainerProps {
  canvasSize: { width: number; height: number };
}

export const MainContainer = ({
  canvasSize,
  children,
}: PropsWithChildren<IMainContainerProps>) => {
  const [characterPosition, setCharacterPosition] = useState({
    x: DEFAULT_X_POS,
    y: DEFAULT_Y_POS,
  });
  const [point, setPoint] = useState(0);
  // const monsterPositions = useRef<IPosition[] | []>([]);
  const monsterPositions = useRef<boolean[][]>(
    Array.from({ length: 1000 }, () => Array(1000).fill(false))
  );
  const updateCharacterPosition = useCallback((x: number, y: number) => {
    setCharacterPosition({
      x: Math.floor(x),
      y: Math.floor(y),
    });
  }, []);

  const addPoint = useCallback((addPoint: number) => {
    setPoint((prevPoint) => prevPoint + addPoint);
  }, []);

  // Texture
  const characterTexture = useMemo(() => Texture.from(characterAsset.src), []);
  const monsterTexture = useMemo(() => Texture.from(demonBatAsset.src), []);
  const playerHPTexture = useMemo(() => Texture.from(hpAsset.src), []);
  const monsterHPTexture = useMemo(() => Texture.from(redHpAsset.src), []);
  const backgroundTexture = useMemo(
    () => Texture.from(backgroundAsset.src),
    []
  );

  const mockCharacter: CharacterT = {
    character_id: "id",
    user_id: "user_id",
    name: "MingPV",
    level: 1,
    current_exp: 5,
    class_id: "class_id",
    inventory_id: "inventory_id",
    class: {
      name: "Swordman",
      description: "just a swordman",
    },
    status: {
      status_point: 20,
      attack_level: 1,
      defense_level: 1,
      hp_level: 1,
      mp_level: 1,
      critical_level: 1,
      attack: 12,
      defense: 10,
      hp: 100,
      mp: 100,
      critical: 0.01,
    },
    equipment_slots: "",
  };

  // MockData
  const mockMonster: Monster = {
    monster_id: "54d9c43d-d5a0-43b8-b1e7-4315d8d2e99d",
    name: "Noob slime",
    description: "weak slime",
    level: 1,
    hp: 10,
    attack: 1,
    defense: 1,
    exp_reward: 10,
    gold_reward: 5,
    monster_type_id: "96faefde-0794-4f62-be5d-4e7e3bf28855",
    monster_type: {
      monster_type_id: "96faefde-0794-4f62-be5d-4e7e3bf28855",
      name: "Slime",
    },
    monster_loots: [],
  };

  // Handle Typing
  const { isTypingMode, command } = useTyping((cmd) => {
    document.dispatchEvent(
      new CustomEvent("player-attack", {
        detail: { word: cmd, character: mockCharacter },
      })
    );
  });

  // Listening to restart game
  useEffect(() => {
    const onRestart = () => {
      setPoint(0);
    };

    document.addEventListener("restart-game", onRestart);
    return () => document.removeEventListener("restart-game", onRestart);
  });

  return (
    <>
      <Container>
        <Sprite
          texture={backgroundTexture}
          width={canvasSize.width}
          height={canvasSize.height}
        />
        {children}
        <Camera characterPosition={characterPosition} canvasSize={canvasSize}>
          <Map1 />
          <MonsterEntity
            texture={monsterTexture}
            hpTexture={monsterHPTexture}
            characterPosition={characterPosition}
            monsterPosition={{
              x: DEFAULT_MONSTER_X_POS_LEFT,
              y: DEFAULT_MONSTER_Y_POS_LEFT,
            }}
            monsterData={mockMonster}
            monsterPositions={monsterPositions.current}
            addPoint={addPoint}
          />
          <MonsterEntity
            texture={monsterTexture}
            hpTexture={monsterHPTexture}
            characterPosition={characterPosition}
            monsterPosition={{
              x: DEFAULT_MONSTER_X_POS_RIGHT,
              y: DEFAULT_MONSTER_Y_POS_RIGHT,
            }}
            monsterData={mockMonster}
            monsterPositions={monsterPositions.current}
            addPoint={addPoint}
          />
          <MonsterEntity
            texture={monsterTexture}
            hpTexture={monsterHPTexture}
            characterPosition={characterPosition}
            monsterPosition={{
              x: DEFAULT_MONSTER_X_POS_CENTER,
              y: DEFAULT_MONSTER_Y_POS_CENTER,
            }}
            monsterData={mockMonster}
            monsterPositions={monsterPositions.current}
            addPoint={addPoint}
          />
          <Character
            texture={characterTexture}
            hpTexture={playerHPTexture}
            onMove={updateCharacterPosition}
            isTypingMode={isTypingMode}
            characterData={mockCharacter}
          />
          {isTypingMode && (
            <CommandBox
              command={command}
              characterPosition={characterPosition}
            />
          )}
        </Camera>
        <Text
          text={`Point: ${point}`}
          style={
            new TextStyle({
              fill: "#ffffff",
              fontSize: 16,
              fontWeight: "bold",
              fontFamily: "Arial",
            })
          }
          x={canvasSize.width / 2 - 15}
          y={canvasSize.height - 50}
        />
      </Container>
    </>
  );
};

export default MainContainer;
