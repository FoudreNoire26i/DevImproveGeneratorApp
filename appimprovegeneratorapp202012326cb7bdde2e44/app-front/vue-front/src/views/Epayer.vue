<template >
<div>
<v-card class="card__box" v-if="!payer">
  <div class="box">
    <h1 >Adresse de livraison</h1>
    <div class="box__item">
      <span><strong> Rue :</strong>  </span>
      <span><strong>Ville:</strong> </span>
      <span><strong>Code postal:</strong> </span>
      <span><strong>Province:</strong> </span>
      <span><strong>Pays: </strong></span>
    </div>
    <h1>Vos produits </h1>
    <div class="box__resum">
      <template v-for="(name,value,i) in products">
        <div :key="i" class="box__item">
         <span></span>
          <v-divider></v-divider>
        </div>
      </template>
      </div>
    </div>
    <h1>Vérification et paiement</h1>
    <div class="box__pay">
      <v-text-field hide-details v-model="NumCart" label="Numéro de la carte" ></v-text-field>
      <v-text-field hide-details v-model="ExpCart"  label="Date expiration" ></v-text-field>
      <v-text-field hide-details v-model="CrypCart"  label="Cryptogramme" ></v-text-field>
    </div>
    <v-btn class="box__btn" @click.prevent="ReqPay">Payer</v-btn>
</v-card>
   <template v-if="payer">
      <v-card>
        <h1>Merci pour votre achat</h1>
        <div class="box__item">
          <span>Votre commande n°45896266 est validé</span>
          <span>Détail paiement :</span>
          <span>Carte n°</span>
          <span>Date d'expiration : </span>
          <v-btn class="box__btn" router :to="{path: '/'}">Home</v-btn>
        </div>
      </v-card>
   </template>
  </div>
</template>


<script>
import store from "../Store"
export default {
  store:store,
  data:()=>({
    payer:false,
    adresse:store.getters.adresse[0],
    products:store.getters.products,
    NumCart:null,
    ExpCart:null,
    CrypCart:null

  }),
  methods:{
    ReqPay:function(){
      this.payer = true
    }
  }
}
</script>

<style scoped>
.card__box{
  margin: 25px;
  padding: 25px;

}
h1{
  font-size: 24px;
  padding: 15px;
  font-weight: 200;
  color: #2E90B5;
}
.box__item{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 5px;
}
.box__resum{
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
}
.box__btn{
  background: #2E90B5!important;
  margin-top: 25px;
  color: #ffff;
}
@media only screen and (min-width: 400px) {
.card__box{
  width: 50%;
  transform: translateX(50%);
  margin: 25px
  }
}
</style>