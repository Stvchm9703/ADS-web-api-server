<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-centered
        p.title {{deptName}}
        p.subtitle {{deptId}}
    .hero-foot
      nav.tabs.is-boxed.is-fullwidth
        .container
          ul
            li.is-active
              nuxt-link(:to='"/dept/" + this.objId ')  Overview
            li
              nuxt-link(:to='"/dept/" + this.objId + "/course" ') Course List
  section.section
    nuxt
</template>