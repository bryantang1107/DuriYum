/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      flex: {
        1: "1 1 0%",
        4: "4 4 0%",
      },
      height: {
        "1/10": "10%",
        "8/10": "80%",
      },
    },
  },
  plugins: [],
};
