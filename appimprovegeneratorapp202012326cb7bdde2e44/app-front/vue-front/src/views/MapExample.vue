<template>
  <Map :items="infos" />
</template>


<script>
import constants from "@/constants";
import Map from "@/components/Map";
export default {
  components: {
    Map
  },
  data() {
    return {
      infos: [
        {
          id: 1,
          coordonnee_geo_lat: 48.8366986,
          coordonnee_geo_long: 2.3116954
        },
        {
          id: 2,
          coordonnee_geo_lat: 48.8646986,
          coordonnee_geo_long: 2.3116954
        }
      ]
    };
  }
};
</script>