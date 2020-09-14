<template>
  <div id="appview">
  <!-- upper bar---------------------------------------- -->
  <v-app-bar color="gray accent-8" dense dark fixed >
      <v-toolbar-title>
          <v-img src='assets/switter-logo.png' height="48px" width="144px">
          </v-img>
          
      </v-toolbar-title>
      <v-spacer>
        <v-btn color="green" id="new-message" v-on:click="newMessageDialog=true">New Message</v-btn>
      </v-spacer>
      <v-btn color="gray" id="logout" v-on:click="Logout">Logout</v-btn>
  </v-app-bar>
<!-- ------------------------------- -->

  <v-card height="48"></v-card>
  <div class="content-center"> 
  <!-- side navigation menu---------------------------------------- -->
  <!--
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
      -->
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
              <div>
                <div class="overline mb-4">{{message.Username}} posted at: {{message.Date }}</div>
                <v-list-item-title class="headline mb-1">{{message.Text}}</v-list-item-title>
              </div>
          </v-card>
        </v-card >
      </v-card>
      <!-- right panel---------------------------------------- -->
      <!--
      <v-card class="d-flex pa-2" height="415" >
        <v-app-bar
          height="400"
          width="256"
        >
        </v-app-bar>
      </v-card>
      -->
      
  <!----------------------------------------------- -->
  </div>
  <!-- end ---------------------------------------- -->
  <v-dialog 
    name ="create-message-modal" 
    v-model="newMessageDialog"
    max-width=50%
  >
    <v-card class="modal"> 
      <div class="overline mb-4">type your literals</div> 
      <v-textarea
        :value="newMessageBody" 
        @change="newMessageBody = $event"
        autofocus="true"
        full-width="true"
        flat="true"
        dark 
        outlined="true"
      >
      </v-textarea>
      <!-- <div class="overline mb-4" v-if="creatingerror" >error while posting...</div> -->
      <v-btn v-on:click="CreateMessage">Create</v-btn>
    </v-card>
  </v-dialog>
</div>

</template>

<script>
export default {
  name: 'appview',
  data(){
    return{ 
      appmessage: "",
      newMessageDialog: false,
      newMessageBody:"",
      creatingerror: false,
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
   
    this.getMessages();
  },
  methods:{
    getMessages:function(){
      console.log("RARARARARARARARARARA HSOTNAEM:", this.$hostname);
      this.$axios
        .get(this.$hostname + '/api/getmessages', 
              {headers:{"Authorization":"Bearer "+localStorage.getItem("switterJWT")
            }})
        .then(response => (this.appmessage = response.data));
    },
    CreateMessageModal:function(){
      //console.log("CreateMessageModal()");
      this.$modal.show('create-message-modal');
    },
    CreateMessage:function(){
      let messageData = new Object();
      messageData.Text = this.newMessageBody;
      messageData.UserID = parseInt(localStorage.getItem("switterUserID") );
      
      this.$axios
        .post( this.$hostname + '/api/createmessage',
        messageData,
        {headers:{"Authorization":"Bearer "+localStorage.getItem("switterJWT")  }}
        ).then(response => {
          if(response.status == 200){
            this.newMessageDialog = false;
            this.newMessageBody = "";
            this.getMessages();
          } else {
            this.creatingerror = true;
          }
        });
    },
    Logout:function(){
      localStorage.removeItem("switterJWT");
      this.$router.push({name:'login'});
    }
  }
}

</script>

<style scoped>
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
  .modal{
    
  }
</style>
