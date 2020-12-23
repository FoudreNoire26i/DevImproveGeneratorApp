<template>
  <v-container fluid pa-0>
    <v-layout align-center column>
      <v-card width="100%" max-width="500" align-center>
        <Carousel v-if="imgs[0]" :imgs="imgs" />
        <v-card-title align-center primary-title>
          <Loading :loading="loading" />
          <v-flex v-if="!loading">
            <h4>{{ info.date_date | toLocale }}</h4>
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
    <Edit :apiUrl="'comments'" :route="{name: 'commentsmodifier'}" />
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
  const apicomment = `${constants.apiUrl}comments/`

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
        return []
      }
    },
    mounted: async function () {
      this.$vuetify.goTo(0);
      try {
        const resource = await this.$http.get(apicomment + this.$route.params.id);
        this.info = resource.data;
        this.loading = false;
      } catch (err) {
        throw err;
      }
    }
  }
</script>