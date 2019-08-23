<template lang="pug">
.container
  .section
    .hero.is-primary.is-bold.is-12
      .hero-body
        .container
          h1.title Create Student 
          h2.subtitle 
  StudentForm(
    :student_id="student_id"
    :student_name="student_name"
    :dob="dob"
    :isUpdate="false"
    @onFormSubmit="submit_check"
    @onCancel="$router.back()"
  )
</template>

<script>
import StudentForm from "~/components/StudentForm.vue";
export default {
  name: "CreateStudentPage",
  components: { StudentForm },
  data: () => ({
    student_id: "",
    student_name: "",
    dob: ""
  }),
  methods: {
    async submit_check(e) {
      console.log(e);
      let ip = await this.$axios.$post("/api/v1/c/student/", e);
      console.log(ip);
      if (ip.status == 0) {
        this.$buefy.toast.open({
          message: "Student Create Success",
          type: "is-success",
          position: "is-bottom"
        });
        setTimeout(() => {
          this.$router.push({
            path: "/student/" + ip.data._id
          });
        }, 2500);
      }
    }
  }
};
</script>
