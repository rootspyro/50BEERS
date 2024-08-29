/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
	theme: {
    extend: {
      colors: {
        "dark": "#242424",
        "base": "#FFFDF6",
        "light": "#FAFAFA"
      },
      fontFamily: {
        "title": ["Anton", 'sans-serif'],
        "content": ["Roboto Condensed Variable", "sans-serif"]
      }
    },
	},
	plugins: [],
}
