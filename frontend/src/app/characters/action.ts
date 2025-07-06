export async function fetchCharacters(user_id: string) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/characters/userid/${user_id}`
    );

    if (!res.ok) {
      throw new Error("Failed to fetch characters");
    }

    const data = await res.json();
    return data;
  } catch (err) {
    console.log(err);
  }
}

export async function fetchClasses() {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/classes`
    );

    if (!res.ok) {
      throw new Error("Failed to fetch classes");
    }

    const data = await res.json();
    return data;
  } catch (err) {
    console.log(err);
  }
}

export async function CreateCharacter(
  user_id: string,
  name: string,
  class_id: string
) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/api/v1/characters`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_id: user_id,
          name: name,
          level: 1,
          current_exp: 0,
          class_id: class_id,
        }),
      }
    );

    if (!res.ok) {
      throw new Error("Failed to create character");
    }

    const data = await res.json();
    return data;
  } catch (err) {
    console.log(err);
    throw new Error("Failed to create character.");
  }
}
