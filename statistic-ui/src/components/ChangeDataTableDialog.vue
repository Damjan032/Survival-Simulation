<template>
  <v-dialog
      :value="value"
      @input="$emit('input', $event)"
      max-width="80%"
  >
    <v-card>
      <v-card-title class="headline">Deletion</v-card-title>

      <v-data-table
          :headers="headers"
          :items="mainDates"
          hide-default-footer
          sort-by="calories"
          class="elevation-1"
      >
        <template v-slot:item.load="{ item }">
          <v-icon
              color="success"
              @click="printData(item)"
          >mdi-download
          </v-icon>
        </template>

      </v-data-table>

      <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn
            color="green darken-1"
            text
            @click="$emit('close')"
        >
          Close
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>



<script>
import {mapActions, mapState} from "vuex";

export default {
  name: "ChangeDataTable",
  props: [ 'value'],
  data: () => ({
    dialog: false,
    headers: [
      {
        text: 'Start data',
        align: 'start',
        value: 'startDate',
      },
      {
        text: 'Number of food sources',
        align: 'NumberOfFoodSources',
        value: 'NumberOfFoodSources',
      },
      {
        text: 'Final epoch',
        align: 'FinalEpoch',
        value: 'FinalEpoch',
      },
      {
        text: 'Level',
        align: 'Level',
        value: 'Level',
      },
      {text: 'Load Data', value: 'load', sortable: false, align: 'center'},
    ],
  }),
  computed: {
    ...mapState('simulations', ['fullData', 'mainDates']),
  },
  methods: {
    ...mapActions('simulations', ['getFullFirstData', 'getMainDates', "getFullDataById"]),
    printData(item) {
      this.getFullDataById(item['_id'])
      this.$emit('close')
    },
    onDaDa(){
      this.getFullDataById("611fe15f0584b4b285569860")
    }
  }
}
</script>

<style scoped>

</style>
