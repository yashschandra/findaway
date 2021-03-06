import React from 'react';
import axios from 'axios';

class Signup extends React.Component{
	constructor(props){
		super(props);
		this.handleSignup=this.handleSignup.bind(this);
	}
	handleSignup(){
		var self=this;
		var username=this.refs.username.value.trim();
		var password=this.refs.password.value.trim();
		var email=this.refs.email.value.trim();
		console.log(username, password, email);
		axios.post('/auth/signup', {username: username, password: password, email: email}).then(function(response){
			self.props.loginUser(response.data);
		});
	}
	render(){
		return(
			<div>
				Signup
				<input type="text" ref="username" />
				<input type="text" ref="password" />
				<input type="text" ref="email" />
				<button type="button" onClick={this.handleSignup}>Submit</button>
			</div>
		);
	}
}

export default Signup;
