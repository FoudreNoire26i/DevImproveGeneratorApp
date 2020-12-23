<template>
  <v-container fill-height fluid pa-0 grid-list-md>
    <v-container pt-1 mb-4>
      <v-layout pa-2 mb-2 column>
        <v-flex v-for="(msg, i) in feed" :key="i">
          <div v-if="msg.id == $id_usr" class="speech-bubble speech-bubble-right">
            <div class="body-2" v-html="msg.content"></div>
          </div>
          <div v-else class="speech-bubble speech-bubble-left">
            <div class="body-2" v-html="msg.content"></div>
          </div>
        </v-flex>
      </v-layout>
    </v-container>
    <v-container fluid pa-1 class="send">
      <v-layout row fill-height>
        <v-text-field
          v-model="content"
          @click:append="send()"
          @click:prepend-inner="''"
          hide-details
          solo
          @keyup.native.enter="send()"
          append-icon="send"
          prepend-inner-icon="hi"
          placeholder="Aa"
        ></v-text-field>
      </v-layout>
    </v-container>
  </v-container>
</template>

<style>
.v-content__wrap {
  background-color: #ECE5DD;
}

.send {
  float: left;
  position: fixed;
  bottom: 0;
}

.speech-bubble {
  max-width: 30ch;
  padding: 0.5em;
  position: relative;
  text-align: left;
  box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
  word-wrap: break-word;
  border-radius: 0.4em;
}

.speech-bubble-left {
  float: left;
  background: white;
}

.speech-bubble-right {
  float: right;
  background: #dcf8c6;
}

.speech-bubble::after {
  content: "";
  position: absolute;
  top: 50%;
  width: 0;
  height: 0;
  border: 11px solid transparent;
  border-bottom: 0;
  box-shadow: 0px 1px rgba(0, 0, 0, 0.1);
  margin-top: -5.5px;
}

.speech-bubble-right::after {
  right: 0;
  border-left-color: #dcf8c6;
  border-right: 0;
  margin-right: -11px;
}

.speech-bubble-left::after {
  left: 0;
  border-right-color: white;
  border-left: 0;
  margin-left: -11px;
}
</style>


<script>
import constants from "@/constants";
const api = `${constants.apiUrl}chats?per_page=50`;
const everything = `${constants.apiUrl}chats?per_page=1000`;
const absolute = `${constants.apiUrl}chats/`;

export default {
  props: {
    room: String,
    defaultFeed: {
      type: Array,
      required: false,
      default: function() {
        return [];
      }
    },
    mock: {
      type: Boolean,
      default: false,
      required: false
    }
  },
  data: function() {
    return {
      content: "",
      user: "Utilisateur",
      polling: null,
      feed: this.defaultFeed
    };
  },
  methods: {
    check: function() {
      if (/^\s*$/.exec(this.content) == null) return true;
      return false;
    },
    pollData: function() {
      this.polling = setInterval(() => {
        this.get();
      }, 4000);
    },
    get: function() {
      this.$http
        .get(api)
        .then(res => {
          this.feed = res.data.filter(x => x.chatroom_str === this.room);
          this.feed = this.feed.map(x => {
            return {
              id: x.id_int,
              name: x.user_str,
              content: x.text_str
            };
          });
          this.$vuetify.goTo(9999, "easeOutQuad");
        })
        .catch(err => {
          throw err;
        });
    },
    post: function() {
      this.$http
        .post(api, {
          chatroom_str: this.room,
          id_int: this.$id_usr,
          text_str: this.content,
          user_str: this.user
        })
        .then(res => {
          this.display({
            id: this.$id_usr,
            name: this.user,
            content: this.content
          });
        })
        .catch(err => {
          throw err;
        });
    },
    display: function() {
      this.feed.push({ id: this.$id_usr, name: this.user, content: this.content });
      this.content = "";
      this.$vuetify.goTo(9999, "easeOutQuad");
    },
    send: function() {
      if (this.check()) {
        this.post();
      }
    }
  },
  created() {
    console.log(this.$id_usr);
    if (!this.mock) {
      this.get();
      this.pollData();
    }
  },
  beforeDestroy() {
    clearInterval(this.polling);
  }
};
</script>