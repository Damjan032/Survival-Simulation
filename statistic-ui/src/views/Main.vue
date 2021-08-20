<template>
  <v-app>
    <NavBar/>

    <v-content class="content">
      <v-container fluid  class="pa-0 ma-0">
        <div style="width: 100%">
          <Chart/>
          <Chart2/>
          <v-card>
          </v-card>

        </div>

      </v-container>
    </v-content>
  </v-app>
</template>
<script>

import {mapActions, mapState} from "vuex";
import NavBar from "../components/NavBar";
import Chart from "../components/Chart";
import Chart2 from "../components/ChartNumberOfPopulation";

export default {
  name: "Main",
  components: {Chart2, Chart, NavBar},
  data: () => ({
    miniDrawer: true,
  }),
  mounted() {
    this.getFullFirstData();
    this.getMainDates();
  },
  computed: {
    ...mapState('simulations', ['fullData']),

  },
  methods: {
    ...mapActions('simulations', ['getFullFirstData', 'getMainDates']),
    printData() {
      console.log(this.fullData)
    },
  }
}
</script>

<style scoped>
.content {
  background-color: #f0f0f0;
}
</style>
