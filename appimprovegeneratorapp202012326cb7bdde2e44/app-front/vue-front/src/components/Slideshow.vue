<template>
  <v-container align-center white my-1 pa-0>
    <v-layout column align-center>
      <v-carousel height="30%" hide-delimiters>
        <v-carousel-item v-for="(info, i) in infos" :key="i">
          <v-card class="white pa-0" flat router :to="resource + info.id">
            <v-img :src="isNull(info.img[0])" :aspect-ratio="4/3"></v-img>
            <v-card-title>
              <v-flex justify-center>
                <h3 class="headline grey--text text--darken-4">{{info.title[0]}}</h3>
              </v-flex>
            </v-card-title>
          </v-card>
        </v-carousel-item>
      </v-carousel>
    </v-layout>
  </v-container>
</template>


<script>
export default {
  props: {
    resource: String,
    infos: Array
  }
};
</script>
