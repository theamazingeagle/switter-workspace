<template>
  <div id="login">
    <div class="center">
      <v-img src='assets/switter-logo.png'  ></v-img>
      
      <!-- <h2 class="header">Login to SWITTER </h2> -->
      <v-form class="form">
        <v-row >
          <v-text-field required type="text"  v-model="email" label="E-mail"></v-text-field>
        </v-row>
        <v-row >
          <v-text-field  required type="password" v-model="password" label="Password"></v-text-field>
        </v-row>
        <v-row>
          <v-btn class="btn-login" v-on:click="login">Sign In</v-btn>
        </v-row>
        <v-row class="to-register">
          
          <router-link class="nav-link-register" to="/register">No Account? Sign Up Right Now!</router-link>
          
        </v-row>
      </v-form>
    </div>
    
  </div>
</template>

<script>
  export default {
    name: 'login',
    methods: {
      login:function(){
        // let postBody = new URLSearchParams(); 
        // postBody.append("userEmail", this.email)
        // postBody.append("userPassword", this.password);
        let postBody = {
          "userEmail": this.email,
          "password": this.password,
        };
        this.$axios
          .post(
            this.$hostname + '/auth/login', 
            postBody,
            {headers:{'Content-Type':'application/json'}}
          ).then(response=>{
            if( response.data != null){
              if( response.status == 200){
                localStorage.setItem("switterJWT",  response.data.jwt);
                localStorage.setItem("switterRT",  response.data.rt);
                this.$router.push({name:'appview'});
              } else {
                this.$router.push({name:'register'});
              }

            }
          });
      },
    },
  }
</script>

<style scoped>
#login{
  
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
.to-register{
  margin-left: auto;
  margin-right: auto;
}
.nav-link-register{
  text-decoration-line: none;
}
label.v-label.theme--dark {
  top: auto;
  margin-top: 5px;
  margin-left: 5px;
}
</style>
