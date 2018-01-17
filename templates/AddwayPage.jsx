import React from 'react';
import axios from 'axios';

class AddwayPage extends React.Component{
	constructor(props){
		super(props);
		this.state={
			fromPlace: 0,
			toPlace: 0,
			modeOptions: [],
			searchResultsFrom: [],
			searchResultsTo: [],
			query: ''
		}
		this.addPlace=this.addPlace.bind(this);
		this.addPlaceFrom=this.addPlaceFrom.bind(this);
		this.addPlaceTo=this.addPlaceTo.bind(this);
		this.addRoute=this.addRoute.bind(this);
		this.querySearch=this.querySearch.bind(this);
		this.handleSearch=this.handleSearch.bind(this);
		this.querySearchFrom=this.querySearchFrom.bind(this);
		this.querySearchTo=this.querySearchTo.bind(this);
		this.setPlaceFrom=this.setPlaceFrom.bind(this);
		this.setPlaceTo=this.setPlaceTo.bind(this);
	}
	componentDidMount(){
		var self=this;
		axios.post('/modes/get', {}).then(function(response){
			if(response.data.status==="OK"){
				self.setState({modeOptions: response.data.modes});
			}
		});
	}
	addPlaceFrom(){
		var place=this.refs.from.value.trim();
		var self=this;
		this.addPlace(place, function(response){
			self.setState({fromPlace: response.id});
		});
	}
	addPlaceTo(){
		var place=this.refs.to.value.trim();
		var self=this;
		this.addPlace(place, function(response){
			self.setState({toPlace: response.id});
		});
	}
	addPlace(placeName, callback){
		axios.post('/places/add', {placeName: placeName, userId: localStorage.getItem("userId")}).then(function(response){
			console.log(response.data);
			callback(response.data);
		});
	}
	querySearchFrom(){
		var self=this;
		var searchTerm=this.refs.from.value.trim();
		this.handleSearch(searchTerm, function(response){
			console.log(response);
			self.setState({searchResultsFrom: response});
		});
	}
	querySearchTo(){
		var self=this;
		var searchTerm=this.refs.to.value.trim();
		this.handleSearch(searchTerm, function(response){
			console.log(response);
			self.setState({searchResultsTo: response});
		});
	}
	querySearch(searchTerm, callback){
		console.log("searchterm", searchTerm);
		if(searchTerm.length>0){
			axios.post('/searchplaces', {search: searchTerm}).then(function(response){
				console.log(response);
				if(response.data.status==='OK'){
					callback(response.data.hits.hits);
				}
			});
		}
	}
	handleSearch(searchTerm, callback){
		this.setState({searchResults: []});
		var self=this;
		window.clearTimeout(this.state.query);
		console.log("cleared query");
		this.setState({query: window.setTimeout(function(){
			console.log(searchTerm, "handling");
			self.querySearch(searchTerm, function(response){
				callback(response);
			});
		}, 1000)});
	}
	componentWillUnmount(){
		window.clearTimeout(this.state.query);	
	}
	setPlaceFrom(place){
		this.setState({fromPlace: String(place.Id)});
		this.refs.from.value=place.Name;
		this.setState({searchResultsFrom: []});
	}
	setPlaceTo(place){
		this.setState({toPlace: String(place.Id)});
		this.refs.to.value=place.Name;
		this.setState({searchResultsTo: []});
	}
	addRoute(){
		var from=this.state.fromPlace;
		var to=this.state.toPlace;
		var distance=this.refs.distance.value.trim();
		var mode=this.refs.mode.value;
		var cost=this.refs.cost.value.trim();
		console.log(from, to, distance, mode, cost);
		axios.post('/routes/add', {fromPlace: from, toPlace: to, distance: distance, userId: localStorage.getItem("userId"), cost: cost, mode: mode}).then(function(response){
			console.log(response);
		});
	}
	render(){
		return(
			<div>
				<input type="text" ref="from" onChange={this.querySearchFrom} />
				{this.state.searchResultsFrom.length>0? (this.state.searchResultsFrom.map((searchResult, i)=><SearchResult key={i} data={searchResult._source} handleClick={this.setPlaceFrom} />)): null}
				<button type="button" onClick={this.addPlaceFrom}>Add place</button>
				<input type="text" ref="to" onChange={this.querySearchTo} />
				{this.state.searchResultsTo.length>0? (this.state.searchResultsTo.map((searchResult, i)=><SearchResult key={i} data={searchResult._source} handleClick={this.setPlaceTo} />)): null}
				<button type="button" onClick={this.addPlaceTo}>Add place</button>
				<input type="number" ref="distance" />
				<select ref="mode">
					{this.state.modeOptions.map((modeOption, i)=><ModeOption key={i} data={modeOption} />)}
				</select>
				<input type="number" ref="cost" />
				<button type="button" onClick={this.addRoute}>Add Route</button>
			</div>
		);
	}
}

class ModeOption extends React.Component{
	constructor(props){
		super(props);
	}
	render(){
		return(
			<option value={this.props.data.Id}>{this.props.data.Name}</option>
		);
	}
}

class SearchResult extends React.Component{
	constructor(props){
		super(props);
		this.handleClick=this.handleClick.bind(this);
	}
	handleClick(){
		this.props.handleClick(this.props.data);
	}
	render(){
		return(
			<div onClick={this.handleClick}>
				{this.props.data.Name}
			</div>
		);
	}
}

export default AddwayPage;
