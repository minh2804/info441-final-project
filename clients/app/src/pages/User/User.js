import React, { useState } from 'react';
import {
	Container,
	Form,
	Input
} from 'reactstrap';
import { useHistory } from "react-router-dom";
import { requestPatchUser } from '../../api/users'
import { TextInputGroup } from '../SignIn/SignIn'

export default function User({ user, onChange, ...rest }) {
	const history = useHistory();
	const [meta, setMeta] = useState({
		id: user.id,
		username: user.username
	})
	const [form, setForm] = useState({
		firstName: user.firstName,
		lastName: user.lastName
	});
	const handleTextInput = (event) => setForm({ ...form, [event.target.name]: event.target.value });
	const handleSubmit = (event) => {
		event.preventDefault();
		async function invokeAPIRequest() {
			const newUser = await requestPatchUser(form);
			onChange(newUser);
			history.push('/home');
		}
		invokeAPIRequest();
	};
	return (
		<Container {...rest}>
			<Form onSubmit={handleSubmit} {...rest}>
				<div>ID: {meta.id}</div>
				<div>Username: {meta.username}</div>
				<TextInputGroup label="First Name" name="firstName" onChange={handleTextInput} placeholder={form.firstName} />
				<TextInputGroup label="Last Name" name="lastName" onChange={handleTextInput} placeholder={form.lastName} />
				<Input type="submit" value="Update information" />
			</Form >
		</Container>
	);
}
