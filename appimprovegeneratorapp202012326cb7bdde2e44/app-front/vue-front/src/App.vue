<template>
<v-app id="inspire" style="background-color:white;;">
  <div>
    <ToolbarBis v-if="$route.meta.displayToolbar === 0" />
    <Toolbar v-else-if="$route.meta.displayToolbar === 1" />
    <Playstore v-else-if="$route.meta.displayToolbar === 2" />
    <BottomNavigation v-else-if="$route.meta.displayToolbar === 3"/>
    <DrawerNavigation v-else-if="$route.meta.displayToolbar === 4"/>
  </div>
  <Notification :info="info" :snackbar="snackbar" @closed="onClose()"/> 
  <v-content>
    <transition mode="out-in" name="slide-y-reverse-transition">
      <router-view @notification="onNotification($event)"/>
    </transition>
  </v-content>
</v-app>
</template>

<script>
import Toolbar from "@/components/Toolbar";
import ToolbarBis from "@/components/ToolbarBis"
import BottomNavigation from "@/components/BottomNavigation"
import DrawerNavigation from "@/components/DrawerNavigation"
import Playstore from "@/components/ToolbarPlayStore"
import Notification from "@/components/Notification"
import constants from "@/constants"

export default {
  name: "App",
  components: {
    Playstore,
    Toolbar,
    ToolbarBis,
    DrawerNavigation,
    BottomNavigation,
    Notification
  },
  data() {
    return {
      info: {},
      snackbar: false,
      idNotif: 0
    };
  },
  methods: {
    async getRelevantNotif(id) {
      try {
        const data = (await this.$http.get(`${constants.apiUrl}notifications`))
          .data;
        const notif = data.find(x => x.id_int === id);
        if (notif) this.info = notif;
      } catch (err) {
        throw err;
      }
    },
    onNotification(e) {
      this.snackbar=true;
      if (Number.isInteger(e)) this.getRelevantNotif(e); 
      // si l'id de la notif est recu
      else this.info = e; // si le body de la notif est recu
    },
    async onClose() {
      this.snackbar = false;
      const localNotifParsed = JSON.parse(localStorage.getItem("notification"));
      
      if (localNotifParsed == undefined || localNotifParsed.id_int !== "none") {
        // Reset de la notification en local
        const notification = {
          id_int: "none",
          text_str: "",
          url_str: "",
          activated_int: 0
        };
        localStorage.setItem("notification", JSON.stringify(notification));
      } else {
        // Reset de la notification en bdd
        const deactivate = await this.$http.put(
          `${constants.apiUrl}notifications/${this.idNotif}`,
          {
            activated_int: 0
          }
        );
      }
    }
  },
  mounted() {
    if (!constants.$debug)
      constants.idNotifInterval = setInterval(async () => {
        try {
          const res = await this.$http.get(`${constants.apiUrl}notifications`);
          //On récupère les éventuelles notifications
          const data = res.data.find(x => x.activated_int > 0);
          const localNotif = localStorage.getItem("notification");

          //Dans le cas d'une notification de bdd
          if (data) {
            this.idNotif = data.id;
            if (data.activated_int === 1) this.onNotification(data);
            if (data.activated_int === 2) {
              this.$router.push({ path: data.url_str }, async () => {
                await this.$http.put(
                  `${constants.apiUrl}notifications/${data.id}`,
                  {
                    activated_int: 0
                  }
                );
              });
            }
          } //Dans le cas d'une notification locale
          else if (localNotif !== null) {
            const localNotifParsed = JSON.parse(localNotif);
            if (localNotifParsed.id_int !== "none") {
              this.onNotification(localNotifParsed);
            }
          } //Dans le cas où il n'y a pas de notification
          else {
            console.warn("Pas de notification à montrer !");
          }
        } catch (err) {
          console.error(err);
        }
      }, 3000);
    else {
      console.warn("Mode debug actif, pas d'appels à l'API inutiles");
    }
    // on set le type de la liste dans local storage pour
    // pouvoir se souvenir du choix de l'utilisateur
    localStorage.setItem("listType", constants.listView);
  }
};
</script>

<style media="screen">
#inspire {
overscroll-behavior-y: contain;
font-family: 'Avenir', Helvetica, Arial, sans-serif;
-webkit-font-smoothing: antialiased;
-moz-osx-font-smoothing: grayscale;
text-align: center;
color: #2c3e50;
}
</style>
