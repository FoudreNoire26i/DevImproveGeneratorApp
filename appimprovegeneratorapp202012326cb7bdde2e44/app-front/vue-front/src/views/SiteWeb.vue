<template>
<div class="contain">
  <div class="bar">
    <div class="bar__right">
       <div class="bar__mobile">
         <v-btn icon>
        <Home/>
         </v-btn>
      </div>
    <div class="bar__pc">
      <v-btn icon>
      <ArrowLeft/>
      </v-btn>
      <v-btn icon>
      <ArrowRight/>
      </v-btn>
      <v-btn icon>
      <Refresh/>
      </v-btn>
      </div>
    </div>
    <div class="bar__middle">
        <input class="bar__search" type="text" placeholder="https://my-new-service.net">  
    </div>
    <div class="bar__left">
      <v-btn icon>
      <More/>
    </v-btn>
    </div>
  </div>
  <div class="nav">
    <div class="nav__bar">
      <div class="nav__logo">
       <!-- Corriger l'export des images <img src="../assets/tooap.png">-->
      </div>
      <div class="nav__menu__mobile">
        <v-btn icon>
        <Menu/>
        </v-btn>
      </div>
      <div class="nav__menu">
          <a class="nav__item" href="#">APPLICATION MOBILE</a>
          <a class="nav__item" href="#">IOT</a>
          <a class="nav__item" href="#">WEB</a>
          <a class="nav__item" href="#">CONTACT</a>
          <a class="nav__item" href="#">FR</a>
      </div>
    </div>
  </div>
  <div class="body__container">
    <div class="body__left">
      <div class="contain__text">
        <h1>My new service</h1>
        <h3>Your sales assistant in your web browser.</h3>
      </div>
      <button @click="$router.push('/playstore/install')"><Play class="box__icon"/>INSTALL</button>
    </div>
    <div class="body__rigth">
    <img class="contain__img" src="https://tooap.com/wp-content/uploads/2019/07/map-application-mobile-tooap3-sml-1.jpg">
    </div>
  </div>
  <div>
    <h1>Plus de 75 000 entreprises font confiance à My new service</h1>
    <div class="logos-list">
    <div class="container" id="testimonial-logos-container">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Brentome%20Consulting@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Brentome%20Consulting@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/bureaux%20%C3%A0%20partager@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/bureaux%20%C3%A0%20partager@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Cleany@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Cleany@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/comet@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/comet@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Cubyn@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Cubyn@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/epicery@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/epicery@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/feed@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/feed@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/happn@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/happn@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Kymono@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Kymono@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/lunchr@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/lunchr@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/maddyness@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/maddyness@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/mon-avocat@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/mon-avocat@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/OuiHelp@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/OuiHelp@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/proprio@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/proprio@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Zenchef@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Zenchef@2x.png" data-was-processed="true">

      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/TooGoodToGo@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/TooGoodToGo@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/trusk@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/trusk@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/ornikar@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/ornikar@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Tanker@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Tanker@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/heetch@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/heetch@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/alan@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/alan@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Urban%20Massage@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Urban%20Massage@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/wavy@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/wavy@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Forest@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Forest@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/livementor@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/livementor@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Stootie@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Stootie@2x.png" data-was-processed="true">
      <img class="lazy loaded" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/appie@2x.png" src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/appie@2x.png" data-was-processed="true">
      <img class="lazy" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Brentome%20Consulting@2x.png">
      <img class="lazy" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/bureaux%20%C3%A0%20partager@2x.png">
      <img class="lazy" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Cleany@2x.png">
      <img class="lazy" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/comet@2x.png">
      <img class="lazy" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/Cubyn@2x.png">
      <img class="lazy" data-src="https://qonto-assets.s3.amazonaws.com/homepage/home/clients/epicery@2x.png">
    </div>
  </div>
  </div>
  <div class="swipper-res">
  <div class="swiper-container">
        <div class="container-card">
          <div class="card-swiper">
            <p>"Très efficace et très bonne réactivité, sans parler du coût. Enfin une solution efficace pour les petits entrepreneurs. Bref un service par des pros pour des pros."</p>
            <div class="author">
              <img class="lazy image" src="https://qonto-assets.s3.amazonaws.com/homepage/home/testimonials/testimonials-photo-ludovic-c-fake.jpg" alt="Photo Ludovic C.">
            <span>
                Ludovic C., commentaire sur le Google Play Store
            </span>
          </div>
        </div>
      </div>
  </div>
   <div class="swiper-container">
        <div class="container-card">
          <div class="card-swiper">
            <p>"Un compte pro facile à ouvrir, et des conseillers (agréables, localisés en France et très compétents) très réactifs pour aplanir les petites inquiétudes ou incompréhensions !"</p>
            <div class="author">
              <img class="lazy image" src="https://qonto-assets.s3.amazonaws.com/homepage/home/testimonials/testimonials-photo-jean-charles-samuelian.jpg" alt="Photo Ludovic C.">
            <span>
               Thomas G. commentaire Facebook
            </span>
          </div>
        </div>
      </div>
  </div>
   <div class="swiper-container">
        <div class="container-card">
          <div class="card-swiper">
            <p>"My new service, c'est 100% en phase avec ce que nous faisons chez Alan : un service simple d'utilisation, une interface vraiment pensée pour le client et un service client ultraréactif !"</p>
            <div class="author">
              <img class="lazy image" src="https://qonto-assets.s3.amazonaws.com/homepage/home/testimonials/testimonials-photo-thomas-g-fake.jpg" alt="Photo Ludovic C.">
            <span>
                Jean-Charles Samuelian,fondateur @Alan
            </span>
          </div>
        </div>
      </div>
  </div>
  </div>
  <div class="container__price">
    <div class="price__title">
      <h1>Simple pricing</h1>
      <h3>Start FREE, get PRO as your business grows</h3>
    </div>
    <div class="price__contain">
      <div class="price__box">
        <div class="price__items">
          <table>
          <tr class="box__type"><td>FREE</td></tr>
          <tr class="box__price"><td>0€ /month /user</td></tr>
          <tr><td>Access to All the features</td></tr>
          <tr><td>Be a part of the community</td></tr>
          <tr><td>20 Contacts + Discussions</td></tr>
          <tr><td>100 emails & documents</td></tr>
          </table>
          <button class="box__button btn1">Contact</button>
        </div>
      </div>
      <div class="price__box">
        <div class="price__items">
          <table>
          <tr class="box__type"><td>PRO</td></tr>
          <tr class="box__price"><td>4€ /month /user</td></tr>
          <tr><td>Access to All the features</td></tr>
          <tr><td>Be a part of the community</td></tr>
          <tr><td>200 Contacts + Discussions</td></tr>
          <tr><td>1000 emails & documents</td></tr>
          <tr><td>+2€ = Extra 400 Contacts + Discussions + 2 000 Emails</td></tr>
          </table>
          <button class="box__button btn2">Contact</button>
        </div>
      </div>
      <div class="price__box">
        <div class="price__items">
          <table>
          <tr class="box__type"><td>ENTREPRISE</td></tr>
          <tr class="box__price"><td>49€ /month /user - min. 10 users</td></tr>
          <tr><td>Access to All the features</td></tr>
          <tr><td>Be a part of the community</td></tr>
          <tr><td>20 000 Contacts + Discussions</td></tr>
          <tr><td>100 000 emails & documents</td></tr>
          <tr><td>Dedicated server instance</td></tr>
          <tr><td>Premium support 24/7</td></tr>
          </table>
          <button class="box__button">Discuss with our sales</button>
        </div>
      </div>
    </div>
  </div>
  <div class="container mobile">
    <div class="container-image-desktop">
        <img class="contimg" src="https://qonto-assets.s3.amazonaws.com/homepage/home/promises/promise-1.jpg" >
    </div>
    <div class="container-text">
      <div class="block-center">
        <h2 class="">Le compte pro idéal, seul ou en équipe</h2>
        <p class="subtitle">Un compte professionnel idéal pour les entreprises de 1 à 250 personnes.</p>
        <p class="paragraph">Profitez d'un outil complet pour gérer les finances de votre entreprise, que vous soyez indépendant ou que vous travailliez en équipe. Les avantages d'un compte courant professionnel complet en ligne ? Des cartes Mastercard business et tous les outils pour simplifier la gestion financière de votre société.</p>
        <a href="https://welcome.qonto.eu/" data-track="website_cta_signup_clicked" data-location="home_benefits_team" class="cta btn-primary-small btn-dark">
           Ouvrir un compte
        </a>
      </div>
    </div>
  </div>
  <!-- FOOTER START -->
