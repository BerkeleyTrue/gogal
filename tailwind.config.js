import { variants } from "@catppuccin/palette";

const colors = variants.frappe;

/** @type {import('tailwindcss').Config} */
const config = {
  content: ["./web/views/**/*.html"],
  theme: {
    extend: {},
  },
  plugins: [
    require("@catppuccin/tailwindcss")({
      defaultFlavour: "frappe",
    }),
    require("daisyui"),
  ],
  daisyui: {
    logs: false,
    themes: [
      {
        catppuccin: {
          primary: colors.rosewater.hex,
          secondary: colors.mauve.hex,
          accent: colors.pink.hex,
          neutral: colors.crust.hex,
          "base-100": colors.base.hex,
          info: colors.blue.hex,
          success: colors.green.hex,
          warning: colors.yellow.hex,
          error: colors.red.hex,
        },
      },
    ],
  },
};

export default config;
