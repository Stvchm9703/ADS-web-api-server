<template lang="pug">
section.section
  h1.subtitle Course Information
  section
    b-field( label="Course ID")
      b-input( 
        name="course_id"
        v-model="f_course_id" 
        required)
    b-field( label="Course Title")
      b-input( 
        name="title"
        v-model="f_title" 
        required)
    b-field( label="Level")
      b-numberinput( 
        min="1"
        max="13"
        name="level"
        v-model="f_level" 
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
export default {
  name: "CourseForm",
  props: {
    course_id: String,
    title: String,
    level: Number,
    isUpdate: Boolean
  },
  data() {
    return {
      f_course_id: this.course_id || "",
      f_title: this.title || "",
      f_level: this.level || 1
    };
  },
  methods: {
    checkValue() {
      if (this.f_course_id == "" || this.f_title == "" || this.f_level == "") {
        this.$buefy.toast.open({
          message: "Form is not valid! Please check the fields.",
          type: "is-danger",
          position: "is-bottom"
        });
      } else
        this.$emit("onFormSubmit", {
          course_id: this.f_course_id,
          title: this.f_title,
          level: this.f_level
        });
    }
  },
  watch: {
    course_id() {
      this.f_course_id = this.course_id;
    },
    title() {
      this.f_title = this.title;
    },
    level() {
      this.f_level = this.level;
    }
  }
};
</script>
