<!-- 
Summary  	: 	Button which redirects user to basket
Props 	 	: 	None 
Editable 	: 	Redirection's path, 
							Button's position,  
-->

<template>
  <div>
    <v-btn id="button" @click="$vuetify.goTo(0)" router :to="'/panier'" icon>
      <v-badge v-model="show" color="red" overlap>
        <v-icon>{{icon}}</v-icon>
        <span slot="badge">{{basketLength}}</span>
      </v-badge>
    </v-btn>
  </div>
</template>
<script>
import store from "../Store";
export default {
  props: {
    icon: {
      type: String,
      default: "shopping_cart"
    }
  },
  data() {
    return {
      show: true
    };
  },
  computed: {
    // Get basket length
    basketLength() {
      if (!store.getters.products.length) this.show = false;
      return store.getters.products.length;
    }
  },
  watch: {
    // Allow animation of the badge's arrival when basket's length changes
    basketLength() {
      if (this.basketLength > 0) {
        this.show = false;
        setTimeout(() => (this.show = true), 100);
      }
    }
  }
};
</script>
<style scoped>
#button {
  margin-left: 30px;
}
</style>