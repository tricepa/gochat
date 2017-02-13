import React from 'react';
import ReactDOM from 'react-dom';
import Form from './Form.jsx';

class Chatroom extends React.Component {
	constructor(props) {
		super(props);
    //since only new messages are shown by default, display button to show all messages
    this.state = { messages: [], showButton: true }; 
    this.ws;
    this.initSocket = this.initSocket.bind(this);
    this.sendMessage = this.sendMessage.bind(this);
    this.showAllMessages = this.showAllMessages.bind(this);
    this.onUnload = this.onUnload.bind(this);
	}

  initSocket () {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.onmessage = (msg) => {
      this.state.messages.push(msg.data);
      this.setState({ messages: this.state.messages });
    }
  }

  //trigger logout to update user's last seen timestamp on tab close
  onUnload() {
    fetch("/logout", {
      method: 'POST',
      headers: { "Content-type": "application/x-www-form-urlencoded; charset=UTF-8" },
      body: 'username='+this.props.user.Username
    })
  }

  componentDidMount () {
    window.addEventListener("beforeunload", this.onUnload)
    //show only new messages by default
    fetch("/newMessages", {
      method: 'POST',
      headers: { "Content-type": "application/x-www-form-urlencoded; charset=UTF-8" },
      body: 'username='+this.props.user.Username
    })
    .then((response) => { return response.json() })
    .then((data) => { this.setState({ messages: data.messages || [] }); });
    this.initSocket();

  }
  componentWillUnmount () {
    window.removeEventListener("beforeunload", this.onUnload) 
  }

  getTimestamp () {
    var iso = new Date().toISOString();
    return iso.split("T")[1].split(".")[0];
  }

	sendMessage(message) {
    this.ws.send(
      JSON.stringify({
        username: this.props.user.Username,
        content: (this.getTimestamp() + " [" + this.props.user.Username + "] " + message)
      })
    );
	}

  showAllMessages() {
    fetch("/allMessages")
    .then((response) => { return response.json() })
    .then((data) => { this.setState({ messages: data.messages || [], showButton: false }); });
  }

	render () {
    const showButton = this.state.showButton;
    return (
      <div>
        <h3>Go Forth and React</h3>
        <pre className="chatroom">
          {this.state.messages.join('\n')}
        </pre>
        <Form placeholder="Go on..." onSubmit={this.sendMessage} />
        { showButton && <button className="show-all-messages" onClick={this.showAllMessages}>Show me ALL messages</button> }
      </div>
    )
  }
}

module.exports = Chatroom;