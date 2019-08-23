<template lang="pug">
section.section
  h1.subtitle Student Information
  section
    b-field( label="Student ID")
      b-input( 
        name="student_id"
        v-model="f_student_id" 
        required)
    b-field( label="Student Name")
      b-input( 
        name="student_name"
        v-model="f_student_name" 
        required)
    b-field( label="Day Of Birth")
      b-datepicker( 
        name="dob"
        placeholder="Click to select..."
        icon="calendar-today"
        v-model="f_dob" 
        required)
    hr
    b-field(grouped group-multiline v-if="isUpdate")
      b-button(@click="$emit('onCancel')") Cancel
      b-button.is-primary(
        type="submit"
        @click="checkValue"
        ) Update

    b-field(grouped group-multiline v-if="!isUpdate")
      b-button(@click="$emit('onCancel')") Cancel
      b-button.is-primary(
        type="submit"
        @click="checkValue"
        ) Create
    
</template>
<script>
import moment from "moment";
export default {
  name: "DeptForm",
  props: {
    student_id: String,
    student_name: String,
    dob: String,
    isUpdate: Boolean
  },
  data() {
    return {
      f_student_id: this.student_id || "",
      f_student_name: this.student_name || "",
      f_dob: new Date(),
    };
  },
  methods: {
    checkValue() {
      if (
        this.f_student_id == "" ||
        this.f_student_name == "" ||
        this.f_dob == null
      ) {
        this.$buefy.toast.open({
          message: "Form is not valid! Please check the fields.",
          type: "is-danger",
          position: "is-bottom"
        });
      } else
        this.$emit("onFormSubmit", {
          student_id: this.f_student_id,
          student_name: this.f_student_name,
          dob: moment( this.f_dob).format("YYYY-MM-DDTHH:mm:ssZ")
        });
    }
  },
  watch: {
    student_id() {
      this.f_student_id = this.student_id;
    },
    student_name() {
      this.f_student_name = this.student_name;
    },
    dob() {
      this.f_dob = moment(this.dob, "YYYY-MM-DDTHH:mm:ssZ").toDate();
    }
  }
};
</script>
