import Vue from "vue";
import Vuex from "vuex"
import router from "./router"
import App from "@/App.vue";
import Vuetify from "vuetify";
import * as VueGoogleMaps from "vue2-google-maps";
import axios from "axios";
import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);
Vue.use(Vuex)
Vue.use(VueGoogleMaps, {
  load: {
    key: "AIzaSyBegbbLBgNrS_2mgtqFE4h__a3WFAIcQZQ",
    libraries: "directions"
  }
});
Vue.mixin({
  methods: {
    isNull(img) {
      return img === "https://jetanswer.com/filekitten/null" || img === null
        ? "data:image/gif;base64,R0lGODlhAQABAAD/ACwAAAAAAQABAAACADs%3D" // gif fond blanc de 8bit
        : img;
    }
  }
});
Vue.filter('formatLabel', value => value.slice(0, value.length-1));
Vue.filter('uppercase', value => value.toUpperCase());
Vue.filter('toLocale', value => new Date(value).toLocaleDateString('fr-FR'));
Vue.filter('formatFileName', value => {
  // on récupère la partie significative d'un fichier (après 'NAME')
  if (!value) return '';
  const file = value.toString().match(/NAME-.*/);
  return file ? file[0].substr(5,) : value;
})
Vue.filter('formatLink', value => {
  // on rajoute 'http://' si pas dans le lien
  value = value.toString();
  if (value.includes("http://"))
    return value;
  else return `http://${value}`; 
})
Vue.filter('capitalize', value => {
  if (!value) return '';
  return value[0].toUpperCase() + value.slice(1,);
})

Vue.config.productionTip = false;
Vue.prototype.$id_usr = Math.floor(Math.random() * 10000); // id de l'utilisateur courant
Vue.prototype.$http = axios;
/* eslint-disable */
new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
