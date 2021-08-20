import Vue from 'vue'
import Vuex from 'vuex'

import VuexPersist from 'vuex-persist';
import SimulationDataModule from "./SimulationDataModule";


Vue.use(Vuex)

const vuexPersist = new VuexPersist({
  key: 'my-app',
  storage: window.localStorage,
});

export default new Vuex.Store({
  plugins: [vuexPersist.plugin],
  namespaced: true,
  actions: {
  },
  modules: {
    simulations : SimulationDataModule
  }
})
