<template>
   <div>
     <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card style="padding:30px">
            <v-form method="POST" @submit.prevent="onSubmit" lazy-validation>
              <div>
                <v-text-field :label="'text_' | formatLabel" v-model="TextStr" append-icon="create"></v-text-field>
                <v-text-field :label="'user_' | formatLabel" v-model="UserStr" append-icon="create"></v-text-field>
                <v-text-field :label="'chatroom_' | formatLabel" v-model="ChatroomStr" append-icon="create"></v-text-field>
                <v-text-field :label="'url_' | formatLabel" v-model="UrlStr" append-icon="create"></v-text-field>
                <v-text-field :label="'id_' | formatLabel" v-model.number="IDInt" append-icon="create"></v-text-field>
                <v-menu :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px">
                  <v-text-field slot="activator" v-model="localeDate6" :label="'date_' | formatLabel" prepend-icon="event" readonly></v-text-field>
                  <v-date-picker :color="$route.meta.toolbarColor" locale="fr-FR" v-model="DateDate" @input="menu6=false"></v-date-picker>
                </v-menu>
              </div>
              <v-card-actions>
                <v-layout justify-center row wrap>
                  <v-btn color="success" type="submit">Valider</v-btn>
                </v-layout>
              </v-card-actions>
            </v-form>
          </v-card>
        </v-flex>
      </v-layout>
   </div>
</template>

<script>
import constants from "@/constants"
const ownerFile = constants.ownerFile;
const apiUrl = `${constants.apiUrl}interviews/`
const UrlFileDev = constants.urlFileDev;

export default {
  data: () => ({
    TextStr: null,
    UserStr: null,
    ChatroomStr: null,
    UrlStr: null,
    IDInt: null,
    DateDate: null,
    menu6: false,
    Spare1Str: null,
    Spare2Str: null,

  }),
  computed: {
    localeDate6() {
      return this.DateDate ? new Date(this.DateDate).toLocaleDateString('fr-FR') : "";
    },
  },
  mounted() {
    this.$vuetify.goTo(0);
  },
  methods: {
      pickFile(ref) {
        this.$refs[ref].click()
      },
    onSubmit() {
      var fileLink = UrlFileDev;
      this.$http.post(apiUrl, {
        text_str: this.TextStr,
        user_str: this.UserStr,
        chatroom_str: this.ChatroomStr,
        url_str: this.UrlStr,
        id_int: this.IDInt,
        date_date: new Date(this.DateDate).toISOString(),
        spare1_str: this.Spare1Str,
        spare2_str: this.Spare2Str,
      }).then(response => {
        console.log(response.data)
        this.$router.go(-1)
      }).catch(e => {
        throw e;
      });
    },
  }
}
</script>
