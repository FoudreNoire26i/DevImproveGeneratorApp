
<template>
<div>
  <div class="nav__container">
    <div>
    <div class="nav__left">
         <v-toolbar-side-icon class="gray--text text--lighten-1"></v-toolbar-side-icon>
      </div>
    </div>
    <div class="nav__middle">
      <v-icon class="nav__bot">search</v-icon>
      <input class="nav__search" type="text" placeholder="Rechercher dans les messages">
      </div>
      <div class="nav__right">
        <div>
          <v-btn class="ma-2" text icon >
             <v-icon color="dark-grey" size="28">help</v-icon>
          </v-btn>
        </div>
        <div>
           <v-btn class="ma-2" text icon>
             <v-icon color="dark-grey" size="24">dialpad</v-icon>
          </v-btn>
        </div>
        <div class="nav__avatar"> <v-avatar size="32" color="grey">
          <span>B</span>
    </v-avatar></div>
        </div>
  </div>
  <div class="body__container">
    <div class="body__right">
      <div class="btn__container">
        <div class="btn__newMessage">
          <p class="btn__para">Nouveau message</p>
        </div>
      </div>
    <div class="menu__container">
       <div class="menu__item">
        <v-icon class="menu__icon" color="dark-grey" size="24">email</v-icon>
        <div>Boîte de réception</div>
      </div>
      <div class="menu__item">
          <v-icon class="menu__icon" color="dark-grey" size="24">star</v-icon>
          <div>Messages suivis</div>
      </div>
      <div class="menu__item">
          <v-icon class="menu__icon" color="dark-grey" size="24">alarm</v-icon>
         <div>En attentes</div>
      </div>
      <div class="menu__item">
        <v-icon class="menu__icon" color="dark-grey" size="24">send</v-icon>
        <div>Messages envoyés</div>
      </div>
    </div>
    </div>
    <router-view></router-view>
  </div>
 </div>
</template>

<style scoped>
.nav__container{
  display: flex;
  background: white;
  height: 64px;
}
.nav__left{
  display: flex;
  flex-direction: row;
  padding: 10px;
}
.nav__bot{
  margin: 10px;
}
.nav__logo{
  width: 100px;
  height: 40px;
  margin: 5px;
}

.nav__avatar{
  margin:15px;
}
.nav__middle{
    display: flex;
    margin-left: 70px;
    background: #f1f3f4;
    height: 48px;
    margin-top: 10px;
    border-radius: 8px;
}
.nav__search{
  width: 700px;
  outline: none;
}

.nav__right{
    display: flex;
    align-items: center;
    width: 600px;
    justify-content: flex-end;
}

.body__container{
  display:flex;
  width: 100%;
  height: 100%;
  background:#fff;
}

.body__right{
  width: 25%;
  height: 600px;
}

.menu__item{
  display: flex;
  cursor: pointer;
  padding: 5px;
  border-radius: 8px;
}
.menu__item:hover{
  background-color: #f1f3f4;
}
.menu__container{
  display: block;
  float: left;
  margin: 25px
}
.menu__icon{
  padding-right: 15px;
}

.btn__container{
    box-shadow: 0 1px 2px 0 rgba(60,64,67,0.302), 0 1px 3px 1px rgba(60,64,67,0.149);
     cursor: pointer;
    float: left;
    align-items: center;
    background-color: #fff;
    border-radius: 24px;
    color: #3c4043;
    display: inline-flex;
    margin-top: 15px;
    margin-left: 15px;
    height: 48px;
    width: 200px;
    padding: 0 24px 0 0;
}
.btn__container:hover{
      box-shadow: 0 1px 3px 0 rgba(60,64,67,0.302), 0 4px 8px 3px rgba(60,64,67,0.149);
    background-color: #fafafb;
}
.btn__newMessage{
  display: flex;
  align-items: center;
}
.btn__para{
  width: 140px;
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 0.5px;
  text-align: start;
  margin-top: 15px;

}
.btn__newMessage::before{
  background-image: url(https://www.gstatic.com/images/icons/material/colored_icons/1x/create_32dp.png);
    background-position: center;
    background-repeat: no-repeat;
    background-size: 32px;
    content: '';
    display: block;
    height: 48px;
    min-width: 56px;
}
</style>
