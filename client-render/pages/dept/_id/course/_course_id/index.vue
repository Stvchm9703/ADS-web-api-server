<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title.is-left {{course_title}} ({{course_id}})
        p.subtitle {{deptName}} ({{deptId}})
    .hero-foot
      nav.tabs.is-boxed.is-fullwidth
        .container
          ul
            li.is-active
              nuxt-link(:to='"/dept/" + deptObjId+ "/course/" + course_obj_id ')  Overview
            li
              nuxt-link(:to='"/dept/" + deptObjId+ "/course/" + course_obj_id  + "/offer" ')  Offers List
            li

  section.section
    .container
      h1.title
          | Information
      h2.subtitle
          | Basic information of the Course
      .columns.is-multiline
          .column
            h3  Title 
          .column
            p {{course_title}}
      .columns.is-multiline
          .column
            h3  Course Id Code 
          .column
            p {{course_id}}
     
      .columns.is-multiline
          .column.is-half-tablet.is-full-mobile
            h3  Total Offers Count: 
          .column.is-half-tablet.is-full-mobile
            p {{offers? offers.length : 0}}

      .columns
          .column.is-full
          hr
      
      .columns.is-multiline
          .column.is-one-third-fullhd.is-one-third-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
            h3 Created at : {{createdAt|timeFormat}}  
          .column.is-one-third-fullhd.is-one-third-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
            h3 Last Update at :  {{lastUpdated|timeFormat}}  
      
      .columns.is-vcentered.is-multiline
          .column.is-half-fullhd.is-half-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
            nuxt-link.button.is-primary(
              :to='"/dept/" + deptObjId + "/course/" + course_obj_id + "/update"'
            )  Update Course
          .column.is-half-fullhd.is-half-widescreen.is-half-desktop.is-full-tablet.is-full-mobile
            b-button.is-danger(
              @click='warnDelete()'
            )  Delete Course
      
</template>

<script>
import moment from "moment";
import Info from "~/components/DeptInfo.vue";

export default {
  components: { Info },
  data: () => ({
    deptName: "",
    deptId: "",
    deptObjId: "",
    course_obj_id: "",
    course_id: "",
    course_title: "",
    lastUpdated: "",
    createdAt: "",
    offers: []
  }),
  filters: {
    timeFormat(i) {
      return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss");
    }
  },
  methods: {
    async fetchDept() {
      let ip = await this.$axios.$get(
        "/api/v1/g/dept/" + this.$route.params["id"]
      );
      console.log(ip);
      this.deptName = ip.data.dept_name;
      this.deptId = ip.data.dept_id;
      this.deptObjId = ip.data._id;
    },
    async fetchCourse() {

      let ip = await this.$axios.$get(
        "/api/v1/g/dept/" +
          this.$route.params["id"] +
          "/course/" +
          this.$route.params["course_id"]
      );
      console.log(ip);
      this.course_obj_id = ip.data._id;
      this.course_id = ip.data.course_id;
      this.course_title = ip.data.title;
      this.lastUpdated = ip.data.updated_at;
      this.createdAt = ip.data.created_at;
      this.offers = ip.data.offers;
    },
    warnDelete() {
      this.$buefy.dialog.confirm({
        title: "Deleting Course",
        message:
          "Are you sure you want to <b>delete</b> this Cousre? This action cannot be undone.",
        confirmText: "Delete Cousre",
        type: "is-danger",
        hasIcon: true,
        onConfirm:()=>{ this.deleteCourse()}
      });
    },
    async deleteCourse() {
      let requet_pack = {
        _id: this.objId
      };
      try {
        let ip = await this.$axios.$post(
          "/api/v1/d/dept/" + this.deptObjId + "/course/" + this.course_obj_id
        );
        this.$buefy.toast.open({
          duration: 5000,
          message: `Delete Success !`,
          position: "is-bottom",
          type: "is-success"
        });
        this.$router.push({
          path : '/dept/' + this.deptObjId + '/course'
        })
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          duration: 5000,
          message: `Delete Fail !`,
          position: "is-bottom",
          type: "is-danger"
        });
      } 
    }
  },
  beforeMount() {
    this.fetchDept();
    this.fetchCourse();
  }
};
</script>
