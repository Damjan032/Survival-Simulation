import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import Axios from "axios";
import './styles/App.css';
import VueAxios from "vue-axios";
import store from "./store";
//import './assets/js/main'
//import VueToast from "vue-toast-notification";
//import "vue-toast-notification/dist/theme-sugar.css";
import Fragment from 'vue-fragment';
Vue.use(Fragment.Plugin);

Vue.config.productionTip = false
//Vue.use(VueToast);
Vue.use(VueAxios, Axios);

Vue.axios.defaults.baseURL = "http://localhost:8080";
new Vue({
  vuetify,
  store,
  render: h => h(App)
}).$mount('#app')
