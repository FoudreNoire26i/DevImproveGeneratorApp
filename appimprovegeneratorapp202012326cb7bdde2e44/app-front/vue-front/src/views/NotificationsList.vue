<template>
  <v-container fluid pa-0 ma-0>
    <v-snackbar
      color="white"
      :timeout="2000"
      v-model="hasSentNotif"
      top
    >
      <v-icon>info</v-icon>
      <span class="body-2 grey--text text--darken-4">Notification envoyée</span>
    </v-snackbar>
    <v-layout row align-center>
      <v-flex>
        <v-card class="mb-2 pa-0 pt-2" flat>
        <v-divider></v-divider>
          <v-list class="pa-0" three-line>
            <v-list-tile v-for="(inf,i) in info" :key="i" append :to="`edit/${inf.id}`">
              <v-list-tile-content class="mx-3">
                <v-list-tile-title>{{inf.title}}</v-list-tile-title>
                <v-list-tile-sub-title><span class="font-italic">url: </span>{{inf.url}}</v-list-tile-sub-title>
                <v-spacer></v-spacer>
              </v-list-tile-content>
              <v-list-tile-action>
                <v-btn icon large ripple @click.prevent.stop="sendNotif(inf.id, 1)">
                  <v-icon x-large :color="$route.meta.toolbarColor">notification_important</v-icon>
                </v-btn>
              </v-list-tile-action>
              <!-- espacement -->
              <div style="width: 20px"></div>
              <!---------------->
              <v-list-tile-action>
                <v-btn icon large v-if="inf.url !== null && inf.url !== ''" ripple @click.prevent.stop="sendNotif(inf.id, 2)">
                  <v-icon x-large color="black">mobile_screen_share</v-icon>
                </v-btn>
              </v-list-tile-action>
            </v-list-tile>
          </v-list>
        </v-card>
      </v-flex>
    </v-layout>
    <v-btn v-if="$route.meta.displayToolbar !== 3" fab bottom right dark fixed :color="$route.meta.toolbarColor" router
      :to="{ name: 'notificationsnew' }">
      <v-icon>add</v-icon>
    </v-btn>
    <v-btn v-else fab fixed depressed class="add" dark color="transparent" right top router
      :to="{ name: 'notificationsnew' }">
      <v-icon>add</v-icon>
    </v-btn>
  </v-container>
</template>

<style>
.add {
  top: 0;
  transform: translate(10px, -15px);
}
</style>

<script>
  import constants from "@/constants"
  const apiUrl = `${constants.apiUrl}notifications/`

  export default {
    data() {
      return {
        data: [],
        info: [],
        hasSentNotif: false
      };
    },
    methods: {
      async sendNotif(id, type) {
        try {
          await this.$http.put(`${apiUrl}${id}`, {
            activated_int: type
          });
          this.hasSentNotif = true;
        } catch (err) {
          throw err;
        }
      },
      async deleteNotif(id) {
        try {
          await this.$http.delete(`${apiUrl}${id}`);
        } catch (err) {
          throw err;
        }
      }
    },
    async mounted() {
      clearInterval(constants.idNotifInterval);
      try {
        const response = await this.$http.get(apiUrl);
        this.data = response.data;
        this.info = this.data.map(d => ({ id: d.id, id_notif: d.id_int || "none", title: d.text_str, url: d.url_str, type: d.activated_int }));
      } catch (error) {
        throw error;
      }
    }
  };
</script>