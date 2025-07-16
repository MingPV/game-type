"use client";

import { useState, useMemo, PropsWithChildren, useCallback } from "react";
import { Texture } from "pixi.js";
import { Container, Sprite } from "@pixi/react";
import { Character } from "./entities/character/character";
import { Map1 } from "./maps/map1";
import { Camera } from "./Camera";
import backgroundAsset from "@/gameAssets/black.jpg";
import characterAsset from "@/gameAssets/hero.png";
import redHpAsset from "@/gameAssets/redHp50.png";
import { Monster as MonsterEntity } from "./entities/monster/monster";
import { Monster } from "@/types/monster";
import { CommandBox } from "./entities/inputBox/inputBox";
import { useTyping } from "./useTyping";

interface IMainContainerProps {
  canvasSize: { width: number; height: number };
}

export const MainContainer = ({
  canvasSize,
  children,
}: PropsWithChildren<IMainContainerProps>) => {
  const [characterPosition, setCharacterPosition] = useState({ x: 0, y: 0 });
  const updateCharacterPosition = useCallback((x: number, y: number) => {
    setCharacterPosition({
      x: Math.floor(x),
      y: Math.floor(y),
    });
  }, []);

  // Texture
  const characterTexture = useMemo(() => Texture.from(characterAsset.src), []);
  const monsterTexture = useMemo(() => Texture.from(characterAsset.src), []);
  const monsterHPTexture = useMemo(() => Texture.from(redHpAsset.src), []);
  const backgroundTexture = useMemo(
    () => Texture.from(backgroundAsset.src),
    []
  );

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
      new CustomEvent("player-attack", { detail: { word: cmd } })
    );
  });

  return (
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
          monsterPosition={{ x: 150, y: 150 }}
          monsterData={mockMonster}
        />
        <MonsterEntity
          texture={monsterTexture}
          hpTexture={monsterHPTexture}
          characterPosition={characterPosition}
          monsterPosition={{ x: 250, y: 250 }}
          monsterData={mockMonster}
        />
        <MonsterEntity
          texture={monsterTexture}
          hpTexture={monsterHPTexture}
          characterPosition={characterPosition}
          monsterPosition={{ x: 150, y: 350 }}
          monsterData={mockMonster}
        />
        <Character
          texture={characterTexture}
          onMove={updateCharacterPosition}
          isTypingMode={isTypingMode}
        />
        {isTypingMode && (
          <CommandBox command={command} characterPosition={characterPosition} />
        )}
      </Camera>
    </Container>
  );
};

export default MainContainer;
