<template>
  <div id="login">
    <div class="center container">
      <img class="logo" src="../assets/switter-logo-2.png">
      
      <!-- <h2 class="header">Login to SWITTER </h2> -->
      <b-form class="form">
        <b-row >
          <b-form-input  required type="text"  v-model="email" label="E-mail"></b-form-input >
        </b-row>
        <b-row >
          <b-form-input   required type="password" v-model="password" label="Password"></b-form-input >
        </b-row>
        <b-row>
          <b-btn class="btn-login" v-on:click="login" >Sign In</b-btn>
        </b-row>
        <b-row class="to-register">
          
          <router-link class="nab-link-register" to="/register">No Account? Sign Up Right Now!</router-link>
          
        </b-row>
      </b-form>
    </div>
    
  </div>
</template>

<script>
export default {
    name: 'login',
    methods: {
        login:function(){
            let postBody = {
                "email": this.email,
                "password": this.password,
            };
            let response = this.$service.Login(postBody);
            if(response !== null){
                localStorage.setItem("switterJWT", response['jwt']);
                localStorage.setItem("switterRT", response['refresh_token']);
                this.$router.push({name:'appview'});
            } else {
                this.$router.push({name:'register'});
            }
        },
    },
}
</script>

<style scoped>
#login {
  align-content: center;
}
.header{
  text-align: center;
}
.center{
  display: block;
  margin-left: auto;
  margin-right: auto;
  width:20%;
}
.btn-login{
  width: 100%;
}
.logo{
  width: 100%;
}
.to-register{
  margin-left: auto;
  margin-right: auto;
}
.nab-link-register{
  text-decoration-line: none;
}
label.b-label.theme--dark {
  top: auto;
  margin-top: 5px;
  margin-left: 5px;
}
</style>
