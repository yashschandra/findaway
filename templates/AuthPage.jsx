import React from 'react';
import {browserHistory} from 'react-router';
import Signin from './Signin.jsx';
import Signup from './Signup.jsx';

class AuthPage extends React.Component{
	constructor(props){
		super(props);
		this.state={
			authState: 'signin'
		}
		this.handleAuthState=this.handleAuthState.bind(this);
		this.loginUser=this.loginUser.bind(this);
	}
	componentWillMount(){
		console.log(localStorage.getItem("verified")==="1");
		if(localStorage.getItem("verified")!==undefined && localStorage.getItem("verified")==="1"){
			browserHistory.replace("/search");
		}
		else{
			localStorage.setItem("verified","0");
		}
	}
	loginUser(response){
		if(response.status==='OK'){
			localStorage.setItem("verified","1");
			localStorage.setItem("user",response.user);
			localStorage.setItem("userId",response.id);
			browserHistory.replace("/search");
			window.location.reload();
		}
	}
	handleAuthState(val){
		this.setState({authState: val});
	}
	render(){
		return(
			<div>
				<button type="button" onClick={()=>this.handleAuthState('signin')} >Signin</button>
				<button type="button" onClick={()=>this.handleAuthState('signup')} >Signup</button>
				{this.state.authState==="signin"?<Signin loginUser={this.loginUser} />:null}
				{this.state.authState==="signup"?<Signup loginUser={this.loginUser} />:null}
			</div>
		);
	}
}

export default AuthPage;
