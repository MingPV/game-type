"use client";

import {
  GAME_HEIGHT,
  GAME_WIDTH,
  OFFSET_X,
  OFFSET_Y,
} from "@/constants/gameWorld";
import { Sprite } from "@pixi/react";
import mapAsset from "@/gameAssets/tilemap.png";

export const Map1 = () => {
  return (
    <>
      <Sprite
        image={mapAsset.src}
        width={GAME_WIDTH}
        height={GAME_HEIGHT + OFFSET_Y}
        scale={1}
        x={OFFSET_X}
        y={OFFSET_Y}
      />
    </>
  );
};