<div class="footer">
  <div class="contain">
  <div class="col">
    <h1>Company</h1>
    <ul>
      <li>About</li>
      <li>Mission</li>
      <li>Services</li>
      <li>Social</li>
      <li>Get in touch</li>
    </ul>
  </div>
  <div class="col">
    <h1>Products</h1>
    <ul>
      <li>About</li>
      <li>Mission</li>
      <li>Services</li>
      <li>Social</li>
      <li>Get in touch</li>
    </ul>
  </div>
  <div class="col">
    <h1>Accounts</h1>
    <ul>
      <li>About</li>
      <li>Mission</li>
      <li>Services</li>
      <li>Social</li>
      <li>Get in touch</li>
    </ul>
  </div>
  <div class="col">
    <h1>Resources</h1>
    <ul>
      <li>Webmail</li>
      <li>Redeem code</li>
      <li>WHOIS lookup</li>
      <li>Site map</li>
      <li>Web templates</li>
      <li>Email templates</li>
    </ul>
  </div>
  <div class="col">
    <h1>Support</h1>
    <ul>
      <li>Contact us</li>
      <li>Web chat</li>
      <li>Open ticket</li>
    </ul>
  </div>
  <div class="col social">
    <h1>Social</h1>
    <ul>
      <li><img src="https://svgshare.com/i/5fq.svg" width="32" style="width: 32px;"></li>
      <li><img src="https://svgshare.com/i/5eA.svg" width="32" style="width: 32px;"></li>
      <li><img src="https://svgshare.com/i/5f_.svg" width="32" style="width: 32px;"></li>
    </ul>
  </div>
