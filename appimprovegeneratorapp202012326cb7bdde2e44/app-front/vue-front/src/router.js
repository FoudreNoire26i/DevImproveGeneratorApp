import Vue from "vue";
import Router from "vue-router";
import { menuLayout } from "./constants";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "*",
      component: () => import("@/views/Home.vue"),
      meta: { displayToolbar: menuLayout.main, toolbarColor: "#0000FF" }
    },
    {
      path: "/",
      name: "mainpage",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Home.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor:"#0000FF", views: true }
    },
    {
      path: "/profile",
      name: "profile",
      component: () =>
      import("@/views/Profile.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/mails",
      component: () => import("@/views/Mails.vue"),
      children: [
        {
          path: "list",
          component: () => import("@/views/MailList.vue"),
          meta: { toolbar: 1 }
        },
        {
          path: "list/:id",
          component: () => import("@/views/Mail.vue"),
          meta: { toolbar: 2 }
        }
      ]
    },
    {
      path: "/gmails",
      component: () => import("@/views/gmail.vue"),
      children: [
        {
          path: "list",
          component: () => import("@/views/gmailList.vue")
        }
      ]
    },
    {
      path:"/site",
      name:"site",
      component: () => import("@/views/SiteWeb.vue")
    },
    {
      path:"/panier",
      name:"Mon panier",
      component: () => import("@/views/Pannier.vue"),
      meta: { displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/adresse",
      name:"Votre adresse",
      component: () => import("@/views/Eadresse.vue"),
      meta:{ displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF", views: true}
    },
    {
      path: "/payer",
      name:"Paiement",
      component: () => import("@/views/Epayer.vue"),
      meta:{ displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF", views: true}
    },
    {
      path: "/p1",
      name: "p1",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Page1.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/p2",
      name: "p2",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Page2.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/p3",
      name: "p3",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Page3.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/p4",
      name: "p4",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Page4.vue"),
      meta: {displayToolbar: 0, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/p5",
      name: "p5",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Page5.vue"),
      meta: {displayToolbar: 0, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/p6",
      name: "p6",
      component: () =>
      import(/* webpackChunkName: "about" */ "@/views/Page6.vue"),
      meta: {displayToolbar: 0, views: true }
    },
    {
      path: "/chat/private",
      name: "message.private",
      component: () => import("@/components/Chat.vue"),
      props: { room: "private" },
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/chat/global",
      name: "message.global",
      component: () => import("@/components/Chat.vue"),
      props: { room: "global" },
      meta: { displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF", views: true }
    },
    {
      path: "/chat/ami",
      name: "Ami",
      component: () => import("@/components/Chat.vue"),
      props: {
        room: "global",
        defaultFeed: [
          { id: 1, content: "Salut" },
          {
            id: 1,
            content: `J'ai trouvé cette app super utile. Voici le <a href="/playstore/install">lien</a>.`
          }
        ],
        mock: true
      },
      meta: {
        displayToolbar: menuLayout.main,
        toolbarColor: "#0000FF"
      }
    },
    {
      path: "/facebook",
      name: "facebook",
      component: () => import("@/views/Facebook.vue")
    },
    {
      path: "/playstore",
      name: "playstore",
      component: () => import("@/views/Playstore.vue"),
      meta: {displayToolbar: 2, toolbarColor: "green darken-3"}
    },
    {
      path: "/playstore/install",
      name: "install",
      component: () => import("@/views/Install.vue"),
      meta: {displayToolbar: 2}
    },
    {
      path: "/settings",
      name: "settings",
      component: () => import("@/views/Settings.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/bateauids",
      name: "bateauidsliste",
      component: () => import("@/views/BateauidsList.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF" },
    },
    {
      path: "/bateauids/new",
      name: "bateauidsnew",
      component: () => import("@/views/BateauidsNew.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/bateauids/map",
      name: "bateauidsmap",
      component: () => import("@/views/BateauidsMap.vue"),
      meta: {displayToolbar: menuLayout.secondary ,map: true , toolbarColor: "#0000FF" },
    },
    {
      path: "/bateauids/:id",
      name: "bateauidsdetail",
      component: () => import("@/views/BateauidsShow.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/bateauids/:id/edit",
      name: "bateauidsmodifier",
      component: () => import("@/views/BateauidsEdit.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/comments",
      name: "commentsliste",
      component: () => import("@/views/CommentsList.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF" },
    },
    {
      path: "/comments/new",
      name: "commentsnew",
      component: () => import("@/views/CommentsNew.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/comments/map",
      name: "commentsmap",
      component: () => import("@/views/CommentsMap.vue"),
      meta: {displayToolbar: menuLayout.secondary ,map: true , toolbarColor: "#0000FF" },
    },
    {
      path: "/comments/:id",
      name: "commentsdetail",
      component: () => import("@/views/CommentsShow.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/comments/:id/edit",
      name: "commentsmodifier",
      component: () => import("@/views/CommentsEdit.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/interviews",
      name: "interviewsliste",
      component: () => import("@/views/InterviewsList.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF" },
    },
    {
      path: "/interviews/new",
      name: "interviewsnew",
      component: () => import("@/views/InterviewsNew.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/interviews/map",
      name: "interviewsmap",
      component: () => import("@/views/InterviewsMap.vue"),
      meta: {displayToolbar: menuLayout.secondary ,map: true , toolbarColor: "#0000FF" },
    },
    {
      path: "/interviews/:id",
      name: "interviewsdetail",
      component: () => import("@/views/InterviewsShow.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/interviews/:id/edit",
      name: "interviewsmodifier",
      component: () => import("@/views/InterviewsEdit.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/placeportids",
      name: "placeportidsliste",
      component: () => import("@/views/PlaceportidsList.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF" },
    },
    {
      path: "/placeportids/new",
      name: "placeportidsnew",
      component: () => import("@/views/PlaceportidsNew.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/placeportids/map",
      name: "placeportidsmap",
      component: () => import("@/views/PlaceportidsMap.vue"),
      meta: {displayToolbar: menuLayout.secondary ,map: true , toolbarColor: "#0000FF" },
    },
    {
      path: "/placeportids/:id",
      name: "placeportidsdetail",
      component: () => import("@/views/PlaceportidsShow.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF"},
    },
    {
      path: "/placeportids/:id/edit",
      name: "placeportidsmodifier",
      component: () => import("@/views/PlaceportidsEdit.vue"),
      meta: {displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/help",
      name: "help",
      component: () => import("@/views/Help.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF"}
    },
    {
      path: "/qrcode",
      name: "qrcode",
      component: () => import("@/views/Qrcode.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF"}
    },
    {
      path: "/visio",
      name: "visio",
      component: () => import("@/views/Visio.vue"),
      meta: {displayToolbar: menuLayout.main, toolbarColor: "#0000FF"}
    },
    {
      path: "/notifications",
      name: "remote notifications",
      component: () => import("@/views/NotificationsList.vue"),
      meta: { displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/notifications/new",
      name: "notificationsnew",
      component: () => import("@/views/NotificationsNew.vue"),
      meta: { displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    },
    {
      path: "/notifications/edit/:id",
      name: "notificationsedit",
      component: () => import("@/views/NotificationsEdit.vue"),
      meta: { displayToolbar: menuLayout.secondary, toolbarColor: "#0000FF" }
    }
  ]
});
