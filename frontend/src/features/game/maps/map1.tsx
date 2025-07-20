"use client";

import {
  GAME_HEIGHT,
  GAME_WIDTH,
  OFFSET_X,
  OFFSET_Y,
} from "@/constants/gameConstants";
import { Sprite } from "@pixi/react";
import mapAsset from "@/gameAssets/dungeonMap.png";

export const Map1 = () => {
  return (
    <>
      <Sprite
        image={mapAsset.src}
        width={GAME_WIDTH}
        height={GAME_HEIGHT + OFFSET_Y}
        scale={0.38}
        x={OFFSET_X}
        y={OFFSET_Y}
      />
    </>
  );
};
