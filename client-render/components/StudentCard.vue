<template lang="pug">

section.column.is-one-third
  b-collapse.card(:aria-id='"contentId" + objId' :open='false')
    .card-header(
      slot='trigger' 
      slot-scope='props' 
      role='button' 
      :aria-controls='"contentId" + objId')
      p.card-header-title
        | {{student_name}} | {{student_id}}
      a.card-header-icon
        b-icon(:icon="props.open ? 'menu-down' : 'menu-up'")
    .card-content
      .content
        p Student Name: {{student_name}}
        p Student id : {{student_id}}
        p Day of Birth : {{dob|timeFormat}}
        p Total Enrolled Count : {{enrolled? enrolled.length:0}}
        p Created at : {{createdAt|timeFormat}}
        p Last updated : {{lastUpdated | timeFormat}}
    footer.card-footer
      nuxt-link.card-footer-item(
        :to='"/student/" + objId '
        ) Student Detail

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
    student_name: { type: String },
    student_id: { type: String },
    dob: { type: String },
    enrolled: { type: Array, default: function() { return []; } },
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
        title: "Deleting Student",
        message:
          "Are you sure you want to <b>delete</b> this Student? This action cannot be undone.",
        confirmText: "Delete Course",
        type: "is-danger",
        hasIcon: true,
        onConfirm: () => {
          this.deleteStudent();
        }
      });
    },
    async deleteStudent() {
      let requet_pack = {
        _id: this.objId
      };
      try {
        let ip = await this.$axios.$post(
          "/api/v1/d/student/" , requet_pack
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
