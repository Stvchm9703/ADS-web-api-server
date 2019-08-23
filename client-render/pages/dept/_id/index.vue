<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-left
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
    Info(
      :objId="objId"
      :deptName="deptName"
      :deptId="deptId"
      :location="location"
      :course="courseList"
      :createdAt="createdAt"
      :lastUpdated="lastUpdated"
    )
  
</template>

<script>

import Info from '~/components/DeptInfo.vue'

export default {
  components: { Info },
  data : () => ({
    objId : "",
    deptName : "",
    deptId : "",
    courseList : [],
    location : "",
    lastUpdated: "",
    createdAt : ""
  }),
  methods: {
    async fetchDept(){
      let ip = await this.$axios.$get('/api/v1/g/dept/' + this.$route.params["id"] )
      this.objId = ip.data._id
      this.deptName = ip.data.dept_name
      this.deptId = ip.data.dept_id
      this.courseList = ip.data.courses == null ? [] : ip.data.courses
      this.location = ip.data.location
      this.lastUpdated = ip.data.updated_at
      this.createdAt = ip.data.created_at
    },
  },
  beforeMount () {
    this.fetchDept()
  } 

}
</script>
