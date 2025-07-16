"use client";

import { useState, useEffect, useCallback } from "react";

export function useTyping(onCommand: (cmd: string) => void) {
  const [isTypingMode, setIsTypingMode] = useState(false);
  const [command, setCommand] = useState("");

  const handleCommand = useCallback(
    (cmd: string) => {
      const lower = cmd.toLowerCase();
      onCommand(lower);
    },
    [onCommand]
  );

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === "Tab") {
        e.preventDefault();
        setIsTypingMode((prev) => !prev);
        setCommand("");
      }

      if (isTypingMode) {
        if (e.key === "Enter") {
          handleCommand(command);
          setCommand("");
          setIsTypingMode(false);
        } else if (e.key === "Backspace") {
          setCommand((prev) => prev.slice(0, -1));
        } else if (e.key.length === 1) {
          setCommand((prev) => prev + e.key);
        }
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [isTypingMode, command, handleCommand]);

  return { isTypingMode, command };
}
