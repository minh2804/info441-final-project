import React, {
	useEffect,
	useState
} from 'react';
import {
	Link,
	Redirect,
	Route,
	Switch,
	BrowserRouter as Router,
} from "react-router-dom";
import {
	NavItem,
	Nav as Nav_
} from 'reactstrap';
import Home from './pages/Home/Home'
import SignIn from './pages/SignIn/SignIn'
import SignOut from './pages/SignOut/SignOut'
import SignUp from './pages/SignUp/SignUp'
import User from './pages/User/User'
import { requestGetTodoList } from './api/tasks'
import './App.css';

export default function App() {
	const [sessionState, setSessionState] = useState({
		user: null,
		todoList: []
	});
	const handleUserState = (newUser) => setSessionState({ ...sessionState, user: newUser });
	const handleTodoListState = (newTodoList) => setSessionState({ ...sessionState, todoList: newTodoList });
	useEffect(() => {
		async function invokeAPIRequest() {
			const responseBody = await requestGetTodoList();
			handleTodoListState(responseBody);
		}
		invokeAPIRequest();
	}, []);
	return (
		<Router>
			<Nav user={sessionState.user} />
			<Switch>
				<Route exact path="/">
					<Home user={sessionState.user} todoList={sessionState.todoList} onChange={handleTodoListState} />
				</Route>
				<Route exact path="/home">
					<Home user={sessionState.user} todoList={sessionState.todoList} onChange={handleTodoListState} />
				</Route>
				<Route exact path="/signin">
					{sessionState.user ? <Redirect push to="/" /> : <SignIn onSignIn={setSessionState} />}
				</Route>
				<Route exact path="/signup">
					<SignUp onSignUp={setSessionState} />
				</Route>
				<Route exact path="/signout" render={() => <SignOut onSignOut={setSessionState} />} />
				<Route exact path="/user/edit">
					{sessionState.user ? <User user={sessionState.user} onChange={handleUserState} /> : <Redirect push to="/signin" />}
				</Route>
			</Switch>
		</Router>
	);
}

export function Nav({ user, ...rest }) {
	return user ? <AuthenticatedNav user={user} {...rest} /> : <UnAuthenticatedNav {...rest} />;
}

export function AuthenticatedNav({ user, ...rest }) {
	return (
		<Nav_ {...rest}>
			<HomeLink />
			<SignOutLink />
			<UserLink user={user} />
		</Nav_ >
	);
}

export function UnAuthenticatedNav(props) {
	return (
		<Nav_ {...props}>
			<HomeLink />
			<SignInLink />
			<SignUpLink />
		</Nav_>
	);
}

export function SignInLink(props) {
	return (
		<NavItem {...props}>
			<Link className="btn" to="/signin">Sign in</Link>
		</NavItem>
	);
}

export function SignUpLink(props) {
	return (
		<NavItem {...props}>
			<Link className="btn" to="/signup">Sign up</Link>
		</NavItem>
	);
}
export function SignOutLink(props) {
	return (
		<NavItem {...props}>
			<Link className="btn" to="/signout">Sign out</Link>
		</NavItem>
	);
}

export function UserLink({ user, ...rest }) {
	return (
		<NavItem {...rest}>
			<Link className="btn" to="/user/edit">{'Edit user: ' + user.firstName + ' ' + user.lastName}</Link>
		</NavItem>
	);
}

export function HomeLink(props) {
	return (
		<NavItem {...props}>
			<Link className="btn" to="/home">Home</Link>
		</NavItem>
	);
}
