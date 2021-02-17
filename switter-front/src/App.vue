<template>
  <div class="root"> 
   <router-view></router-view>
</div>
</template>

<script>
export default {
  name: 'App',
  data(){
    return{ 
    }
  },
  props:{
  },
  components: {
  },
  computed:{
    
  },
  created () {

  },
  mounted() {
  },
}
</script>

<style>
body {
  background-color: black;
  background-image: url(assets/switter-background-1.png);
  background-size: 20vh;
  background-repeat: repeat;
  padding: 0;
  margin: 0;
}
</style>
