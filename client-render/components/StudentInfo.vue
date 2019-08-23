<template lang="pug">
.container
  h1.title
    | Information
  h2.subtitle
    | Basic information of the Student
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Student name 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{student_name}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Student Id Code 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{student_id}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Day of Birth 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{dob | dobForm}}
  .columns
    .column.is-one-quarter-fullhd.is-one-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      h3  Total Enrolled Count: 
    .column.is-third-quarter-fullhd.is-two-third-widescreen.is-full-desktop.is-full-tablet.is-full-mobile
      p {{enrolled? enrolled.length : 0}}

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
      nuxt-link.button.is-primary(:to='"/student/" + objId +"/update" ')   Update information
    .column.is-half-fullhd.is-half-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
      b-button.is-danger(@click="warnDelete()")  Delete information
</template>

<script>
import moment from "moment";
export default {
  props: {
    objId: String,
    student_name: { type: String },
    student_id: { type: String },
    dob: String,
    enrolled: { type: Array, default: function() { return []; } },
    createdAt: String,
    lastUpdated: String
  },
  filters: {
    timeFormat(i) {
      return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss");
    }, 
    dobForm(i){
       return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD");
    }
  },
  methods: {
    warnDelete() {
      this.$buefy.dialog.confirm({
        title: "Deleting Student",
        message:
          "Are you sure you want to <b>delete</b> this Student? This action cannot be undone.",
        confirmText: "Delete Student",
        type: "is-danger",
        hasIcon: true,
        onConfirm: () => {
          this.deleteStudent();
        }
      });
    },
    async deleteStudent() {
      let ip = await this.$axios.$post("/api/v1/d/student/", {
        _id: String(this.objId)
      });
      console.log(ip);
      this.$buefy.toast.open({
        duration: 5000,
        message: `deleted !`,
        position: "is-bottom",
        type: "is-success"
      });
      this.$router.push({
        path: "/student"
      });
    }
  }
};
</script>

<style>
</style>