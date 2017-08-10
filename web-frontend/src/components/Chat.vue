<!--  npm run build -->
<!--  -->

<template>
<div class="main">
  <div class="row">
        <div class="col s12">
            <div class="card horizontal">
                <div id="chat-messages" class="card-content" v-html="chatContent">
                </div>
            </div>
        </div>
    </div>
  <div class="row" v-if="joined">
    <div class="input-field col s8">
      <input type="text" v-model="currentMsg" @keyup.enter="send">
    </div>
    <div class="input-field col s4">
      <button class="waves-effect waves-light btn" @click="send">
                <i class="material-icons right">chat</i>
                Send
            </button>
    </div>
  </div>
  <div class="row" v-if="!joined">
    <div class="input-field col s8">
      <input type="email" v-model.trim="email" placeholder="Email">
    </div>
    <div class="input-field col s8">
      <input type="text" v-model.trim="username" placeholder="Username">
    </div>
    <div class="input-field col s4">
      <button class="waves-effect waves-light btn" @click="join()">
                Join
            </button>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'chat',
  data() {
    return {
      ws: null,
      currentMsg: '',
      chatContent: '',
      username: '',
      email: '',
      joined: false
    }
  },
  created: function() {
    var self = this;
    this.ws = new WebSocket('ws://' + window.location.host + '/ws');
    this.ws.addEventListener('message', function(e) {
      var msg = JSON.parse(e.data);
      self.chatContent +=  msg.username + ': ' + msg.message + '<br/>';

      // var element = document.getElementById('chat-messages');
      // element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
    });
  },
  methods: {
    send: function() {
      if (this.currentMsg != '') {
        this.ws.send(
          JSON.stringify({
            email: this.email,
            username: this.username,
            message: this.currentMsg // Strip out html
          }));
        this.currentMsg = ''; // Reset newMsg
      }
    },
    join: function() {
      // if (!this.email) {
      //   Materialize.toast('You must enter an email', 2000);
      //   return
      // }
      // if (!this.username) {
      //   Materialize.toast('You must choose a username', 2000);
      //   return
      // }
      this.email = this.email;
      this.username = this.username;
      this.joined = true;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
