<template>
  <v-container fluid pa-0 ma-0>
    <template v-if="$route.meta.toolbar === 1">
      <v-toolbar flat app color="white" extended>
        <v-toolbar-side-icon class="orange--text text--lighten-1"></v-toolbar-side-icon>
        <v-toolbar-title>
          <div class="font-weight-black subheading text-xs-left">INBOX</div>
          <!-- <div class="grey--text text--darken-3 caption">example@gmail.com</div> -->
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn small flat color="orange lighten-1" class="caption">Modifier</v-btn>
          <v-btn small flat color="orange lighten-1" class="caption">Trier</v-btn>
        </v-toolbar-items>
        <template slot="extension">
          <v-input append-icon="mic" hide-details class="search">
            <template>
              <input placeholder="Chercher" type="text" />
            </template>
          </v-input>
        </template>
      </v-toolbar>
    </template>
    <template v-else-if="$route.meta.toolbar === 2">
      <v-toolbar scroll-off-screen :scroll-threshold="50" flat app color="white">
        <v-toolbar-side-icon class="orange--text text--lighten-1" @click="$router.go(-1)">
          <v-icon>chevron_left</v-icon>
        </v-toolbar-side-icon>
        <v-spacer></v-spacer>
        <v-btn icon>
          <v-icon color="orange lighten-1">keyboard_arrow_up</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon color="orange lighten-1">keyboard_arrow_down</v-icon>
        </v-btn>
        <v-btn icon router to="/">
          <v-icon color="orange lighten-1">more_vert</v-icon>
        </v-btn>
      </v-toolbar>
    </template>

    <router-view></router-view>

    <template v-if="$route.meta.toolbar === 2">
      <v-bottom-nav app :value="true" fixed>
        <template v-for="item in items">
          <v-btn :key="item.title" light>
            <span>{{ item.title | capitalize }}</span>
            <v-icon small>{{ item.icon }}</v-icon>
          </v-btn>
        </template>
      </v-bottom-nav>
    </template>
  </v-container>
</template>

<script>
  export default {
    name: "Mails",
    data: () => ({
      items: [
        { title: "Répondre", icon: "reply" },
        { title: "Répondre à tous", icon: "reply_all" },
        { title: "Transférer", icon: "forward" },
        { title: "Supprimer", icon: "delete" }
      ]
    })
  };
</script>

<style>
  .v-toolbar__content {
    padding: 0 10px;
  }

  .v-toolbar__extension {
    padding: 0;
  }

  .v-input__slot {
    width: 100%;
    padding: 5px 0;
  }

  .v-input {
    padding: 0;
    background-color: #dfd2d2;
    border-radius: 3px;
    padding: 5px 10px;
  }
</style>
