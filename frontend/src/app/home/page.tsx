"use client";

import React, { useEffect, useState, useRef } from "react";

const wordsPool = [
  "apple",
  "banana",
  "cat",
  "dog",
  "elephant",
  "fish",
  "grape",
  "hat",
];

type FallingWord = {
  id: number;
  text: string;
  top: number;
  left: number;
};

export default function Home() {
  const [fallingWords, setFallingWords] = useState<FallingWord[]>([]);
  const [input, setInput] = useState("");
  const gameAreaRef = useRef<HTMLDivElement>(null);
  const wordIdRef = useRef(0);

  useEffect(() => {
    const addWordInterval = setInterval(() => {
      if (!gameAreaRef.current) return;
      const gameWidth = gameAreaRef.current.clientWidth;

      const newWord: FallingWord = {
        id: wordIdRef.current++,
        text: wordsPool[Math.floor(Math.random() * wordsPool.length)],
        top: 0,
        left: Math.floor(Math.random() * (gameWidth - 100)),
      };
      setFallingWords((w) => [...w, newWord]);
    }, 2000);

    return () => clearInterval(addWordInterval);
  }, []);

  useEffect(() => {
    const fallInterval = setInterval(() => {
      setFallingWords((words) => {
        if (!gameAreaRef.current) return words;
        const gameHeight = gameAreaRef.current.clientHeight;

        const newWords = words
          .map((word) => ({
            ...word,
            top: word.top + 5,
          }))
          .filter((word) => {
            if (word.top > gameHeight - 30) {
              console.log("fail");
              return false;
            }
            return true;
          });

        return newWords;
      });
    }, 100);

    return () => clearInterval(fallInterval);
  }, []);

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!input) return;

    setFallingWords((words) =>
      words.filter((word) => word.text !== input.trim().toLowerCase())
    );
    setInput("");
  };

  return (
    <div
      ref={gameAreaRef}
      className="min-h-screen min-w-screen border border-black mx-auto mt-5 overflow-hidden bg-gray-100 font-sans"
    >
      {fallingWords.map((word) => (
        <div
          key={word.id}
          style={{
            top: word.top,
            left: word.left,
          }}
          className="absolute text-blue-600 font-bold text-lg select-none pointer-events-none"
        >
          {word.text}
        </div>
      ))}

      <form
        onSubmit={onSubmit}
        className="absolute bottom-0 w-full bg-white px-4 py-2"
      >
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="พิมพ์คำที่กำลังตกลงมาแล้วกด Enter"
          className="w-full p-2 text-base border border-gray-300 rounded"
          autoComplete="off"
        />
      </form>
    </div>
  );
}
