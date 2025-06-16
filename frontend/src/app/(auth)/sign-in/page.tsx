"use client";

import React, { useEffect, useState } from "react";
import { validateEmail } from "@/lib/utils/validateEmail";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { IoMdArrowRoundBack } from "react-icons/io";

export default function Page() {
  const [emailOrUsername, setEmailOrUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isSigningIn, setIsSigningIn] = useState(false);
  const [error, setError] = useState("");

  const router = useRouter();

  useEffect(() => {}, []);

  const handleSignIn = async () => {
    setError("");
    setIsSigningIn(true);

    let isSignInWithEmail = true;

    if (!validateEmail(emailOrUsername)) {
      isSignInWithEmail = false;
    }

    try {
      let res;

      if (isSignInWithEmail) {
        res = await fetch(
          `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/auth/signin`,
          {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              email: emailOrUsername,
              password: password,
            }),
            credentials: "include",
          }
        );
      } else {
        res = await fetch(
          `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/auth/signin/username`,
          {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              username: emailOrUsername,
              password: password,
            }),
            credentials: "include",
          }
        );
      }

      const data = await res.json();

      if (!res.ok) {
        setError(data.message || "Sign-in failed. Please try again.");
      } else {
        console.log("Sign-in success", data);
      }
    } catch (err) {
      console.error(err);
      setError("Something went wrong. Please try again later.");
    } finally {
      setIsSigningIn(false);
      setEmailOrUsername("");
      setPassword("");
    }

    router.push("/characters");
  };

  return (
    <div className="flex flex-col justify-center items-center h-full w-full">
      <div className="text-8xl font-bold text-stone-400 mb-10">GameName.io</div>
      <div className="flex flex-row gap-2 w-[40vw]">
        <Link
          className="mr-2 flex flex-row gap-2 items-center h-fit mt-2 text-stone-300/80 w-fit font-bold bg-white/20 hover:bg-white/10 transition-all duration-200 py-2 px-4 rounded-md"
          href={"/"}
        >
          <span className="font-bold">
            <IoMdArrowRoundBack />
          </span>
          Back{" "}
        </Link>
        <div className="bg-white/20 p-10 flex flex-col items-center w-[25vw]">
          <input
            type="text"
            placeholder="Email or username"
            value={emailOrUsername}
            onChange={(e) => setEmailOrUsername(e.target.value)}
            className=" w-full pb-1 px-2 pt-2 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-400"
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full pb-1 px-2 pt-2 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-400"
          />
          <div className="flex flex-row w-full justify-between text-xs text-stone-900">
            <div className="flex flex-row gap-1">
              <input type="checkbox" />
              <div>remember me</div>
            </div>
            <div className="text-stone-900 text-xs border-b-1 border-stone-800 w-fit cursor-pointer hover:text-stone-400 hover:border-stone-400 transition-all duration-200">
              forgot password?
            </div>
          </div>
          <button
            className="p-2 bg-stone-300 cursor-pointer hover:bg-stone-100 rounded-xl mt-6 w-full"
            onClick={handleSignIn}
          >
            {isSigningIn ? "Logging In" : "Play"}
          </button>
          {error && <div className="text-red-500 font-semibold">{error}</div>}
        </div>
        <div className="px-4 flex flex-row gap-2 items-center h-fit mt-2 text-transparent cursor-default">
          <span className="font-bold">
            <IoMdArrowRoundBack />
          </span>
          Back{" "}
        </div>

        {/* <Link href={"/sign-up"} className="text-white/80">
          Sign up
        </Link>
        <Link href={"/"} className="text-white/80">
          Back to home
        </Link> */}
      </div>
      <div className="w-[25vw] mt-2 text-stone-100 text-sm flex flex-row gap-4 justify-center">
        <div className="text-stone-100/30">{"Don't have an account?"}</div>
        <Link href={"/sign-up"} className="text-white/80 text-sm">
          Create an account
        </Link>
      </div>
    </div>
  );
}
