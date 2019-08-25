<template lang="pug">
section.section
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title Scope
    //- .hero-foot
    //-   nav.tabs.is-boxed.is-fullwidth
    //-     .container
    //-       ul
    //-         li.is-active
    //-           nuxt-link(:to='"/dept/" + this.objId ')  Overview
    //-         li
    //-           nuxt-link(:to='"/dept/" + this.objId + "/course" ') Course List
  
  .columns.is-multiline
    //- most popular
    
    
    //- card(
    //-   v-for="k in deptList"
    //-   :deptName='k.dept_name'
    //-   :deptId='k.dept_id'
    //-   :objId='k._id'
    //-   :courseCount='k.courses? k.courses.length: 0'
    //-   )
</template>
<script>
export default {
  data: () => ({
    hello: "#Hello World"
  }),
  methods: {
    async fetchPost() {
      try {
        let ip = await this.$axios.$get("/md/info.md");
        console.log(ip);
        this.hello = ip
      } catch (e) {
        console.warn(e);
      }
    }
  },
  created() {
    this.fetchPost()
  }
};
</script>