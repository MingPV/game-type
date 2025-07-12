"use client";

import { deleteCharacter } from "@/app/characters/action";
import { Character } from "@/app/types/character";
import Link from "next/link";
import React, { useState } from "react";
import { RiDeleteBin6Line } from "react-icons/ri";

type CharacterDisplayProps = {
  character: Character;
  characters: Character[] | null;
  setCharacters: (characters: Character[]) => void;
};

export default function CharacterDisplay({
  character,
  characters,
  setCharacters,
}: CharacterDisplayProps) {
  const [isDeleteModalOpen, setIsDeleteModelOpen] = useState(false);

  const handleDeleteCharacter = async () => {
    deleteCharacter(character.character_id)
      .then((data) => {
        console.log(data);
        setIsDeleteModelOpen(false);
        if (characters) {
          setCharacters(
            characters.filter((c) => c.character_id !== character.character_id)
          );
        }
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <div className="w-full h-full bg-stone-900/50 flex flex-col justify-start items-center">
      <div className="w-full flex justify-end">
        <span
          className="mr-6 mt-6 cursor-pointer p-1 text-xl text-white/60 hover:text-white/80"
          onClick={() => setIsDeleteModelOpen(true)}
        >
          <RiDeleteBin6Line />
        </span>
      </div>
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
      {isDeleteModalOpen && (
        <div
          className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
          onClick={() => setIsDeleteModelOpen(false)}
        >
          <div
            className="bg-white rounded-2xl shadow-lg p-6 max-w-sm w-full"
            onClick={(e) => e.stopPropagation()}
          >
            <h2 className="text-xl font-semibold mb-4 text-gray-800">
              Confirm Delete
            </h2>
            <p className="text-gray-600 mb-6">
              Are you sure you want to delete this character? This action cannot
              be undone.
            </p>
            <div className="flex justify-end gap-3">
              <button
                onClick={() => setIsDeleteModelOpen(false)}
                className="px-4 py-2 rounded bg-gray-200 text-gray-800 hover:bg-gray-300"
              >
                Cancel
              </button>
              <button
                onClick={handleDeleteCharacter}
                className="px-4 py-2 rounded bg-red-600 text-white hover:bg-red-700"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
