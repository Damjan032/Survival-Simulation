<template>



  <v-container ma-0 pa-0 fill-height>
    <Chart/>
    <Chart2/>
  </v-container>
</template>

<script>
import Chart from "./Chart";
import Chart2 from "./ChartNumberOfPopulation";
export default {
  name: 'HelloWorld',
  components: {Chart2, Chart},
  props: {
    msg: String
  },

}
</script>
