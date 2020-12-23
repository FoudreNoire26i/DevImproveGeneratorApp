export default {
  "listView": 1, // la vue par défaut dans le listing des ressources
  "apiUrl": "https://jetanswer.com/apis/appimprovegeneratorapp202012326cb7bdde2e44/",
  "ownerFile": "appimprovegeneratorapp202012326cb7bdde2e44",
  "urlFileDev": "https://jetanswer.com/filekitten/",
  "idNotifInterval": 0,
  "$debug": false // mettre à 'false' pour sortir du mode debug
}

const layouts = {
  BACK: 0,
  TOOLBAR: 1,
  PLAYSTORE: 2,
  BOTTOM: 3,
  BURGER: 4 
};

export const menuLayout = {
  main: layouts.TOOLBAR,
  secondary: layouts.BACK
};
