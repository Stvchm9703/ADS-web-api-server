<template lang="pug">
.section
  .content.is-1(v-html="$md.render(hello)")
  .footer
    
</template>
<script>
export default {
  name:"docMDRender",
  data: () => ({
    hello: "#Hello World"
  }),
  computed: {
    topic() {
        let a = this.$route.hash ?   this.$route.hash.replace("#","") : ""
        let b = this.$route.hash? a.split("/") : []
        return b
    }
  },
  watch: {
    fhash(){
      this.fetchPost()
    }
  },
  methods: {
    async fetchPost() {
      try {
        let filename = 'info'
        if (this.topic[1]){
            filename = this.topic[0] + "/" +this.topic[1];
        } else if (this.topic[0]){
            filename = this.topic[0] + "/index"
        }
        let ip = await this.$axios.$get("/md/" + filename + ".md");
        console.log(ip);
        this.hello = ip
      } catch (e) {
        console.warn(e);
        this.$router.push({
          path:"/doc"
        })
      }
    }
  },
  created() {
    this.fetchPost()
  }
};
</script>