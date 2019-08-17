<template lang="pug">
.column.is-one-third
  nuxt-link(:to="toLink") 
    .card
      header.card-header
        p.card-header-title
          | {{title}} | {{courseId}}
      .card-content
        .content
          p
          | level: {{level}}
          | Department : 

</template>

<script>
export default {
  props: {
    objId :{
        type: String
    },
    deptId : {
        type : String
    },
    deptName : {
        type  : String
    },
    deptObjId : {
        type : String
    },
    courseId:{
        type : String
    },
    title : {
        type : String
    },
    level : {
        type : Number
    },
    offerCount :{
        type : Number
    },
    Offers : {
        type : Array
    },
    updatedAt : {
        type : String
    }
  },
  computed : {
    toLink : function () {
      return '/course/' + this.objId
    },
    latestOffer : function () {
        return this.Offers[-1]
    }
  }
}
</script>
