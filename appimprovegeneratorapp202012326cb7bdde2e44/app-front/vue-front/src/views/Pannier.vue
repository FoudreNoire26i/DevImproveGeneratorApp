<template>
  <v-container class="card__container" v-if="infos.length > 0">
    <v-card>
      <template v-for="(name,i) in infos">
        <div class="card__info" :key="i">
          <div class="card__item">
            <span></span>
          <v-btn class="card__btnSupp" icon @click="supp(name)"><v-icon>delete</v-icon></v-btn>
          <v-divider></v-divider>
        </div>
        </div>
    </template>
    <div>
      <div class="box__payer">
      <div class="box__item">
        <h1>Sous-total ( <span v-if="article > 1">articles</span> <span v-else>article</span></h1>
      </div>
      <v-btn class="card__btn" router :to="{path: '/adresse'}" >Passer commande</v-btn>
    </div>
    </div>
    </v-card>
  </v-container>
  <v-container v-else>
    <template>
      <v-card class="card__vide">
        <div>
          <h1>Votre panier est vide</h1>
        <removeCard size="45"/>
        </div>
      </v-card>
    </template>
  </v-container>
</template>

<script>
import Vuex from "vuex"
import removeCard from "vue-material-design-icons/CartRemove"
import CardSearch from "vue-material-design-icons/CardSearch"
import store from "../Store"


export default {
  components:{
    removeCard,
    CardSearch
  },
  store:store,
  data:() => ({
    infos:store.getters.products
  }),
  methods:{
    supp(item){
      store.commit("SUPP_PRODUCT",item)
    }
  },computed:{
    article: function(){
      return store.getters.products.length
    }
  }
}
</script>
<style scoped>
.card__btn{
  background: #2E90B5!important;
  color: #fff;
  font-weight: 600;
}
.card__btnSupp{
  background: rgb(179, 31, 31);
  color: #fff;

}
.box__payer{
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 15px;
}
.card__item{
  display: flex;
  align-items: center;
  justify-content: space-around;
  padding: 15px;

}
.card__item > h2{
  color: #464040;
  font-size: 18px;
}

.card__vide{
  height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9c9c9b;
}
.card__btnDet{
  background: #2E90B5;
  color:#fff;
}
.box__item > h1 , span {
  font-size: 18px;
  color: #676767;
  font-style: italic;
}
.box__item > h1  {
  padding-right: 25px;
}

@media only screen and (max-width: 400px) {
  .box__payer{
    flex-direction: column;
  }
  .card__item > span {
  font-size: 14px;
}
}

</style>
