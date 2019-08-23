<template lang="pug">
section.section
  .columns.is-multiline
    .column
      nuxt-link.button.is-primary(to="/create/student") Create Student
  
  .columns.is-multiline
    card(
      v-for="k in List"
      :student_name='k.student_name'
      :student_id='k.student_id'
      :objId='k._id'
      :enrolled='k.enrolled'
      @onReload="fetchStudent()"
      )
</template>

<script>
import Card from '~/components/StudentCard.vue'
export default {
  name: 'StudentPage',
  components: { Card },
  data : () => ({
    List: [],
  }),
 
  methods: {
    async fetchStudent(){
      let qwe = await this.$axios.$get('/api/v1/l/student')
      this.List = qwe.data
    }
  },
  beforeMount () {
    this.fetchStudent()
  } 

}
</script>
