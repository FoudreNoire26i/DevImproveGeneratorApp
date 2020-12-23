<template>
  <v-container fluid pa-0>
    <v-layout column>
      <v-card tile class="mt-2">
        <v-list two-line dense>
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <!-- <img src="https://randomuser.me/api/portraits/women/51.jpg" /> -->
              <img :src="post.user.img" />
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>
                <v-layout justify-space-between>
                  <!-- Emilia Kyllo -->
                  {{post.user.name}}
                  <v-icon>more_horiz</v-icon>
                </v-layout>
              </v-list-tile-title>
              <v-list-tile-sub-title>
                {{post.ellapsed}} min &centerdot;
                <v-icon small>public</v-icon>
              </v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
        <v-card-text>
          <div class="text-xs-left" v-html="post.content"></div>
          <v-layout justify-space-between mt-4>
            <span class="font-weight-thin caption">
              {{post.metrics.likes}}
              <v-icon small color="blue darken-4">thumb_up</v-icon>{{post.metrics.loves}}
              <v-icon small color="pink">favorite</v-icon>
            </span>
            <span class="font-weight-thin caption">{{post.metrics.comments}} commentaires &centerdot; 2 partages</span>
          </v-layout>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-layout justify-center>
            <v-btn flat color="grey darken-1">
              <v-icon small class="mx-1">thumb_up</v-icon>
              <span class="caption text-none">J'aime</span>
            </v-btn>
            <v-btn flat color="grey darken-1">
              <v-icon small class="mx-1">comment</v-icon>
              <span class="caption text-none">Commenter</span>
            </v-btn>
            <v-btn flat color="grey darken-1">
              <v-icon small class="mx-1">share</v-icon>
              <span class="caption text-none">Partager</span>
            </v-btn>
          </v-layout>
        </v-card-actions>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "FacebookPost",
  props: {
    post: Object
  }
};
</script>

<style>
</style>
