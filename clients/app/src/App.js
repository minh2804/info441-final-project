import React, {
	useEffect,
	useState
} from 'react';
import {
	NavLink,
	Redirect,
	Route,
	Switch,
	BrowserRouter as Router
} from "react-router-dom";
import {
	NavItem,
	Nav as Nav_
} from 'reactstrap';
import { requestGetTodoList } from './api/tasks'
import Home from './pages/Home/Home'
import SignIn from './pages/SignIn/SignIn'
import SignOut from './pages/SignOut/SignOut'
import SignUp from './pages/SignUp/SignUp'
import User from './pages/User/User'
import './App.css';

const NAV_ACTIVE_CLASS_NAME = "border-bottom border-3"
const NAV_CLASS_NAME = "h4 p-3 text-decoration-none"

export default function App() {
	const [sessionState, setSessionState] = useState({
		user: null,
		todoList: []
	});

	const handleUserState = (newUser) => setSessionState({ ...sessionState, user: newUser });
	const handleTodoListState = (newTodoList) => setSessionState({ ...sessionState, todoList: newTodoList });

	useEffect(() => {
		async function invokeAPIRequest() {
			const newTodoList = await requestGetTodoList();
			handleTodoListState(newTodoList);
		}
		invokeAPIRequest();
	}, []);

	const renderSignInPage = () => sessionState.user ? <Redirect push to="/" /> : <SignIn onSignIn={setSessionState} />;
	const renderSignOutPage = () => sessionState.user ? <SignOut onSignOut={setSessionState} /> : <Redirect push to="/" />;
	const renderSignUpPage = () => sessionState.user ? <Redirect push to="/" /> : <SignUp onSignUp={setSessionState} />;
	const renderUserEditPage = () => sessionState.user ? <User user={sessionState.user} onChange={handleUserState} /> : <Redirect push to="/signin" />;

	return (
		<Router>
			<Nav className="p-3" user={sessionState.user} />
			<Switch>
				<Route exact path="/">
					<Home user={sessionState.user} todoList={sessionState.todoList} onChange={handleTodoListState} />
				</Route>
				<Route exact path="/home">
					<Home user={sessionState.user} todoList={sessionState.todoList} onChange={handleTodoListState} />
				</Route>
				<Route exact path="/signin" render={renderSignInPage} />
				<Route exact path="/signout" render={renderSignOutPage} />
				<Route exact path="/signup" render={renderSignUpPage} />
				<Route exact path="/user/edit" render={renderUserEditPage} />
			</Switch>
		</Router>
	);
}

function Nav({ user, ...rest }) {
	return user ? <AuthenticatedNav user={user} {...rest} /> : <UnauthenticatedNav {...rest} />;
}

function AuthenticatedNav({ user, ...rest }) {
	return (
		<Nav_ {...rest}>
			<HomeLink />
			<SignOutLink />
			<UserLink user={user} />
		</Nav_ >
	);
}

function UnauthenticatedNav(props) {
	return (
		<Nav_ {...props}>
			<HomeLink />
			<SignInLink />
			<SignUpLink />
		</Nav_>
	);
}

function HomeLink(props) {
	return (
		<NavItem {...props}>
			<NavLink activeClassName={NAV_ACTIVE_CLASS_NAME} className={NAV_CLASS_NAME} to="/home">Home</NavLink>
		</NavItem>
	);
}

function SignInLink(props) {
	return (
		<NavItem {...props}>
			<NavLink activeClassName={NAV_ACTIVE_CLASS_NAME} className={NAV_CLASS_NAME} to="/signin">Sign in</NavLink>
		</NavItem>
	);
}

function SignUpLink(props) {
	return (
		<NavItem {...props}>
			<NavLink activeClassName={NAV_ACTIVE_CLASS_NAME} className={NAV_CLASS_NAME} to="/signup">Sign up</NavLink>
		</NavItem>
	);
}
function SignOutLink(props) {
	return (
		<NavItem {...props}>
			<NavLink activeClassName={NAV_ACTIVE_CLASS_NAME} className={NAV_CLASS_NAME} to="/signout">Sign out</NavLink>
		</NavItem>
	);
}

function UserLink({ user, ...rest }) {
	return (
		<NavItem {...rest}>
			<NavLink activeClassName={NAV_ACTIVE_CLASS_NAME} className={NAV_CLASS_NAME} to="/user/edit">{'Edit user: ' + user.firstName + ' ' + user.lastName}</NavLink>
		</NavItem>
	);
}
