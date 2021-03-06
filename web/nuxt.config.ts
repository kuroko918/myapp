import { NuxtConfig } from '@nuxt/types'
// @ts-ignore
import colors from 'vuetify/es5/util/colors'

const config: NuxtConfig = {
  head: {
    titleTemplate: '%s - myapp-kuroko918',
    title: 'myapp-kuroko918',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  css: [
  ],
  plugins: [
    './plugins/axios',
    './plugins/firebase',
    './plugins/persistedState',
  ],
  components: true,
  buildModules: [
    ['@nuxtjs/date-fns', { locales: ['ja'], methods: ['format'] }],
    '@nuxtjs/dotenv',
    '@nuxt/typescript-build',
    '@nuxtjs/vuetify',
  ],
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/dotenv',
    'vue-scrollto/nuxt',
  ],
  axios: {},
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        }
      }
    }
  },
  build: {
    filenames: {
      app: ({ isDev }) => isDev ? '[name].[hash].js' : '[chunkhash].js',
      chunk: ({ isDev }) => isDev ? '[name].[hash].js' : '[chunkhash].js'
    }
  },
  router: {
    middleware: 'authenticate'
  },
  server: {
    port: 3000,
    host: '0.0.0.0',
    timing: false,
  },
}

export default config
