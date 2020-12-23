<!-- 
Summary  	: 	SearchBar which filters list according to the searched element

Props 	 	: 	.placeholder  : (String) placeholder of the selector
              .attribute   	: (String) attribute on which the filter is apply
              .items        : (Array) list of items to filter
              .init         : (String) initial value of the SearchBar 

Emits     :   .updateItems  : (Array) initial list filtered

Editable 	: 	filter on multiple attributes
			
Tips      :   
-->
<template>
  <div id="search">
    <v-text-field v-model="searched" :label="placeholder" solo></v-text-field>
  </div>
</template>

<script>
export default {
  props: {
    placeholder: {
      type: String,
      default: "Rechercher",
    },
    init: {
      type: String,
      required: false,
      default: "",
    },
    attribute: {
      type: String,
      required: true,
		},
		items: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      searched: "",
    };
  },
  methods: {
    filterItems() {
      return this.items.filter(item => {
        if(this.searched === null || this.searched === "") return true;
        return item[this.attribute][0].toLowerCase().includes(this.searched.toLowerCase());
      }) 
    },
  },
  watch: {
    searched() {
      if(this.searched === undefined) this.searched = "";
      this.$emit('updateItems', this.filterItems(this.searched));
    },
    init() {
      this.searched = this.init;
    }
  },
};
</script>

<style lang="scss" scoped>
#search {
  z-index: 2;
  margin-left: 2%;
  margin-right: 2%;
}
</style>