<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <Chart/>
    <h1>{{ msg }}</h1>
    <Chart2/>
  </div>
</template>

<script>
import Chart from "./Chart";
import Chart2 from "./Chart2";
export default {
  name: 'HelloWorld',
  components: {Chart2, Chart},
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
