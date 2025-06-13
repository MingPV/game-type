"use client";

import React, { useEffect, useState } from "react";
import { validateEmail } from "@/lib/utils/validateEmail";
import Link from "next/link";
import { useRouter } from "next/navigation";

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

    router.push("/home");
  };

  return (
    <div className="bg-stone-800 h-[100vh] w-[100vw] flex justify-center items-center">
      <div className="bg-white/20 p-10 flex flex-col gap-2">
        <input
          type="text"
          placeholder="Email or username"
          value={emailOrUsername}
          onChange={(e) => setEmailOrUsername(e.target.value)}
          className="pb-1 px-2 pt-2 border-b-1 border-black/20"
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="pb-1 px-2 pt-2 border-b-1 border-black/20"
        />
        <button
          className="p-2 bg-black/20 cursor-pointer hover:bg-black/30"
          onClick={handleSignIn}
        >
          {isSigningIn ? "Signing in" : "Sign in"}
        </button>
        {error && <div className="text-red-500 font-semibold">{error}</div>}
        <Link href={"/sign-up"} className="text-white/80">
          Sign up
        </Link>
        <Link href={"/"} className="text-white/80">
          Back to home
        </Link>
      </div>
    </div>
  );
}
