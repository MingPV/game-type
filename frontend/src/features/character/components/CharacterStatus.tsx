"use client";

import { Character } from "@/app/types/character";
import React from "react";

type CharacterSelectionProps = {
  character: Character;
};

export default function CharacterStatus({
  character,
}: CharacterSelectionProps) {
  return (
    <div className="w-full h-full bg-stone-700/50 flex flex-col items-center">
      <div className="border-b-1 border-stone-200/20 w-1/2 mt-24 mb-2"></div>
      <div className="text-stone-200/70 mb-2">Status</div>
      <div className="border-b-1 border-stone-200/20 w-1/3 mb-6"></div>
      <div
        className="w-72 h-80 bg-stone-100/20"
        style={{
          clipPath:
            "polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%)",
        }}
      ></div>
      <div className="columns-2 w-1/2 mt-12">
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Health</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">1200</div>
        </div>
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Armor</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">1200</div>
        </div>
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Damage</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">1200</div>
        </div>
        <div className="p-2  flex flex-col items-center gap-1">
          <div className="text-stone-200/70">Critical</div>
          <div className="border-b-1 border-stone-200/20 w-full"></div>
          <div className="text-stone-200">1200</div>
        </div>
      </div>
    </div>
  );
}
