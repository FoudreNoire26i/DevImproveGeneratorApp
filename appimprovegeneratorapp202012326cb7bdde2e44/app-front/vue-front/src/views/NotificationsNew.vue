<template>
   <div>
     <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card style="padding:30px">
            <v-form id="form" method="POST" @submit.prevent="onSubmit" lazy-validation>
              <div>
                <v-text-field label="text_str" v-model="TextStr" append-icon="create"></v-text-field>
                <v-text-field label="url_str" v-model="UrlStr" append-icon="create"></v-text-field>
                <v-select :items="options" disabled v-model.number="ActivatedInt" label="activated_int" append-icon="create"></v-select>
                <v-text-field label="start_int" v-model.number="StartInt" append-icon="create"></v-text-field>
                <v-text-field label="repeat_int" v-model.number="RepeatInt" append-icon="create"></v-text-field>
                <v-text-field label="id_int" v-model.number="IDInt" append-icon="create"></v-text-field>
                <v-menu v-model="menu9" :close-on-content-click="false" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px">
                  <v-text-field slot="activator" v-model="localeDate" label="date_date" prepend-icon="event" readonly></v-text-field>
                  <v-date-picker :color="$route.meta.toolbarColor" locale="fr-FR" v-model="DateDate" @input="menu9=false"></v-date-picker>
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
const apiUrl = `${constants.apiUrl}notifications/`
const UrlFileDev = constants.urlFileDev;

export default {
  data: () => ({
    CreatedAt: null,
    UpdatedAt: null,
    TextStr: null,
    UrlStr: null,
    ActivatedInt: 0,
    StartInt: null,
    RepeatInt: null,
    IDInt: null,
    DateDate: null,
    menu9: false,
    Spare1Str: null,
    Spare2Str: null,
    options: [
      { value: 0, text: "Désactivé" },
      { value: 1, text: "Notification" },
      { value: 2, text: "Changer de page" }
    ],
  }),
  computed: {
    localeDate() {
      return this.DateDate ? new Date(this.DateDate).toLocaleDateString('fr-FR') : "";
    }
  },
  mounted() {
    this.$vuetify.goTo(0);
    clearInterval(constants.idNotifInterval);
  },
  methods: {
    onSubmit() {
      var fileLink = UrlFileDev;
      this.$http.post(apiUrl, {
        created_at: this.CreatedAt,
        updated_at: this.UpdatedAt,
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
        console.log(response.data)
        this.$router.go(-1)
      }).catch(e => {
        throw e;
      });
    }
  }
};
</script>
