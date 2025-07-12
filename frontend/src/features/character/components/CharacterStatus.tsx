"use client";

import { Character } from "@/app/types/character";
import React from "react";

type CharacterSelectionProps = {
  character: Character;
};

export default function CharacterStatus({
  character,
}: CharacterSelectionProps) {
  function getInnerPentagonPoints() {
    const centerX = 50;
    const centerY = 50;

    const points = [
      [50, 0, character.status.attack_level / 200 + 0.5], // top (atk)
      [100, 38, character.status.defense_level / 200 + 0.5], // right (def)
      [82, 100, character.status.critical_level / 200 + 0.5], // bottom-right (crit)
      [18, 100, character.status.mp_level / 200 + 0.5], // bottom-left (mp)
      [0, 38, character.status.hp_level / 200 + 0.5], // left (hp)
    ];

    // คำนวณแต่ละจุดให้เข้าใกล้ center ตามค่าสถานะ
    return points
      .map(([x, y, v]) => {
        const dx = x - centerX;
        const dy = y - centerY;
        const scaledX = centerX + dx * v;
        const scaledY = centerY + dy * v;
        return `${scaledX}% ${scaledY}%`;
      })
      .join(", ");
  }

  return (
    <div className="w-full h-full bg-stone-700/50 flex flex-col items-center">
      <div className="border-b-1 border-stone-200/20 w-1/2 mt-24 mb-2"></div>
      <div className="text-stone-200/70 mb-2">Status</div>
      <div className="border-b-1 border-stone-200/20 w-1/3 mb-6"></div>
      <div className="relative w-80 h-72">
        {/* Pentagon Shape */}
        <div
          className="w-full h-full bg-stone-100/20"
          style={{
            clipPath: "polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%)",
          }}
        ></div>

        {/* Character status */}
        <div className="absolute inset-0">
          <div
            className="w-full h-full bg-white/30"
            style={{
              clipPath: `polygon(${getInnerPentagonPoints()})`,
            }}
          ></div>
        </div>

        {/* Labels outside the corners */}
        <span className="absolute left-1/2 top-0 -translate-x-1/2 -translate-y-full text-red-600/80">
          atk
        </span>
        <span className="absolute right-0 top-[38%] translate-x-full -translate-y-1/2 text-cyan-500/80">
          def
        </span>
        <span className="absolute right-[18%] bottom-0 translate-x-1/2 translate-y-full text-amber-600/80">
          crit
        </span>
        <span className="absolute left-[18%] bottom-0 -translate-x-1/2 translate-y-full text-sky-300">
          mp
        </span>
        <span className="absolute left-0 top-[38%] -translate-x-full -translate-y-1/2 text-lime-400">
          hp
        </span>
      </div>

      <div className="columns-2 w-1/2 mt-12">
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Health</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">{character.status.attack}</div>
        </div>
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Defense</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">{character.status.defense}</div>
        </div>
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Damage</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">{character.status.attack}</div>
        </div>
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Mana</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">{character.status.mp}</div>
        </div>
      </div>
    </div>
  );
}
