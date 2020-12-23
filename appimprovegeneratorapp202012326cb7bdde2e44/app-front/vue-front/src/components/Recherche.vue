<template>
  <v-container fluid mt-1 class="search">
    <v-layout justify-center>
      <v-flex xs12 md6 sm6>
        <v-combobox solo :search-input.sync="search" :placeholder="placeholder" :items="sortedItems"></v-combobox>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  export default {
    props: {
      items: {
        type: Array,
        required: false,
        default: function () {
          return [];
        }
      },
      placeholder: {
        type: String,
        required: false,
        default: "Où allez-vous ?"
      }
    },
    data: () => ({
      search: ""
    }),
    computed: {
      sortedItems() {
        return this.items.sort();
      }
    }
  };
</script>

<style>
  .search {
    position: fixed;
    z-index: 2;
  }
</style>