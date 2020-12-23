<!-- This component is used to carry out a loading -->
<template>
  <v-flex v-if="loading">
    <v-progress-circular
    indeterminate
    :color="$route.meta.toolbarColor"
  ></v-progress-circular>
  </v-flex>
</template>

<script>
export default {
  name: "Loading",
  props: {
    loading: Boolean,
    default: true
  }
}
</script>
