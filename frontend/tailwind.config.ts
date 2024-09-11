import type { Config } from "tailwindcss";

export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],

  theme: {
    extend: {
      colors: {
        "go": "#00ADD8",
        "mango": "#F4BB44"
      }
    }
  },

  plugins: []
} as Config;
