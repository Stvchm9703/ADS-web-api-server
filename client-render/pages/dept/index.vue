<template lang="pug">
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
import Card from '~/components/DeptCard.vue'
export default {
  name: 'deptPage',
  components: { Card },
  methods: {
    hello () {
      return 'hello'
    }
  }
}
</script>
