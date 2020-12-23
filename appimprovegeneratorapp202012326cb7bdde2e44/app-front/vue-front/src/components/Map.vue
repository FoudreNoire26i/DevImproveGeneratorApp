<!-- 
Summary  	: 	Map which displays markers and return selected item id and position
Props 	 	: 	.items        : (Array) items displayed on the map (id and localisation required)
							.selectedItem : (Object) pre-selected item from items array 
              .markersIcon  : (String) icon of markers 
              .startPosition: (Object) position on map at the start (lat: latitude, lng: longitude, zoom: zoom)

Emits     :   .updateSelectedMarker : (Object) contains selected object and its position
              .updateSelectedItem   : (null) reset updateSelectedItem value
Editable 	: 	
			
Tips      :   

-->
<template>
  <div>
    <div>
      <gmap-map
        ref="map"
        :center="center"
        :options="options"
        :zoom="zoom"
        style="width: 100%;height: 93vh"
        @zoom_changed="zoomChanged"
        @center_changed="centerChanged"
      >
        <gmap-marker
          v-for="marker in markers"
          :key="marker.id"
          :position="{
                      lat: parseFloat(marker.coordonnee_geo_lat),
                      lng: parseFloat(marker.coordonnee_geo_long),
                    }"
          :icon="getIcon(marker.id)"
          @click="markerSelected(marker.id)"
        />
      </gmap-map>
    </div>
  </div>
</template>

<style lang="scss" scoped>
</style>

<script>
import constants from "@/constants";
import store from "@/Store";
const apiUrl = `${constants.apiUrl}arbreids/`;

export default {
  components: {},
  props: {
    selectedItem: {
      type: Object,
      default: null
    },
    items: {
      type: Array,
      default: null
    },
    markersIcon: {
      type: String,
      default:
        "https://cdn0.iconfinder.com/data/icons/small-n-flat/24/678111-map-marker-512.png"
    },
    startPosition: {
      type: Object,
      default: () => ({
        lat: 48.8366986,
        lng: 2.3116954,
        zoom: 12.17
      })
    }
  },
  data: () => ({
    data: null,
    markers: [],
    selectedMarker: null,

    // Map
    center: { lat: 0, lng: 0 },
    centerBis: { lat: 0, lng: 0 },
    zoom: 0,
    // options of the map
    options: {
      zoomControl: true,
      mapTypeControl: false,
      scaleControl: false,
      streetViewControl: false,
      rotateControl: false,
      fullscreenControl: false,
      disableDefaultUi: false
      // You can customize map's style by adding :
      // styles : (generate style JSON from https://mapstyle.withgoogle.com/)
    }
  }),
  async mounted() {
    if (this.items) {
      this.markers = this.items.map(d => {
        return {
          id: d.id || "",
          coordonnee_geo_lat: d.coordonnee_geo_lat || "",
          coordonnee_geo_long: d.coordonnee_geo_long || "",
          icon: {
            url: this.markersIcon,
            scaledSize: { width: 32, height: 32 }
          }
        };
      });
    }
    // Initialize map center and zoom
    this.center = { lat: this.startPosition.lat, lng: this.startPosition.lng };
    this.zoom = this.startPosition.zoom;
  },
  methods: {
    markerSelected(selectedId) {
      this.selectedMarker = this.markers.find(
        marker => marker.id === selectedId
      );
      this.$emit("updateSelectedMarker", this.selectedMarker);
      this.zoomOnSelected();
    },
    zoomOnSelected() {
      this.center = this.centerBis;
      setTimeout(() => {
        // Source de beug
        this.center = {
          lat: this.selectedMarker.coordonnee_geo_lat,
          lng: this.selectedMarker.coordonnee_geo_long
        };
      }, 10);
    },
    zoomChanged(value) {
      this.zoom = value;
    },
    centerChanged(value) {
      this.centerBis = value;
    },
    getIcon(markerId) {
      var icon = this.markers.find(marker => marker.id === markerId).icon;
      if (this.selectedMarker)
        if (this.selectedMarker.id === markerId) {
          return { ...icon, scaledSize: { width: 64, height: 64 } };
        }
      return icon;
    }
  },
  computed: {},
  watch: {
    selectedItem() {
      if (this.selectedItem !== null) {
        this.selectedMarker = this.selectedItem;
        this.zoomOnSelected();
        this.$emit("updateSelectedItem", null);
      }
    },
    items() {
      if (this.items) {
        this.markers = this.items.map(d => {
          return {
            id: d.id || "",
            coordonnee_geo_lat: d.coordonnee_geo_lat || "",
            coordonnee_geo_long: d.coordonnee_geo_long || "",
            marker: {
              url: this.markersIcon,
              scaledSize: { width: 32, height: 32 }
            }
          };
        });
      }
    }
  }
};
</script>