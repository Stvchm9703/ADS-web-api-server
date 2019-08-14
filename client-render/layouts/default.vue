<template lang="pug">
div
  nav.navbar.header.has-shadow.is-primary(role='navigation' aria-label='main navigation')
    .navbar-brand
      a.navbar-item(href='/')
        h2.title.has-text-white SCOPE ADS
      .navbar-burger
        span
        span
        span
  section.main-content.columns
    aside.column.is-2.section
      p.menu-label.is-hidden-touch
        | General
      ul.menu-list
        li(v-for='(item, key) of items' :key='key')
          nuxt-link(:to='item.to' exact-active-class='is-active')
            b-icon(:icon='item.icon')
            | {{ item.title }}
    .container.column.is-10
      nuxt

</template>

<script>
export default {
  data () {
    return {
      items: [{
        title: 'Home',
        icon: 'home',
        to: {
          name: 'index'
        }
      },
      {
        title: 'Inspire',
        icon: 'lightbulb',
        to: {
          name: 'inspire'
        }
      }
      ]
    }
  }
}

</script>
