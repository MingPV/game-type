"use client";

import { Character } from "@/app/types/character";
import React from "react";
import Image from "next/image";

type CharacterSelectionProps = {
  isLoadingCharacters: boolean;
  characters: Character[];
  selectedIndex: number | null;
  setSelectedIndex: (index: number) => void;
};

export default function CharacterSelection({
  isLoadingCharacters,
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
            className={`bg-black/50 border-2 border-stone-200/10  cursor-pointer flex flex-col justify-center items-center text-white/20 ${
              selectedIndex === i
                ? "ring-1 ring-stone-200/50"
                : "hover:bg-black/20"
            }`}
            onClick={() => setSelectedIndex(i)}
          >
            <div className="flex justify-center items-center">
              {c.class.name == "Swordman" ? (
                <Image
                  className={`h-fit ${
                    selectedIndex === i ? "opacity-100" : "opacity-40"
                  }`}
                  src={"/images/swordmanCharacter.png"}
                  alt="swordmanCharacter"
                  width={150}
                  height={150}
                />
              ) : c.class.name == "Magician" ? (
                <Image
                  className={`h-fit ${
                    selectedIndex === i ? "opacity-100" : "opacity-40"
                  }`}
                  src={"/images/magicianCharacter.png"}
                  alt="magicianCharacter"
                  width={100}
                  height={100}
                />
              ) : c.class.name == "Assassin" ? (
                <Image
                  className={`h-fit ${
                    selectedIndex === i ? "opacity-100" : "opacity-40"
                  }`}
                  src={"/images/assassinCharacter.png"}
                  alt="assassinCharacter"
                  width={150}
                  height={150}
                />
              ) : c.class.name == "Hunter" ? (
                <Image
                  className={`h-fit ${
                    selectedIndex === i ? "opacity-100" : "opacity-40"
                  }`}
                  src={"/images/hunterCharacter.png"}
                  alt="hunterCharacter"
                  width={120}
                  height={120}
                />
              ) : c.class.name == "Dark Knight" ? (
                <Image
                  className={`h-fit ${
                    selectedIndex === i ? "opacity-100" : "opacity-40"
                  }`}
                  src={"/images/darkKnightCharacter.png"}
                  alt="darkKnightCharacter"
                  width={150}
                  height={150}
                />
              ) : null}
            </div>

            <span
              className={`${
                selectedIndex === i ? "text-white/70" : "text-white/40"
              }`}
            >
              {characters[i].name}
            </span>
            <span
              className={`${
                selectedIndex === i ? "text-white/90" : "text-white/50"
              }`}
            >
              Level {characters[i].level}
            </span>
          </div>
        ))}
        {Array.from({ length: 6 - characters.length }, (_, i) => i).map(
          (c, i) => (
            <div
              key={i}
              className={`bg-black/50 border-2 border-stone-200/10  cursor-pointer flex flex-col justify-center items-center text-white/20 ${
                selectedIndex === characters.length + i
                  ? "ring-1 ring-stone-200/50"
                  : "hover:bg-black/20"
              }`}
              onClick={() => setSelectedIndex(characters.length + i)}
            >
              <Image
                className={`h-fit opacity-0`}
                src={"/images/swordmanCharacter.png"}
                alt="swordmanCharacter"
                width={150}
                height={150}
              />
              <div>{isLoadingCharacters ? "loading" : "empty"}</div>
            </div>
          )
        )}
      </div>
      <div className="w-full bg-stone-900/50 py-8 text-center text-stone-200/50"></div>
    </div>
  );
}
