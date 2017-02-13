import React from 'react';
import ReactDOM from 'react-dom';
import Form from './components/Form.jsx';
import Chatroom from './components/Chatroom.jsx';


class App extends React.Component {
	constructor(props) {
		super(props);
		this.state = { user: null };
		this.setUsername = this.setUsername.bind(this);
	}

	setUsername(username) {
		fetch('/login', {
			method: 'POST',
			headers: { "Content-type": "application/x-www-form-urlencoded; charset=UTF-8" },
			body: 'username='+username
		})
		.then((response) => { return response.json(); })
		.then((responseJson) => { this.setState({ user: responseJson.user }); });
	}

	render() {
		let loginForm = (
			<div>
				<h1>Join the conversation!</h1>
				<Form placeholder = "Enter username" onSubmit = {this.setUsername} />
			</div>
		)
		let chatroom = <Chatroom user = {this.state.user} />
		if (this.state.user) {
			return chatroom;
		} else {
			return loginForm;
		}
	}
}

ReactDOM.render( < App / > ,
	document.getElementById('root')
)