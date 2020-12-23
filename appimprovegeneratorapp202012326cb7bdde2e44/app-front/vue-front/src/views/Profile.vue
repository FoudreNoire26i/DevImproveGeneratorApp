<template>
  <v-container pa-0 ma-0 mt-2>
    <v-layout>
      <v-flex xs12 offset-sm3 sm6 text-xs-center>
        <v-avatar size="200" :color="$route.meta.toolbarColor">
          <v-icon size="150" color="white">account_circle</v-icon>
        </v-avatar>
      </v-flex>
    </v-layout>
    <br>
    <v-layout>
      <v-flex xs12 offset-sm3 sm6 text-xs-center>
        <v-card>
          <v-list subheader three-line>
            <v-subheader>User Controls</v-subheader>
            <v-list-tile>
              <v-list-tile-content>
                <v-list-tile-title>Content filtering</v-list-tile-title>
                <v-list-tile-sub-title>Set the content filtering level to restrict appts that can be downloaded
                </v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-divider></v-divider>
            <v-list-tile>
              <v-list-tile-content>
                <v-list-tile-title>Password</v-list-tile-title>
                <v-list-tile-sub-title>Require password for purchase or use password to restrict purchase
                </v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list>
          <v-divider></v-divider>
          <v-list subheader three-line>
            <v-subheader>General</v-subheader>
          </v-list>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>