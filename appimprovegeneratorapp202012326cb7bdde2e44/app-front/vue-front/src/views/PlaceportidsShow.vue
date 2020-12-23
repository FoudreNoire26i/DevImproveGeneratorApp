<template>
  <v-container fluid pa-0>
    <v-layout align-center column>
      <v-card width="100%" max-width="500" align-center>
        <Carousel v-if="imgs[0]" :imgs="imgs" />
        <v-card-title align-center primary-title>
          <Loading :loading="loading" />
          <v-flex v-if="!loading">
            <h3 class="headline">{{ info.nomport_title }}</h3>
            <v-btn outline :color="$route.meta.toolbarColor" router :to="{ name: 'placeportidsmap'}" append>
              <v-icon :color="$route.meta.toolbarColor">location_on</v-icon>&nbsp;carte
            </v-btn>            
            <h4>{{ info.numeroplace_subtitle }}</h4>
            <v-flex v-if="info.presentationvideo_vid">
              <video controls width="90%">
                <source :src="info.presentationvideo_vid">
              </video>
            </v-flex>
            <!-- <v-btn outline :color="$route.meta.toolbarColor" router :to="{ name: 'message.private'}" append>
              <v-icon :color="$route.meta.toolbarColor">chat</v-icon>&nbsp;propriétaire
            </v-btn> -->
            <!-- <Payer v-on="$listeners" :payment="prompt" /> -->
          </v-flex>
        </v-card-title>
        <!--<v-divider></v-divider>
        <v-card-text>
          <Comments :rating="4" />
        </v-card-text>-->
      </v-card>
    </v-layout>
    <Edit :apiUrl="'placeportids'" :route="{name: 'placeportidsmodifier'}" />
  </v-container>
</template>

<style>
a {
  text-decoration: none;
}
</style>

<script>
  import Comments from '@/components/Comments';
  import Carousel from '@/components/Carousel';
  import Edit from "@/components/Edit";
  import Payer from '@/components/Payer';
  import Loading from '@/components/Loading';
  import constants from "@/constants"
  const apiplaceportid = `${constants.apiUrl}placeportids/`

  export default {
    components: {
      Comments,
      Carousel,
      Payer,
      Edit,
      Loading
    },
    data: () => ({
      info: {},
      loading: true,
      prompt: {
        btn: "Réserver ce produit",
        title: "Résumé de votre achat",
        text: ``
      },
    }),
    computed: {
      imgs() {
        return [this.info.photo_img,]
      }
    },
    mounted: async function () {
      this.$vuetify.goTo(0);
      try {
        const resource = await this.$http.get(apiplaceportid + this.$route.params.id);
        this.info = resource.data;
        this.prompt.text = `Votre achat: <i>${ this.info.nomport_title }</i>`;
        this.loading = false;
      } catch (err) {
        throw err;
      }
    }
  }
</script>