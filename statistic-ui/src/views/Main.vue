<template>
  <v-app>
<!--    <v-navigation-drawer-->
<!--        dark-->
<!--        class="primary elevation-5"-->
<!--        color="teal lighten-3"-->
<!--        :mini-variant="mini"-->
<!--        permanent-->
<!--        app-->
<!--    >-->
<!--    </v-navigation-drawer>-->
    <NavBar/>

    <v-content class="content">
      <v-container fluid style="width: 90%" class="justify-center">
        <div >
          <MainInfo/>
          <v-divider></v-divider>
          <ChartPercentOfPopulation/>
          <v-divider></v-divider>
          <ChartNumberOfPopulation/>

        </div>

      </v-container>
    </v-content>
  </v-app>
</template>
<script>

import {mapActions, mapState} from "vuex";
import NavBar from "../components/NavBar";
import ChartPercentOfPopulation from "../components/ChartPercentOfPopulation";
import ChartNumberOfPopulation from "../components/ChartNumberOfPopulation";
import MainInfo from "../components/MainInfo";

export default {
  name: "Main",
  components: {MainInfo, ChartNumberOfPopulation, ChartPercentOfPopulation, NavBar},
  data: () => ({
    miniDrawer: true,
    mini: {
      type: Boolean,
      default: true
    },
  }),
  mounted() {
    this.getFullFirstData();
    this.getMainDates();
  },
  computed: {
    ...mapState('simulations', ['fullData']),

  },

  methods: {
    ...mapActions('simulations', ['getFullFirstData', 'getMainDates', "getFullDataById"]),

  }
}
</script>

<style scoped>
.content {
  background-color: #f0f0f0;
}
</style>
