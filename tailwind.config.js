const formsPlugin = require('@tailwindcss/forms')

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    '**/*.{html,templ}',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography')
  ],
}

