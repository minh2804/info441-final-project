import React, { useState } from 'react';
import {
	Container,
	Form,
	FormGroup,
	Input,
	Label
} from 'reactstrap';
import { requestGetTodoList } from '../../api/tasks'
import { requestPostUser } from '../../api/users'
import {
	PasswordInputGroup,
	UsernameInputGroup
} from '../SignIn/SignIn'

export default function SignUp({ onSignUp, ...rest }) {
	const [form, setForm] = useState({
		username: "",
		password: "",
		passwordConf: "",
		firstName: "",
		lastName: "",
		isTemporary: false
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
		<Container {...rest}>
			<Form onSubmit={handleSubmit}>
				<UsernameInputGroup onChange={handleTextInput} />
				<PasswordInputGroup onChange={handleTextInput} />
				<PasswordConfInputGroup onChange={handleTextInput} />
				<FirstNameInputGroup onChange={handleTextInput} />
				<LastNameInputGroup onChange={handleTextInput} />
				<Input type="submit" value="Sign up" />
			</Form>
		</Container>
	);
}

export function PasswordConfInputGroup({ onChange, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for="passwordConf">Password Confirm</Label>
			<Input id="passwordConf" name="passwordConf" type="password" onChange={onChange} />
		</FormGroup>
	);
}

export function FirstNameInputGroup({ onChange, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for="firstName">First Name</Label>
			<Input id="firstName" name="firstName" type="text" onChange={onChange} />
		</FormGroup>
	);
}

export function LastNameInputGroup({ onChange, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for="lastName">Last Name</Label>
			<Input id="lastName" name="lastName" type="text" onChange={onChange} />
		</FormGroup>
	);
}
