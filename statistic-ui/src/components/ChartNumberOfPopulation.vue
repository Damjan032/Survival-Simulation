<template>
  <div class="" id="vue-container2" >
    <div class="" id="chartContainer2" style="height: 360px; width: 100%;"></div>
  </div>

</template>

<script>
import * as CanvasJS from '../CanvasJS/canvasjs.min'
import {mapState} from "vuex";

export default {
  // el: '#vue-container2',
  name: "Chart2",
  data: function () {
    return {
      counter: 0,
      chartOptions: {
        animationEnabled: true,
        title: {
          text: "Total population graph",
        },
        axisX: {
          title: "Epoch"
        },
        axisY: {
          title: "Population",
          includeZero: true
        },
        data: []
      },
      chart: null
    };
  },
  mounted: function () {
    this.chart = new CanvasJS.Chart("chartContainer2", this.chartOptions);
    this.drawGraph()
    this.chart.render();
  },
  computed: {
    ...mapState('simulations', ['fullData', 'mainDates']),

  },
  methods: {
    drawGraph() {
      console.log(this.fullData['Epoches'])
      let dataPointsGood = []
      let dataPointsBad = []
      let allPopulation = []
      console.log(this.fullData['Epoches'].length)
      let step = parseInt(this.fullData['Epoches'].length / 20)
      dataPointsGood.push({
        x: this.fullData['Epoches'][0]['CurrentEpoch'],
        y: this.fullData['Epoches'][0]['NumberOfGood']
      })
      dataPointsBad.push({
        x: this.fullData['Epoches'][0]['CurrentEpoch'],
        y: this.fullData['Epoches'][0]['NumberOfBad']
      })
      allPopulation.push({
        x: this.fullData['Epoches'][0]['CurrentEpoch'],
        y: (this.fullData['Epoches'][0]['NumberOfBad'] + this.fullData['Epoches'][0]['NumberOfGood'])
      })
      console.log()
      for (let i = step; i < this.fullData['Epoches'].length - 1; i = i + step) {
        dataPointsBad.push({
          x: this.fullData['Epoches'][i]['CurrentEpoch'],
          y: this.fullData['Epoches'][i]['NumberOfBad']
        })
        dataPointsGood.push({
          x: this.fullData['Epoches'][i]['CurrentEpoch'],
          y: this.fullData['Epoches'][i]['NumberOfGood']
        })
        allPopulation.push({
          x: this.fullData['Epoches'][i]['CurrentEpoch'],
          y: (this.fullData['Epoches'][i]['NumberOfBad'] + this.fullData['Epoches'][i]['NumberOfGood'])
        })
      }
      dataPointsGood.push({
        x: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['CurrentEpoch'],
        y: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfGood']
      })
      dataPointsBad.push({
        x: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['CurrentEpoch'],
        y: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfBad']
      })
      allPopulation.push({
        x: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['CurrentEpoch'],
        y: (this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfBad'] +
            this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfGood'])
      })
      this.chartOptions.data = [{
        showInLegend: true,
        type: "spline",
        color: 'blue',
        name: "Number of Good",
        connectNullData: true,
        dataPoints: dataPointsGood
      },
        {
          showInLegend: true,
          type: "spline",
          color: 'red',
          name: "Number of Bad",
          connectNullData: true,
          dataPoints: dataPointsBad
        },
        {
          showInLegend: true,
          type: "spline",
          color: 'green',
          name: "Full Population",
          connectNullData: true,
          dataPoints: allPopulation
        }]
      this.chart.render()
    },
  }
}
</script>

<style scoped>

</style>
