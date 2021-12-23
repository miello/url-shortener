module.exports = {
  mode: "jit",
  content: ["./public/index.html", "./src/**/*.svelte"],
  darkMode: "media",
  theme: {
    fontFamily: {
      display: ["Poppins", "sans-serif"],
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
