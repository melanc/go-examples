new Vue({
    el: '#app',

    data: {
        ws: null, // websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        host: null, // host address
        username: null, // username
        joined: false // True if email and username have been filled in
    },
	this.ws.addEventListener('message', function(e) {
		var msg = JSON.parse(e.data);
		self.chatContent += '<div class="chip">'
				+ '<img src="' + self.gravatarURL(msg.username) + '">' // Avatar
				+ msg.username
			+ '</div>'
			+ emojione.toImage(msg.message) + '<br/>'; // Parse emojis

		var element = document.getElementById('chat-messages');
		element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
	});
    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + self.host);
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.username) + '">' // Avatar
                    + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
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
			
			// connect websocket
			this.ws = new WebSocket('ws://' + this.host);
			if(this.ws){
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
        },

        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});