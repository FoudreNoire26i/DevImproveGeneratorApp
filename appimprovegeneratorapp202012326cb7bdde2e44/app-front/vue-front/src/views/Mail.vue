<template>
  <v-container fluid ma-0 pa-0>
    <v-card class="px-2">
      <v-card-title>
        <v-layout column align-start>
          <div class="title">App à tester</div>
          <div class="grey--text text--darken-2 caption">
            {{ new Date("2019-07-22T14:38:01.635026Z").toLocaleString('fr-FR') }}</div>
        </v-layout>
      </v-card-title>
      <v-card-actions>
        <v-chip class="ml-1">Tooap</v-chip>
        <v-spacer></v-spacer>
        <v-btn flat color="orange lighten-1">Détails</v-btn>
      </v-card-actions>
    </v-card>

    <v-card class="px-2" min-height="60vh">
      <v-card-text>
        <div class="text-xs-left">
          Bonjour,
          <br />comment ça va ? J'ai trouvé une app super utile. Clique sur ce
          <router-link to="/playstore/install">lien</router-link>&nbsp; si tu veux la tester.
        </div>
      </v-card-text>
    </v-card>
  </v-container>
</template>
