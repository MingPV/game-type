"use client";

import Link from "next/link";
import React, { useState, useEffect } from "react";

import { useAuth } from "@/contexts/AuthContext";
import { fetchCharacters, fetchClasses, signOut } from "./action";
import CharacterDisplay from "@/features/character/components/CharacterDisplay";
import CharacterSelection from "@/features/character/components/CharacterSelection";
import CharacterStatus from "@/features/character/components/CharacterStatus";
import CharacterCreate from "@/features/character/components/CharacterCreate";
import ClassDetail from "@/features/character/components/ClassDetail";
import { Character } from "../../types/character";

export default function Page() {
  const { user } = useAuth();
  const [classes, setClasses] = useState([]);
  const [characterIndex, setCharacterIndex] = useState(-1);
  // const [isLoadingClasses, setIsLoadingClasses] = useState(true);
  const [charcaters, setCharacters] = useState(Array<Character>);
  const [isLoadingCharacter, setIsLoadingCharacter] = useState(true);

  useEffect(() => {
    console.log(user);
    if (user) {
      fetchCharacters(user.id)
        .then((data) => {
          setCharacters(data);
          setIsLoadingCharacter(false);
        })
        .catch((err) => {
          console.error("Error fetching characters:", err);
          setCharacters([]);
        });
    }
  }, [user]);

  useEffect(() => {
    fetchClasses()
      .then((data) => {
        setClasses(data);
        // setIsLoadingClasses(false);
        console.log(data);
      })
      .catch((err) => {
        console.error("Error fetching classes:", err);
        setCharacters([]);
      });
  }, []);

  return (
    <div
      className="w-full h-full flex flex-col justify-center items-center"
      style={{
        backgroundImage: "url('/images/game3.jpg')",
        backgroundSize: "cover",
        backgroundPosition: "center",
        minHeight: "100vh",
      }}
    >
      <div className="absolute inset-0 bg-black opacity-50 z-0" />
      <div className="w-full h-full flex flex-col justify-center items-center z-10">
        <div className="w-full py-6 text-stone-200/70 bg-black/50 flex flex-row justify-between items-center">
          <Link
            href="/sign-in"
            className="ml-8 hover:underline text-xl bg-white/20 px-4 pt-2 pb-1 rounded-lg"
          >
            Back
          </Link>
          <div className="text-xl mt-2">TypeQuest</div>
          <Link
            href={"/"}
            onClick={() => signOut()}
            className="mr-8 px-4 py-2 bg-red-700/70 text-white rounded shadow cursor-pointer hover:underline"
          >
            Logout
          </Link>
        </div>
        <div className="flex flex-row h-full w-full">
          <CharacterSelection
            isLoadingCharacters={isLoadingCharacter}
            characters={charcaters}
            selectedIndex={characterIndex}
            setSelectedIndex={setCharacterIndex}
          />
          {charcaters[characterIndex] ? (
            <CharacterDisplay
              character={charcaters[characterIndex]}
              characters={charcaters}
              setCharacters={setCharacters}
            />
          ) : (
            <CharacterCreate
              classes={classes}
              characters={charcaters}
              setCharacters={setCharacters}
              setCharacterIndex={setCharacterIndex}
            />
          )}

          {charcaters[characterIndex] ? (
            <CharacterStatus character={charcaters[characterIndex]} />
          ) : (
            <ClassDetail classes={classes} />
          )}
        </div>
      </div>
    </div>
  );
}
