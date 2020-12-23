<!-- PAGE DE VISIO
    /!\ CAM et MICRO marche pas en local car http ; uniquement fctel sur le lien jetanswer car https
    *- bien laisser la class fill sur le container et le flex pour que la la page jitsi prenne tout l'écran
    *- roomId aleatoire et unique
-->



<template>
<v-container pa-0 ma-0 class="fill">
  <v-flex id="meet" class="fill"></v-flex>
</v-container>
</template>

<style>
.fill{
  display:block;
  height:100%;
}
</style>
<script>
  import constants from "@/constants";
  //import sfu from "@/sfu";
  const api = `${constants.apiUrl}s`
  export default {
    data(){
      return {
        domain : 'meet.jit.si',
        options : {
          roomName: 'StartAppHelper',
          width: "100%",
          height: "100%",
          parentNode: document.getElementById('meet'),
          configOverwrite: { disableDeepLinking: true, 
                            enableWelcomePage: false
                            },
          interfaceConfigOverwrite: { SHOW_JITSI_WATERMARK: false,
                                      TOOLBAR_BUTTONS: [
                                        'microphone', 'camera', 'closedcaptions', 'desktop', 'embedmeeting', 'fullscreen',
                                        'fodeviceselection', 'hangup', 'profile', 'chat', 'recording',
                                        'etherpad', 'sharedvideo', 'settings', 'raisehand',
                                        'videoquality', 'filmstrip', 'invite', 'feedback', 'stats',
                                        'tileview', 'videobackgroundblur', 'download', 'mute-everyone', 'security'
                                        ]
                                    }
          },
        apiJ : null
        }
    },
    mounted(){
      this.openOrCloseConf()
    },
    methods:{
      openOrCloseConf(){
          this.options.roomName = this.uuidv4()
          this.options.parentNode = document.getElementById('meet')
          this.apiJ = new JitsiMeetExternalAPI(this.domain, this.options)
          //connectedlistener
          /*
          this.apiJ.addEventListener("videoConferenceJoined", () => {
                    console.log("Local User Joined");
          })*/
          //hang-up listener
          this.apiJ.addEventListener("videoConferenceLeft", () => {
              this.$router.go(-1)
          })
      },
      //UUID generator for room name
      uuidv4() {
          return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
            var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
      }
    },
  }
  
</script> 
