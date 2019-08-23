<template lang="pug">
.container
  .section
    .hero.is-primary.is-bold.is-12
      .hero-body
        .container
          h1.title Update Department 
          h2.subtitle 
  DeptForm(
    :dept_id="dept_id"
    :dept_name="dept_name"
    :location="location"
    :isUpdate="true"
    @onFormSubmit="submit_check"
    @onCancel="$router.push({path:'/dept/' + objId})"
  )
</template>

<script>
import DeptForm from "~/components/DeptForm.vue";
export default {
  name: "UpdateDeptPage",
  components: { DeptForm },
  data: () => ({
    objId: "",
    dept_id: "",
    dept_name: "",
    location: ""
  }),
  methods: {
    async fetchDept() {
      try {
        let ip = await this.$axios.$get(
          "/api/v1/g/dept/" + this.$route.params["id"]
        );
        if (ip.status == 0) {
          console.log(ip);
          this.objId = ip.data._id;
          this.dept_name = ip.data.dept_name;
          this.dept_id = ip.data.dept_id;
          this.location = ip.data.location;
        }
      } catch (e) {
        console.log("fail fetch department");
        this.$buefy.toast.open({
          message: "Unknown Department Record",
          type: "is-warning",
          position: "is-bottom"
        });
        this.$route.push({
          path: "/dept"
        });
      }
    },
    async submit_check(e) {
      console.log(e);
      e["_id"] = this.objId;
      try {
        let ip = await this.$axios.$post("/api/v1/u/dept/", e);
        console.log(ip);
        if (ip.status == 0) {
          this.$buefy.toast.open({
            message: "Department Update Success",
            type: "is-success",
            position: "is-bottom"
          });
          setTimeout(() => {
            this.$router.push({
              path: "/dept/" + ip.data._id
            });
          }, 2500);
        }
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          message: "Department Update Fail",
          type: "is-danger",
          position: "is-bottom"
        });
      }
    }
  },
  // life cycle
  beforeMount() {
    this.fetchDept();
  }
};
</script>
