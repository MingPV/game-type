"use client";

import { useState, useMemo, PropsWithChildren, useCallback } from "react";
import { Texture } from "pixi.js";
import { Container, Sprite } from "@pixi/react";
import { Character } from "./entities/character/character";
import { Map1 } from "./maps/map1";
import { Camera } from "./Camera";
import backgroundAsset from "@/gameAssets/black.jpg";
import characterAsset from "@/gameAssets/hero.png";

interface IMainContainerProps {
  canvasSize: { width: number; height: number };
}

export const MainContainer = ({
  canvasSize,
  children,
}: PropsWithChildren<IMainContainerProps>) => {
  const [characterPosition, setCharacterPosition] = useState({ x: 0, y: 0 });
  const updateCharacterPosition = useCallback((x: number, y: number) => {
    console.log(x, y);
    setCharacterPosition({
      x: Math.floor(x),
      y: Math.floor(y),
    });
  }, []);

  const characterTexture = useMemo(() => Texture.from(characterAsset.src), []);
  // const coinTextureRed = useMemo(() => Texture.from(coinRedAsset), [])
  // const coinTextureGold = useMemo(() => Texture.from(coinGoldAsset), [])
  const backgroundTexture = useMemo(
    () => Texture.from(backgroundAsset.src),
    []
  );

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
        <Character
          texture={characterTexture}
          onMove={updateCharacterPosition}
        />
        {/* <Coin texture={coinTextureRed} x={5} y={10} />
        <Coin texture={coinTextureGold} x={6} y={11} /> */}
      </Camera>
    </Container>
  );
};

export default MainContainer;
