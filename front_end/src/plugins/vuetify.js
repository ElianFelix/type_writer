/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Composables
import { createVuetify } from 'vuetify'

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

const light = {
  dark: false,
  colors: {
    primary: '#82FFE2',
    secondary: '#74AEFF',
    error: '#FD3A0F',
    'on-primary': '#001A14',
    'on-secondary': '#000B1A',
    'on-error': '#190500',
    background: '#E6E6E6',
},
}

const dark = {
  dark: true,
  colors: {
    primary: '#82FFE2',
    secondary: '#74AEFF',
    error: '#FD3A0F',
    'on-primary': '#001A14',
    'on-secondary': '#00204D',
    'on-error': '#190500',
    surface: '#2F2F2F',
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
