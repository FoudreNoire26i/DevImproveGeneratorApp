<template>
  <div>
    <v-alert v-model="alert" type="success" dismissible>{{ log }}</v-alert>
    <v-layout>
      <v-flex>
        <v-subheader class="justify-center subheading font-weight-black">GESTION PROJET</v-subheader>
        <v-btn color="primary" large router :to="{path: '/notifications'}">télécommande</v-btn>
        <v-btn color="primary" large router :to="{name: 'message.global'}">chat global</v-btn>
        <v-divider></v-divider>
        <v-btn router :to="{name: 'playstore'}">playstore</v-btn>
        <v-btn router :to="{path: '/mails/list'}">mails</v-btn>
        <v-btn router :to="{path: '/chat/ami'}">chat ami</v-btn>
        <v-btn router :to="{path: '/facebook'}">facebook</v-btn>        
        <v-btn router :to="{path: '/interviews'}">interviews</v-btn>
        <v-btn router :to="{path: '/comments'}">comments</v-btn>
        <v-expansion-panel>
          <v-expansion-panel-content>
            <template slot="header">DEBUG</template>
            <v-btn router :to="{name: 'bateauidsliste'}">bateauids</v-btn>
            <v-btn router :to="{name: 'placeportidsliste'}">placeportids</v-btn>
            <v-btn router :to="{name: 'mainpage'}">mainpage</v-btn>
            <v-btn router :to="{name: 'profile'}">profile</v-btn>
            <v-btn router :to="{name: 'settings'}">settings</v-btn>
            <v-btn router :to="{name: 'install'}">install</v-btn>
            <v-btn color="error" @click="dialog=true">delete messages</v-btn>
            <v-btn color="error" @click="deleteNotifications()">delete notifications</v-btn>
            <v-dialog v-model="dialog" max-width="290">
              <v-card>
                <v-card-text class="body-2">Voulez-vous supprimer tous les messages de l'API ?</v-card-text>
                <v-card-actions>
                  <v-layout justify-center row>
                    <v-btn flat color="success" @click="deleteMessages">oui</v-btn>
                    <v-btn flat color="error" @click="dialog=false">non</v-btn>
                  </v-layout>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-flex>
    </v-layout>
    <p class="footer">POC created by - TOOAP <img class="logo" src="./../assets/tooapSquare.png"></p>
  </div>
</template>

<script>
  import constants from "@/constants";
  const api = `${constants.apiUrl}notifications/`;
  const chats = `${constants.apiUrl}chats/`;
  const chatsAll = `${constants.apiUrl}chats?per_page=1000`;

  export default {
    data: () => ({
      info: null,
      value: 0,
      time: null,
      id: null,
      interval: {},
      dialog: false,
      log: "",
      alert: false
    }),
    methods: {
      deleteMessages() {
        this.$http
          .get(chatsAll)
          .then(res => {
            res.data.map((x, i, arr) => {
              this.$http
                .delete(chats + x.id)
                .then(res => {
                  this.dialog = false;
                  this.alert = true;
                  this.log = `${arr.length} messages ont été supprimés`;
                })
                .catch(err => {
                  throw err;
                });
            });
            this.dialog = false;
          })
          .catch(err => {
            throw err;
          });
      },
      deleteNotifications() {
        this.$http.get(api).then(res => {
          res.data.map((x, _, arr) => {
            this.$http.delete(`${api}${x.id}`).then(() => {
              this.alert = true;
              this.log = `${arr.length} notifications ont été supprimés`;
            })
          })
        }).then(() => {
          this.alert = false;
        }).catch(err => {
          throw err;
        })
      }
    },
    mounted() {
      this.$http.get(api).then(response => {
        this.info = response.data;
      })
        .catch(error => {
          throw error;
        })
    }
  }
</script>

<style>
.footer {
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  text-align: center;
}
.logo {
  width: 18px;
}
</style>