<template>
  <div id="register">
    <h2>Register to SWITTER </h2>
    <form>
        <input type="text">
        <input type="password">
        <input type="password">
        <button/>
    </form>

  </div>
</template>

<script>
  export default {
    name: 'register',
    props: {
        email: String,
        password: String
    },
  }
</script>

<style scoped></style>