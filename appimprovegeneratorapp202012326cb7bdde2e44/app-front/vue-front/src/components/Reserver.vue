<template>
  <v-container>
    <v-layout wrap align-center column>
      <v-flex xs12 sm6 md6>
        <v-form ref="form">
          <v-radio-group
            row
            v-model="trip.type"
            :rules="[() => trip.type !== '0' || 'Séléctionner un type de vol !']"
          >
            <v-radio label="Aller Simple" :color="$route.meta.toolbarColor" value="1"></v-radio>
            <v-radio label="Aller & Retour" :color="$route.meta.toolbarColor" value="2"></v-radio>
          </v-radio-group>

          <v-select
            ref="depDate"
            :color="$route.meta.toolbarColor"
            label="Au départ de"
            v-model="trip.dep"
            :items="items"
            :rules="fieldRules"
            @change="$refs.arrDate.validate()"
            outline
            clearable
            required
          ></v-select>
          <v-select
            ref="arrDate"
            :color="$route.meta.toolbarColor"
            label="À destination de"
            v-model="trip.dest"
            :items="items"
            :rules="fieldRules"
            @change="$refs.depDate.validate()"
            outline
            clearable
            required
          ></v-select>

          <v-menu
            full-width
            transition="scale-transition"
            :close-on-content-click="false"
            :nudge-right="40"
            offset-y
            min-width="290px"
            v-model="menus.dep"
            lazy
          >
            <v-text-field
              slot="activator"
              label="Date de départ"
              v-model="formattedDates.dep"
              :rules="[...rules, () => trip.type === '1' || 'Veuillez rens']"
              append-icon="event"
              :color="$route.meta.toolbarColor"
              clearable
              outline
              readonly
              required
            ></v-text-field>
            <v-date-picker
              :min="new Date().toISOString()"
              :max="trip.type === '1' ? '' : dates.dest"
              v-model="dates.dep"
              locale="fr-FR"
              @input="menus.dep=false"
              :color="$route.meta.toolbarColor"
            ></v-date-picker>
          </v-menu>
          <v-menu
            full-width
            transition="scale-transition"
            :close-on-content-click="false"
            :nudge-right="40"
            offset-y
            min-width="290px"
            v-model="menus.dest"
            :disabled="trip.type == 1 ? true : false"
            lazy
          >
            <v-text-field
              slot="activator"
              label="Date de retour"
              v-model="formattedDates.dest"
              :rules="[() => trip.type === '1' || 'Veuillez renseigner ce champ !']"
              append-icon="event"
              :disabled="trip.type == 1 ? true : false"
              :color="$route.meta.toolbarColor"
              clearable
              outline
              readonly
              required
            ></v-text-field>
            <v-date-picker
              :min="dates.dep || new Date().toISOString()"
              v-model="dates.dest"
              locale="fr-FR"
              @input="menus.dest=false"
              :color="$route.meta.toolbarColor"
            ></v-date-picker>
          </v-menu>
          <v-btn @click="handleValidate" :color="$route.meta.toolbarColor" block>Recherche</v-btn>
        </v-form>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "Reserver",
  props: {
    to: {
      type: String,
      required: false,
      default: "/"
    }
  },
  data: () => ({
    menus: {
      dep: false,
      dest: false
    },
    dates: {
      dep: "",
      dest: ""
    },
    trip: {
      dep: "",
      dest: "",
      type: "1"
    },
    items: ["Lyon", "Marseille", "Paris", "Strasbourg", "Bordeaux"],
    rules: [v => (v && v.length > 0) || "Veuillez renseigner ce champ !"]
  }),
  computed: {
    fieldRules() {
      const equality = () =>
        this.trip.dep !== this.trip.dest || "Départ et arrivée identiques";

      return [...this.rules, equality];
    },
    formattedDates() {
      return {
        dep: this.dates.dep
          ? new Date(this.dates.dep).toLocaleDateString("fr-FR")
          : "",
        dest: this.dates.dest
          ? new Date(this.dates.dest).toLocaleDateString("fr-FR")
          : ""
      };
    }
  },
  methods: {
    handleValidate() {
      if (this.$refs.form.validate()) this.$router.push({ path: this.to });
    }
  }
};
</script>

<style>
</style>
