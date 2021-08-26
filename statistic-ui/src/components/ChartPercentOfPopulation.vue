<template>
  <div class="" id="vue-container">
    <div class="" id="chartContainer" style="height: 360px; "></div>
  </div>
</template>

<script>
import * as CanvasJS from '../CanvasJS/canvasjs.min'
import {mapState} from "vuex";

export default {
  //el: '#vue-container',
  name: "Chart",
  data: function () {
    return {
      chartOptions: {
        animationEnabled: true,
        title: {
          text: "Percent of population graph",
        },
        axisX: {
          title: "Epoch"
        },
        axisY: {
          title: "Percentage",
          suffix: "%",
          includeZero: true
        },
        data: []
      },
      chart: null
    };
  },
  mounted: function () {
    this.chart = new CanvasJS.Chart("chartContainer", this.chartOptions);
    this.drawGraph()
    this.chart.render();
  },
  computed: {
    ...mapState('simulations', ['fullData', 'mainDates']),

  },
  watch: {
    'fullData': function () {
      this.drawGraph()
    }
  },
  methods: {

    percentage(partialValue, totalValue) {
      return parseFloat(((100 * partialValue) / totalValue).toFixed(2));
    },
    drawGraph() {
      console.log(this.fullData['Epoches'])
      let dataPointsGood = []
      let dataPointsBad = []
      console.log(this.fullData['Epoches'].length)
      let step = parseInt(this.fullData['Epoches'].length / 20)
      dataPointsGood.push({
        x: this.fullData['Epoches'][0]['CurrentEpoch'],
        y: this.percentage(this.fullData['Epoches'][0]['NumberOfGood'],
            (this.fullData['Epoches'][0]['NumberOfBad'] + this.fullData['Epoches'][0]['NumberOfGood']))
      })
      dataPointsBad.push({
        x: this.fullData['Epoches'][0]['CurrentEpoch'],
        y: this.percentage(this.fullData['Epoches'][0]['NumberOfBad'],
            (this.fullData['Epoches'][0]['NumberOfBad'] + this.fullData['Epoches'][0]['NumberOfGood']))
      })
      for (let i = step; i < this.fullData['Epoches'].length - 1; i = i + step) {
        dataPointsGood.push({
          x: this.fullData['Epoches'][i]['CurrentEpoch'],
          y: this.percentage(this.fullData['Epoches'][i]['NumberOfGood'],
              (this.fullData['Epoches'][i]['NumberOfBad'] + this.fullData['Epoches'][i]['NumberOfGood']))
        })
        dataPointsBad.push({
          x: this.fullData['Epoches'][i]['CurrentEpoch'],
          y: this.percentage(this.fullData['Epoches'][i]['NumberOfBad'],
              (this.fullData['Epoches'][i]['NumberOfBad'] + this.fullData['Epoches'][i]['NumberOfGood']))
        })
      }
      dataPointsGood.push({
        x: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['CurrentEpoch'],
        y: this.percentage(this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfGood'],
            (this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfBad'] +
                this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfGood']))
      })
      dataPointsBad.push({
        x: this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['CurrentEpoch'],
        y: this.percentage(this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfBad'],
            (this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfBad'] +
                this.fullData['Epoches'][this.fullData['Epoches'].length - 1]['NumberOfGood']))
      })
      console.log("PRECENT")
      console.log(dataPointsBad)
      console.log(dataPointsGood)
      this.chartOptions.data = [{
        showInLegend: true,
        type: "spline",
        color: 'blue',
        name: "Number of Good",
        yValueFormatString: "#,##0.##\"%\"",
        connectNullData: true,
        dataPoints: dataPointsGood
      },
        {
          showInLegend: true,
          type: "spline",
          color: 'red',
          name: "Number of Bad",
          connectNullData: true,
          yValueFormatString: "#,##0.##\"%\"",
          dataPoints: dataPointsBad
        }]
      this.chart.render()
    },
  }

}
</script>

<style scoped>

</style>
