import React from 'react';
import { Redirect } from "react-router-dom";
import { requestDeleteSession } from '../../api/users'
import { requestGetTodoList } from '../../api/tasks'

export default function SignOut({ onSignOut }) {
	async function invokeAPIRequest() {
		await requestDeleteSession();
		const todoListBody = await requestGetTodoList();
		onSignOut({ user: null, todoList: todoListBody });
	}
	invokeAPIRequest();
	return <Redirect push to="/" />;
}
