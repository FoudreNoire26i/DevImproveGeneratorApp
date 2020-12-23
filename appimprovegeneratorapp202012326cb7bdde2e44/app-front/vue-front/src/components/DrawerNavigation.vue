<template>
  <v-flex xs12 sm8 offset-sm>
    <!-- <v-toolbar
      clipped-left
      app
      color="transparent"
      flat
      fixed
    >
      <v-toolbar-side-icon @click.stop="isDrawerActive=!isDrawerActive"></v-toolbar-side-icon>
      <v-toolbar-title>tout</v-toolbar-title>
    </v-toolbar> -->
    <v-btn @click.stop="isDrawerActive=!isDrawerActive" fab fixed depressed color="transparent" left><v-icon>menu</v-icon></v-btn>
    <v-navigation-drawer v-model="isDrawerActive" width="200" absolute temporary>
      <v-list>
        <!-- <v-list-tile>
          <v-list-tile-avatar>
            <img src="https://randomuser.me/api/portraits/men/1.jpg">
          </v-list-tile-avatar>
        </v-list-tile>
        <v-divider></v-divider> -->
        <v-list-tile v-for="(item, i) in items" router :to="item.route" :key="i">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.title | capitalize }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
  </v-flex>
</template>

<script>
export default {
  data: () => ({
    isDrawerActive: false,
    items: [
      { title: "Home", route: "/", icon: "home" },
      { title: "bateauids", route: "/bateauids", icon: "apps" },
      { title: "placeportids", route: "/placeportids", icon: "apps" },
      { title: "comments", route: "/comments", icon: "comment" },
      { title: "settings", route: "/settings", icon: "settings" },
      { title: "help", route: "/help", icon: "help"}
    ]
  })
};
</script>
