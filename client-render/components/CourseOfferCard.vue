<template lang="pug">

section.column.is-one-third
  b-collapse.card(:aria-id='"contentId" + courseId' :open='false')
    .card-header(
      slot='trigger' 
      slot-scope='props' 
      role='button' 
      :aria-controls='"contentId" + courseId')
      p.card-header-title
        | {{title}} | {{courseId}}
      a.card-header-icon
        b-icon(:icon="props.open ? 'menu-down' : 'menu-up'")
    .card-content
      .content
        p Year: {{year}}
        p level: {{level}}
        p Number of Student : {{num_of_stud}}
        p Department : 
          nuxt-link(:to='"/dept/" + deptObjId ') {{deptName}} ({{deptId}})
    footer.card-footer
      nuxt-link.card-footer-item(
        :to='"/dept/" + deptObjId + "/course/" + objId '
        ) Cousre Detail

      .card-footer-item.has-text-danger(
        @click="warnDelete()"
      ) Delete

</template>

<script>
import moment from "moment";
export default {
  name: "CourseOfferCard",
  props: {
    objId: { type: String },
    deptId: { type: String },
    deptName: { type: String },
    deptObjId: { type: String },
    courseId: { type: String },
    title: { type: String },
    level: { type: Number },
    num_of_stud:{ type: Number },
    year: { type: Number },
  },
  filters: {
    timeFormat(i) {
      return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss");
    }
  },
  methods: {
    warnDelete() {
      this.$buefy.dialog.confirm({
        title: "Deleting Course",
        message:
          "Are you sure you want to <b>delete</b> this Course? This action cannot be undone.",
        confirmText: "Delete Course",
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
          "/api/v1/d/dept/" + this.deptObjId + "/course/" + this.objId
        );
        this.$buefy.toast.open({
          duration: 5000,
          message: `Delete Success !`,
          position: "is-bottom",
          type: "is-success"
        });
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          duration: 5000,
          message: `Delete Fail !`,
          position: "is-bottom",
          type: "is-danger"
        });
      } finally {
        this.$emit("onReload");
      }
    }
  }
};
</script>
