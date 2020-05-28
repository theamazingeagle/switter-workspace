<template>
  <div id="login">
    <h2>Login to SWITTER </h2>
    <form>
        <input type="text" v-bind="email">
        <input type="password" v-bind="password">
        <button/>
    </form>

  </div>
</template>

<script>
  export default {
    name: 'login',
    props: {
        email: String,
        password: String
    },
  }
</script>

<style scoped></style>