<template>
  <div class="appview " id="appview" >
    <!-- upper bar---------------------------------------- -->
    
      <nav class="navbar fixed-top navbar-expand-lg">
          <a class="navbar-brand">
            <router-link class="nav-link-register" to="/">
              <img class="logo" src="../assets/switter-logo.png" height="48px" width="144px">
            </router-link>
          </a>
          <ul class="navbar-nav">
              <li class="nav-item">
                  <b-button color="black" id="new-message" v-on:click="newMessageDialog=true">New Message</b-button>
              </li>
              <li class="nav-item">
                  <b-button v-if="accessToken" color="gray" id="logout" v-on:click="Logout">Logout</b-button>
              </li>
          </ul>
      </nav>
   
  <!-- ------------------------------- -->
  <main class="container">
    <div class="thread container">
      <div class="message container" 
            color="gray accent-2"    
            v-for="message in appmessage" :key="message.ID"
      >
        <div>
          <div class="msg-title">
            <img class="avatar" src="../assets/user.png">
              <div class="msg-username" >{{message['username']}}</div>
              <div class="msg-date" overline mb-4> {{message['date'] }}</div>
            </div>
            <div class="msg-content mb-1">{{message['text']}}</div>
          </div>
      </div>
    </div>
  </main>
      
    <!-- end ---------------------------------------- -->
    <!-- <v-dialog 
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
          <b-button v-on:click="CreateMessage">Create</b-button>
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
      
    </v-dialog> -->
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
      this.$http
        .get(this.$hostname + '/api/message/all?page='+this.msgListPage,
          {headers:{
            "Authorization":"Bearer "+localStorage.getItem("switterJWT")}}
          )
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
      //console.log("JJJJJJJJJJJWT: ", localStorage.getItem("switterJWT"));
      let messageData = new Object();
      messageData.Text = this.newMessageBody;
      //messageData.UserID = parseInt(localStorage.getItem("switterUserID") );
      
      this.$http
        .post( this.$hostname + '/api/message/create',
          messageData,
          {headers:{
            "Authorization":"Bearer "+localStorage.getItem("switterJWT"),
            'Content-Type':'application/json',
          }}
        ).then((response) => {
          console.log("~~~ reading response ...");
          if(response.status == 200){
            this.newMessageDialog = false;
            this.newMessageBody = "";
            //this.getMessages();
            document.location.reload();
          }
        },
        (response)=>{
          console.log("~~~ trying update...");
            let message = {
                  "jwt":localStorage.getItem("switterJWT"),
                  "rt":localStorage.getItem("switterRT"),
                };

            this.$http
              .post( this.$hostname + '/api/auth/refresh',
                message,
                {
                  headers:{'Content-Type':'application/json',
                }}
            ).then( (response) => {

              if (response.status == 200){
                console.log("~~~ update success");
                console.log("OOOOHMYYYYY : ",response.data);
                localStorage.setItem("switterJWT",  response.data.jwt);
                localStorage.setItem("switterRT",  response.data.rt);
                this.CreateMessage();
              }
            },
            (response) => {
              console.log("~~~ update fail");
              let message = {
                  "jwt":localStorage.getItem("switterJWT"),
                  "rt":localStorage.getItem("switterRT"),
                };
                this.$http
                  .post( this.$hostname + '/api/auth/logout',
                  message,
                  {
                    headers:{'Content-Type':'application/json',
                  }}
                  );
                //localStorage.removeItem("switterJWT");
                this.$router.push({name:'login'});
            });
        });
    },
    Logout:function(){
      this.$http
        .post( this.$hostname + '/api/auth/logout',
        {
          "jwt":localStorage.getItem("switterJWT"),
          "rt":localStorage.getItem("switterRT"),
        },
        {
          headers:{'Content-Type':'application/json',
        }}
        );
      localStorage.removeItem("switterJWT");
      this.$router.push({name:'login'});
    },
    onScroll: function () {
      if( Math.max(window.pageYOffset, document.documentElement.scrollTop, document.body.scrollTop) + window.innerHeight === document.documentElement.offsetHeight ) {
        this.getMessages();
      }
    }
  }
}

</script>

<style lang="scss" scoped>
$app-width: 1360px;
$thread-block-width: 800px;

.appview {
    background-color: #110000;
    
  nav{
    background-color: #430000;
    padding:0;
  }
  main {
    display: flex;
    justify-content: space-between;
    padding: 5px;
    .thread {
      margin-top: 50px;
      height: 100%;
      width: $thread-block-width;
      .message {
        background-color: #110000;
        color: #ffc;
        margin-top: 5px;
        padding: 5px;
      }
    }
    .list {
      background-color: #500000;
      width: 300px;
      height: 500px;
    }
  }
}
</style>