<div class="clearfix"></div>
</div>
</div>
<!-- END OF FOOTER -->
</div>
</template>

<script>
import ArrowLeft from 'vue-material-design-icons/ArrowLeft'
import ArrowRight from 'vue-material-design-icons/ArrowRight'
import Refresh from 'vue-material-design-icons/Refresh'
import More from 'vue-material-design-icons/DotsVertical'
import Home from 'vue-material-design-icons/Home'
import Menu from 'vue-material-design-icons/Menu'
import Play from 'vue-material-design-icons/GooglePlay'
export default {
  components:{
    ArrowLeft,
    ArrowRight,
    Refresh,
    More,
    Home,
    Menu,
    Play
  }
}
</script>

<style scoped>
.contain{
  margin:0;
  background: #fff;
}
.bar{
  display: flex;
  height: 40px;
  width: 100%;
  padding-top: 15px;
}
.bar__right{
  display: flex;
  align-items: center;
  padding: 10px;
  margin: 10px;
  justify-content: space-around;
  width: 10%;
  min-width: 10%;
}
.bar__pc{
  display:flex;
}
.bar__middle{
  display: flex;
  align-items: center;
  justify-content: center;
  width: 70%;
}
.bar__mobile{
  display: none;
}
.bar__search{
  width: calc(100% - 30px);
  background: rgb(233, 231, 231);
  padding: 5px;
  border-radius: 40px;
  outline: none;
  padding-left: 30px;
}
.bar__left{
  display: flex;
  align-items: center;
  justify-content: flex-end;
  width: 20%;
}
.bar__left > button{
  margin-top: 18px;
}

