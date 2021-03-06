import React from 'react';
import {Link} from 'react-router';

class Layout extends React.Component{
	constructor(props){
		super(props);
		this.state={
			loggedIn: false
		}
		this.logoutUser=this.logoutUser.bind(this);
	}
	logoutUser(){
		localStorage.setItem("verified","0");
		localStorage.setItem("user",null);
		localStorage.setItem("userId",null);
		this.setState({loggedIn: false});
		window.location.reload();
	}
	componentWillMount(){
		console.log(localStorage.getItem("verified")==="1");
		if(localStorage.getItem("verified")!==undefined && localStorage.getItem("verified")==="1"){
			this.setState({loggedIn: true});
		}
		else{
			localStorage.setItem("verified","0");
		}
	}
	render(){
		return(
			<div>
				<TopBar loggedIn={this.state.loggedIn} logoutUser={this.logoutUser}/>
				{this.props.children}
				<BottomBar />
			</div>
		);
	}
}

class TopBar extends React.Component{
	constructor(props){
		super(props);
	}
	render(){
		return(
			<div>
				Top
				{this.props.loggedIn?<LoggedInHead logoutUser={this.props.logoutUser} />: <AuthHead />}
			</div>
		);
	}
}

class BottomBar extends React.Component{
	constructor(props){
		super(props);
	}
	render(){
		return(
			<div>
				Bottom
			</div>
		);
	}
}

class LoggedInHead extends React.Component{
	constructor(props){
		super(props);
	}
	render(){
		return(
			<div>
				<Link to="/addway">Add Route</Link>
				<Link to="/addmode">Add Mode</Link>
				<Link to="/search">Search</Link>
				<button type="button" onClick={this.props.logoutUser}>Logout</button>
			</div>
		);
	}
}

class AuthHead extends React.Component{
	constructor(props){
		super(props);
	}
	render(){
		return(
			<div>
				<Link to="/auth">Auth</Link>
				<Link to="/search">Search</Link>
			</div>
		);
	}
}

export default Layout;
