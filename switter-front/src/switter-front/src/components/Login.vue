<template>
  <div id="login">
    <h2>Login to SWITTER </h2>
    <form>
        <input type="text" v-bind="email">
        <input type="password" v-bind="password">
        <button v-on:click="login"/>
    </form>

  </div>
</template>

<script>
  export default {
    name: 'login',
    methods: {
      login:function(event){
        postBody: {   // пример данных для отправки(позже они преобразуются в json)
        bEmail : email,
        bPassword : password
        };
        this.$axios
          .post('http://172.18.0.1/api/login', JSON.stringify(postBody));
      };
    },
    props: {
        email: String,
        password: String
    },
  }
</script>

<style scoped></style>