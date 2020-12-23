<template>
  <div>
    <BasketButton icon="shopping_cart"/>
    <Ebutton :infos="basketInfos" color="blue" icon="add_shopping_cart" />
  </div>
</template>


<script>
import constants from "@/constants";
import Ebutton from "@/components/Ebutton";
import BasketButton from "@/components/BasketButton"
export default {
  components: {
    Ebutton,
    BasketButton
  },
  data() {
    return {
      basketInfos: {}
    };
  },
  mounted() {
    this.basketInfos = {
      id: 1, // id of the item (required)
      // img: img, // image of the item
      title: `${"info1"} ${"info2"}`, // title of the item
      row1: `${"info3"} ${"info4"}`, // first information row
      row2: `${"info5"} ${"info6"}`, // second information row
      price: "price" // price of this item
    };
  }
};
</script>