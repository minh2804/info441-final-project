import React, { useState } from 'react';
import {
	Container,
	Form,
	FormGroup,
	Input,
	Label
} from 'reactstrap';
import { requestUpdateUser } from '../../api/users'

export default function User({ user, onChange, ...rest }) {
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
			const newUserBody = await requestUpdateUser(form);
			onChange(newUserBody);
		}
		invokeAPIRequest();
	};
	return (
		<Container {...rest}>
			<Form onSubmit={handleSubmit} {...rest}>
				<div>ID: {meta.id}</div>
				<div>Username: {meta.username}</div>
				<FormGroup>
					<Label for="firstName">First Name</Label>
					<Input id="firstName" name="firstName" type="text" placeholder={form.firstName} onChange={handleTextInput} />
				</FormGroup>
				<FormGroup>
					<Label for="lastName">Last Name</Label>
					<Input id="lastName" name="lastName" type="text" placeholder={form.lastName} onChange={handleTextInput} />
				</FormGroup>
				<Input type="submit" value="Update user information" />
			</Form >
		</Container>
	);
}
