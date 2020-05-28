import Vue from 'vue'
import App from '@/App.vue'
/**/
import Router from 'vue-router'
Vue.use(Router)

import Vuetify from 'vuetify'
Vue.use(Vuetify)
/**/
import axios from 'axios'
//import VueAxios from 'vue-axios'
Vue.use( axios)
/**/
import Login from '@/components/Login.vue'
import Register from '@/components/Register.vue'
import AppView from '@/components/AppView.vue'
/**/
import '@/plugins/axios'
import vuetify from '@/plugins/vuetify';
/**/
Vue.config.productionTip = false


const router = new Router({
  routes: [
    {
      path: '/login',
      name:'login',
      component: Login,
    },
    {
      path: '/register',
      name:'register',
      component: Register,
    },
    {
      path: '/appview',
      name:'appview',
      component: AppView,
    },
  ]
 })

new Vue({
  el:"#app",
  render: h => h(App),
  vuetify,
  axios,
  router,
}).$mount('#app')
