<template lang="pug">
.container
  .section
    .hero.is-primary.is-bold.is-12
      .hero-body
        .container
          h1.title Update Course 
          h2.subtitle 
  CourseForm(
    :course_id="course_id"
    :title="title"
    :level="level"
    :isUpdate="true"
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
    obj_id: "",
    course_id: "",
    title: "",
    level: 1,
  }),
  mounted() {
    this.fetchCourse();
  },
  methods: {
    async fetchCourse() {
      try {
        let ip = await this.$axios.$get(
          "/api/v1/g/dept/" +
            this.$route.params["id"] +
            "/course/" +
            this.$route.params["course_id"]
        );
        this.dept_obj_id = this.$route.params["id"];
        this.course_id = ip.data.course_id;
        this.title = ip.data.title;
        this.level = ip.data.level;
        this.obj_id = ip.data._id;
        console.log(ip);
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
          "/api/v1/u/dept/" + this.dept_obj_id + "/course/" + this.obj_id,
          e
        );
        this.$buefy.toast.open({
          message: "Course Update Success",
          type: "is-success",
          position: "is-bottom"
        });
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          message: "Course Update Fail",
          type: "is-danger",
          position: "is-bottom"
        });
      } finally {
        this.$router.push({
          path: "/dept/" + this.dept_obj_id + "/course/" + this.obj_id
        });
      }
    }
  }
};
</script>
