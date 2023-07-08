/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*.html", "./pages/*.html"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
