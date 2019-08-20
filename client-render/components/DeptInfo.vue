<template lang="pug">
.container
  h1.title
    | Information
  h2.subtitle
    | Basic information of the Depatment
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Depatment name 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{deptName}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Depatment Id Code 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{deptId}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Depatment Office Location 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{location}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Total Courses Count: 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{course.length}}

  .columns
    .column.is-full
      hr
  .columns
    .column.is-one-third-fullhd.is-one-third-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
      h3 Created at : {{createdAt|timeFormat}}  
    .column.is-one-third-fullhd.is-one-third-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
      h3 Last Update at :  {{lastUpdated|timeFormat}}  
  .columns.is-vcentered
    .column.is-half-fullhd.is-half-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
      button.button.is-primary  Update information
    .column.is-half-fullhd.is-half-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
      button.button.is-danger  delete information
</template>

<script>
import moment from 'moment'
export default {
  props: {
    deptName : { type  : String, },
    deptId : { type : String, },
    location : String,
    course : Array,
    createdAt : String,
    lastUpdated : String,
  },
  filters: {
    timeFormat (i){
      return moment(i , "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss")
    }
  },
}
</script>

<style>

</style>