<template>
  <v-flex class="justify-center">
    <h3 class="title">Note et avis</h3>
    <v-rating class="mt-2" readonly background-color="dark" hover color="yellow accent-4" half-increments size="18"
      :value="rating"></v-rating>
    <v-list three-line>
      <template v-for="(item, index) in comments">
        <v-subheader v-if="item.header" :key="item.header">{{ item.header }}</v-subheader>
        <v-divider v-else-if="item.divider" :inset="item.inset" :key="index"></v-divider>
        <v-list-tile v-else :key="item.title" avatar>
          <v-list-tile-avatar>
            <img :src="item.avatar">
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title v-html="item.title"></v-list-tile-title>
            <v-list-tile-sub-title v-html="item.subtitle"></v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>
      </template>
    </v-list>
  </v-flex>
</template>

<script>
  export default {
    data: () => {
      return {
        comments: [
          {
            avatar: "https://cdn.vuetifyjs.com/images/lists/1.jpg",
            title: "Meilleure application",
            subtitle:
              "<span class='text--primary'>Ali Connors</span> &mdash; Je recommande"
          },
          { divider: true, inset: true },
          {
            avatar: "https://cdn.vuetifyjs.com/images/lists/2.jpg",
            title: "Rapide",
            subtitle:
              "<span class='text--primary'>John</span> &mdash; Simplement parfait"
          }
        ]
      };
    },
    props: {
      rating: Number
    }
  };
</script>