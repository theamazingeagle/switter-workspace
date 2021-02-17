<template>
<div id="threadlist" class="container">
  
  <div class="thread rounded" 
          color="gray accent-2"    
          v-for="thread in threads" :key="thread.ID"
    >      
    <div class="msg-content mb-1 rounded">{{ thread["text"] }}</div>
    <div class="msg-title">
      <div class="answer-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil" viewBox="0 0 16 16">
          <path d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5L13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175l-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"></path>
        </svg>
      </div>
      <div class="user-info">
        <img class="title-avatar" src="../../assets/user.png" />
        <div class="title-username">{{ thread["username"] }}</div>
        <div class="title-date" overline mb-4>{{ thread["date"] }}</div>   
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'threadlist',
  data(){
    return{ 
        threads:[{"id":6,"text":"plane","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":5,"text":"birch","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":4,"text":"milk","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":3,"text":"horse","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":2,"text":"death","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":1,"text":"first","date":"16 Jan 2021 16:37","username":"first","user_id":1}]
    }
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
      });
    },
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

<style lang="scss" scoped>
.thread {
  background-color: #110000;
  border: 1px solid;
  border-color: #aac;
  color: #ffc;
  margin-top: 5px;
  padding: 5px;
  .msg-title {
    height: 100%;
    display: flex;
    padding: 5px;
    justify-content: space-between;
    .user-info{
      display: flex;
      .title-avatar {
          height: 16px;
          width: 16px;
          vertical-align: middle;
          border-radius: 50%;
          margin-right: 5px;
          margin-left:5px;
      }
      .title-username {
          display: flex;
          margin-right: 5px;
          margin-left:5px;
          color:#aa0;
      }
      .title-date {
          color: grey;
          display: flex;
          margin-right: 5px;
          margin-left:5px;
      }
    }
  }
  .msg-content {
    padding: 5px;
    border: 1px solid;
    border-color: #8a0000;
    background-color: #4d0001;
    font-size: 2em;
  }
}
.thread:hover{
  background-color: #220000;
}
</style>
