<template lang="pug">
.container
  .section
    .hero.is-primary.is-bold.is-12
      .hero-body
        .container
          h1.title Create Department 
          h2.subtitle 
  DeptForm(
    :dept_id="dept_id"
    :dept_name="dept_name"
    :location="location"
    :isUpdate="false"
    @onFormSubmit="submit_check"
    @onCancel="$router.back()"
  )
</template>

<script>
import DeptForm from "~/components/DeptForm.vue";
export default {
  name: "CreateDeptPage",
  components: { DeptForm },
  data: () => ({
    dept_id: "",
    dept_name: "",
    location: ""
  }),
  methods: {
    async submit_check(e) {
      console.log(e);
      let ip = await this.$axios.$post("/api/v1/c/dept/", e);
      console.log(ip);
      if (ip.status == 0) {
        this.$buefy.toast.open({
          message: "Department Create Success",
          type: "is-success",
          position: "is-bottom"
        });
        setTimeout(() => {
          this.$router.push({
            path: "/dept/" + ip.data._id
          });
        }, 2500);
      }
    }
  }
};
</script>
