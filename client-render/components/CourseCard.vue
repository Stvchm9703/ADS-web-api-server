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
        p level: {{level}}
        p Department : 
          nuxt-link(:to='"/dept/" + deptObjId ') {{deptName}} ({{deptId}})
        p Total Offers Count : {{Offers? Offers.length:0}}
        p Created at : {{createdAt|timeFormat}}
        p Last updated : {{lastUpdated | timeFormat}}
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
  name: "CourseCard",
  props: {
    objId: { type: String },
    deptId: { type: String },
    deptName: { type: String },
    deptObjId: { type: String },
    courseId: { type: String },
    title: { type: String },
    level: { type: Number },
    Offers: { type: Array, default: function() { return []; } },
    createdAt: { type: String },
    lastUpdated: { type: String }
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
