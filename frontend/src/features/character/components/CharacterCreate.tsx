"use client";

import React, { useState } from "react";
import clsx from "clsx";
import { CharacterClass } from "@/types/characterClass";
import { createCharacter } from "@/app/characters/action";
import { useAuth } from "@/contexts/AuthContext";
import { Character } from "@/types/character";
import Image from "next/image";

type CharacterCreateProps = {
  classes: CharacterClass[];
  characters: Character[] | null;
  setCharacters: (characters: Character[]) => void;
  setCharacterIndex: (index: number) => void;
};

export default function CharacterCreate({
  classes,
  characters,
  setCharacters,
  setCharacterIndex,
}: CharacterCreateProps) {
  const { user } = useAuth();
  const [name, setName] = useState("");
  const [selectedClassName, setSelectedClassName] = useState("");
  const [classID, setClassID] = useState("");
  const [isOpenForm, setIsOpenForm] = useState(false);
  const [error, setError] = useState("");

  const handleCreateCharacter = () => {
    if (user) {
      createCharacter(user.id, name, classID)
        .then((data) => {
          setIsOpenForm(false);
          if (characters) {
            setCharacterIndex(characters.length);
            setCharacters([...characters, data]);
          } else {
            setCharacters(data);
            setCharacterIndex(0);
          }
        })
        .catch((err) => {
          setError("Failed to create character. Please try again");
          console.error(err);
        });
    }
  };

  return (
    <div className="w-full h-full bg-stone-900/50 text-white/80 flex flex-col justify-center items-center">
      <div className="text-xl text-white/70">{`Time to enter the world. Design your hero!`}</div>
      <>
        <div
          className={clsx(
            "overflow-hidden transition-all duration-800 ease-in-out mb-8 flex flex-col",
            isOpenForm ? "max-h-[70vh]" : "max-h-0"
          )}
        >
          <div className="w-full flex flex-col justify-end h-[200px] items-center mt-8 ">
            <Image
              className={`${
                selectedClassName != "Swordman" ? "hidden" : ""
              } h-fit`}
              src={"/images/swordmanCharacter.png"}
              alt="swordmanCharacter"
              width={200}
              height={200}
            />
            <Image
              className={`${
                selectedClassName != "Magician" ? "hidden" : ""
              } h-fit`}
              src={"/images/magicianCharacter.png"}
              alt="swordmanCharacter"
              width={120}
              height={120}
            />
            <Image
              className={`${
                selectedClassName != "Assassin" ? "hidden" : ""
              } h-fit`}
              src={"/images/assassinCharacter.png"}
              alt="swordmanCharacter"
              width={200}
              height={200}
            />
            <Image
              className={`${
                selectedClassName != "Hunter" ? "hidden" : ""
              } h-fit`}
              src={"/images/hunterCharacter.png"}
              alt="swordmanCharacter"
              width={150}
              height={150}
            />
            <Image
              className={`${
                selectedClassName != "Dark Knight" ? "hidden" : ""
              } h-fit`}
              src={"/images/darkKnightCharacter.png"}
              alt="swordmanCharacter"
              width={210}
              height={210}
            />
          </div>
          <div className="text-xl my-1">Name</div>
          <input
            type="text"
            placeholder="Enter your character name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full px-2 py-2 mb-4 text-xl bg-white/10 placeholder:font-mono text-stone-100 text-center placeholder-black/40 placeholder:text-sm focus:ring-1 focus:ring-stone-600 focus:outline-none font-mono font-bold border border-transparent hover:border-stone-60 rounded-md"
          />
          <div className="text-xl my-1">Select class</div>
          <div className="grid grid-cols-3 gap-2">
            {classes.map((c) => (
              <div
                key={c.class_id}
                onClick={() => {
                  setClassID(c.class_id);
                  setSelectedClassName(c.name);
                }}
                className={`p-2 flex justify-center items-center rounded-md cursor-pointer hover:bg-white/40 ${
                  classID == c.class_id ? "bg-white/40" : "bg-white/20"
                }`}
              >
                {c.name}
              </div>
            ))}
          </div>
        </div>
      </>
      {isOpenForm ? (
        <div className="w-full flex flex-col items-center">
          <div
            className="w-fit bg-white/50 px-5 pt-4 pb-3 mb-4 rounded-lg text-xl hover:bg-white/20 hover:text-white/30 cursor-pointer"
            onClick={handleCreateCharacter}
          >
            Create
          </div>
          <div className="text-red-700 mb-16">{error}</div>
        </div>
      ) : (
        <div
          className="bg-white/50 px-5 pt-4 pb-3 rounded-lg text-xl mb-12 hover:bg-white/20 hover:text-white/30 cursor-pointer"
          onClick={() => setIsOpenForm(!isOpenForm)}
        >
          + Create character
        </div>
      )}
    </div>
  );
}
