import React, { useState, useEffect } from 'react';
import { Switch, Route, Link, Redirect, BrowserRouter as Router } from "react-router-dom";
import { Button, Container, Form, FormGroup, Label, Input, NavItem, Nav as Nav_, } from 'reactstrap';
import Popup from 'reactjs-popup';
import './App.css';

function App() {
	const [sessionState, setSessionState] = useState({
		user: null,
		todoList: [],
	});
	const handleUser = (newUser) => setSessionState({ ...sessionState, user: newUser });
	const handleTodoList = (newTodoList) => setSessionState({ ...sessionState, todoList: newTodoList });

	useEffect(() => {
		async function invokeAPIRequest() {
			const responseBody = await requestGetTodoList();
			handleTodoList(responseBody);
		}
		invokeAPIRequest();
	}, []);

	return (
		<Router>
			<Nav user={sessionState.user} />
			<Switch>
				<Route exact path="/">
					<TodoList tasks={sessionState.todoList} onChange={handleTodoList} />
				</Route>
				<Route exact path="/home">
					<TodoList tasks={sessionState.todoList} onChange={handleTodoList} />
				</Route>
				<Route exact path="/signin">
					{sessionState.user ? <Redirect push to="/" /> : <SignIn onSignIn={setSessionState} />}
				</Route>
				<Route exact path="/signup">
					<SignUp onSignUp={setSessionState} />
				</Route>
				<Route exact path="/signout" render={() => <SignOut onSignOut={setSessionState} />} />
			</Switch>
		</Router>
	);
}

function Nav({ user }) {
	return (
		<Nav_>
			<NavItem>
				<Link to="/home">Home</Link>
			</NavItem>
			{user ? (
				<div>
					<NavItem>
						<Link to="/signout">Sign out</Link>
					</NavItem>
					Welcome {user.firstName + ' ' + user.lastName}!
				</div>
			) : (
				<div>
					<NavItem>
						<Link to="/signin">Sign in</Link>
					</NavItem>
					<NavItem>
						<Link to="/signup">Sign up</Link>
					</NavItem>
				</div>
			)}
		</Nav_>
	);
}

function SignIn({ onSignIn, ...rest }) {
	const [form, setForm] = useState({
		username: "",
		password: "",
	});
	const handleTextInput = (event) => setForm({ ...form, [event.target.name]: event.target.value });
	const handleSubmit = (event) => {
		event.preventDefault();
		async function invokeAPIRequest() {
			const userBody = await requestPostSession(form);
			const todoListBody = await requestGetTodoList();
			onSignIn({ user: userBody, todoList: todoListBody });
		}
		invokeAPIRequest();
	};

	return (
		<Form onSubmit={handleSubmit} {...rest}>
			<FormGroup>
				<Label for="username">Username</Label>
				<Input id="username" name="username" type="text" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup>
				<Label for="password">Password</Label>
				<Input id="password" name="password" type="password" onChange={handleTextInput} />
			</FormGroup>
			<Input type="submit" value="Submit" />
		</Form>
	);
}

function SignUp({ onSignUp, ...rest }) {
	const [form, setForm] = useState({
		username: "",
		password: "",
		passwordConf: "",
		firstName: "",
		lastName: "",
		isTemporary: false,
	});
	const handleTextInput = (event) => setForm({ ...form, [event.target.name]: event.target.value });
	const handleSubmit = (event) => {
		event.preventDefault();
		async function invokeAPIRequest() {
			const userBody = await requestPostUser(form);
			const todoListBody = await requestGetTodoList();
			onSignUp({ user: userBody, todoList: todoListBody });
		}
		invokeAPIRequest();
	};

	return (
		<Form onSubmit={handleSubmit} {...rest}>
			<FormGroup>
				<Label for="username">Username</Label>
				<Input id="username" name="username" type="text" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup>
				<Label for="password">Password</Label>
				<Input id="password" name="password" type="password" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup>
				<Label for="passwordConf">Password Confirm</Label>
				<Input id="passwordConf" name="passwordConf" type="password" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup>
				<Label for="firstName">First Name</Label>
				<Input id="firstName" name="firstName" type="text" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup>
				<Label for="lastName">Last Name</Label>
				<Input id="lastName" name="lastName" type="text" onChange={handleTextInput} />
			</FormGroup>
			<Input type="submit" value="Submit" />
		</Form>
	);
}

function SignOut({ onSignOut }) {
	async function invokeAPIRequest() {
		await requestDeleteSession();
		const todoListBody = await requestGetTodoList();
		onSignOut({ user: null, todoList: todoListBody });
	}
	invokeAPIRequest();
	return <Redirect push to="/" />;
}

function TodoList({ tasks, onChange }) {
	const handleAddTask = (task) => {
		async function invokeAPIRequest() {
			const responseBody = await requestPostTask(task);
			onChange([...tasks, responseBody]);
		}
		invokeAPIRequest();
	};
	const handleDeleteTask = (event) => {
		const taskID = parseInt(event.target.value);
		async function invokeAPIRequest() {
			await requestDeleteTask(taskID);
			onChange(tasks.filter(task => task.id !== taskID));
		}
		invokeAPIRequest();
	};
	const renderedRows = tasks.map(task => <Task key={task.id} content={task} onDelete={handleDeleteTask} />);
	return (
		<Container>
			<table>
				<thead>
					<tr>
						<th></th>
						<th>ID</th>
						<th>Name</th>
						<th>Description</th>
						<th>Complete</th>
						<th>Hidden</th>
						<th>Created At</th>
						<th>Edited At</th>
					</tr>
				</thead>
				<tbody>
					{renderedRows}
				</tbody>
			</table>
			<AddTaskPopUpButton onSubmit={handleAddTask}>Add Task</AddTaskPopUpButton>
		</Container>
	);
}

