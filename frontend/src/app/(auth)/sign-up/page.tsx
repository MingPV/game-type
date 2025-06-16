"use client";

import { validateEmail } from "@/lib/utils/validateEmail";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { FaCheck } from "react-icons/fa";
import { ImCross } from "react-icons/im";

export default function Page() {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [isSigningUp, setIsSigningUp] = useState(false);
  const [isUsernameValid, setIsUsernameValid] = useState(true);
  const [checkingUsername, setCheckingUsername] = useState(false);
  const [error, setError] = useState("");

  const router = useRouter();

  // Check username unique
  useEffect(() => {
    if (!username) {
      setIsUsernameValid(true);
      return;
    }

    setCheckingUsername(true);

    const timeout = setTimeout(async () => {
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/users/username/${username}`
        );

        setIsUsernameValid(!res.ok);
      } catch (err) {
        console.error("Error checking username", err);
        setIsUsernameValid(true);
      } finally {
        setCheckingUsername(false);
      }
    }, 700);

    return () => clearTimeout(timeout);
  }, [username]);

  useEffect(() => {}, []);

  const handleSignUp = async () => {
    setError("");
    setIsSigningUp(true);

    if (!validateEmail(email)) {
      setError("Invalid email.");
      setIsSigningUp(false);
      return;
    }

    if (password.length < 6) {
      setError("Password should contain more than 6 characters.");
      setIsSigningUp(false);
      return;
    }

    if (password != confirmPassword) {
      setError("Passwords do not match.");
      setIsSigningUp(false);
      return;
    }

    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/auth/signup`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            email,
            password,
          }),
        }
      );

      const data = await res.json();

      if (!res.ok) {
        setError(data.message || "Sign-in failed. Please try again.");
        setIsSigningUp(false);
        return;
      } else {
        console.log("Sign-in success", data);
      }
    } catch (err) {
      console.error(err);
      setError("Something went wrong. Please try again later.");
    }

    setIsSigningUp(false);
    setEmail("");
    setUsername("");
    setPassword("");
    setConfirmPassword("");

    router.push("/sign-in");
  };

  return (
    <div className="h-full w-full flex justify-center items-center">
      <div className="bg-white/20 p-10 flex flex-col gap-2">
        <input
          type="text"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="pb-1 px-2 pt-2 border-b-1 border-black/20"
        />
        <div className="flex flex-row gap-2 items-center border-b-1 border-black/20">
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="pb-1 px-2 pt-2 "
          />
          {username ? (
            checkingUsername ? (
              "checking"
            ) : isUsernameValid ? (
              <span className="mt-2 text-lime-400/50">
                <FaCheck />
              </span>
            ) : (
              <span className="mt-2 text-red-600/50">
                <ImCross />
              </span>
            )
          ) : null}
        </div>
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="pb-1 px-2 pt-2 border-b-1 border-black/20"
        />
        <input
          type="password"
          placeholder="Confirm Password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          className="pb-1 px-2 pt-2 border-b-1 border-black/20"
        />
        <button
          className="p-2 bg-black/20 cursor-pointer hover:bg-black/30"
          onClick={handleSignUp}
        >
          {isSigningUp ? "Signing up" : "Sign up"}
        </button>
        {error && <div className="text-red-500 font-semibold">{error}</div>}
        <Link href={"/sign-in"} className="text-white/80">
          Sign in
        </Link>
        <Link href={"/"} className="text-white/80">
          Back to home
        </Link>
      </div>
    </div>
  );
}
