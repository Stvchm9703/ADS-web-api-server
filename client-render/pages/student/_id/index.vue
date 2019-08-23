<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title {{student_name}}
        p.subtitle {{student_id}}
    .hero-foot
      nav.tabs.is-boxed.is-fullwidth
        .container
          ul
            li.is-active
              nuxt-link(:to='"/student/" + this.objId ')  Overview
            li
              nuxt-link(:to='"/student/" + this.objId + "/enrolled" ') Enrolled Course List
  section.section
    Info(
      :objId="objId"
      :student_name="student_name"
      :student_id="student_id"
      :dob="dob"
      :enrolled="enrolled"
      :createdAt="createdAt"
      :lastUpdated="lastUpdated"
    )
  
</template>

<script>

import Info from '~/components/StudentInfo.vue'

export default {
  components: { Info },
  data : () => ({
    objId : "",
    student_name : "",
    student_id : "",
    enrolled : [],
    dob : "",
    lastUpdated: "",
    createdAt : ""
  }),
  methods: {
    async fetchDept(){
      let ip = await this.$axios.$get('/api/v1/g/student/' + this.$route.params["id"] )
      this.objId = ip.data._id
      this.student_name = ip.data.student_name
      this.student_id = ip.data.student_id
      this.enrolled = ip.data.enrolled == null ? [] : ip.data.enrolled
      this.dob = ip.data.dob
      this.lastUpdated = ip.data.updated_at
      this.createdAt = ip.data.created_at
    },
  },
  beforeMount () {
    this.fetchDept()
  } 

}
</script>
