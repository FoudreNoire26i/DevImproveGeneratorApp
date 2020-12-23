<!-- 
Summary  	: 	SearchBar which proposes list of options and filters list according
              to the selected option or searched element
Props 	 	: 	.placeholder  : (String) placeholder of the selector
              .options      : (Array) different selectable items
              .position   	: (String) css position of the component
              .attribute   	: (String) attribute on which the filter is apply
              .items        : (Array) list of items to filter
              .init         : (String) initial value of the SearchBar 
Emits     :   .updateInfos  : (Array) initial list filtered
Editable 	: 	filter on multiple attributes
			
Tips      :   
-->
<template>
  <div id="search">
    <v-combobox 
      solo 
      :search-input.sync="search" 
      :label="placeholder" 
      :items="defaultOptions"
      v-model="selected"
      :disabled="disabled"
      clearable
      hide-details
    ></v-combobox> 
  </div>
</template>

<script>
export default {
  props: {
    options: {
      type: Array,
      required: false,
      default: function() {
        return [];
      }
    },
    placeholder: {
      type: String,
      required: false,
      default: "Où allez-vous ?"
    },
    position: {
      type: String,
      default: "static"
    },
    attribute: {
      type: String,
      required: true
    },
    items: {
      type: Array,
      required: true
    },
    init: {
      type: String,
      required: false,
      default: ""
    },
    disabled: {
      type: Boolean,
      required: false,
      default: false
    },
    sort: {
      type: Boolean,
      required: false,
      default: false
    },
    rules: {
      type: String,
      required: false,
      default: null
    }
  },
  data: () => ({
    search: "",
    selected: ""
  }),
  computed: {
    sortedItems() {
      if (this.sort) return this.options.sort();
      return this.options;
    },
    defaultOptions(){ //----------changer sorted pour default
      let isAlreadyNull=false
      if ( this.options.length > 0){
        let itemsAttributes = []
        this.items.forEach((element, i) => {
          if (element[this.attribute][0] !== null && element[this.attribute][0] !== "" && element[this.attribute][0] !== undefined)
            itemsAttributes.push(element[this.attribute][0])
           else if (!isAlreadyNull) {
             isAlreadyNull= true
             itemsAttributes.push("Non renseigné")
             this.items[i][this.attribute] = ["Non renseigné"]
           } else {
             this.items[i][this.attribute] = ["Non renseigné"]
           }
        });
        return itemsAttributes
      } else return this.options
    }
  },
  mounted() {
    document.getElementById("search").style.position = this.position;
  },
  watch: {
    selected() {
      if (this.selected === undefined) this.selected = "";
      if (this.selected !== null)
        this.$emit(
          "updateInfos",
          this.filterItems(this.selected),
          this.selected
        );
    },
    init() {
      this.selected = this.init;
    }
  },
  methods: {
    filterItems(selected) {
      return this.items.filter(item => {
        if (selected === null || selected === "") return true;
        return (
          item[this.attribute][0].toLowerCase().includes(selected.toLowerCase()) ||
          this.evalInContext(this.rules, {
            attribute: this.attribute,
            selected: selected,
            item: item,
            options: this.options,
          })
        );
      });
    },
    evalInContext(js, context) {
      //# Return the results of the in-line anonymous function we .call with the passed context
      return function() {
        return eval(js);
      }.call(context);
    }
  }
};
</script>

<style>
  #search {
    z-index: 2;
    margin: 2%;
  }
</style>