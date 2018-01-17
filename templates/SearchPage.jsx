import React from 'react';
import axios from 'axios';
import {Link} from 'react-router';
import {browserHistory} from 'react-router';

class SearchPage extends React.Component{
	constructor(props){
		super(props);
		this.state={
			fromPlace: 0,
			toPlace: 0,
			searchResultsFrom: [],
			searchResultsTo: [],
			query: '',
			pathFound: false,
			path: []
		}
		this.querySearch=this.querySearch.bind(this);
		this.handleSearch=this.handleSearch.bind(this);
		this.querySearchFrom=this.querySearchFrom.bind(this);
		this.querySearchTo=this.querySearchTo.bind(this);
		this.setPlaceFrom=this.setPlaceFrom.bind(this);
		this.setPlaceTo=this.setPlaceTo.bind(this);
		this.searchWay=this.searchWay.bind(this);
	}
	querySearchFrom(){
		var self=this;
		var searchTerm=this.refs.from.value.trim();
		this.handleSearch(searchTerm, function(response){
			self.setState({searchResultsFrom: response});
		});
	}
	querySearchTo(){
		var self=this;
		var searchTerm=this.refs.to.value.trim();
		this.handleSearch(searchTerm, function(response){
			self.setState({searchResultsTo: response});
		});
	}
	querySearch(searchTerm, callback){
		console.log("searchterm", searchTerm);
		if(searchTerm.length>0){
			axios.post('/searchplaces', {search: searchTerm}).then(function(response){
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
		this.setState({query: window.setTimeout(function(){
			self.querySearch(searchTerm, function(response){
				callback(response);
			});
		}, 1000)});
	}
	componentWillUnmount(){
		window.clearTimeout(this.state.query);	
	}
	setPlaceFrom(place){
		this.setState({fromPlace: place.Id});
		this.refs.from.value=place.Name;
		this.setState({searchResultsFrom: []});
	}
	setPlaceTo(place){
		this.setState({toPlace: place.Id});
		this.refs.to.value=place.Name;
		this.setState({searchResultsTo: []});
	}
	searchWay(){
		var self=this;
		axios.post('/way', {from: this.state.fromPlace, to: this.state.toPlace}).then(function(response){
			console.log("way", response);
			self.setState({pathFound: true, path: response.data.path});
		});
	}
	render(){
		return(
			<div>
				<input type="text" ref="from" onChange={this.querySearchFrom} />
				{this.state.searchResultsFrom.length>0? (this.state.searchResultsFrom.map((searchResult, i)=><SearchResult key={i} data={searchResult._source} handleClick={this.setPlaceFrom} />)): null}
				<input type="text" ref="to" onChange={this.querySearchTo} />
				{this.state.searchResultsTo.length>0? (this.state.searchResultsTo.map((searchResult, i)=><SearchResult key={i} data={searchResult._source} handleClick={this.setPlaceTo} />)): null}
				<button type="button" onClick={this.searchWay}>Search</button>
				{this.state.pathFound? <ShowPath data={this.state.path} />: null}
			</div>
		);
	}
}

class ShowPath extends React.Component{
	constructor(props){
		super(props);
	}
	componentWillReceiveProps(nextProps){
		console.log("show path", nextProps, this.props==nextProps);
	}
	render(){
		return(
			<div>
				{this.props.data!==null? <ShowRoute data={this.props.data} />: <div>No route.</div>}
			</div>
		);
	}
}

class ShowRoute extends React.Component{
	constructor(props){
		super(props);
	}
	componentWillReceiveProps(nextProps){
		console.log("show route", nextProps, this.props==nextProps);
	}
	render(){
		return(
			<div>
				{this.props.data.map((place, i)=><Place key={i} data={place} />)}
			</div>
		);
	}
}

class Place extends React.Component{
	constructor(props){
		super(props);
	}
	render(){
		return(
			<div>
				{this.props.data.FromPlace} -> {this.props.data.ToPlace} @ {this.props.data.Cost} | {this.props.data.Distance} by {this.props.data.Mode}
			</div>
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

export default SearchPage;
