"use client";

import { GameCanvas } from "@/features/game/GameCanvas";
import Link from "next/link";

export default function Page() {
  return (
    <div>
      <Link
        href="/home"
        className="absolute top-10 left-10 hover:underline text-xl bg-white/80 px-4 pt-2 pb-1 rounded-lg"
      >
        Back
      </Link>
      <GameCanvas />
    </div>
  );
}
