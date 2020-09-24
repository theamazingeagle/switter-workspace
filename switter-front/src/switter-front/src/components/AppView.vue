<template>
  <div id="appview" >
  <!-- upper bar---------------------------------------- -->
  <v-app-bar color="gray accent-8" dense dark fixed >
      <v-toolbar-title>
          <router-link class="nav-link-register" to="/">
            <v-img src='assets/switter-logo.png' height="48px" width="144px">
            </v-img>
          </router-link>
      </v-toolbar-title>
      <v-spacer>
        <v-btn color="black" id="new-message" v-on:click="newMessageDialog=true">New Message</v-btn>
      </v-spacer>
      <v-btn v-if="accessToken" color="gray" id="logout" v-on:click="Logout">Logout</v-btn>
  </v-app-bar>
<!-- ------------------------------- -->

  <v-card height="48"></v-card>
  <div class="content-center" > 
    <v-card  
      color="gray accent-2"
      class="mx-auto msg"
      min-width="540"
      min-height="160"
      outlined
      v-for="message in appmessage" :key="message.ID">
        <div >
          <div class="msg-title">
            <img class="avatar" src="assets/user.png"></img>
            <div class="msg-username" >{{message.Username}}</div>
            <div class="msg-date" overline mb-4> {{message.Date }}</div>
          </div>
          <div class="msg-content mb-1">{{message.Text}}</div>
        </div>
    </v-card>
     
  </div>
  <!-- end ---------------------------------------- -->
  <v-dialog 
    name ="create-message-modal" 
    v-model="newMessageDialog"
    max-width=50%
  >
    <v-card  class="modal" min-height="160">
      <div v-if="accessToken"> 
        <div class="overline mb-4">type your literals</div> 
        <textarea 
          class="modal-textarea"
          v-model="newMessageBody" 
          autofocus
        >
        </textarea>
        <!-- <div class="overline mb-4" v-if="creatingerror" >error while posting...</div> -->
        <v-btn v-on:click="CreateMessage">Create</v-btn>
      </div>

      <div v-else class="modal" >
        <div class="not-authorized">
          Not Authorized,
          <router-link class="nav-link-register" to="/login"> login</router-link>
          or
          <router-link class="nav-link-register" to="/register">register</router-link>
        </div>
      </div >
    </v-card>
    
  </v-dialog>
</div>

</template>

<script>
export default {
  name: 'appview',
  data(){
    return{
      msgListPage:0, 
      appmessage: [],
      newMessageDialog: false,
      newMessageBody:"",
      creatingerror: false,
      accessToken: localStorage.getItem('switterJWT'),
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
      this.$vuetify.theme.dark = true;
      window.addEventListener('scroll', this.onScroll);
  },
  destroyed () {
    window.removeEventListener('scroll', this.onScroll);
  },
  mounted() {
   
    this.getMessages();
  },
  methods:{
    getMessages:function(){
      this.$axios
        .get(this.$hostname + '/api/getmessages?page='+this.msgListPage)
        .then(response => {
          
          this.appmessage = this.appmessage.concat(response.data);
          this.msgListPage = this.appmessage.length;
          //console.log("###### this.appmessage.length : ", this.appmessage.length);
        });
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
          } else if( responce.status == 401){
            localStorage.removeItem("switterJWT");
            this.$router.push({name:'login'});
            
          }
        });
    },
    Logout:function(){
      localStorage.removeItem("switterJWT");
      this.$router.push({name:'login'});
    },
    onScroll: function () {
      if( Math.max(window.pageYOffset, document.documentElement.scrollTop, document.body.scrollTop) + window.innerHeight === document.documentElement.offsetHeight ) {
        this.getMessages();
      }
      //console.log("... scroll-scroll-scroll ... ");
      //console.log("window.pageYOffset : ", window.pageYOffset);
      //console.log("document.documentElement.offsetHeight : ", document.documentElement.offsetHeight);
      //console.log("document.documentElement.scrollTop : ", document.documentElement.scrollTop);
      //console.log("document.body.scrollTop : ", document.body.scrollTop);
      //console.log("window.innerHeight : ", window.innerHeight);
    }
  }
}

</script>

<style scoped>
  .content-center {
    color:#ffffff;
    display: flex;
    justify-content: center;
    flex-direction: column;
  }
  button {
    background: #009435;
    border: 1px solid #009435;
  }

  .small-container {
    max-width: 680px;
  }
  .not-authorized{
    text-align: center;
    padding: 40px;
    font-size: 130%;
  }
  .modal{

  }
  .msg-title{
    padding-left: 15px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid #4a4140 ;
    background-color: #363636;
    height: 24px;
  }
  .avatar{
    border-radius: 50%;
    height: 16px;
    margin-left,margin-right: 10px;
  }
  .msg-date{
    font-size: 85%;
    padding: 5px;
    color: gray;
  }
  .msg-content{
    display: flex;
    font-size: 150%;
    margin: 10px;
    
  }
  .msg-username {
    font-size: 110%;
    margin: 10px
  }
  .msg{
    margin: 5px;
    background-color: #000;
  }
  .modal-textarea {
    width: 100%;
    height: 240px;
    background-color: #000;
    color: #fff;
    border: solid #343244;
    resize: none;
    font-size: 24px;
  }
</style>
