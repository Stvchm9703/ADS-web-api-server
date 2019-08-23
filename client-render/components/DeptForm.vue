<template lang="pug">
section.section
  h1.subtitle Department Information
  section
    b-field( label="Department ID")
      b-input( 
        name="dept_id"
        v-model="f_dept_id" 
        required)
    b-field( label="Department Name")
      b-input( 
        name="dept_name"
        v-model="f_dept_name" 
        required)
    b-field( label="Office Location")
      b-input( 
        name="location"
        v-model="f_location" 
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
  name: "DeptForm",
  props: {
    dept_id: String,
    dept_name: String,
    location: String,
    isUpdate: Boolean
  },
  data() {
    return {
      f_dept_id:  this.dept_id ||"",
      f_dept_name: this.dept_name|| "",
      f_location: this.location ||""
    };
  },
  methods: {
    checkValue() {
      if (
        this.f_dept_id == "" ||
        this.f_dept_name == "" ||
        this.f_location == ""
      ) {
        this.$buefy.toast.open({
          message: "Form is not valid! Please check the fields.",
          type: "is-danger",
          position: "is-bottom"
        });
      } else
        this.$emit("onFormSubmit", {
          dept_id: this.f_dept_id,
          dept_name: this.f_dept_name,
          location: this.f_location
        });
    }
  },
  watch: {
    dept_id () {
      this.f_dept_id=this.dept_id
    }, 
    dept_name () {
      this.f_dept_name = this.dept_name
    },
    location () {
      this.f_location = this.location
    }
  },
};
</script>
