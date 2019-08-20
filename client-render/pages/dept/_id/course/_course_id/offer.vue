<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title.is-left {{course_title}} ({{course_id}})
        p.subtitle {{deptName}} ({{deptId}})
    .hero-foot
      nav.tabs.is-boxed.is-fullwidth
        .container
          ul
            li
              nuxt-link(:to='"/dept/" + deptObjId+ "/course/" + course_obj_id ')  Overview
            li.is-active
              nuxt-link(:to='"/dept/" + deptObjId+ "/course/" + course_obj_id  + "/offer" ')  Offers List
            li

  section.section
    .container
      
      
</template>

<script>

import Info from '~/components/DeptInfo.vue'

export default {
  components: { Info },
  data : () => ({
    deptName : "",
    deptId : "",
    deptObjId : "",
    course_obj_id : "",
    course_id : "",
    course_title :"",
    lastUpdated : "",
    createdAt : "",
    offers_count : [],
  }),
  methods: {
    async fetchDept(){
      let ip = await this.$axios.$get(
        '/api/v1/g/dept/' + 
        this.$route.params["id"]   
      )
      console.log(ip)
      this.deptName = ip.data.dept_name
      this.deptId = ip.data.dept_id
      this.deptObjId = ip.data._id
    },
    async fetchCourse(){
      let   ip = await this.$axios.$get(
        '/api/v1/g/dept/' + 
        this.$route.params["id"] +
        '/course/' + 
        this.$route.params["course_id"]    
      )
      console.log(ip)
      this.course_obj_id = ip.data._id
      this.course_id = ip.data.course_id
      this.course_title = ip.data.title
      this.lastUpdated = ip.data.updated_at
      this.createdAt = ip.data.created_at
      this.offers_count = ip.data.offers
    }
  },
  beforeMount () {
    this.fetchDept()
    this.fetchCourse()
  } 

}
</script>
