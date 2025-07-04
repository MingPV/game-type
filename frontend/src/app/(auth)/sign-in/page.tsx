"use client";

import React, { useEffect, useState } from "react";
import { validateEmail } from "@/lib/utils/validateEmail";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { IoMdArrowRoundBack } from "react-icons/io";

export default function Page() {
  const [guestUsername, setGuestUsername] = useState("");
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
      <div className="flex flex-col justify-center items-center h-full w-full z-10">
        <Link
          className="fixed left-10 top-10 mr-2 flex items-center flex-row gap-2 h-fit mt-2 text-stone-300/80 w-fit font-bold bg-white/20 hover:bg-white/10 transition-all duration-200 py-2 px-4 rounded-md z-11"
          href={"/"}
        >
          <span className="font-bold text-xl">
            <IoMdArrowRoundBack />
          </span>
          <span className="mt-1 text-xl">Back </span>
        </Link>
        <div className="bg-white/40 px-14 py-10 flex flex-row shadow-md shadow-stone-700 rounded-xl">
          <div className="p-10 w-[30vw] flex flex-col border-r pr-20 border-stone-100 justify-center">
            <div className="text-stone-800 text-4xl text-center">
              Play as guest
            </div>
            <input
              type="text"
              placeholder="Enter your username"
              value={guestUsername}
              onChange={(e) => setGuestUsername(e.target.value)}
              className=" w-full px-2 py-2 mt-8 mb-4 text-xl bg-white/10 placeholder:font-mono text-stone-100 text-center placeholder-black/40 placeholder:text-sm focus:ring-1 focus:ring-stone-600 focus:outline-none font-mono font-bold border border-transparent hover:border-stone-60 rounded-md"
            />
            <button
              className="p-2 bg-stone-300 cursor-pointer hover:bg-stone-100 rounded-xl mt-6 w-full border-r-4 border-b-4 border-stone-800 text-xl text-stone-800"
              // onClick={handleSignIn}
            >
              {isSigningIn ? "Logging In" : "Play"}
            </button>
            <div className="text-start mt-2 text-white/80 text-lg">
              ** Progress may be lost when playing as a guest.
            </div>
          </div>
          <div className=" p-10 flex flex-col items-center w-[30vw] pl-20">
            <div className="text-stone-800 text-4xl text-center">
              Play with account
            </div>
            <input
              type="text"
              placeholder="Email or username"
              value={emailOrUsername}
              onChange={(e) => setEmailOrUsername(e.target.value)}
              className="w-full py-3 px-2 mt-6 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-300 font-mono placeholder:text-sm placeholder:text-stone-500 font-bold pl-4 focus:ring-1 focus:ring-stone-300 focus:outline-none rounded-sm"
            />
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full py-3 px-2 mt-2 mb-4 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-300 font-mono placeholder:text-sm placeholder:text-stone-500 font-bold pl-4 focus:ring-1 focus:ring-stone-300 focus:outline-none rounded-sm"
            />
            <div className="flex flex-row w-full justify-between text-stone-800">
              <div className="flex flex-row gap-1">
                <input type="checkbox" />
                <div>remember me</div>
              </div>
              <div className="text-stone-800 border-stone-800 w-fit cursor-pointer hover:underline transition-all duration-200">
                forgot password?
              </div>
            </div>
            <button
              className="p-2 bg-stone-300 cursor-pointer hover:bg-stone-100 rounded-xl mt-6 w-full border-r-4 border-b-4 border-stone-800 text-xl text-stone-800"
              onClick={handleSignIn}
            >
              {isSigningIn ? "Logging In" : "Play"}
            </button>
            {error && <div className="text-red-500 font-semibold">{error}</div>}
          </div>
        </div>
        <div className="w-[25vw] mt-2 text-stone-100 text-sm flex flex-row gap-4 justify-center">
          <div className="text-stone-100/30 text-xl">
            {"Don't have an account?"}
          </div>
          <Link
            href={"/sign-up"}
            className="text-white/80 text-xl hover:underline"
          >
            Create an account
          </Link>
        </div>
      </div>
    </div>
  );
}
