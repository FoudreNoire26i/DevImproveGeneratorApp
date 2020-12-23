<template>
  <div>
    <v-tabs centered slider-color="blue darken-4">
      <v-tab>
        <v-icon>view_agenda</v-icon>
      </v-tab>
      <v-tab>
        <v-icon>group</v-icon>
      </v-tab>
      <v-tab>
        <v-icon>account_circle</v-icon>
      </v-tab>
      <v-tab>
        <v-icon>ondemand_video</v-icon>
      </v-tab>
      <v-tab>
        <v-icon>notifications</v-icon>
      </v-tab>
      <v-tab>
        <v-badge overlap color="blue darken-4">
          <template slot="badge">&nbsp;</template>
          <v-icon>menu</v-icon>
        </v-badge>
      </v-tab>
    </v-tabs>
    <template v-for="(post, i) in posts">
      <FacebookPost :key="i" :post="post" />
    </template>
  </div>
</template>

<script>
import FacebookPost from "@/views/FacebookPost";

export default {
  components: { FacebookPost },
  data: () => ({
    active: false,
    posts: [
      {
        user: {
          img: "https://randomuser.me/api/portraits/women/51.jpg",
          name: "Emilia Kyllo"
        },
        content: `J'ai découvert cette app super. Voici le
            <a href="/playstore/install">lien</a>.`,
        metrics: {
          likes: 5,
          loves: 2,
          comments: 4
        },
        ellapsed: 45
      },
      {
        user: {
          img: "https://randomuser.me/api/portraits/men/11.jpg",
          name: "Dwayne Fletcher"
        },
        content: `Je suis à Paris en vacances ! <br><img width="100%" height="250px" src="/img/eiffel.jpg" />`,
        metrics: {
          likes: 2,
          loves: 1,
          comments: 1
        },
        ellapsed: 59
      }
    ]
  })
};
</script>

<style>
.v-badge__badge {
  width: 8px;
  height: 8px;
  transform: translate(-8px, 5px);
}
</style>
