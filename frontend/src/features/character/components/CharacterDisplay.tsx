import { Character } from "@/app/types/character";
import Link from "next/link";
import React from "react";

type CharacterDisplayProps = {
  character: Character;
};

export default function CharacterDisplay({ character }: CharacterDisplayProps) {
  return (
    <div className="w-full h-full bg-stone-900/50 flex flex-col justify-end items-center">
      <div className="bg-white/10 h-[40vh] w-[10vw] rounded-full mb-12"></div>
      <div className="text-3xl text-stone-200">{character.name}</div>
      <div className="text-stone-200/20 mb-4">{character.class.name}</div>
      <div className="text-stone-200/50">Level {character.level}</div>
      <div className="flex flex-row w-[90%] my-3">
        <div className="border border-stone-200 w-1/3"></div>
        <div className="border border-stone-100/10 w-2/3"></div>
      </div>
      <div className="mb-8 text-stone-200/80">
        {character.current_exp}/ 1,060,000
      </div>
      <Link
        href={"/home"}
        className="bg-stone-200/70 px-8 py-2 rounded-lg hover:bg-stone-100/80 transition-all duration-200 mb-16"
      >
        Play
      </Link>
    </div>
  );
}
