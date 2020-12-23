<template>
   <div>
     <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card style="padding:30px">
            <v-form id="form" method="PUT" @submit.prevent="onSubmit" lazy-validation>
              <div>
                <v-text-field label="text_str" v-model="TextStr" append-icon="create"></v-text-field>
                <v-text-field label="url_str" v-model="UrlStr" append-icon="create"></v-text-field>
                <v-select :items="options" v-model.number="ActivatedInt" label="activated_int" append-icon="create"></v-select>
                <v-text-field label="start_int" v-model.number="StartInt" append-icon="create"></v-text-field>
                <v-text-field label="repeat_int" v-model.number="RepeatInt" append-icon="create"></v-text-field>
                <v-text-field label="id_int" v-model.number="IDInt" append-icon="create"></v-text-field>
                <v-menu v-model="menu9" :close-on-content-click="false" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px">
                  <v-text-field slot="activator" label="date_date" v-model="localeDate" prepend-icon="event" readonly></v-text-field>
                  <v-date-picker :color="$route.meta.toolbarColor" v-model="DateDate" @input="menu9=false"></v-date-picker>
                </v-menu>
              </div>
              <v-card-actions fixed height="auto">
                <v-layout justify-center row wrap>
                  <v-btn outline color="error" @click="deleteNotif()">
                    <v-icon color="error" large>delete_outline</v-icon>  
                  </v-btn>
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
const apiUrl = `${constants.apiUrl}notifications/`

export default {
  data: () => ({
    info: null,
    TextStr: null,
    UrlStr: null,
    ActivatedInt: null,
    StartInt: null,
    RepeatInt: null,
    IDInt: null,
    DateDate: null,
    menu9: false,
    Spare1Str: null,
    Spare2Str: null,
    options: [
      { value: 0, text: "Desactivé" },
      { value: 1, text: "Notification" },
      { value: 2, text: "Changer de page" }
    ],
  }),
  computed: {
    localeDate() {
      return this.DateDate ? new Date(this.DateDate).toLocaleDateString('fr-FR') : "";
    }
  },
  methods: {
    onSubmit() {
      this.$http.put(apiUrl+this.$route.params.id, {
        text_str: this.TextStr,
        url_str: this.UrlStr,
        activated_int: this.ActivatedInt,
        start_int: this.StartInt,
        repeat_int: this.RepeatInt,
        id_int: this.IDInt,
        date_date: new Date(this.DateDate).toISOString(),
        spare1_str: this.Spare1Str,
        spare2_str: this.Spare2Str,
      }).then(response => {
        this.$router.go(-1)
      }).catch(e => {
        throw e;
      });
    },
    async deleteNotif() {
      try {
        await this.$http.delete(apiUrl+this.$route.params.id);
        this.$router.go(-1);
      } catch (error) {
        throw error;
      }
    }
  },
  mounted () {
    clearInterval(constants.idNotifInterval);
    this.$http.get(apiUrl+this.$route.params.id).then(response => {
      this.TextStr = response.data.text_str;
      this.UrlStr = response.data.url_str;
      this.ActivatedInt = response.data.activated_int;
      this.StartInt = response.data.start_int;
      this.RepeatInt = response.data.repeat_int;
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
