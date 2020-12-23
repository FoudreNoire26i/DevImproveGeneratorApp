<template>
  <v-layout column>
    <v-flex xs12>
      <v-list two-line>
        <template v-for="mail in mails">
          <v-list-tile class="py-2" ripple :key="mail.id" router append :to="mail.route">
            <v-list-tile-content>
              <v-list-tile-title v-html="mail.title"></v-list-tile-title>
              <v-list-tile-sub-title v-html="mail.subtitle"></v-list-tile-sub-title>
              <v-list-tile-sub-title class="grey--text" v-html="mail.excerpt"></v-list-tile-sub-title>
            </v-list-tile-content>
            <v-list-tile-action>
              <v-list-tile-action-text></v-list-tile-action-text>
              <v-icon color="grey lighten-1">star</v-icon>
            </v-list-tile-action>
          </v-list-tile>
        </template>
      </v-list>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {
    data: () => ({
      mails: [
        {
          id: 1,
          title: "Keira Smith",
          subtitle: "Fwd: Project proposal for new model",
          excerpt: "Hey I am sending the document",
          date: "10:32"
        },
        {
          id: 2,
          title: '<span class="font-weight-black black--text">Tooap</span>',
          subtitle:
            '<span class="font-weight-black black--text">App à tester</span>',
          excerpt: `<span class="font-weight-black">Salut, j'ai trouvé une nouvelle application super ! Tu dois la tester.</span>`,
          date: "9:52",
          get route() {
            return this.id.toString();
          }
        },
        {
          id: 3,
          title: "Christine Jean",
          subtitle: "Project draft",
          excerpt: "Test shot for upcoming event this friday",
          date: "9:50"
        },
        {
          id: 4,
          title: "Marissa Hero",
          subtitle: "How are you doing? It is fantastic season",
          excerpt: "Hey I am sending the document",
          date: "9:31"
        },
        {
          id: 5,
          title: "Keira Smith",
          subtitle: "Android report 2",
          excerpt: "I've attached the latest file for Android report",
          date: "9:11"
        },
        {
          id: 6,
          title: "Keira Smith",
          subtitle: "Android report 2",
          excerpt: "I've attached the latest file for Android report",
          date: "9:11"
        },
        {
          id: 7,
          title: "Keira Smith",
          subtitle: "Android report 2",
          excerpt: "I've attached the latest file for Android report",
          date: "9:11"
        }
      ]
    })
  };
</script>