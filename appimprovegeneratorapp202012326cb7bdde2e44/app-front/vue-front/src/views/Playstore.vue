<template>
  <v-card class="mx-2 mb-2">
    <v-list three-line v-for="item in items" :key="item.title">
      <v-list-tile router :to="item.route" avatar>
        <v-list-tile-avatar size="64" tile>
          <v-img class="mx-4 my-3" :src="item.src"></v-img>
        </v-list-tile-avatar>
        <v-list-tile-content class="mx-3">
          <v-list-tile-title v-html="item.title"></v-list-tile-title>
          <v-list-tile-sub-title v-html="item.sub"></v-list-tile-sub-title>
          <v-list-tile-sub-title class="my-3"><v-icon small>star</v-icon>
          </v-list-tile-sub-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-icon :color="item.active ? 'teal' : 'grey'">more_vert</v-icon>
        </v-list-tile-action>
      </v-list-tile>
      <v-divider v-if="item.title !='Liligo'"></v-divider>
    </v-list>
  </v-card>
</template>

<script>
  export default {
    data() {
      return {
        items: [
          { rating: "3.3", title: "Kayak: vols, hôtels, voitures", sub: "KAYAK.com", src: "https://tooap.com/platform/projects/clients/planes-demo/example-s4s5f6c8-b1a89c8d-ac42-4a68-a6a9-e0e43bafa115-NAME-item3.png" },
          { rating: "5", title: "improvegeneratorapp", sub: "improvegeneratorapp.com", src: "https://tooap.com/platform/projects/clients/planes-demo/example-s4s5f6c8-b36c9242-34c5-49fb-b070-5005b6bc6414-NAME-avatar.jpg", route: "/playstore/install" },
          { rating: "1.3", title: "Skyscanner", sub: "Skyscanner Ltd", src: "https://tooap.com/platform/projects/clients/planes-demo/example-s4s5f6c8-6906a1dd-1dc9-4351-87cc-74f845f6c5a9-NAME-item4.jpg" },
          { rating: "2", title: "JetCost", sub: "JetCost", src: "https://tooap.com/platform/projects/clients/planes-demo/example-s4s5f6c8-7f59272b-f43a-4316-8a0e-edf6838217a1-NAME-item2.jpeg" },
          { rating: "3", title: "Liligo", sub: "Liligo.com", src: "https://tooap.com/platform/projects/clients/planes-demo/example-s4s5f6c8-b1a89c8d-ac42-4a68-a6a9-e0e43bafa115-NAME-item3.png" },
        ],
      };
    }
  };
</script>

<style>
  .clearfix {
    overflow: auto;
  }
</style>