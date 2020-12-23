<!-- This component is used to make a carousel with images -->
<template>
  <v-container pa-0>
    <v-carousel height="40vh" hide-delimiters>
      <v-carousel-item v-for="(img, i) in imgsNotNull" :key="i" :aspect-ratio="4/3" :src="isNull(img)"></v-carousel-item>
    </v-carousel>
  </v-container>
</template>

<script>
  export default {
    computed: {
      imgsNotNull() {
        return this.imgs.filter(img => img !== "https://jetanswer.com/filekitten/null");
      }
    },
    props: {
      imgs: Array
    }
  };
</script>