.nav{
  width: 100%;
  height: 60px;
  margin-top: 10px;
}
.nav__bar{
  height: 100%;
  width: 80%;
  margin-left: auto;
  margin-right: auto;
}
.nav__logo{
  float: left;
}
.nav__menu{
  display: flex;
  height: 100%;
  align-items: center;
  justify-content: flex-end;
}
.nav__menu__mobile{
  display: none;
  height: 100%;
  align-items: center;
  justify-content: flex-end;
}
.nav__item{
  padding: 25px;
  font-size: 12px;
  color: #424242;
  text-decoration: none;
}
.body__container{
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 30px;
}
.body__left{
  margin-bottom: 100px;
  margin-right: 45px;
}
.body__rigth > img{
  height: 400px;
  margin-bottom: 30px;
}
.contain__text{
  text-align: start;
  font-weight: 400;
  padding-bottom: 15px;
  line-height: 64px;
}
.contain__text > h1{
  line-height: 64px;
  color: #2dc86e;
  font-weight: 400;
  padding-bottom: 15px;
  font-size: 60px;
}
.contain__text > h3{
  font-size: 30px;
  line-height: 33px;
}
.body__left > button{
  background: #745cf9;
  line-height: 24px;
  min-width: 200px;
  padding-bottom: 13px;
  padding-top: 13px;
  color: #fff;
  font-weight: 600;
  float: left;
  font-size: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.body__left > button:hover{
  background: #553be4
}
.box__icon{
  padding-right:10px;
  padding-top: 5px;
}
.container__price{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center
}

.price__contain{
  display: flex;
  align-items: center;
  min-width: 1000px;
}
.price__title{
  padding: 15px;
}

.price__box{
  height: 420px;
  width: 100%;
  padding: 25px;
}

.price__items{
  display: flex;
  height: 100%;
  flex-direction: column;
  border: 1px solid #c8bebe;
  font-size: 12px;
}
.btn1{
  margin-top: 85px!important;
}
.btn2{
  background: #e73838!important;
}
.box__price{
  background:  #e6e2e2;
  font-size: 14px;
}
.box__type{
  background: #b8b8b8db;
  color: black;
  font-weight: 600;
}
.box__button{
  height: 40px;
  background: #44a2d0;
  margin: 5px;
  color: white;
  font-weight: 600;
  border-radius: 5px;
}
table {
  border-collapse: collapse;
  line-height: 40px;
}

.logos-list {
    width: 100%;
    max-width: 1048px;
    margin: auto;
    display: flex;
    overflow: hidden;
    position: relative;
    margin-bottom: 80px;
    margin-top: 80px
}

.logos-list::before,
.logos-list::after {
    background: linear-gradient(to right, #fff 0%, rgba(255, 255, 255, 0) 100%);
    content: "";
    height: 100px;
    position: absolute;
    width: 200px;
    z-index: 2
}

 .logos-list::after {
    right: 0;
    top: 0;
    transform: rotateZ(180deg)
}

 .logos-list::before {
    left: 0;
    top: 0
}

@keyframes scroll {
    0% {
        transform: translateX(0)
    }
    100% {
        transform: translateX(calc(-88px * 27))
    }
}

@keyframes scroll-es {
    0% {
        transform: translateX(0)
    }
    100% {
        transform: translateX(calc(-88px * 11))
    }
}

.logos-list.es .container {
    width: calc(88px * 11);
    animation: scroll-es 20s linear infinite
}

.logos-list .container {
    animation: scroll 40s linear infinite;
    display: flex;
    width: calc(88px * 27);
    padding-top: 5px;
    padding-bottom: 5px
}

 .logos-list .container>img {
    width: 56px;
    height: 56px;
    margin-right: 32px;
    flex-shrink: 0;
    border-radius: 50%;
    box-shadow: 0px 0px 5px 3px rgba(0, 0, 0, 0.03)
}


.swiper-container  {
    width: 100vw
}

 .swiper-container  .container-card {
    width: 100vw
}

 
.swiper-container  .container-card .card-swiper {
    display: flex;
    flex-flow: column nowrap;
    width: 320px;
    height: 250px;
    margin: auto;
    margin-bottom: 45px;
    padding: 24px;
    background: white;
    box-shadow: 0 15px 50px 0 rgba(28, 15, 125, 0.05);
    border-radius: 8px
}
 .swiper-container .container-card .card-swiper:not(:last-child) {
    margin-right: 32px
}

 .swiper-container  .container-card .card-swiper p {
    color: #414660;
    margin-bottom: auto;
    font-size: 16px;
    line-height: 24px
}

 .swiper-container  .container-card .card-swiper .author {
    display: flex;
    height: 48px;
    width: 100%
}

 .swiper-container  .container-card .card-swiper .author .image {
    height: 48px;
    width: 48px;
    border-radius: 50%;
    margin-right: 16px
}

.swiper-container .container-card .card-swiper .author span {
    color: #414660;
    font-size: 16px;
    line-height: 24px;
    font-weight: 700
}

.container-text{
  margin-left: 45px;
}
.block-center > h2 {
  font-size: 52px;
  line-height: 64px;
  letter-spacing: 0px;
  font-weight: 100;
  color: #1b1e2b;
}
.block-center .subtitle{
  font-size: 20px;
  line-height: 1.6;
  letter-spacing: 0px;
  color: #82879e;
  font-weight: 300;
  margin-top: 32px!important;
}
.block-center .paragraph{
  font-size: 16px;
  line-height: 24px;
  color: #414660;
  font-weight: 300;
  margin-top: 32px!important;
  margin-bottom: 32px!important;
}
.block-center > a {
  width: 320px;
    border-radius: 4px;
    border-style: none;
    transition: all 0.3s ease;
    background-color: #745cf9;
    padding: 12px 8px 12px 8px;
    text-align: center;
    font-size: 16px;
    line-height: 24px;
    letter-spacing: 0px;
    text-decoration: none;
    color: #ffffff;
}


@media screen and (min-width: 450px) {
  .swipper-res{
  display: flex;
  margin-left: auto;
  margin-right: auto;
  width: 70%;
}
.container-card{
  width: 250px!important;
}
.swiper-container{
  margin-right: 45px;
}
}
.container{
padding: 0px;
margin: 0px;
-webkit-font-smoothing: antialiased;
box-sizing: border-box;
width: 100%;
margin-top: 80px;
margin-bottom: 80px;
display: flex;
align-items: center;
}
/* STYLES SPECIFIC TO FOOTER  */
.footer {
  width: 100%;
  position: relative;
  height: auto;
  background-color: #070617;
}
.footer .col {
  width: 190px;
  height: auto;
  float: left;
  box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  padding: 0px 20px 20px 20px;
}
.footer .col h1 {
  margin: 0;
  padding: 0;
  font-family: inherit;
  font-size: 12px;
  line-height: 17px;
  padding: 20px 0px 5px 0px;
  color: rgba(255,255,255,0.2);
  font-weight: normal;
  text-transform: uppercase;
  letter-spacing: 0.250em;
}
.footer .col ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
}
.footer .col ul li {
  color: #999999;
  font-size: 14px;
  font-family: inherit;
  font-weight: bold;
  padding: 5px 0px 5px 0px;
  cursor: pointer;
  transition: .2s;
  -webkit-transition: .2s;
  -moz-transition: .2s;
}
.social ul li {
  display: inline-block;
  padding-right: 5px !important;
}

.footer .col ul li:hover {
  color: #ffffff;
  transition: .1s;
  -webkit-transition: .1s;
  -moz-transition: .1s;
}
.clearfix {
  clear: both;
}
@media only screen and (min-width: 1280px) {
  .contain {
    margin: 0 auto;
  }
}
@media only screen and (max-width: 1139px) {
  .contain .social {
    width: 1000px;
    display: block;
  }
  .social h1 {
    margin: 0px;
  }
}
@media only screen and (max-width: 950px) {
  .footer .col {
    width: 33%;
  }
  .footer .col h1 {
    font-size: 14px;
  }
  .footer .col ul li {
    font-size: 13px;
  }
}
@media only screen and (max-width: 500px) {
    .footer .col {
      width: 50%;
    }
    .footer .col h1 {
      font-size: 14px;
    }
    .footer .col ul li {
      font-size: 13px;
    }
}
@media only screen and (max-width: 340px) {
  .footer .col {
    width: 100%;
  }
}

@media screen and (max-width: 450px) {
  .bar__mobile{
    display: flex;
  }
  .container.mobile{
  flex-direction: column!important;
}
.container-text{
  margin-left: 0px;
  padding: 15px;
}
.contimg{
  height: 300px;
}
  .bar__pc{
    display: none;
  }
  .nav__menu__mobile{
    display: flex;
  }
  .nav__menu{
    display: none;
  }
  .body__container{
    flex-direction: column;
  }
  .contain__text{
    text-align: center;
  }
  .body__left{       
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-right:0px;
  }
  .body__rigth > img{
  height: 250px;
  margin-bottom: 30px;
}
  .price__contain{
    flex-direction: column;
    max-width: 370px;
    min-width: 200px;
  }
  .btn2{
    margin-top: 45px!important;
  }
}
</style>
