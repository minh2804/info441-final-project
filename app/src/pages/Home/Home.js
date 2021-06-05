import React, { useState } from 'react';
import Popup from 'reactjs-popup';
import {
	Button,
	Container,
	Form,
	FormGroup,
	Input,
	Label
} from 'reactstrap';
import {
	requestDeleteTask,
	requestPostTask,
	requestUpdateTask
} from '../../api/tasks'

export default function Home({ user, todoList, onChange, ...rest }) {
	const handleAddTask = (task) => {
		async function invokeAPIRequest() {
			const responseBody = await requestPostTask(task);
			onChange([...todoList, responseBody]);
		}
		invokeAPIRequest();
	};
	const handleDeleteTask = (event) => {
		const taskID = parseInt(event.target.value);
		async function invokeAPIRequest() {
			await requestDeleteTask(taskID);
			onChange(todoList.filter(task => task.id !== taskID));
		}
		invokeAPIRequest();
	};
	const renderedRows = todoList.map(task => <Task className="border-bottom" key={task.id} content={task} onDelete={handleDeleteTask} />);
	return (
		<Container {...rest}>
			<table className="w-100 mb-3">
				<thead>
					<tr className="border-bottom">
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
			<div className="d-flex justify-content-center" >
				<ShareListPopUpButton className="m-2" user={user}>Share</ShareListPopUpButton>
				<AddTaskPopUpButton className="m-2" onSubmit={handleAddTask}>Add Task</AddTaskPopUpButton>
			</div>
		</Container>
	);
}

export function Task({ content, onDelete, ...rest }) {
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
		<tr {...rest}>
			<td><button type="button" className="btn btn-outline-danger btn-sm mb-1" onClick={onDelete} value={taskContent.id} >Delete task: {taskContent.id}</button></td>
			<td>{taskContent.name}</td>
			<td>{taskContent.description}</td>
			<td><input type="checkbox" checked={taskContent.isComplete} onChange={toggleComplete} /></td>
			<td><input type="checkbox" checked={taskContent.isHidden} onChange={toggleHidden} /></td>
			<td>{taskContent.createdAt}</td>
			<td>{taskContent.editedAt}</td>
		</tr>
	);
}

export function ShareListPopUpButton({ user, children, ...rest }) {
	const [open, setOpen] = useState(false);
	const toggleOpen = () => setOpen(!open);
	const closeModal = () => setOpen(false);
	return (
		<div {...rest}>
			<button type="button" className="btn btn-outline-success" onClick={toggleOpen}>{children}</button>
			<Popup open={open} closeOnDocumentClick onClose={closeModal} >
				<div onCancel={closeModal} className="container bg-white border border-secondary p-5">
					Shareable link: <a href="url">{'https://api.thenightbeforeitsdue.de/tasks/import/' + user.id}</a>
				</div>
			</Popup >
		</div >
	);
}

export function AddTaskPopUpButton({ onSubmit, children, ...rest }) {
	const [open, setOpen] = useState(false);
	const toggleOpen = () => setOpen(!open);
	const closeModal = () => setOpen(false);
	const handleSubmit = (newTask) => {
		onSubmit(newTask);
		closeModal();
	}
	return (
		<div {...rest}>
			<button type="button" className="btn btn-outline-primary" onClick={toggleOpen}>{children}</button>
			<Popup open={open} closeOnDocumentClick onClose={closeModal} >
				<NewTaskForm onSubmit={handleSubmit} onCancel={closeModal} className="container bg-white border border-secondary p-5" />
			</Popup >
		</div >
	);
}

export function NewTaskForm({ onSubmit, onCancel, ...rest }) {
	const [form, setForm] = useState({
		name: "",
		description: "",
		isHidden: false,
		isComplete: false
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
			<button type="button" className="btn btn-outline-secondary w-100 mb-2" onClick={onCancel}>Cancel</button>
			<Input className="btn btn-outline-primary" type="submit" value="Add task" />
		</Form>
	);
}
