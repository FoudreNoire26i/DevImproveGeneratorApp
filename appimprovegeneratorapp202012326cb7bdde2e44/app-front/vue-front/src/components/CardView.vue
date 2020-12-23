<!-- This component is used to display the cards -->
<template>
  <v-layout>
    <v-flex xs12 sm6 offset-sm3>
      <v-card
        v-for="(inf,i) in info"
        :key="i"
        append
        @click="$vuetify.goTo(0)"
        router
        :to="inf.id"
        class="pa-0"
      >
        <v-img :src="isNull(inf.img[0])" v-if="inf.img.length > 0 && inf.img[0]" height="300px">
          <!--<template v-slot:placeholder>
            <v-layout fill-height align-center justify-center ma-0>
              <v-progress-circular indeterminate color="grey lighten-2"></v-progress-circular>
            </v-layout>
          </template>-->
        </v-img>
        <v-card-title primary-title>
          <v-flex v-if="inf.interview">
            <div class="headline">{{inf.user}}</div>
            <p>{{inf.text}}</p>
          </v-flex>
          <v-flex v-else>
            <div class="headline">{{inf.title[0]}}</div>
            <span>{{inf.subtitle[0]}}</span>
            <br>
            <span class="grey--text">{{inf.price}}{{inf.cur}}</span>
            <!--<v-rating
              :value="4"
              background-color="dark"
              color="yellow accent-4"
              half-increments
              dense
              size="15"
              v-if="inf.rating"
            ></v-rating>-->
          </v-flex>
        </v-card-title>
        <v-divider inset class="my-2"></v-divider>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  props: {
    info: Array
  }
};
</script>

