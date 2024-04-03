/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,mustache}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["winter", "dim"],
    darkTheme: "dim",
  },
}

