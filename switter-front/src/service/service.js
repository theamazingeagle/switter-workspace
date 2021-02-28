import Axios from 'axios';

export default class Service {
    constructor(baseUrl){
        this.baseUrl = baseUrl;
    }
    async Login(payload){
        let response = await Axios.post(
            this.baseUrl + '/api/auth/login',
            payload,
            {headers:{'Content-Type':'application/json'}}
        );

        if (data.status == 200) {
            return response.data;
        }

        return null;
    }
    async Register(payload) {
        let response = await Axios.post(
            this.baseUrl + '/api/auth/register', 
            payload,
            {headers:{'Content-Type':'application/json'}}
        );
        if (data.status == 200) {
            return response.data;
        }
        return null;
    }
    async RefreshToken (payload) {
        let response = await Axios.post(this.baseUrl + '/api/auth/refresh');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
    async Logout (payload){
        let response = await Axios.get(this.baseUrl + '/api/auth/logout');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
    async GetThreadsPage(payload, token) {
        let response = await Axios.get(this.baseUrl + '/api/thread/getpage');
        if (data.status == 200) {
            return response.data;
        }
        return null;
    }
    async CreateThread(payload, token){
        let response = await Axios.post(this.baseUrl + '/api/thread/create');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
    async DeleteThread (payload, token) {
        let response = await Axios.get(this.baseUrl + '/api/thread/delete');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
    async GetCommentsPage(payload, token) {
        let response = await Axios.get(this.baseUrl + '/api/comment/getpage');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
    async CreateComment(payload, token){
        let response = await Axios.post(this.baseUrl + '/api/comment/create');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
    async DeleteComment(payload, token) {
        let response = await Axios.get(this.baseUrl + '/api/comment/delete');
        if (data.status == 200) {
            return response.data;
        }
        
        return null;
    }
}