<!-- This component is used to make a payment -->
<template>
  <v-dialog v-model="paypal" width="500">
    <v-btn slot="activator" :color="$route.meta.toolbarColor" dark>{{payment.btn}}</v-btn>

    <v-card>
      <v-card-title class="headline grey lighten-5 justify-center" primary-title>{{payment.title}}</v-card-title>

      <v-card-text v-html="payment.text"></v-card-text>

      <v-card-actions class="justify-center">
        <v-btn color="primary" raised @click="handlePayment()">Payer avec paypal</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data() {
    return {
      paypal: false
    };
  },
  props: {
    payment: Object,
  },
  methods: {
    handlePayment() {
      this.paypal = false;
      this.$emit("notification", 1);
    }
  }
};
</script>