import React from 'react';
import axios from 'axios';

class Signin extends React.Component{
	constructor(props){
		super(props);
		this.handleSignin=this.handleSignin.bind(this);
	}
	handleSignin(){
		var self=this;
		var username=this.refs.username.value.trim();
		var password=this.refs.password.value.trim();
		axios.post('/auth/signin', {username: username, password: password}).then(function(response){
			self.props.loginUser(response.data);
		});
	}
	render(){
		return(
			<div>
				Signin
				<input type="text" ref="username" />
				<input type="text" ref="password" />
				<button type="button" onClick={this.handleSignin}>Submit</button>
			</div>
		);
	}
}

export default Signin;
