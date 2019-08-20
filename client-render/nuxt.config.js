
export default {
  mode: 'spa',

  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  loading: { color: '#fff' },
  css: [
  ],
  plugins: [
    '~/plugins/axios'
  ],
  devModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    // '@nuxtjs/eslint-module'
  ],
  modules: [
    'nuxt-buefy',
    '@nuxtjs/axios',
    '@nuxtjs/pwa'
  ],
  axios: {
    proxy : true,
  },
   proxy: {
       '/api': {
           target: 'http://0.0.0.0:8080',
           pathRewrite: {
               '^/api': '/api'
           }
       }
   },
  generate: { 
    dir : 'dist',
  }
}
