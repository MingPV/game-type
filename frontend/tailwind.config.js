/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",

    // Or if using `src` directory:
    "./src/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        gameFont: ["GameFont"],
        pixel: ["pixel"],
      },
      colors: {
        brandBlue: "#1DA1F2",
        softPink: "#FFB6C1",
        customGreen: {
          light: "#6EE7B7",
          DEFAULT: "#10B981",
          dark: "#047857",
        },
      },
    },
  },
  plugins: [],
};
