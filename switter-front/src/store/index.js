import Vue from 'vue';
import Vuex from 'vuex';
import Axios from 'axios';

import Conf from '../conf'

Vue.use(Vuex);

export const store = new Vuex.Store({
    state:{
        jwt: "",
        refreshToken: "",
        messageList:[],
    },
    getters: {},
    actions:{
        Login: async (context, payload) => {
            let data = await Axios.get(Conf.baseUrl + '/api/auth/login');

            // .then(response=>{
            //     if( response.data != null){
            //       if( response.status == 200){
            //         localStorage.setItem("switterJWT", response.data['jwt']);
            //         localStorage.setItem("switterRT", response.data['refresh_token']);
            //         this.$router.push({name:'appview'});
            //       } else {
            //         this.$router.push({name:'register'});
            //       }
    
            //     }
            //   });

            if (data.status == 200) {
                context.commit('SET_NAME', name);
              }
            context.commit('SetJWT', data);
        },
        Register: async (context, payload) => {
            let {data} = await Axios.get(Conf.baseUrl + '/api/auth/register');
            context.commit('SetJWT', data);
        },
        RefreshToken: async (context, payload) => {
            let {data} = await Axios.get(Conf.baseUrl + '/api/auth/refresh');
            context.commit('SetJWT', data);
        },
        Logout: async (context, payload) => {
            let {data} = await Axios.get(Conf.baseUrl + '/api/auth/logout');
            context.commit('SetJWT', data);
        },
        GetMessagesPage: async (context, payload) => {
            let {data} = await Axios.get(Conf.baseUrl + '/api/message/getmessagepage');
            context.commit('SetJWT', data);
        },
        CreateMessage: async (context, payload) => {
            let {data} = await Axios.get(Conf.baseUrl + '/api/message/create');
            context.commit('SetJWT', data);
        },
        DeleteMessage: async (context, payload) => {
            let {data} = await Axios.get(Conf.baseUrl + '/api/auth/delete');
            context.commit('SetJWT', data);
        },
    },
    mutations:{
        SetMessageList :(state, payload) => {
            state.messageList.push(payload);
        },
        SetJWT :(state, payload) => {
            state.jwt = payload;
        },
        SetRefreshToken :(state, payload) => {
            state.refreshToken = payload;
        }
    }
});
