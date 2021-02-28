import Conf from './conf/conf.js'

import Vue from 'vue'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

import Router from 'vue-router'
Vue.use(Router)
// import axios from 'axios';
// Vue.use(axios);

import Service from './service/service.js'

import App from './App.vue'
import Login from './components/Login.vue'
import Register from './components/Register.vue'
import AppView from './components/AppView.vue'
import ThreadView from './components/AppView/ThreadView.vue'
import ThreadList from './components/AppView/ThreadList.vue'

//Vue.prototype.$http = axios;
Vue.prototype.$service = new Service(Conf.BaseUrl);
Vue.prototype.$hostname = Conf.BaseUrl;
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
      children:[
        {
            path:'',
            component: ThreadList
        },
        {
          path:'thread/:threadID',
          component: ThreadView,
          props: true,
        },
      ],
    },
    {
      path: "*",
      component: { render: (h) => h("div", ["404! Page Not Found!"]) },
    },
  ]
 });

 
new Vue({
  render: h => h(App),
  router,
  Conf,
}).$mount('#appl');
