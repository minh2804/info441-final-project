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
			const newUser = await requestPostSession(form);
			const newTodoList = await requestGetTodoList();
			onSignIn({ user: newUser, todoList: newTodoList });
		}
		invokeAPIRequest();
	};
	return (
		<Container {...rest}>
			<Form onSubmit={handleSubmit}>
				<TextInputGroup label="Username" name="username" onChange={handleTextInput} />
				<PasswordInputGroup label="Password" name="password" onChange={handleTextInput} />
				<Input type="submit" value="Sign in" />
			</Form >
		</Container>
	);
}

export function TextInputGroup({ label, name, onChange, placeholder, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for={name}>{label}</Label>
			<Input type="text" name={name} onChange={onChange} placeholder={placeholder} />
		</FormGroup>
	);
}

export function PasswordInputGroup({ label, name, onChange, placeholder, ...rest }) {
	return (
		<FormGroup {...rest}>
			<Label for={name}>{label}</Label>
			<Input type="password" name={name} onChange={onChange} placeholder={placeholder} />
		</FormGroup>
	);
}
