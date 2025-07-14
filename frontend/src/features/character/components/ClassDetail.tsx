"use client";

import React from "react";
import { CharacterClass } from "@/types/characterClass";

type ClassDetailProps = {
  classes: CharacterClass[];
};

export default function ClassDetail({ classes }: ClassDetailProps) {
  return (
    <div className="w-full h-full bg-stone-700/50 flex flex-col gap-8">
      <div className="text-white/70 text-xl flex w-full justify-center mt-6">
        Choose your path and begin your journey.
      </div>
      {classes.map((c, index) => (
        <div className="mr-8" key={index}>
          <div className="ml-6 text-white/80">{c.name}</div>
          <div className="ml-8 text-white/50">- {c.description}</div>
        </div>
      ))}
    </div>
  );
}
