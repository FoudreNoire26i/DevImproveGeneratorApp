<!-- This component is used to build the toolbar of the applications-->
<template>
  <v-layout justify-center>
    <v-flex xs12 sm8 offset-sm>
      <v-toolbar :scroll-threshold="200" scroll-off-screen app tabs dark :color="$route.meta.toolbarColor">
        <v-toolbar-title>improvegeneratorapp</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn @click="$vuetify.goTo(0)" router :to="'/profile'" icon>
          <v-icon>account_circle</v-icon>
        </v-btn>
        <BasketButton icon="shopping_cart"/>
        <v-menu top left transition="slide-x-reverse-transition">
          <v-btn @click="$vuetify.goTo(0)" slot="activator" icon>
            <v-icon>more_vert</v-icon>
          </v-btn>
          <v-list>
            <v-list-tile v-for="(more, i) in mores" :key="i" router :to="more.route">
              <v-list-tile-title>{{ more.title }}</v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-menu>
        <v-tabs slot="extension" slider-color="white" grow color="transparent">
          <v-tab
            v-for="(item, i) in items"
            :key="i"
            @click="$vuetify.goTo(0)"
            router
            :to="item.route"
          >{{ item.title }}</v-tab>
        </v-tabs>
      </v-toolbar>
    </v-flex>
  </v-layout>
</template>

<script>
import BasketButton from "@/components/BasketButton"
   export default {
     components: {
       BasketButton
     },
     data() {
       return {
         mores: [
         { title: 'Paramètres', route: '/settings'},
         { title: 'A propos', route: '/help'}
         ],
         items: [
           {title: "home", route: "/"},
           { title: "bateauids", route: "/bateauids" },
           { title: "placeportids", route: "/placeportids" },
           {title: "qrcode", route: "/qrcode"},
           {title: "visio", route: "/visio"},
           {title: "p1", route: "/p1"},
           {title: "p2", route: "/p2"},
           {title: "p3", route: "/p3"}
         ]
       };
     }
   };
</script>
