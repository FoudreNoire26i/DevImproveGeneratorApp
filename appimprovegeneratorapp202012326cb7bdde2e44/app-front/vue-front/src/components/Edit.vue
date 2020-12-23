<!-- This component is used to edit the pages -->
<template>
  <v-speed-dial fab bottom right v-model="fab" fixed transition="slide-y-reverse-transition">
    <v-btn slot="activator" v-model="fab" :color="$route.meta.toolbarColor" dark fab>
      <v-icon>more_vert</v-icon>
      <v-icon>close</v-icon>
    </v-btn>
    <v-btn @click="$vuetify.goTo(0)" router :to="route" fab dark small color="green">
      <v-icon>edit</v-icon>
    </v-btn>
    <v-btn fab dark small color="red" @click.stop="dialog=true">
      <v-icon>delete</v-icon>
    </v-btn>
    <v-dialog width="500" :value="dialog">
      <v-card>
        <v-card-title class="headline" primary-title>Attention</v-card-title>
        <v-card-text>Voulez vous vraiment supprimer cette page ?</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="success" flat @click="deletePage(); dialog=false">Oui</v-btn>
          <v-btn color="error" @click="dialog=false" flat>Non</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-speed-dial>
</template>

<script>
import constants from "@/constants"

export default {
  props: {
    route: Object,
    apiUrl: String
  },
  data() {
    return {
      fab: false,
      dialog: false
    };
  },
  methods: {
    async deletePage() {
      try {
        const response = await this.$http.delete(`${constants.apiUrl}/${this.apiUrl}/${this.$route.params.id}`);
        this.$router.go(-1);
      } catch (err) {
        throw err;
      }
    }
  }
};
</script>
