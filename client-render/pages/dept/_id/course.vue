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
            li
              a Overview
            li.is-active
              a Course List

  section.section
    .columns
      card(
        deptName='Computer Science'
        deptId='CS'
        objId='1234y9qwhk'
        courseCount=13
        )
</template>

<script>
import moment from 'moment'
import Card from '~/components/DeptCard.vue'
export default {
  components: { Card },
  data : () => ({
    deptName : "Computer Science",
    deptId : "CS",
    courseList : [{
      "_id" : "asbdivbdsv",
      "course_id" : "scope_12350",
      "title" : "dept testing",
      "level" : 1
    }],
    location : "green zone",
    lastUpdated: "2018/09/12",
    createdAt : "1889/09/12"
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
      const ip = await this.$axios.$get('/api/v1/g/dept/' + this.$route.params["id"] )
      console.log(ip)
    }
  },
  beforeMount () {
    this.fetchDept()
  } 

}
</script>
