<!-- 
Summary  	: 	Display a list of items
Props 	 	: 	.infos    : (Array) list
Editable 	: 	
			
Tips      :   
-->
<template>
  <v-layout>
    <v-flex xs12 sm12>
      <v-card class="pa-0">
        <v-list v-for="(info,i) in infos" :key="i" class="pa-0" three-line>
          <v-list-tile ripple append router @click="itemSelected(info.id)">
            <v-avatar color="blue" >
              <v-img v-if="info.img !== 'Séléctionner une image'" :src="isNull(info.img[0])" width="60%" min-width="48"></v-img>
            </v-avatar>
            <v-list-tile-content class="pa-3">
              <!--Ici remplacer title subtitle par le type voulu dans la liste-->
              <v-list-tile-title>
                {{info.title}}
              </v-list-tile-title>
              <v-list-tile-sub-title>
                {{info.subtitle}}
              </v-list-tile-sub-title>
              <v-list-tile-sub-title>
                {{info.desc}}
              </v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
          <v-divider></v-divider>
        </v-list>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {
    props: {
      infos: Array
    },
    methods: {
      itemSelected(selectedId){
        this.$emit("updateSelectedItemId", selectedId)
      }
    }
  };
</script>