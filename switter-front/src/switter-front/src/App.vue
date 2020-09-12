<template>
  <v-app>
    <router-view></router-view>
    
  </v-app>
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
      //this.$vuetify.theme.dark = true
  },
  mounted() {
    
  },
}

</script>

<style>
  .content-center {
    
    display: flex;
    justify-content: center;
    flex-direction: row;
  }
  button {
    background: #009435;
    border: 1px solid #009435;
  }

  .small-container {
    max-width: 680px;
  }
</style>
