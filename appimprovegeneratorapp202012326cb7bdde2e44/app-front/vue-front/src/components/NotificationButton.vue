<!-- 
Summary  	: 	Button which activates a local notification
Props 	 	: 	.placeholder : (String) placeholder of the button
              .title   	 : (String) title of the notification
              .url   	 	 : (String) redirection's path of the notification
              .start   	 : (Number) displaying delay (seconds) of the notification
              .load      : (Boolean) allows loading display 
Emits     	:  
Editable 	  : 	
			
Tips      	:   
-->
<template>
  <div>
    <v-btn v-if="!stop" :color="color" @click.prevent.stop="displayNotif()" dark>
      <div v-if="!loading"></div>
      <v-progress-circular
        v-else-if="!stop"
        color="white"
        indeterminate
      ></v-progress-circular>
    </v-btn>
    <v-btn v-if="stop" color="green" dark>
      <v-icon dark>done</v-icon>
    </v-btn>
  </div>
</template>

<script>
import store from "../Store";
export default {
  props: {
    color: String,
    placeholder: {
      type: String,
      default: "Notification"
    },
    title: {
      type: String,
      required: true
    },
    url: {
      type: String,
      required: true
    },
    start: {
      type: Number,
      required: false,
      default: 0
    },
    load: {
      type: Boolean,
      required: false,
      default: false,
    }
  },
  data () {
      return {
        interval: {},
        value: 0,
        loading: false,
        stop: false,
      }
    },
    
    
  methods: {
    displayNotif() {
      this.loading = true;
      const notification = {
        id_int: 1,
        text_str: this.title,
        url_str: this.url,
        activated_int: 1
      };
      setTimeout(
        () => {
          localStorage.setItem("notification", JSON.stringify(notification));
          this.stop = true;
        }, this.start * 1000);
    },
    beforeDestroy () {
      clearInterval(this.interval)
    },
  }
};
</script>