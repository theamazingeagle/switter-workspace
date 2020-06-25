
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
import '@/plugins/router'
import vuetify from '@/plugins/vuetify';
/**/
Vue.config.productionTip = false
/**/
import VModal from 'vue-js-modal'
Vue.use(VModal)
/**/
//--- Auth Hook -------------
// const ifNotAuthenticated = (to, from, next) => {
//   if (localStorage.getItem("switterJWT") !=="") {
//     next()
//     return
//   }
//   next('/login')
// }
//---------------------------
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
      meta: { 
        requiresAuth: true
      },
      //beforeEnter: ifNotAuthenticated
    },
  ]
 });

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (localStorage.getItem("switterJWT") === null ) {
        //console.log("here1");
        next({ path: '/login' });
        return;
    } else {
        next(); 
        //console.log("here2");
        return;
      } 
  } else {
    next(); 
    //console.log("here3");
    return;
  }    
});

new Vue({
  el:"#app",
  render: h => h(App),
  vuetify,
  axios,
  router,
}).$mount('#app');