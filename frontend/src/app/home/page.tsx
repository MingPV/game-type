"use client";

import Link from "next/link";
import React from "react";

export default function page() {
  return (
    <div
      className="w-full h-full flex flex-col justify-center items-center"
      style={{
        backgroundImage: "url('/images/gameBackground1.jpg')",
        backgroundSize: "cover",
        backgroundPosition: "center",
        minHeight: "100vh",
      }}
    >
      <div className="absolute inset-0 bg-black opacity-30 z-0" />
      <div className="w-full h-full flex justify-center items-center z-10">
        <Link
          href={"/"}
          className="absolute top-2 right-2 px-4 py-2 bg-red-700/70 text-white rounded shadow cursor-pointer"
        >
          Logout
        </Link>
        <div className="flex flex-col gap-4 items-center">
          <div className="w-32 h-56 bg-white/50 rounded-full"></div>
          <div className="text-white/80">MingPV Lv.14</div>
          <Link
            className="bg-white/50 py-2 px-4 cursor-pointer rounded-xl"
            href={"/game"}
          >
            Play
          </Link>
        </div>
      </div>
    </div>
  );
}
