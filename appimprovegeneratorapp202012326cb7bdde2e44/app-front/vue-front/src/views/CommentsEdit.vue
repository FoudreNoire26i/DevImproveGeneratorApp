<template>
   <div>
     <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card style="padding:30px">
            <v-form method="PUT" @submit.prevent="onSubmit" lazy-validation>
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
              <v-card-actions fixed height="auto">
                <v-layout justify-center row wrap>
                  <v-btn type="submit" color="success">Valider</v-btn>
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
const UrlFileDev = constants.urlFileDev;
const apiUrl = `${constants.apiUrl}comments/`

export default {
  data: () => ({
    info: null,
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
  methods: {
      pickFile(ref) {
        this.$refs[ref].click()
      },
    onSubmit() {
      this.$http.put(apiUrl+this.$route.params.id, {
        text_str: this.TextStr,
        user_str: this.UserStr,
        chatroom_str: this.ChatroomStr,
        url_str: this.UrlStr,
        id_int: this.IDInt,
        date_date: this.DateDate ? new Date(this.DateDate).toISOString() : null,
        spare1_str: this.Spare1Str,
        spare2_str: this.Spare2Str,
      }).then(response => {
        this.$router.go(-1)
      }).catch(e => {
        throw e;
      });
    }
  },
  mounted () {
    this.$http.get(apiUrl+this.$route.params.id).then(response => {
      this.TextStr = response.data.text_str;
      this.UserStr = response.data.user_str;
      this.ChatroomStr = response.data.chatroom_str;
      this.UrlStr = response.data.url_str;
      this.IDInt = response.data.id_int;
      this.DateDate = response.data.date_date.substr(0, 10);
      this.Spare1Str = response.data.spare1_str;
      this.Spare2Str = response.data.spare2_str;
    }).catch(err => {
      throw err;
    })
  }
}
</script>
