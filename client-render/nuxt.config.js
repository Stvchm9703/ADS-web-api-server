
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
    '@nuxtjs/pwa',
    '@nuxtjs/markdownit'
  ],
  markdownit: {
    preset: 'commonmark',
    linkify: true,
    breaks: false,
    use: [
      'markdown-it-div',
      'markdown-it-attrs',
      'markdown-it-multimd-table',
      'markdown-it-task-lists',
      // 'markdown-it-github-headings',
      'markdown-it-github-preamble',
      'markdown-it-table-of-contents',

    ],
    injected: true
  },
  axios: {
    proxy: true,
  },
  proxy: {
    '/api': {
      target: 'http://0.0.0.0:8080',

    },
    '/md': {
      target: 'http://0.0.0.0:8080'
    }
  },
  generate: {
    dir: 'dist',
  },
  build: {
    publicPath: '/static/',
    indicator: false,
    extend(config, ctx) {
      config.module.rules.push({
        test: /\.md$/,
        use: ['raw-loader']
      });
    },
    extractCSS: true,
    vendor: ['vuex', 'axios']
  },
}
