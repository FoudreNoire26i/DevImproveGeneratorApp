<template>
  <div>
    <v-card class="card__box">
    <h1>Renseigner votre adresse de livraison</h1>
    <div class="box__items">
      <v-text-field hide-details v-model="Num" class="box__input" label="N°" ></v-text-field>
      <v-text-field hide-details  v-model="Rue" class="box__input" label="Rue" ></v-text-field>
      <v-text-field hide-details v-model="Ville" class="box__input" label="Ville" ></v-text-field>
      <v-text-field hide-details v-model="Province" class="box__input" label="Province/Etat" ></v-text-field>
      <v-text-field hide-details v-model="Cp" class="box__input" label="Code postal" ></v-text-field>
      <v-text-field hide-details v-model="Pays" class="box__input" label="Pays" ></v-text-field>
    </div>
    <v-btn class="box__btn" @click="add" router :to="{path:'/payer'}">Procédé au paiement </v-btn>
    </v-card>
  </div>
</template>

<script>
import store from "../Store"

export default {
  store:store,
  data:()=>({
    Num: 15,
    Rue: "avenue des champs",
    Ville: "Marseille",
    Cp: 13005,
    Province: "PACA",
    Pays: "France"
  }),
  methods:{
    add:function(){
      let adresse = {
        num:this.Num,
        rue:this.Rue,
        ville:this.Ville,
        cp:this.Cp,
        province:this.Province,
        pays:this.Pays
      }
      store.commit("ADD_ADRESSE",adresse)
    }
  }

}
</script>

<style  scoped>
.card__box{
  padding: 25px;
  margin: 25px;
}
.card__box > h1{
  font-size: 22px;
  color: #2E90B5;
  font-weight: 300;
}
.box__btn{
  margin-top: 25px;
  background: #2E90B5!important;
  color: #fff;
}

@media only screen and (min-width: 400px) {
.card__box{
  width: 50%;
  transform: translateX(50%) translateY(10%);
  margin: 0px
}
}
</style>