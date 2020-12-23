<!-- 
Summary  	: 	Button which add item to the basket
Props 	 	: 	.infos    : (Object) item to add to the basket
							.color 		: (String) color of the button
              .outline  : (Boolean) makes the button outlined 
              .text     : (String) text put in the button
              .icon     : (String) icon put in the button
Editable 	: 	
			
Tips      :   See an example at the bottom of the page
-->

<template>
  <div>
    <v-btn v-if="!IsInBasket" @click="add" :color="color" :outline="outline" :dark="!outline">
      
      <v-icon>{{icon}}</v-icon>
    </v-btn>
    <v-btn v-else :color="color" :outline="outline" :dark="!outline" id="position">
      <v-icon>done</v-icon>
    </v-btn>
  </div>
</template>

<script>
import Vuex from "vuex";
import store from "../Store";
import check from "vue-material-design-icons/CheckCircle";

export default {
  components: {
    check
  },
  store: store,
  props: {
    infos: {
      type: Object,
      required: true
    },
    color: {
      type: String,
      default: "",
      required: false
    },
    outline: {
      type: Boolean,
      default: false,
      required: false
    },
    text: {
      type: String,
      required: false,
      default: ""
    },
    icon: {
      type: String,
      required: false,
      default: "add"
    }
  },
  methods: {
    // Add item in the backet
    add() {
      store.commit("ADD_PRODUCT", this.infos);
    }
  },
  computed: {
    // Check if this item is in the basket
    IsInBasket() {
      const products = store.getters.products;
      for (let index = 0; index < products.length; index++) {
        if (products[index].id === this.infos.id) {
          return true;
        }
      }
      return false;
    }
  }
};
</script>

<style scoped>
.check {
  color: #1a971a;
  padding: 5px;
}
</style>

<!-- 
In HTML : 
  <Ebutton :infos="basketInfos" :color="$route.meta.toolbarColor" />

In Script :
basketInfos = {
    id: id, // id of the item (required)
    img: img, // image of the item
    title: `${info1} ${info2}`, // title of the item
    row1: `${info3} ${info4}`, // first information row
    row2: `${info5} ${info6}`, // second information row
    price: price, // price of this item
  }
 -->