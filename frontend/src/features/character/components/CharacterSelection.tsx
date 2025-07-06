"use client";

import { Character } from "@/app/types/character";
import React from "react";

type CharacterSelectionProps = {
  characters: Character[];
  selectedIndex: number | null;
  setSelectedIndex: (index: number) => void;
};

export default function CharacterSelection({
  characters,
  selectedIndex,
  setSelectedIndex,
}: CharacterSelectionProps) {
  return (
    <div className="w-full h-full bg-stone-700/50 flex flex-col">
      <div className="w-full bg-stone-900/50 px-4 py-4 text-center text-stone-200/50">
        Select a character
      </div>
      <div className="h-full grid grid-cols-3 gap-2 m-2">
        {characters.map((c, i) => (
          <div
            key={i}
            className={`bg-black/50 border-2 border-stone-200/10  cursor-pointer flex justify-center items-center text-white/20 ${
              selectedIndex === i
                ? "ring-1 ring-stone-200/50"
                : "hover:bg-black/20"
            }`}
            onClick={() => setSelectedIndex(i)}
          >
            {characters[i] ? `${characters[i].name}` : "empty"}
          </div>
        ))}
        {Array.from({ length: 6 - characters.length }, (_, i) => i).map(
          (c, i) => (
            <div
              key={i}
              className={`bg-black/50 border-2 border-stone-200/10  cursor-pointer flex justify-center items-center text-white/20 ${
                selectedIndex === characters.length + i
                  ? "ring-1 ring-stone-200/50"
                  : "hover:bg-black/20"
              }`}
              onClick={() => setSelectedIndex(characters.length + i)}
            >
              empty
            </div>
          )
        )}
      </div>
      <div className="w-full bg-stone-900/50 py-8 text-center text-stone-200/50"></div>
    </div>
  );
}
