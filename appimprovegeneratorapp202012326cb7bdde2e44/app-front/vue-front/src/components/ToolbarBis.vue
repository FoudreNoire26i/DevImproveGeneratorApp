<!-- This component is used to realize the toolbar of 
the applications when you are in a page -->
<template>
  <v-toolbar dark app :color="$route.meta.toolbarColor">
    <v-btn icon dark @click="$router.go(-1); $vuetify.goTo(0)">
      <v-icon>arrow_back</v-icon>
    </v-btn>
    <v-toolbar-title>{{$route.name}}</v-toolbar-title>
  </v-toolbar>
</template>