"use client";

import { useCallback, useEffect, useState } from "react";
import { Direction } from "@/types/gameWorld";

const DIRECTION_KEYS: Record<string, Direction> = {
  KeyW: "UP",
  KeyS: "DOWN",
  KeyA: "LEFT",
  KeyD: "RIGHT",
  ArrowUp: "UP",
  ArrowDown: "DOWN",
  ArrowLeft: "LEFT",
  ArrowRight: "RIGHT",
};

export const useCharacterControls = () => {
  const [heldDirections, setHeldDirections] = useState<Direction[]>([]);

  // Handle Add and Remove key from heldDirections
  const handleKey = useCallback((e: KeyboardEvent, isKeyDown: boolean) => {
    const direction = DIRECTION_KEYS[e.code];
    if (!direction) return;

    setHeldDirections((prev) => {
      if (isKeyDown) {
        return prev.includes(direction) ? prev : [direction, ...prev];
      }
      return prev.filter((dir) => dir !== direction);
    });
  }, []);

  // Handle pressing key
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => handleKey(e, true);
    const handleKeyUp = (e: KeyboardEvent) => handleKey(e, false);

    window.addEventListener("keydown", handleKeyDown);
    window.addEventListener("keyup", handleKeyUp);

    return () => {
      window.removeEventListener("keydown", handleKeyDown);
      window.removeEventListener("keyup", handleKeyUp);
    };
  }, [handleKey]);

  // Function to get a direction.
  const getControlsDirection = useCallback((): Direction | null => {
    if (heldDirections.length >= 2) {
      const dirSet = new Set(heldDirections.slice(0, 2));
      if (dirSet.has("UP") && dirSet.has("LEFT")) return "UP_LEFT";
      if (dirSet.has("UP") && dirSet.has("RIGHT")) return "UP_RIGHT";
      if (dirSet.has("DOWN") && dirSet.has("LEFT")) return "DOWN_LEFT";
      if (dirSet.has("DOWN") && dirSet.has("RIGHT")) return "DOWN_RIGHT";
    }

    return heldDirections[0] || null;
  }, [heldDirections]);

  return { getControlsDirection };
};
