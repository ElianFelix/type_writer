/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify } from 'vuetify'

const light = {
  dark: false,
  colors: {
    primary: '#82FFE2',
    secondary: '#74AEFF',
  },
}

const dark = {
  dark: true,
  colors: {
    primary: '#82FFE2',
    secondary: '#74AEFF',
  },
}

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  styles: {
    configFile: 'src/styles/settings.scss',
  },
  defaults: {
    global: {
      variant: 'flat',
    },
  },
  theme: {
    layers: true,
    defaultTheme: 'dark',
    themes: {
      light,
      dark,
    },
  },
})
