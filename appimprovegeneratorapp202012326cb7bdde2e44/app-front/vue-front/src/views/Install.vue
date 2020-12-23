<template>
  <v-content>
    <v-layout align-center column>
      <v-card align-center class="mx-2">
        <v-img :src="img" width="100" height="100" style="float: left;"></v-img>
        <v-card-title primary-title>
          <div>
            <h4 class="title">improvegeneratorapp</h4>
            <p class="subheading green--text text--darken-4">
              <b>improvegeneratorapp.com</b>
            </p>
            <v-chip small>Incontournables</v-chip>
          </div>
          <v-btn
            :disabled="show"
            dark
            color="green darken-3"
            @click="goToHome(); calc(); show = !show"
          >{{ text }}</v-btn>
          <v-progress-circular
            :value="value"
            v-if="show === true"
            indeterminate
            color="green darken-3"
          >{{ value }}</v-progress-circular>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text class="text-xs-center">
          <Comments :rating="4"/>
        </v-card-text>
      </v-card>
    </v-layout>
  </v-content>
</template>

<script>
import Comments from "@/components/Comments";

export default {
  components: {
    Comments
  },
  data() {
    return {
      interval: {},
      value: 0,
      show: false,
      rating: 4.5,
      text: "INSTALLER",
      img:
        "https://tooap.com/platform/projects/clients/planes-demo/example-s4s5f6c8-b36c9242-34c5-49fb-b070-5005b6bc6414-NAME-avatar.jpg"
    };
  },
  beforeDestroy() {
    clearInterval(this.interval);
  },
  methods: {
    goToHome() {
      if (this.text == "OUVRIR") {
        this.$router.push({ name: "mainpage" });
      }
    },
    calc() {
      this.interval = setInterval(() => {
        if (this.value === 100) {
          this.show = false;
          return (this.text = "OUVRIR");
        }
        this.value += 10;
      }, 500);
    }
  }
};
</script>
