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
      appmessage: "huu",
    }
  },
  props:{
    //appmessage:String,
  },
  components: {
    //Login,
    //Register,
    //MainPage,
  },
  computed:{
    selectComponent: function(){return "";}
  },
  created () {
      this.$vuetify.theme.dark = true
  },
  mounted() {
    this.$axios
      .get('http://172.18.0.1/api/getmessages')
      .then(response => (this.appmessage = response.data));
  },
}

</script>

<style>
  .content-center {
    color:#ffffff;
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
