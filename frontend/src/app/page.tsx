"use client";

import React, { useEffect } from "react";
import Link from "next/link";

export default function Page() {
  useEffect(() => {}, []);

  return (
    <div className="bg-stone-800 h-[100vh] w-[100vw] flex justify-center items-center">
      <div className="bg-white/20 flex flex-col gap-2">
        <Link href={"/sign-in"}>Play game</Link>
        <Link href={"/sign-up"}>Create an account</Link>
        <Link href={"/sign-in"}>Settings</Link>
      </div>
    </div>
  );
}
