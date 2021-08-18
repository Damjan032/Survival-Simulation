<template>
  <div id="vue-container2">
    <div id="chartContainer2" style="height: 360px; width: 100%;"></div>
  </div>
</template>

<script>
import * as CanvasJS from '../CanvasJS/canvasjs.min'
export default {
  el: '#vue-container2',
  name: "Chart2",
  data: function() {
    return {
      chartOptions:{
        animationEnabled: true,
        title: {
          text: "Hourly Average CPU Utilization",
        },
        axisX: {
          title: "Time"
        },
        axisY: {
          title: "Percentage",
          suffix: "%",
          includeZero: true
        },
        data: [{
          showInLegend: true,
          type: "spline",
          name: "CPU Utilization",
          connectNullData: true,
          //nullDataLineDashType: "solid",
          xValueType: "dateTime",
          xValueFormatString: "DD MMM hh:mm TT",
          yValueFormatString: "#,##0.##\"%\"",
          dataPoints: [
            { x: 1501048673000, y: 35.939 },
            { x: 1501052273000, y: 40.896 },
            { x: 1501055873000, y: 56.625 },
            { x: 1501059473000, y: 26.003 },
            { x: 1501063073000, y: 20.376 },
            { x: 1501066673000, y: 19.774 },
            { x: 1501070273000, y: 23.508 },
            { x: 1501073873000, y: 18.577 },
            { x: 1501077473000, y: 15.918 },
            { x: 1501081073000, y: null }, // Null Data
            { x: 1501084673000, y: 10.314 },
            { x: 1501088273000, y: 10.574 },
            { x: 1501091873000, y: 14.422 },
            { x: 1501095473000, y: 18.576 },
            { x: 1501099073000, y: 22.342 },
            { x: 1501102673000, y: 22.836 },
            { x: 1501106273000, y: 23.220 },
            { x: 1501109873000, y: 23.594 },
            { x: 1501113473000, y: 24.596 },
            { x: 1501117073000, y: 31.947 },
            { x: 1501120673000, y: 31.142 }
          ]
        },
          {
            showInLegend: true,
            type: "spline",
            name: "CPU Utilization",
            connectNullData: true,
            //nullDataLineDashType: "solid",
            xValueType: "dateTime",
            xValueFormatString: "DD MMM hh:mm TT",
            yValueFormatString: "#,##0.##\"%\"",
            dataPoints: [
              { x: 1501048673000, y: 45.939 },
              { x: 1501052273000, y: 30.896 },
              { x: 1501055873000, y: 66.625 },
              { x: 1501059473000, y: 36.003 },
              { x: 1501063073000, y: 30.376 },
              { x: 1501066673000, y: 29.774 },
              { x: 1501070273000, y: 33.508 },
              { x: 1501073873000, y: 28.577 },
              { x: 1501077473000, y: 25.918 },
              { x: 1501081073000, y: null }, // Null Data
              { x: 1501084673000, y: 20.314 },
              { x: 1501088273000, y: 20.574 },
              { x: 1501091873000, y: 24.422 },
              { x: 1501095473000, y: 28.576 },
              { x: 1501099073000, y: 32.342 },
              { x: 1501102673000, y: 32.836 },
              { x: 1501106273000, y: 33.220 },
              { x: 1501109873000, y: 33.594 },
              { x: 1501113473000, y: 34.596 },
              { x: 1501117073000, y: 41.947 },
              { x: 1501120673000, y: 41.142 }
            ]
          }]
      },
      chart : null
    };
  },
  mounted: function () {
    this.chart = new CanvasJS.Chart("chartContainer2", this.chartOptions);
    this.chart.render();
  }

}
</script>

<style scoped>

</style>
