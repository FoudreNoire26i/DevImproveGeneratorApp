<!-- 
Summary  	: 	SearchBar which proposes list of options  (ordered by categories) and filters list according
              to the selected option or searched element

Props 	 	: 	.placeholder  : (String) placeholder of the selector
              .items        : (Array) list of items to filter
              .init         : (String) initial value of the SearchBar 

Emits     :   .updateSelected  : (Item) Item selected in the category

Editable 	: 	filter on multiple attributes
			
Tips      :   See an example at the bottom of the page
-->
<template>
  <div id="search">
    <v-autocomplete
      v-model="selected"
      :items="items"
      filled
      color="blue-grey lighten-2"
      label="Select"
      :filter="filterObject"
      :placeholder="placeholder"
      solo
      clearable
    >
      <template v-slot:selection="data">
        {{data.item.title}}
      </template>
      <template v-slot:item="data">
        <template v-if="typeof data.item !== 'object'">
          <v-list-tile-content v-text="data.item"></v-list-tile-content>
        </template>
        <template v-else>
          <v-list-tile-content>
            <v-list-tile-title v-html="data.item.title"></v-list-tile-title>
            <template>
              <v-list-tile-sub-title v-html="data.item.subtitle"></v-list-tile-sub-title>
            </template>
          </v-list-tile-content>
          
        </template>
      </template>
    </v-autocomplete>
  </div>
</template>

<script>
export default {
  props: {
    placeholder: {
      type: String,
      default: "Rechercher",
    },
    items: {
      type: Array,
      required: true,
    },
    init: {
      type: String,
      required: false,
      default: "",
    },
  },
  data() {
    return {
      selected: "",
    };
  },
  methods: {
    filterObject(item, queryText, itemText) {
      if (item.header !== undefined || item.divider !== undefined) {
        return true;
      } else {
        return (
          item.title.toLocaleLowerCase().indexOf(queryText.toLocaleLowerCase()) > -1 ||
          item.subtitle.toLocaleLowerCase().indexOf(queryText.toLocaleLowerCase()) > -1
        );
      }
    },
  },
  watch: {
      selected() {
        if(this.selected === undefined) this.selected = {title:"", subtitle:"", category: this.placeholder};
        if(this.selected.title === undefined) this.selected.title = "";
        if(this.selected.subtitle === undefined) this.selected.subtitle = "";
        this.$emit('updateSelected', this.selected);
      },
      init() {
        this.selected = this.items.filter(obj => {return obj.title === this.init})[0]
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

<!-- 
In HTML
  <SearchBarCategory placeholder="Ville" :items="items" @updateItems="updateInfosVille" :init="villeInit"/>

In Script  
.Initilize 'items'
      /*
        categoryName: (String) name of the category. Insert attribute's name you wish to be displayed in category
        title : (String) insert attribute's name you wish to be displayed in title
        subtitle : (String) insert attribute's name you wish to be displayed in subtitle
        allowDuplicates : (Boolean) allows multiple occurence of an element
        avoid : (Object) specify in where (title or subtitle) to not display a specific element in what
        extra : (Object) specify in where (title or subtitle) to display extra next to (title or subtitle) in what
      */
      formItemsList(categoryName, title, subtitle, allowDuplicates, avoid = {where:"title", what:""}, extra = {where:"title", what:""} ) {
        //Get all items
        var itemsMap = new Map();
        var itemArray = [];
        this.infos.forEach(element => {
          if(!itemsMap.has(element.category)) itemsMap.set(element.category, [{header: element.category}]);
          if (element[avoid.where] !== avoid.what) {
            itemArray.push({
                title: (extra.where === "title") ? `${element[title]} ${extra.what}` : element[title],
                subtitle: (extra.where === "subtitle") ? `${element[subtitle]} ${extra.what}` : element[subtitle],
                category: element.category,
            });
          } 
        });

        //Remove duplicates elements
        if(!allowDuplicates)
          itemArray = this.removesDuplicates(itemArray);

        itemArray.forEach(element => {
          itemsMap.get(element.category).push(element);
        });

        //Form list from these items
        for (var [key, value] of itemsMap) {
          value.forEach(element => {
            this.items.push(element);
          });
          this.items.push({ divider: true });
        }
      },

      removesDuplicates(array) {
        let seen = false;
        array.forEach(item => {
          seen = false;
          array = array.filter(function(element) {
            if (!seen && element.title === item.title) {
              seen = true;
              return true;
            } else if (seen && element.title === item.title) {
              return false;
            } else return true;
          });
        });
        return array;
      },

  Example with multiple filters :
  filteredInfos() {
        let filteredGroup = [];
        let stop = false;

        filteredGroup = this.filteredInfosCategory.filter(element => {
          return ((this.filteredInfosCity.indexOf(element) > -1) && (element.title !== "mon_animal") && (element.price <= this.price) && (element.race === this.selectedRace || this.selectedRace === undefined ))
        })
      }
-->