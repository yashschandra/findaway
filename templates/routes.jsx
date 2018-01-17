import React from 'react';
import {Route, IndexRoute, IndexRedirect} from 'react-router';
import Layout from './Layout.jsx';
import AuthPage from './AuthPage.jsx';
import SearchPage from './SearchPage.jsx';
import AddwayPage from './AddwayPage.jsx';
import AddModePage from './AddModePage.jsx';

const routes=(
	<Route path="/" component={Layout}>
		<IndexRedirect to="/search" />
		<Route path="/search" component={SearchPage} />
		<Route path="/auth" component={AuthPage} />
		<Route path="/addway" component={AddwayPage} />
		<Route path="/addmode" component={AddModePage} />
	</Route>
);

export default routes;
