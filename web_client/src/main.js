import Vue from 'vue'
import router from './router/index'
import {store} from  './store/store'
import Master from './components/layouts/Master'


window.eventBus = new Vue()
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store: store,
  router:router,
  components: { Master },
  template: '<Master/>'
})
