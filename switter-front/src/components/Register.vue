<template>
  <div id="register">
  <div class="auth-error" b-if="authError">Try Again</div>
    <div class="center">
      <img class="logo" src="../assets/switter-logo-2.png">
      <b-form class="form">
        <b-row >
          <b-form-input type="text"  v-model="username" label="Your Name"></b-form-input>
        </b-row>
        <b-row >
          <b-form-input type="text"  v-model="email" label="E-mail"></b-form-input>
        </b-row>
        <b-row >
          <b-form-input type="password" v-model="password" label="Password"></b-form-input>
        </b-row>
          <b-btn class="btn-register" v-on:click="register">Sign Up</b-btn>
      </b-form>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'register',
    methods: {
      register:function(){
        let postBody = {
          "username": this.username,
          "email": this.email,
          "password": this.password,
        };
        let response = this.$service.Register(postBody);
        
        if(response !== null) {
            this.authError = false;
            localStorage.setItem("switterJWT",  response.jwt);
            localStorage.setItem("switterRT",  response.rt);
            this.$router.push({name:'appview'});
        }
        this.authError = true;
        
      },
    },
    data: function() {
      return {
          authError:false,
          email: '',
          password: '',
          username: '',
      }
    },
  }
</script>

<style scoped>
#register{
  
  align-content: center;
    
}
.header{
  text-align: center;
}
.logo{
  width: 100%;
}
.center{
  display: block;
  margin-top: auto;
  margin-left: auto;
  margin-right: auto;
  width:20%;
}
.btn-register{
  width: 100%;
}
</style>
