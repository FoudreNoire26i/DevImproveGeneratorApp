<template>
   <div>
     <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card style="padding:30px">
            <v-form method="POST" @submit.prevent="onSubmit" lazy-validation>
              <div>
                <v-text-field :label="'nomport_' | formatLabel" v-model="NomportTitle" append-icon="create"></v-text-field>
                <v-text-field :label="'LocalisationGeoLat'" v-model.number="LocalisationGeoLat" ></v-text-field>
                <v-text-field :label="'LocalisationGeoLong'" v-model.number="LocalisationGeoLong" ></v-text-field>
                <v-text-field :label="'numeroplace_' | formatLabel" v-model="NumeroplaceSubtitle" append-icon="create"></v-text-field>
                <v-btn block color="primary" @click='pickFile("vid5")'><v-icon>videocam</v-icon>&nbsp;{{ vidName5 }}</v-btn>
                <input
                  type="file"
                  style="display: none"
                  ref="vid5"
                  accept="video/*"
                  @change="SaveVid5($event)"
                >
                <v-btn block color="primary" @click='pickFile("image6")'><v-icon>image</v-icon>&nbsp;{{ imageName6 }}</v-btn>
                <input
                  type="file"
                  style="display: none"
                  ref="image6"
                  accept="image/*"
                  @change="SavePhoto6($event)"
                >
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
const apiUrl = `${constants.apiUrl}placeportids/`
const UrlFileDev = constants.urlFileDev;

export default {
  data: () => ({
    NomportTitle: null,
    LocalisationGeoLat: null,
    LocalisationGeoLong: null,
    NumeroplaceSubtitle: null,
    vidName5: "Séléctionner une vidéo",
    PresentationvideoVid: null,
    imageName6: "Séléctionner une image",
    PhotoImg: null,
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
  mounted() {
    this.$vuetify.goTo(0);
  },
  methods: {
      pickFile(ref) {
        this.$refs[ref].click()
      },
       SaveVid5(event)
       {
        var vm = this;
        let xhr = new XMLHttpRequest();
        let formData = new FormData();
        let vid = event.target.files[0];
        this.vidName5 = vid.name;
        let owner = ownerFile;

        formData.append("owner", owner);
        formData.append("uploadfile", vid);

        xhr.onreadystatechange = function() {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            var jsonResponse = JSON.parse(xhr.responseText);
            jsonResponse = jsonResponse.FileUrl;
            if (jsonResponse !== null) {
              vm.PresentationvideoVid = UrlFileDev+jsonResponse;
            }
          }
        }
        xhr.open("POST", UrlFileDev+"upload-7b73a1bf-88d1-48dc-96f0-b246be29f860");
        xhr.send(formData);
        },
      SavePhoto6(event)
      {
       var vm = this;
       let xhr = new XMLHttpRequest();
       let formData = new FormData();
       let photo = event.target.files[0];
       this.imageName6 = photo.name;
       let owner = ownerFile;

       formData.append("owner", owner);
       formData.append("uploadfile", photo);

       xhr.onreadystatechange = function() {
       if (xhr.readyState == XMLHttpRequest.DONE) {
           var jsonResponse = JSON.parse(xhr.responseText);
           jsonResponse = jsonResponse.FileUrl;
           if (jsonResponse !== null) {
             vm.PhotoImg = UrlFileDev+jsonResponse;
           }
         }
       }
       xhr.open("POST", UrlFileDev+"upload-7b73a1bf-88d1-48dc-96f0-b246be29f860");
       xhr.send(formData);
       },
    onSubmit() {
      var fileLink = UrlFileDev;
      this.$http.post(apiUrl, {
        nomport_title: this.NomportTitle,
        localisation_geo_lat: this.LocalisationGeoLat,
        localisation_geo_long: this.LocalisationGeoLong,
        numeroplace_subtitle: this.NumeroplaceSubtitle,
        presentationvideo_vid: this.PresentationvideoVid,
        photo_img: this.PhotoImg,
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
        console.log(response.data)
        this.$router.go(-1)
      }).catch(e => {
        throw e;
      });
    },
  }
}
</script>
