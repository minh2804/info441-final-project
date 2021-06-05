import React, { useState } from 'react';
import {
	Container,
	Form,
	FormGroup,
	Input,
	Label
} from 'reactstrap';
import {
	PasswordInputGroup,
	TextInputGroup
} from '../SignIn/SignIn'
import { requestGetTodoList } from '../../api/tasks'
import { requestPostUser } from '../../api/users'

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
			const newUser = await requestPostUser(form);
			const newTodoList = await requestGetTodoList();
			onSignUp({ user: newUser, todoList: newTodoList });
		}
		invokeAPIRequest();
	};
	return (
		<Container {...rest}>
			<Form onSubmit={handleSubmit}>
				<TextInputGroup label="Username" name="username" onChange={handleTextInput} />
				<PasswordInputGroup label="Password" name="password" onChange={handleTextInput} />
				<PasswordInputGroup label="Password Confirmation" name="passwordConf" onChange={handleTextInput} />
				<TextInputGroup label="First Name" name="firstName" onChange={handleTextInput} />
				<TextInputGroup label="Last Name" name="lastName" onChange={handleTextInput} />
				<Input type="submit" value="Sign up" />
			</Form>
		</Container>
	);
}
