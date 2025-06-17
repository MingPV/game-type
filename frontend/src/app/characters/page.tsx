"use client";

import Link from "next/link";
import React, { useEffect } from "react";

export default function Page() {
  useEffect(() => {}, []);

  return (
    <div className="w-full h-full flex flex-col justify-center items-center">
      <div className="w-full py-6 text-stone-200/70 bg-stone-800/10 flex flex-row justify-between items-center">
        <Link
          href="/sign-in"
          className="ml-8 border-b-1 hover:border-stone-200/70 border-stone-800/10"
        >
          Back
        </Link>
        <div>GameName.io</div>
        <div className="mr-8 text-transparent cursor-default">Next</div>
      </div>
      <div className="flex flex-row h-full w-full">
        <div className="w-full h-full bg-stone-600/50 flex flex-col">
          <div className="w-full bg-stone-900/30 px-4 py-4 text-center text-stone-200/50">
            Select a character
          </div>
          <div className="h-full grid grid-cols-3 gap-2 m-2">
            <div className="bg-black/10 border-2 border-stone-200/10 hover:bg-black/20 cursor-pointer"></div>
            <div className="bg-black/10 border-2 border-stone-200/10 hover:bg-black/20 cursor-pointer"></div>
            <div className="bg-black/10 border-2 border-stone-200/10 hover:bg-black/20 cursor-pointer"></div>
            <div className="bg-black/10 border-2 border-stone-200/10 hover:bg-black/20 cursor-pointer"></div>
            <div className="bg-black/10 border-2 border-stone-200/10 hover:bg-black/20 cursor-pointer"></div>
            <div className="bg-black/10 border-2 border-stone-200/10 hover:bg-black/20 cursor-pointer"></div>
          </div>
          <div className="w-full bg-stone-900/30 py-8 text-center text-stone-200/50"></div>
        </div>
        <div className="w-full h-full bg-stone-900/50 flex flex-col justify-end items-center">
          <div className="bg-white/10 h-[40vh] w-[10vw] rounded-full mb-12"></div>
          <div className="text-3xl text-stone-200">MingPV</div>
          <div className="text-stone-200/20 mb-4">swordman</div>
          <div className="text-stone-200/50">Level 24</div>
          <div className="flex flex-row w-[90%] my-3">
            <div className="border border-stone-200 w-1/3"></div>
            <div className="border border-stone-100/10 w-2/3"></div>
          </div>
          <div className="mb-8 text-stone-200/80">398,200 / 1,060,000</div>
          <Link
            href={"/home"}
            className="bg-stone-200/70 px-8 py-2 rounded-lg hover:bg-stone-100/80 transition-all duration-200 mb-16"
          >
            Play
          </Link>
        </div>
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
      </div>
    </div>
  );
}
