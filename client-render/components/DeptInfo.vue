<template lang="pug">
.container
  h1.title
    | Information
  h2.subtitle
    | Basic information of the Department
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Department name 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{deptName}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Department Id Code 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{deptId}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Department Office Location 
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
      nuxt-link.button.is-primary(:to='"/dept/" + objId +"/update" ')   Update information
    .column.is-half-fullhd.is-half-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
      b-button.is-danger(@click="warnDelete()")  delete information
</template>

<script>
import moment from "moment";
export default {
  props: {
    objId: String,
    deptName: { type: String },
    deptId: { type: String },
    location: String,
    course: Array,
    createdAt: String,
    lastUpdated: String
  },
  filters: {
    timeFormat(i) {
      return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss");
    }
  },
  methods: {
    warnDelete() {
      this.$buefy.dialog.confirm({
        title: "Deleting Department",
        message:
          "Are you sure you want to <b>delete</b> this Department? This action cannot be undone.",
        confirmText: "Delete Department",
        type: "is-danger",
        hasIcon: true,
        onConfirm:()=>{ this.deleteDept()}
      });
    },
    async deleteDept() {
      let ip = await this.$axios
        .$post("/api/v1/d/dept/", {
          _id: String(this.objId)
        })
      console.log(ip)
      this.$buefy.toast.open({
        duration: 5000,
        message: `deleted !`,
        position: "is-bottom",
        type: "is-success"
      });
      this.$router.push({
        path : "/dept"
      })
    }
  }
};
</script>

<style>
</style>