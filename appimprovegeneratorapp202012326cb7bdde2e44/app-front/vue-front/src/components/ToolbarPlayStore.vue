<!-- This component is used to create the PlayStore toolbar. -->
<template>
  <v-text-field append-icon="mic" onfocus="this.style.color='#757575'" class="mx-2 my-2" background-color="white"
    color="grey darken-1" label="Google Play" prepend-inner-icon="arrow_back" @click:prepend-inner="$router.go(-1)"
    hide-details solo-inverted></v-text-field>
</template>