function AddTaskPopUpButton({ onSubmit, children, ...rest }) {
	const [open, setOpen] = useState(false);
	const toggleOpen = () => setOpen(!open);
	const closeModal = () => setOpen(false);
	const handleSubmit = (newTask) => {
		onSubmit(newTask);
		closeModal();
	}
	return (
		<div>
			<Button type="button" onClick={toggleOpen} {...rest}>{children}</Button>
			<Popup open={open} closeOnDocumentClick onClose={closeModal} >
				<NewTaskForm onSubmit={handleSubmit} onCancel={closeModal} className="container bg-white" />
			</Popup >
		</div >
	);
}

function NewTaskForm({ onSubmit, onCancel, ...rest }) {
	const [form, setForm] = useState({
		name: "",
		description: "",
		isHidden: false,
		isComplete: false,
	});
	const handleTextInput = (event) => setForm({ ...form, [event.target.name]: event.target.value });
	const handleCheckboxInput = (event) => setForm({ ...form, [event.target.name]: event.target.checked });
	const handleSubmit = (event) => {
		event.preventDefault();
		onSubmit(form);
	};

	return (
		<Form onSubmit={handleSubmit} {...rest}>
			<FormGroup>
				<Label for="name">Name</Label>
				<Input id="name" name="name" type="text" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup>
				<Label for="description">Description</Label>
				<Input id="description" name="description" type="textarea" onChange={handleTextInput} />
			</FormGroup>
			<FormGroup check>
				<Label check>
					<Input type="checkbox" name="isHidden" onChange={handleCheckboxInput} />{' '}Hidden
							</Label>
			</FormGroup>
			<FormGroup check>
				<Label check>
					<Input type="checkbox" name="isComplete" onChange={handleCheckboxInput} />{' '}Complete
							</Label>
			</FormGroup>
			<Button onClick={onCancel}>Cancel</Button>
			<Input type="submit" value="Submit" />
		</Form>
	);
}

function Task({ content, onDelete }) {
	const [taskContent, setTaskContent] = useState(content);

	const toggleComplete = () => {
		async function invokeAPIRequest() {
			const updatedTask = await requestUpdateTask({ ...taskContent, isComplete: !taskContent.isComplete });
			setTaskContent(updatedTask);
		}
		invokeAPIRequest();
	}

	const toggleHidden = () => {
		async function invokeAPIRequest() {
			const updatedTask = await requestUpdateTask({ ...taskContent, isHidden: !taskContent.isHidden });
			setTaskContent(updatedTask);
		}
		invokeAPIRequest();
	}

	return (
		<tr>
			<td><Button onClick={onDelete} value={taskContent.id} >Delete</Button></td>
			<td>{taskContent.id}</td>
			<td>{taskContent.name}</td>
			<td>{taskContent.description}</td>
			<td><input type="checkbox" checked={taskContent.isComplete} onChange={toggleComplete} /></td>
			<td><input type="checkbox" checked={taskContent.isHidden} onChange={toggleHidden} /></td>
			<td>{taskContent.createdAt}</td>
			<td>{taskContent.editedAt}</td>
		</tr>
	);
}

async function requestPostUser(newUser) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/users', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(newUser)
		});
		if (response.status != 201) {
			throw await response.text();
		}
		localStorage.setItem('Authorization', response.headers.get('Authorization'));
		return response.json();
	} catch (error) {
		alert(error);
	}
}

async function requestPostSession(credentials) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/sessions', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(credentials)
		});
		if (response.status != 201) {
			throw await response.text();
		}
		localStorage.setItem('Authorization', response.headers.get('Authorization'));
		return response.json();
	} catch (error) {
		alert(error);
	}
}

async function requestDeleteSession() {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/sessions/mine', {
			method: 'DELETE',
			headers: {
				'Authorization': localStorage.getItem('Authorization')
			},
		});
		if (response.status != 200) {
			throw await response.text();
		}
		localStorage.removeItem('Authorization');
	} catch (error) {
		alert(error);
	}
}

async function requestGetTodoList() {
	try {
		let response
		if (!localStorage.getItem('Authorization')) {
			response = await fetch('https://api.thenightbeforeitsdue.de/tasks', {
				method: 'GET'
			});
			localStorage.setItem('Authorization', response.headers.get('Authorization'));
		} else {
			response = await fetch('https://api.thenightbeforeitsdue.de/tasks', {
				method: 'GET',
				headers: {
					'Authorization': localStorage.getItem('Authorization')
				},
			});
		}
		if (response.status != 200) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

async function requestPostTask(newTask) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/tasks', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('Authorization')
			},
			body: JSON.stringify(newTask)
		});
		if (response.status != 201) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

async function requestDeleteTask(taskID) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/tasks/' + taskID, {
			method: 'DELETE',
			headers: {
				'Authorization': localStorage.getItem('Authorization')
			}
		});
		if (response.status != 200) {
			throw await response.text();
		}
	} catch (error) {
		alert(error);
	}
}

async function requestUpdateTask(modifiedTask) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/tasks/' + modifiedTask.id, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('Authorization')
			},
			body: JSON.stringify(modifiedTask)
		});
		if (response.status != 200) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export default App;
