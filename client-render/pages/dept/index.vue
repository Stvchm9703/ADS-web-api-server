<template lang="pug">
section.section
  .columns.is-multiline
    .column
      nuxt-link.button.is-primary(to="/create/dept") Create Department
  
  .columns.is-multiline
    card(
      v-for="k in deptList"
      :deptName='k.dept_name'
      :deptId='k.dept_id'
      :objId='k._id'
      :courseCount='k.courses? k.courses.length: 0'
      )
</template>

<script>
import Card from '~/components/DeptCard.vue'
export default {
  name: 'deptPage',
  components: { Card },
  data : () => ({
    deptList: [],
  }),
  filters: {
    timeFormat (i){
      return moment(i , "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD")
    }
  },
  methods: {
    async fetchDept(){
      let qwe = await this.$axios.$get('/api/v1/l/dept')
      this.deptList = qwe.data
    }
  },
  beforeMount () {
    this.fetchDept()
  } 

}
</script>
