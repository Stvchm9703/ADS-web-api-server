<template lang="pug">
.column.is-one-third
  nuxt-link(:to="toLink") 
    .card
      header.card-header
        p.card-header-title
          | {{deptName}} | {{deptId}}
      .card-content
        .content
          p
          | Courses: {{courseCount}}
</template>

<script>
export default {
  props: {
    deptName: {
      type: String,
      default: ""
    },
    deptId: {
      type: String,
      default: ""
    },
    courseCount: {
      type: Number,
      default: 0
    },
    objId: {
      type: String,
      default: ""
    }
  },
  computed : {
    toLink : function () {
      return '/dept/' + this.objId
    }
  }
}
</script>
