<template>
  <div>
    <ChangeDataTable
        v-model="dialog"
        @close="onCloseDialog"
    />
   <v-card>
  <v-data-table
      :headers="headers"
      :items="desserts"
      hide-default-footer
      sort-by="calories"
      class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
          flat
      >
        <v-toolbar-title>Main Information</v-toolbar-title>
        <v-divider
            class="mx-4"
            inset
            vertical
        ></v-divider>
        <v-spacer></v-spacer>
        <v-btn
            color="primary"
            dark
            class="mb-2"
            v-on:click="openTable"
        >
          Change Data
        </v-btn>
      </v-toolbar>
    </template>

  </v-data-table>
   </v-card>
  </div>
</template>


<script>
import {mapState} from "vuex";
import ChangeDataTable from "./ChangeDataTableDialog";

export default {
  name: "MainInfo",
  components: {ChangeDataTable},
  data: () => ({
    dialog: false,
    headers: [
      {
        text: 'Start data',
        align: 'start',
        sortable: false,
        value: 'startDate',
      },
      {
        text: 'Number of food sources',
        align: 'NumberOfFoodSources',
        sortable: false,
        value: 'NumberOfFoodSources',
      },
      {
        text: 'Final epoch',
        align: 'FinalEpoch',
        sortable: false,
        value: 'FinalEpoch',
      },
      {
        text: 'Level',
        align: 'Level',
        sortable: false,
        value: 'Level',
      },
    ],
    desserts: [],
    editedIndex: -1,
    editedItem: {
      name: '',
      calories: 0,
      fat: 0,
      carbs: 0,
      protein: 0,
    },
    defaultItem: {
      name: '',
      calories: 0,
      fat: 0,
      carbs: 0,
      protein: 0,
    },
  }),
  computed: {
    ...mapState('simulations', ['fullData']),
    tableDialog: {
        get () {
          return this.$store.state.simulations.tableDialog
        },
        set (newValue) {
          return this.$store.dispatch('simulations/setTableShow', newValue)
        }
    }
  },
  watch: {
    'fullData': function () {
      this.initialize()
    }
  },
  created() {
    this.initialize()
  },
  methods: {
    initialize() {
      this.desserts = [
        this.fullData["MainParams"]
      ]
    },
    openTable() {
      console.log(this.dialog = true)
      this.dialog = true
    },
    onCloseDialog(){
      this.dialog = false
    }



  }
}
</script>

<style scoped>

</style>
