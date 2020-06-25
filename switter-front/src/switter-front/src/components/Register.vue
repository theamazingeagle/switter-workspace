<template>
  <div id="register">
    <div class="center">
      <v-img src='../assets/switter-logo.png'  ></v-img>
      <!-- <h2 class="header">register to SWITTER </h2> -->
      <v-form class="form">
        <v-row >
          <v-text-field type="text"  v-model="username" label="Your Name"></v-text-field>
        </v-row>
        <v-row >
          <v-text-field type="text"  v-model="email" label="E-mail"></v-text-field>
        </v-row>
        <v-row >
          <v-text-field type="password" v-model="password" label="Password"></v-text-field>
        </v-row>
          <v-btn class="btn-register" v-on:click="register">Sign Up</v-btn>
      </v-form>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'register',
    methods: {
      register:function(){
        let postBody = new URLSearchParams(); 
        postBody.append("userName", this.username)
        postBody.append("userEmail", this.email)
        postBody.append("userPassword", this.password);
        
        this.$axios
          .post(
            'http://172.18.0.1/api/register', 
            postBody,
            {headers:{'Content-Type':'application/x-www-form-urlencoded'}}
          ).then(response=>{
            if(response.status == 200) {
              //localStorage.setItem("jwt",  response.data);
              //this.$router.push({name:'appview'});
              this.$router.push({name:'login'});
            }
            
          });
      },
    },
    data: function() {
      return {
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
