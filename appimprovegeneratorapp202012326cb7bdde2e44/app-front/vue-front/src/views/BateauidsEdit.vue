<template>
   <div>
     <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card style="padding:30px">
            <v-form method="PUT" @submit.prevent="onSubmit" lazy-validation>
              <div>
                <v-text-field :label="'nombateau_' | formatLabel" v-model="NombateauTitle" append-icon="create"></v-text-field>
                <v-text-field :label="'modele_' | formatLabel" v-model="ModeleSubtitle" append-icon="create"></v-text-field>
                <v-text-field :label="'LocalisationportGeoLat'" v-model.number="LocalisationportGeoLat" ></v-text-field>
                <v-text-field :label="'LocalisationportGeoLong'" v-model.number="LocalisationportGeoLong" ></v-text-field>
                <v-text-field :label="'spare5_' | formatLabel" v-model="Spare5Str" append-icon="create"></v-text-field>
                <v-text-field :label="'spare6_' | formatLabel" v-model="Spare6Str" append-icon="create"></v-text-field>
                <v-text-field :label="'spare7_' | formatLabel" v-model="Spare7Str" append-icon="create"></v-text-field>
                <v-text-field :label="'spare8_' | formatLabel" v-model="Spare8Str" append-icon="create"></v-text-field>
                <v-text-field :label="'spare9_' | formatLabel" v-model="Spare9Str" append-icon="create"></v-text-field>
                <v-text-field :label="'spare10_' | formatLabel" v-model="Spare10Str" append-icon="create"></v-text-field>
                <v-text-field :label="'spare11_' | formatLabel" v-model.number="Spare11Int" append-icon="create"></v-text-field>
                <v-text-field :label="'spare12_' | formatLabel" v-model.number="Spare12Int" append-icon="create"></v-text-field>
                <v-text-field :label="'spare13_' | formatLabel" v-model.number="Spare13Int" append-icon="create"></v-text-field>
                <v-text-field :label="'spare14_' | formatLabel" v-model.number="Spare14Int" append-icon="create"></v-text-field>
                <v-text-field :label="'spare15_' | formatLabel" v-model.number="Spare15Int" append-icon="create"></v-text-field>
                <v-text-field :label="'spare16_' | formatLabel" v-model.number="Spare16Int" append-icon="create"></v-text-field>
                <v-text-field :label="'spare17_' | formatLabel" v-model="Spare17ID" append-icon="create"></v-text-field>
                <v-text-field :label="'spare18_' | formatLabel" v-model="Spare18ID" append-icon="create"></v-text-field>
                <v-text-field :label="'spare19_' | formatLabel" v-model="Spare19ID" append-icon="create"></v-text-field>
                <v-text-field :label="'spare20_' | formatLabel" v-model="Spare20ID" append-icon="create"></v-text-field>
                <v-text-field :label="'spare21_' | formatLabel" v-model="Spare21ID" append-icon="create"></v-text-field>
                <v-text-field :label="'spare22_' | formatLabel" v-model="Spare22ID" append-icon="create"></v-text-field>
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
const apiUrl = `${constants.apiUrl}bateauids/`

export default {
  data: () => ({
    info: null,
    NombateauTitle: null,
    ModeleSubtitle: null,
    LocalisationportGeoLat: null,
    LocalisationportGeoLong: null,
    Spare1Str: null,
    Spare2Str: null,
    Spare3Str: null,
    Spare4Str: null,
    Spare5Str: null,
    Spare6Str: null,
    Spare7Str: null,
    Spare8Str: null,
    Spare9Str: null,
    Spare10Str: null,
    Spare11Int: null,
    Spare12Int: null,
    Spare13Int: null,
    Spare14Int: null,
    Spare15Int: null,
    Spare16Int: null,
    Spare17ID: null,
    Spare18ID: null,
    Spare19ID: null,
    Spare20ID: null,
    Spare21ID: null,
    Spare22ID: null,
  }),
  computed: {
  },
  methods: {
      pickFile(ref) {
        this.$refs[ref].click()
      },
    onSubmit() {
      this.$http.put(apiUrl+this.$route.params.id, {
        nombateau_title: this.NombateauTitle,
        modele_subtitle: this.ModeleSubtitle,
        localisationport_geo_lat: this.LocalisationportGeoLat,
        localisationport_geo_long: this.LocalisationportGeoLong,
        spare1_str: this.Spare1Str,
        spare2_str: this.Spare2Str,
        spare3_str: this.Spare3Str,
        spare4_str: this.Spare4Str,
        spare5_str: this.Spare5Str,
        spare6_str: this.Spare6Str,
        spare7_str: this.Spare7Str,
        spare8_str: this.Spare8Str,
        spare9_str: this.Spare9Str,
        spare10_str: this.Spare10Str,
        spare11_int: this.Spare11Int,
        spare12_int: this.Spare12Int,
        spare13_int: this.Spare13Int,
        spare14_int: this.Spare14Int,
        spare15_int: this.Spare15Int,
        spare16_int: this.Spare16Int,
        spare17_id: this.Spare17ID,
        spare18_id: this.Spare18ID,
        spare19_id: this.Spare19ID,
        spare20_id: this.Spare20ID,
        spare21_id: this.Spare21ID,
        spare22_id: this.Spare22ID,
      }).then(response => {
        this.$router.go(-1)
      }).catch(e => {
        throw e;
      });
    }
  },
  mounted () {
    this.$http.get(apiUrl+this.$route.params.id).then(response => {
      this.NombateauTitle = response.data.nombateau_title;
      this.ModeleSubtitle = response.data.modele_subtitle;
      this.LocalisationportGeoLat = response.data.localisationport_geo_lat;
      this.LocalisationportGeoLong = response.data.localisationport_geo_long;
      this.Spare1Str = response.data.spare1_str;
      this.Spare2Str = response.data.spare2_str;
      this.Spare3Str = response.data.spare3_str;
      this.Spare4Str = response.data.spare4_str;
      this.Spare5Str = response.data.spare5_str;
      this.Spare6Str = response.data.spare6_str;
      this.Spare7Str = response.data.spare7_str;
      this.Spare8Str = response.data.spare8_str;
      this.Spare9Str = response.data.spare9_str;
      this.Spare10Str = response.data.spare10_str;
      this.Spare11Int = response.data.spare11_int;
      this.Spare12Int = response.data.spare12_int;
      this.Spare13Int = response.data.spare13_int;
      this.Spare14Int = response.data.spare14_int;
      this.Spare15Int = response.data.spare15_int;
      this.Spare16Int = response.data.spare16_int;
      this.Spare17ID = response.data.spare17_id;
      this.Spare18ID = response.data.spare18_id;
      this.Spare19ID = response.data.spare19_id;
      this.Spare20ID = response.data.spare20_id;
      this.Spare21ID = response.data.spare21_id;
      this.Spare22ID = response.data.spare22_id;
    }).catch(err => {
      throw err;
    })
  }
}
</script>
