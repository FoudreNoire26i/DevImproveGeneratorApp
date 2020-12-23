<template>
  <v-container pa-0 ma-0 fluid>
    <v-layout column>
      <v-flex xs12 sm6>
        <v-btn icon color="grey accent-4" @click="viewType()" flat>
          <v-icon>web</v-icon>
        </v-btn>
        <v-btn icon color="red" flat router :to="{name: 'placeportidsmap'}">
          <v-icon>location_on</v-icon>
        </v-btn>
      </v-flex>

      <!--AJOUTER LES FILTRES ICI-->

      <v-flex xs4 sm4>
        <SearchBarList
          placeholder="Catégorie (1)"
          :items="infos"
          :options="attribut1Opts"
          attribute="title"
          sort
          @updateInfos="updateInfosAttribut1"
        />
      </v-flex>
      <v-flex xs4 sm4>
        <SearchBar
          placeholder="Recherche par mots clés"
          :items="infos"
          attribute="title"
          @updateItems="updateInfosAttribut2"
        />
      </v-flex>
      <v-flex xs4 sm4>
        <SearchBarCategory
          placeholder="Catégorie (2)"
          :items="infosCategory"
          @updateSelected="updateInfosAttribut3"
        />
      </v-flex>

      <!--FIN DES FILTRES ICI-->

      <Loading :loading="loading" />
      <div v-if="!loading">
      <v-flex v-if="view === 1">
        <CardView :info="filteredInfos" />
      </v-flex>
      <v-flex v-else>
        <ListView :infos="filteredInfos" />
      </v-flex>
      </div>
    </v-layout>
    <!--BOUTON ADD-->
    <v-btn v-if="$route.meta.displayToolbar !== 3" fab bottom right dark fixed :color="$route.meta.toolbarColor" router
      :to="{ name: 'placeportidsnew'}">
      <v-icon>add</v-icon>
    </v-btn>
    <v-btn v-else fab fixed depressed class="add" color="transparent" right top router
      :to="{ name: 'placeportidsnew'}">
      <v-icon>add</v-icon>
    </v-btn>
  </v-container>
</template>

<style>
.add {
  top: 0;
  transform: translate(10px, -15px);
}
</style>

<script>
  import constants from "@/constants"
  const apiUrl = `${constants.apiUrl}placeportids/`
  import Submenu from "@/components/Submenu";
  import CardView from "@/components/CardView";
  import ListView from "@/components/ListView";
  import Loading from "@/components/Loading";
  import SearchBarList from "@/components/SearchBarList";
  import SearchBar from "@/components/SearchBar";
  import SearchBarCategory from "@/components/SearchBarCategory";

  export default {
    components: {
      Submenu,
      CardView,
      ListView,
      Loading,
      SearchBarList,
      SearchBar,
      SearchBarCategory
    },
    data() {
      return {
        data: null,
        infos: [],
        loading: true,
        attribut1Opts: [],
        filteredListAttribut1: [],
        filteredListAttribut2: [],
        filteredListAttribut3: [],
        infosCategory: [],
        view: parseInt(localStorage.getItem("listType")),
        nViews: 2,
        
      };
    },
    methods: {
      viewType() {
        if (this.view % this.nViews === 0) this.view = 1;
        else this.view++;
      },
      updateInfosAttribut1(filteredListAttribut1) {
      this.filteredListAttribut1 = filteredListAttribut1;
      },
      updateInfosAttribut2(filteredListAttribut2) {
        this.filteredListAttribut2 = filteredListAttribut2;
      },
      updateInfosAttribut3(selectedItem) {
        if (selectedItem['title'] != "" && selectedItem['title'] != undefined){
            this.filteredListAttribut3 = this.infos.filter(element => {
                if (element['title'] != undefined && element['title'][0] != null ) return selectedItem['title'].trim()==element['title'][0].trim();
                return false
            });
          } else {
            this.filteredListAttribut3 = this.infos
          }
      },
      filterGroup(filteredLists) {
        // Return filteredGroup which i composed of identical items of filteredLists (Array of lists)
        var filteredGroup = filteredLists[0].filter(element => {
          return filteredLists[1].indexOf(element) > -1;
        });

        for (let index = 2; index < filteredLists.length; index++) {
          filteredGroup = filteredLists[index].filter(element => {
            return filteredGroup.indexOf(element) > -1;
          });
        }
        return filteredGroup;
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
      formItemsList(categoryName, categoryTitle, categorySubtitle, allowDuplicates, extra = "") {
        //Get all items
        var itemsMap = new Map();
        var itemArray = [];
        itemsMap.set(categoryName, [{ header: categoryName }]);

        this.infos.forEach(element => {
          itemArray.push({
            title: `${element[categoryTitle]} ${extra}`,
            subtitle: element[categorySubtitle],
            category: categoryName
          });
        });

        //Remove duplicates elements
        if (!allowDuplicates) itemArray = this.removesDuplicates(itemArray);

        itemArray.forEach(element => {
          itemsMap.get(categoryName).push(element);
        });

        //Form list from these items
        for (var [key, value] of itemsMap) {
          value.forEach(element => {
            this.infosCategory.push(element);
          });
          this.infosCategory.push({ divider: true });
        }
      }
    },
    mounted() {
      this.$http
        .get(apiUrl)
        .then(response => {
          this.data = response.data;
          this.infos = this.data.map(d => {
            return {
              id: d.id || "",
              nom: d.nom_str || "", //La variable dans le filtre
              title: [d.nomport_title,] || "",
              subtitle: [d.numeroplace_subtitle,] || "",
              img: [d.photo_img,] || "",
              desc: d.desc || "",
            };
          });
          this.infos.forEach(element => {
          if (this.attribut1Opts.indexOf(element.nom) < 0)
            this.attribut1Opts.push(element.nom);
          });
          // Populate SearchBarCategory
          this.formItemsList("Catégorie", "title", "subtitle", false);
          //Initialize filterdList
          this.filteredListAttribut1 = this.infos;
          this.filteredListAttribut2 = this.infos;
          this.filteredListAttribut3 = this.infos;
        }).then(() => {
          this.loading = false;
        })
        .catch(error => {
          throw error;
        });
    },
    watch: {
      view() {
        localStorage.setItem("listType", this.view);
      }
    },
    computed: {
      filteredInfos() {
        var filteredGroup = this.filterGroup([
          this.filteredListAttribut1,
          this.filteredListAttribut2,
          this.filteredListAttribut3
        ]);
        return filteredGroup;
      }
    }
  };
</script>
