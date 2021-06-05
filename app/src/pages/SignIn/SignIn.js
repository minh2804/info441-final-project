import React, { useState } from 'react';
import {
	Container,
	Form,
	FormGroup,
	Input,
	Label
} from 'reactstrap';
import { requestGetTodoList } from '../../api/tasks'
import { requestPostSession } from '../../api/users'

export default function SignIn({ onSignIn, ...rest }) {
	const [form, setForm] = useState({
		username: "",
		password: ""
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
		<Container {...rest}>
			<Form onSubmit={handleSubmit}>
				<UsernameInputGroup onChange={handleTextInput} />
				<PasswordInputGroup onChange={handleTextInput} />
				<Input type="submit" value="Sign in" />
			</Form >
		</Container>
	);
}

export function UsernameInputGroup({ onChange, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for="username">Username</Label>
			<Input id="username" name="username" type="text" onChange={onChange} />
		</FormGroup>
	);
}

export function PasswordInputGroup({ onChange, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for="password">Password</Label>
			<Input id="password" name="password" type="password" onChange={onChange} />
		</FormGroup>
	);
}
