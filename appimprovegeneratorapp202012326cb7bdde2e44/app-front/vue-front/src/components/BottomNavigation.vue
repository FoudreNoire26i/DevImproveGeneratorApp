<!-- This component is used to navigate -->
<template>
  <v-container>
    <v-toolbar app dark :color="$route.meta.toolbarColor">
      <v-spacer></v-spacer>
      <v-toolbar-title>improvegeneratorapp</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-bottom-nav shift app :value="true" fixed>
      <template>
        <v-btn
          v-for="(item, i) in items"
          small
          :color="$route.meta.toolbarColor"
          :key="i"
          router
          :to="{ path: item.route }"
          flat
        >
          <span>{{ item.title | capitalize }}</span>
          <v-icon small>{{ item.icon }}</v-icon>
        </v-btn>
        <v-menu top left transition="slide-x-reverse-transition">
          <v-btn small :color="$route.meta.toolbarColor" flat @click="$vuetify.goTo(0)" slot="activator">
            <span>&nbsp;</span>            
            <v-icon small>more_vert</v-icon>
          </v-btn>
          <v-list>
            <v-list-tile v-for="(more, i) in mores" :key="i" router :to="more.route">
              <v-list-tile-title>{{ more.title }}</v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-menu>
      </template>
    </v-bottom-nav>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    items: [
      { title: "Home", route: "/", icon: "home" },
      { title: "bateauids", route: "/bateauids", icon: "apps" },
      { title: "placeportids", route: "/placeportids", icon: "apps" },
      { title: "comments", route: "/comments", icon: "comment" },
    ],
    mores: [
      { title: 'Paramètres', route: '/settings'},
      { title: 'A propos', route: '/help'}
    ],
  })
};
</script>