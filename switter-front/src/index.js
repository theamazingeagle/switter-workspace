import {BaseUrl} from './conf/conf.js'

import Vue from 'vue'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

import Router from 'vue-router'
Vue.use(Router)
import axios from 'axios';
Vue.use(axios);

//import store from './store'

import App from './App.vue'
import Login from './components/Login.vue'
import Register from './components/Register.vue'
import AppView from './components/AppView.vue'

Vue.prototype.$http = axios;
Vue.prototype.$hostname = BaseUrl;
let router = new Router({
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
      path: '/',
      name:'appview',
      component: AppView,
    },
  ]
 });

 
new Vue({
  render: h => h(App),
  router,
  BaseUrl,
}).$mount('#appl');
