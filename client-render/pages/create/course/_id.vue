<template lang="pug">
.container
  .section
    .hero.is-primary.is-bold.is-12
      .hero-body
        .container
          h1.title Create Course 
          h2.subtitle 
  CourseForm(
    course_id=""
    title=""
    :level="1"
    :isUpdate="false"
    @onFormSubmit="submit_check"
    @onCancel="$router.back()"
  )
</template>

<script>
import CourseForm from "~/components/CourseForm.vue";
export default {
  name: "CreateCoursePage",
  components: { CourseForm },
  data: () => ({
    dept_obj_id: "",
    course_id: "",
    title: "",
    level: ""
  }),
  mounted() {
    this.dept_obj_id = this.$route.params["_id"];
    this.fetchDept();
  },
  methods: {
    async fetchDept() {
      try {
        let ip = await this.$axios.$get(
          "/api/v1/g/dept/" + this.$route.params["id"]
        );
        this.dept_obj_id = this.$route.params["id"];
      } catch (e) {
        console.log(e);
        console.log("the requested dept is not exists");
        console.log("redirect to home");
        this.$buefy.toast.open({
          message: "The Requested Department is not exist",
          type: "is-warning",
          position: "is-bottom"
        });
        this.$router.push({
          path: "/dept"
        });
      }
    },
    async submit_check(e) {
      console.log(e);
      try {
        let ip = await this.$axios.$post(
          "/api/v1/c/dept/" + this.dept_obj_id + "/course",
          e
        );
        this.$buefy.toast.open({
          message: "Course Create Success",
          type: "is-success",
          position: "is-bottom"
        });
        this.$router.push({
          path: "/dept/" + this.dept_obj_id + "/course/" + ip.data._id
        });
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          message: "Course Create Fail",
          type: "is-danger",
          position: "is-bottom"
        });
        this.$router.push({
          path: "/dept"
        });
      }
    }
  }
};
</script>
