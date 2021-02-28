<template>
  <div id="threadview" class="container">
    <div class="thread rounded">
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
    <div class="answer-list container">
      <div
        class="message container rounded"
        v-for="message in appmessage"
        :key="message.ID"
      >
        <div class="msg-title container">
          <div class="title-username">{{ message["username"] }}</div>
          <img class="title-avatar" src="../../assets/user.png" />
          <div class="title-date" overline >{{ message["date"] }}</div>
          <div class="answer-button">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil" viewBox="0 0 16 16">
              <path d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5L13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175l-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"></path>
            </svg>
          </div>
        </div>
        <div class="msg-content mb-1 rounded">{{ message["text"] }}</div>
        
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "threadview",
  data() {
    return {
      thread: {"id":6,"text":"This Is Thread","date":"16 Jan 2021 16:37","username":"first","user_id":1},
      comments:[{"id":6,"text":"plane","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":5,"text":"birch","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":4,"text":"milk","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":3,"text":"horse","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":2,"text":"death","date":"16 Jan 2021 16:37","username":"first","user_id":1},{"id":1,"text":"first","date":"16 Jan 2021 16:37","username":"first","user_id":1}]
    };
  },
  props: {
      threadID: Number,
  },
  components: {},
  computed: {},
  created() {},
  mounted() {
      console.log(this.threadID);
  },
};
</script>

<style lang="scss" scoped>
.thread {
  background-color:#140001;
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
    border-color: #2a0000;
    background-color: #110000;
    font-size: 2em;
  }
}
.thread:hover{
  background-color: #220000;
}
.answer-list {
  margin: 5px;
  height: 100%;
  .message {
    background-color:  #0b0822;
    border: 1px solid;
    border-color: #222222;
    color: #ffc;
    margin-top: 5px;
    padding: 5px;
    .msg-title {
      display: flex;
      padding: 0px;
      height: 32px;
      .title-avatar {
        height: 50%;
        width: auto;
        vertical-align: middle;
        border-radius: 50%;
        margin-right: 5px;
      }
      .title-username {
        display: flex;
        margin-left:3px;
        margin-right: 3px;
      }
      a{
        margin-left:3px;
        margin-right: 3px;
        color: #aa0;
      }
      .title-date {
        display: flex;
        margin-left:3px;
        margin-right: 3px;
        color: grey;
        font-weight: 250;
      }
      
    }
    .msg-content {
      padding: 5px;
      border: 1px solid;
      border-color: #00001d;
      background-color: #010008;
      font-weight: 250;
    }
    
  }
  .message:hover{
    background-color: #010012;
  }
}
.answer-button{
        
        height:100%;
        margin-left:3px;
        margin-right: 3px;
      }
</style>
