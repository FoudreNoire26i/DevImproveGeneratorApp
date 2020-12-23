<!-- This component is used to carry out a notification -->
<template>
  <v-snackbar
    color="white"
    multi-line
    :timeout="0"
    v-touch="{
      up: () => $emit('closed'),
      left: () => $emit('closed'),
      right: () => $emit('closed')
    }"
    @click="handleSend()"
    v-model="snackbar"
    top
  >
    <v-icon>info</v-icon>
    <span class="body-2 grey--text text--darken-4">{{info.text_str}}</span>
  </v-snackbar>
</template>

<script>
import constants from "@/constants";

export default {
  name: "Notification",
  props: {
    snackbar: Boolean,
    info: Object
  },
  data() {
    return {
      // info: {}
    };
  },
  methods: {
    handleSend() {
      if (this.info.url_str.length > 0 && this.$router.history.current.path != this.info.url_str)
        this.$router.push({ path: this.info.url_str }, () => this.$emit('closed'));
      else{
        this.$emit('closed')
      }
    }
  }
};
</script>
