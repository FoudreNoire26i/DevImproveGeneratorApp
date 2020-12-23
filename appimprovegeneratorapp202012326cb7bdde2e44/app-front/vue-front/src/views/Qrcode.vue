<template>
  <v-container>
    <v-alert color="warning" dismissible v-model="showAlert"></v-alert>
    Veuillez scanner un Qr code valide
    <p class="error">{{ error }}</p>

    <qrcode-stream @decode="onDecode" @init="onInit" />
  </v-container>
</template>


<script>
import { QrcodeStream } from "vue-qrcode-reader";
import constants from "@/constants";

const api = `${constants.apiUrl}`;

export default {
  components: { QrcodeStream },

  data() {
    return {
      result: "",
      error: "",
      showAlert: false
    };
  },

  methods: {
    async onDecode(result) {
      this.result = result;
      this.result = this.result.split("/");
      const [resource, id] =
        this.result[0] === "" ? [this.result[1], this.result[2]] : this.result;
      try {
        const response = await this.$http.get(`${api}${resource}/${id}`);
        this.$router.push({ path: `/${resource}/${id}` });
      } catch (err) {
        this.showAlert = true;
      }
    },

    async onInit(promise) {
      try {
        await promise;
      } catch (error) {
        if (error.name === "NotAllowedError") {
          this.error = "ERREUR: vous devez autoriser la caméra";
        } else if (error.name === "NotFoundError") {
          this.error = "ERREUR: pas de caméra";
        } else if (error.name === "NotReadableError") {
          this.error = "ERREUR: caméra utilisée";
        } else if (error.name === "StreamApiNotSupportedError") {
          this.error = "ERREUR: Stream API non supporté";
        }
      }
    }
  }
};
</script>