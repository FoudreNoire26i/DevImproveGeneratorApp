<template>
  <v-parallax height="250" :src="img" light>
    <v-layout align-center column justify-center>
      <h1 class="display-2 font-weight-thin mb-3">improvegeneratorapp</h1>
      <h4 class="subheading">improvegeneratorapp</h4>
    </v-layout>
  </v-parallax>
</template>

<script>
  export default {
    props: {
      img: String
    }
  }
</script>