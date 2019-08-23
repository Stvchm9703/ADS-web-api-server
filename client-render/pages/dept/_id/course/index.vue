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
            li
              nuxt-link(:to='"/dept/" + objId ')  Overview
            li.is-active
              nuxt-link(:to='"/dept/" + objId + "/course" ')  Overview Course List


  section.section
    .columns.is-multiline      
      nuxt-link.button.is-primary(:to='"/create/course/" + objId ') Create Course
    .columns.is-multiline
      card(
        v-for='c in courseList'
          :deptName='deptName'
          :deptId='deptId'
          :deptObjId='objId'
          :objId='c._id'
          :courseId='c.course_id'
          :title='c.title'
          :level='c.level'
          :offers='c.offers'
          :createdAt='c.created_at'
          :lastUpdated='c.updated_at'
          @onReload="fetchDept()"
        )
</template>

<script>
import moment from 'moment'
import Card from '~/components/CourseCard.vue'
export default {
  components: { Card },
  data : () => ({
    objId : "",
    deptName : "",
    deptId : "",
    courseList : [],
    location : "",
    lastUpdated: "",
    createdAt : ""
  }),
  computed: {
    courseCount () {
      return this.courseList.length;
    }
  },
  filters: {
    timeFormat (i){
      return moment(i , "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD")
    }
  },
  methods: {
    async fetchDept(){
      let ip = await this.$axios.$get('/api/v1/g/dept/' + this.$route.params["id"] )
      this.objId = ip.data._id
      this.deptName = ip.data.dept_name
      this.deptId = ip.data.dept_id
      this.courseList = ip.data.courses
      this.location = ip.data.location
      this.lastUpdated = ip.data.updated_at
      this.createdAt = ip.data.created_at
    }
  },
  beforeMount () {
    this.fetchDept()
  } 

}
</script>
