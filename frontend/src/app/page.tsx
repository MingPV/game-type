"use client";

import React, { useEffect } from "react";
import Link from "next/link";

export default function Page() {
  useEffect(() => {}, []);

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
      <div className="flex flex-col justify-center items-center z-10">
        <div className="text-8xl gap-1 text-white mb-10 font-gameFont text-stroke-home flex flex-row">
          <div>T</div>
          <div>Y</div>
          <div>P</div>
          <div>E</div>
          <div>Q</div>
          <div>U</div>
          <div>E</div>
          <div>S</div>
          <div>T</div>
        </div>
        <div className=" flex flex-col gap-4 mb-24 items-center">
          <Link
            href={"/sign-in"}
            className="text-4xl hover:bg-stone-700 px-4 py-2 text-stone-300 transition-all duration-200"
          >
            Play game
          </Link>
          <Link
            href={"/sign-up"}
            className="text-4xl hover:bg-stone-700 px-4 py-2 text-stone-300 transition-all duration-200"
          >
            Create an account
          </Link>
          <Link
            href={"/sign-in"}
            className="text-4xl hover:bg-stone-700 px-4 py-2 text-stone-300 transition-all duration-200"
          >
            Settings
          </Link>
        </div>
      </div>
    </div>
  );
}
