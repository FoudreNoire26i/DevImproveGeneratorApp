<template>
  <v-container fluid pa-0>
    <v-bottom-sheet v-model="sheet">
      <v-list two-line>
        <v-list-tile avatar>
          <v-list-tile-content>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-bottom-sheet>
    <gmap-map ref="mymap" :center="center" :options="options" :zoom="6" style="width: 100vw;height: 89vh">
      <gmap-marker
        v-for="(item, key) in info"
        :key="key"
        :clickable="true"
        @click="getMarker(item.id)"
      />
      <!--
      <template v-if="shouldLoad">
      <gmap-circle
        :center="user"
        :radius="5000"
        :visible="true"
        :options="{ fillColor: 'red', fillOpacity: .5, strokeColor: 'red', strokeOpacity: .75 }"
      ></gmap-circle>
      <gmap-marker :position="user" />
      </template> -->
    </gmap-map>
  </v-container>
</template>

<script>
import constants from "@/constants";
const apiUrl = `${constants.apiUrl}comments/`;
const geoLoc = false; // true si besoin de geoloc

export default {
  data: () => ({
    info: [],
    sheet: false,
    center: { lat: 46.849378, lng: 2.410893 },
    point: {},
    options: {
      zoomControl: true,
      mapTypeControl: false,
      scaleControl: false,
      streetViewControl: false,
      rotateControl: false,
      fullscreenControl: true,
      disableDefaultUi: false
    },
    shouldLoad: false,
    // si on a besoin de mocker la position de l'utilisateur
    user: {},
    // si on veut utiliser une icone, changer l'url
    icon: { url: 'https://static.thenounproject.com/png/95546-200.png', scaledSize: { width: 40, height: 40 } }
  }),
  methods: {
    getMarker(id) {
      this.sheet = !this.sheet;
      this.point = this.info.find(x => x.id === id);
    },
    calculateDirections(a, b) {
      const directionsService = new google.maps.DirectionsService();
      const directionsDisplay = new google.maps.DirectionsRenderer({
        suppressMarkers: true
      });
      directionsDisplay.setMap(this.$refs.mymap.$mapObject);
      directionsService.route(
        {
          origin: a,
          destination: b,
          travelMode: "DRIVING"
        },
        (res, status) => {
          if (status === "OK") {
            directionsDisplay.setDirections(res);
          } else window.alert("Directions request failed due to " + status);
        }
      );
    }
  },
  async mounted() {
    try {
      this.info = (await this.$http.get(apiUrl)).data;
    } catch (error) {
      throw err;
    }
  },
  created() {
    if ("geolocation" in navigator && geoLoc) {
      navigator.geolocation.getCurrentPosition(
        pos => {
          const {
            coords: { latitude: lat },
            coords: { longitude: lng }
          } = pos;
          this.user = { lat, lng };
          this.shouldLoad = true;
        },
        () => {
          this.user = { lat: 43.294111, lng: 5.373739 };
        }
      );
    } else {
      this.user = { lat: 43.294111, lng: 5.373739 };
      this.shouldLoad = true;
    }
  }
};
</script>