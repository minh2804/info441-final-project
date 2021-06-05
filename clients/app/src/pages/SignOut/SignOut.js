import React from 'react';
import { Redirect } from "react-router-dom";
import { requestDeleteSession } from '../../api/users'
import { requestGetTodoList } from '../../api/tasks'

export default function SignOut({ onSignOut }) {
	async function invokeAPIRequest() {
		await requestDeleteSession();
		const newTodoList = await requestGetTodoList();
		onSignOut({ user: null, todoList: newTodoList });
	}
	invokeAPIRequest();
	return <Redirect push to="/" />;
}
