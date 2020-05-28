<template>
  <div id="appview">
  <!-- upper bar---------------------------------------- -->
  <v-app-bar color="gray accent-8" dense dark fixed >
      <v-toolbar-title>
          <v-img src='../assets/switter-logo.png' height="48px" width="144px">
          </v-img>
      </v-toolbar-title>
      <v-spacer></v-spacer>    
  </v-app-bar>
<!-- ------------------------------- -->
  <v-card height="48"></v-card>
  <div class="content-center"> 
  <!-- side navigation menu---------------------------------------- -->
      <v-card class="d-flex pa-2" height="415" >  
        <v-card
          height="400"
          width="256"
        >
          <v-navigation-drawer permanent>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title class="title">
                  Application
                </v-list-item-title>
                <v-list-item-subtitle>
                  subtext
                </v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>

            <v-divider></v-divider>

            <v-list dense nav>
              <v-list-item>
                <v-list-item-content>
                  <li class="nav-item">
                    <router-link class="nav-link" to="/home">Home</router-link>
                  </li>
                </v-list-item-content>
              </v-list-item>

              <v-list-item>
                <v-list-item-content>
                  <li class="nav-item">
                    <router-link class="nav-link" to="/home">User</router-link>
                  </li>
                </v-list-item-content>
              </v-list-item>

              <v-list-item>
                <v-list-item-content>
                  <li class="nav-item">
                    <router-link class="nav-link" to="/home">Feed</router-link>
                  </li>
                </v-list-item-content>
              </v-list-item>

            </v-list>
          </v-navigation-drawer>
        </v-card>
      
      </v-card>
      <!-- message list---------------------------------------- -->
      <v-card class="d-flex pa-2">
        <v-card >
          <v-card  
            color="gray accent-2"
            class="mx-auto"
            min-width="540"
            min-height="160"
            outlined
            v-for="message in appmessage" :key="message.ID">
              <v-card-title>Author</v-card-title>
              <v-card-text >{{message.Text}}</v-card-text>
              <v-card-actions>comment out</v-card-actions>
              
          </v-card>
        </v-card >
      </v-card>
      <!-- right panel---------------------------------------- -->
      <v-card class="d-flex pa-2" height="415" >
        <v-app-bar
          height="400"
          width="256"
        >
        </v-app-bar>
      </v-card>
  <!----------------------------------------------- -->
  </div>
  <!-- end ---------------------------------------- -->
</div>
</template>

<script>
export default {
  name: 'appview',
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
