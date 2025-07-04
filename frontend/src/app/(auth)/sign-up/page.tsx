"use client";

import { validateEmail } from "@/lib/utils/validateEmail";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { FaCheck } from "react-icons/fa";
import { ImCross } from "react-icons/im";
import { IoMdArrowRoundBack } from "react-icons/io";

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
        <div className="bg-white/40 px-14 py-8 flex flex-row shadow-md shadow-stone-700 rounded-xl">
          <div className=" p-10 flex flex-col items-center w-[30vw]">
            <div className="text-stone-800 text-4xl text-center">
              Create an account
            </div>
            <input
              type="text"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full py-3 px-2 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-300 font-mono placeholder:text-sm placeholder:text-stone-500 font-bold pl-4 focus:outline-none rounded-sm"
            />
            <div className="w-full flex flex-row bg-stone-700 my-2">
              <input
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="w-full px-2 py-3 border-b-1 border-black/20 bg-transparent text-stone-300 font-mono placeholder:text-sm placeholder:text-stone-500 font-bold pl-4 focus:ring-0 focus:outline-none rounded-sm"
              />
              {username ? (
                checkingUsername ? (
                  <span className="mr-6 text-stone-600 text-center flex justify-center items-center">
                    Checking
                  </span>
                ) : isUsernameValid ? (
                  <span className="mr-6 text-lime-400/50 text-center flex justify-center items-center">
                    <FaCheck />
                  </span>
                ) : (
                  <span className="mr-6 text-red-600/50 text-center flex justify-center items-center">
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
              className="w-full py-3 px-2 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-300 font-mono placeholder:text-sm placeholder:text-stone-500 font-bold pl-4 focus:outline-none rounded-sm"
            />
            <input
              type="password"
              placeholder="Confirm Password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="w-full py-3 px-2 border-b-1 border-black/20 bg-stone-700 my-2 text-stone-300 font-mono placeholder:text-sm placeholder:text-stone-500 font-bold pl-4 focus:outline-none rounded-sm"
            />
            <button
              className="p-2 bg-stone-300 cursor-pointer hover:bg-stone-100 rounded-xl mt-6 w-full border-r-4 border-b-4 border-stone-800 text-xl text-stone-800"
              onClick={handleSignUp}
            >
              {isSigningUp ? "Signing up" : "Create an account"}
            </button>
            {error && <div className="text-red-500 font-semibold">{error}</div>}
          </div>
        </div>
        <div className="w-[25vw] mt-2 text-stone-100 text-sm flex flex-row gap-4 justify-center">
          <div className="text-stone-100/30 text-xl">
            {"Already have an account?"}
          </div>
          <Link
            href={"/sign-in"}
            className="text-white/80 text-xl hover:underline"
          >
            Sign in
          </Link>
        </div>
      </div>
    </div>
  );
}
