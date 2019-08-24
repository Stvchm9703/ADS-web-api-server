<template lang="pug">
.section
  .content.is-1(v-html="$md.render(hello)")
  .footer
</template>
<script>
export default {
  data: () => ({
    hello: "#Hello World"
  }),
  computed: {
    fhash(){
      return this.$route.hash
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
        let filename = "info"
        if (this.$route.hash){
          filename = this.$route.hash.replace("#","")
        }
        let ip = await this.$axios.$get("/md/" + filename + ".md");
        console.log(ip);
        this.hello = ip
      } catch (e) {
        console.warn(e);
        this.$router.push({
          path:"/"
        })
      }
    }
  },
  created() {
    this.fetchPost()
  }
};
</script>