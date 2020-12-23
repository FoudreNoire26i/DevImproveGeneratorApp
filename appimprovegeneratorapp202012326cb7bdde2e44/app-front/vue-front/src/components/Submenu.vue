<template>
  <v-container pa-0 ma-0>
    <v-menu :close-on-content-click="false" v-model="menu" nudge-width="200" offset-x>
      <v-chip outline slot="activator" :color="$route.meta.toolbarColor">{{ first }}</v-chip>
      <v-card>
        <v-card-text>
          <v-slider
            :thumb-color="$route.meta.toolbarColor"
            thumb-label="always"
            max="4"
            step="1"
            ticks="always"
            tick-size="2"
          ></v-slider>
        </v-card-text>
        <v-btn flat color="error" @click="menu=false">Cancel</v-btn>
        <v-btn :color="$route.meta.toolbarColor" flat @click="menu=false">Save</v-btn>
      </v-card>
    </v-menu>
    <v-menu :close-on-content-click="false" v-model="date" nudge-width="200" offset-x>
      <v-chip outline slot="activator" :color="$route.meta.toolbarColor">Date</v-chip>
      <v-card>
        <v-card-text class="pa-0 ma-0">
          <v-date-picker class="pa-0 ma-0" :color="$route.meta.toolbarColor"></v-date-picker>
        </v-card-text>
        <v-card-actions class="justify-center">
          <v-btn flat color="error" @click="date=false">Cancel</v-btn>
          <v-btn flat :color="$route.meta.toolbarColor" @click="date=false">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-menu>
    <v-chip outline :color="$route.meta.toolbarColor">Filters</v-chip>
  </v-container>
</template>

<script>
export default {
  props: {
    first: String,
    required: true
  },
  data() {
    return {
      date: false,
      menu: false
    };
  }
};
</script>
