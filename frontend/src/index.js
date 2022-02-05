import Vue from 'vue'
import App from './components/App.vue'
import "bootstrap/dist/css/bootstrap.css"

Vue.config.productionTip = false

new Vue({
  el: '#app',
  render: h => h(App)
})
