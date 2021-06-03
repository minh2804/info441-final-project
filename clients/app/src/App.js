import React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";

function App() {
	return (
		<Router>
			<Route exact path="/">
				<Home />
			</Route>
			<Route exact path="/home">
				<Home />
			</Route>
			<Route exact path="/signin">
				<SignIn />
			</Route>
			<Route exact path="/signup">
				<SignUp />
			</Route>
			<Route exact path="/stats">
				<Stats />
			</Route>
		</Router>
	);
}

export default App;
