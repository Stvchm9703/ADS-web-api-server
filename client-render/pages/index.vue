<template lang="pug">
section.section
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title Welcome To Scope
    //- .hero-foot
    //-   nav.tabs.is-boxed.is-fullwidth
    //-     .container
    //-       ul
    //-         li.is-active
    //-           nuxt-link(:to='"/dept/" + this.objId ')  Overview
    //-         li
    //-           nuxt-link(:to='"/dept/" + this.objId + "/course" ') Course List
  section.section
    .columns
      p.title.is-4 Most Popular Course
    .columns
      b-field.column.is-full(grouped)
        b-taginput(
          placeholder='Search...' 
          type='search' 
          icon='magnify'
          :data="keyWordTag" 
          :allow-new="true"
          v-model='searchQ'
          autocomplete
          expanded)
        p.control
          b-button.is-primary(
            @click="fetchPost"
          ) Search

    .columns.is-multiline
      //- most popular
      card(
        v-for="k in deptList"
          :deptObjId='k.dept_obj_id'
          :deptName='k.dept_name'
          :deptId='k.dept_id'
          :objId='k.course_obj_id'
          :title='k.title'
          :year='k.year'
          :level='k.level'
          :courseId='k.course_id'
          :num_of_stud='k.num_of_stud'
        )
</template>
<script>
import cdw from "lodash/cloneDeep";
import Card from "~/components/CourseOfferCard.vue";
export default {
  data: () => ({
    deptList: [],
    searchQ: [],
    keyWordTag: ["year", "title", "level", "course_id", "dept_id","dept_name"]
  }),
  components: { Card },
  methods: {
    async fetchPost() {
      try {
        let query = `sort={"num_of_stud":"desc"}`;
        // + `year={"in":[2016]}&`
        query += this.breakQuery();
        console.log(query)
        
        let ip = await this.$axios.$get(`/api/v1/vpj/l/dept/?${query}`);
        console.log(ip);
        this.deptList = ip.data;
        // this.keyWordTag = Object.keys(ip.data[0]);
      } catch (e) {
        console.warn(e);
      }
    },
    breakQuery() {
      let y = {};
      this.searchQ.forEach(d => {
        let w = d.split(":");
        if (this.keyWordTag.indexOf(w[0]) > -1) {
          console.log("w", w);
          if (y[w[0] + ""] == null) {
            y[w[0] + ""] = [w[1]];
          } else {
            y[w[0] + ""].push(w[1]);
          }
        }
      });
      let exp=""
      Object.keys(y).forEach((key)=> {
        console.log(key, y[key]);
        exp += `&${key}={"in":[${y[key].join(",")}]}`
      });
      console.log(exp);
      return exp;
    }
  },
  created() {
    this.fetchPost();
  }
};

// {
//   class_size: 40
//   course_id: "scope_23010"
//   course_obj_id: "5d5eb78608386435cc05f519"
//   course_offer_id: "scope_23010_2013"
//   dept_id: "CS"
//   dept_name: "Computer Science"
//   dept_obj_id: "5d5d822e0838640430d2fb77"
//   level: 4
//   location: "Blue Zone"
//   num_of_stud: 2
//   title: "advanced developer rule"
//   year: 2013
//   _id: "5d5ef1fa08386444046cc5bf"
// }
</script>