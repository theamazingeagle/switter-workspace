<template>
  <div class="appview container row" id="appview" >
    <!-- upper bar---------------------------------------- -->
    
  <nav class="container col-3 rounded sticky-top">
    <a >
    <router-link  to="/">
        <img  src="../assets/switter-logo-2.png" height="48px" width="144px">
    </router-link>
    </a>
    <div class="nav-item">
        <a href="#"  color="black" id="new-message" v-on:click="newMessageDialog=true">New Thread</a>
    </div>
    <div class="nav-item">
        <a href="#"  v-if="accessToken" color="gray" id="logout" v-on:click="Logout">Logout</a>
    </div>
      
  </nav>
   
  <!-- ------------------------------- -->
  <div class="col"></div>
  <main class="container col-8">
   <router-view></router-view>
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

.appview {

  nav{
    margin: 10px;
    text-align: center;
    border-color: #300000;
    border: 1px solid;
    background-color: #100000;
    color: #556677;
    height: 100%;
    div{
        padding: 5px;
        border-top: 1px gray solid;
        a{
            color: #556677;
            text-decoration: none;
        }
    }
  }
  main {
    display: flex;
    justify-content: space-between;
    padding: 5px;
  }
}
</style>
