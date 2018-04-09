new Vue({
    el: '#app',

    data: {
        ws: null, // websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatUser: '',
        chatContent: '', // A running list of chat messages displayed on the screen
        host: "127.0.0.1:8000/chat", // host address
        username: null, // username
        joined: false // True if email and username have been filled in
    },

    methods: {
        send: function () {
            if (this.newMsg !== '') {
                this.ws.send(
                    JSON.stringify({
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            var self = this;
            if (!this.host) {
                Materialize.toast('You must enter a host', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.host = $('<p>').html(this.host).text();
            this.username = $('<p>').html(this.username).text();
			
			// connect to server
            var url = 'ws://' + this.host;
			this.ws = new WebSocket(url);
			if(this.ws != null){
			    self.chatUser += this.username;
			    // 注册事件，监听消息
				this.ws.addEventListener('message', function(e) {
				var msg = JSON.parse(e.data);
				self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
					+ msg.username
					+ '</div>'
					+ emojione.toImage(msg.message) + '<br/>'; // Parse emojis

                    var element = document.getElementById('chat-messages');
                    element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
                });
				this.joined = true;
			}
			else{
			    console.log("can't connect to", url)
            }
        },

        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});