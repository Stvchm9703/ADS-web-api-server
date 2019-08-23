<template lang="pug">
.container
  .section
    .hero.is-primary.is-bold.is-12
      .hero-body
        .container
          h1.title Update Student 
          h2.subtitle 
  StudentForm(
    :student_id="student_id"
    :student_name="student_name"
    :dob="dob"
    :isUpdate="true"
    @onFormSubmit="submit_check"
    @onCancel="$router.push({path:'/dept/' + objId})"
  )
</template>

<script>
import StudentForm from "~/components/StudentForm.vue";
export default {
  name: "UpdateStudentPage",
  components: { StudentForm },
  data: () => ({
    objId: "",
    student_id: "",
    student_name: "",
    dob: ""
  }),
  methods: {
    async fetchStudent() {
      try {
        let ip = await this.$axios.$get(
          "/api/v1/g/student/" + this.$route.params["id"]
        );
        if (ip.status == 0) {
          console.log(ip);
          this.objId = ip.data._id;
          this.student_id = ip.data.student_id;
          this.student_name = ip.data.student_name;
          this.dob = ip.data.dob;
        }
      } catch (e) {
        console.log("fail fetch student");
        this.$buefy.toast.open({
          message: "Unknown Student Record",
          type: "is-warning",
          position: "is-bottom"
        });
        this.$route.push({
          path: "/student"
        });
      }
    },
    async submit_check(e) {
      console.log(e);
      e["_id"] = this.objId;
      try {
        let ip = await this.$axios.$post("/api/v1/u/student/", e);
        console.log(ip);
        if (ip.status == 0) {
          this.$buefy.toast.open({
            message: "Student Update Success",
            type: "is-success",
            position: "is-bottom"
          });
          setTimeout(() => {
            this.$router.push({
              path: "/student/" + ip.data._id
            });
          }, 2500);
        }
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          message: "Student Update Fail",
          type: "is-danger",
          position: "is-bottom"
        });
      }
    }
  },
  // life cycle
  beforeMount() {
    this.fetchStudent();
  }
};
</script>
