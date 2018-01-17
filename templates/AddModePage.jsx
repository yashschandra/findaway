import React from 'react';
import axios from 'axios';

class AddModePage extends React.Component{
	constructor(props){
		super(props);
		this.addMode=this.addMode.bind(this);
	}
	addMode(){
		var self=this;
		var mode=this.refs.mode.value.trim();
		axios.post('/modes/add', {userId: localStorage.getItem("userId"), modeName: mode}).then(function(response){
			console.log(response.data);
		});
	}
	render(){
		return(
			<div>
				<input type="text" ref="mode" />
				<button type="button" onClick={this.addMode}>Add</button>
			</div>
		);
	}
}

export default AddModePage;
