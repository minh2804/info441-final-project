import React, { useState } from 'react';
import {
	Container,
	Form,
	FormGroup,
	Input,
	Label
} from 'reactstrap';
import { Link } from "react-router-dom";
import Popup from 'reactjs-popup';
import {
	requestDeleteTask,
	requestGetImportTasks,
	requestPatchTask,
	requestPostTask
} from '../../api/tasks'
import { TextInputGroup } from '../SignIn/SignIn'
import ENDPOINTS from '../../constants/api-endpoints'

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
				<ImportPopUpButton className="btn btn-outline-dark m-2" onSubmit={onChange} >Import</ImportPopUpButton>
				<SharePopUpButton className="btn btn-outline-success m-2" user={user}>Share</SharePopUpButton>
				<AddTaskPopUpButton className="btn btn-outline-primary m-2" onSubmit={handleAddTask}>Add task</AddTaskPopUpButton>
			</div>
		</Container>
	);
}

function Task({ content, onDelete, ...rest }) {
	const [taskContent, setTaskContent] = useState(content);

	const toggleComplete = () => {
		async function invokeAPIRequest() {
			const updatedTask = await requestPatchTask({ ...taskContent, isComplete: !taskContent.isComplete });
			setTaskContent(updatedTask);
		}
		invokeAPIRequest();
	}
	const toggleHidden = () => {
		async function invokeAPIRequest() {
			const updatedTask = await requestPatchTask({ ...taskContent, isHidden: !taskContent.isHidden });
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

function ImportPopUpButton({ onSubmit, children, ...rest }) {
	const [open, setOpen] = useState(false);
	const toggleOpen = () => setOpen(!open);
	const closeModal = () => setOpen(false);
	const handleSubmit = (newTodoList) => {
		onSubmit(newTodoList);
		closeModal();
	}
	return (
		<div>
			<button type="button" onClick={toggleOpen} {...rest}>{children}</button>
			<Popup open={open} closeOnDocumentClick onClose={closeModal} >
				<ImportForm onSubmit={handleSubmit} onCancel={closeModal} className="container bg-white border border-secondary p-5" />
			</Popup >
		</div >
	);
}

function ImportForm({ onSubmit, onCancel, ...rest }) {
	const [form, setForm] = useState({ importLink: '' });
	const handleTextInput = (event) => setForm({ ...form, [event.target.name]: event.target.value });
	const handleSubmit = (event) => {
		event.preventDefault();
		if (form.importLink.startsWith(ENDPOINTS.base + ENDPOINTS.handlers.import)) {
			async function invokeAPIRequest() {
				const responseBody = await requestGetImportTasks(form.importLink);
				onSubmit(responseBody);
			}
			invokeAPIRequest();
		} else {
			alert('Invalid link provided');
		}
	};
	return (
		<Form onSubmit={handleSubmit} {...rest}>
			<TextInputGroup label="Import Link" name="importLink" onChange={handleTextInput} />
			<button type="button" className="btn btn-outline-secondary w-100 mb-2" onClick={onCancel}>Cancel</button>
			<Input className="btn btn-outline-primary" type="submit" value="Import list" />
		</Form>
	);
}

function SharePopUpButton({ user, children, ...rest }) {
	const [open, setOpen] = useState(false);
	const toggleOpen = () => setOpen(!open);
	const closeModal = () => setOpen(false);
	const renderedShareLink = user ? <AuthenticatedShare user={user} /> : <UnauthenticatedShare />;
	return (
		<div>
			<button type="button" onClick={toggleOpen} {...rest}>{children}</button>
			<Popup open={open} closeOnDocumentClick onClose={closeModal} >
				<div onCancel={closeModal} className="container bg-white border border-secondary p-4">
					{renderedShareLink}
				</div>
			</Popup >
		</div >
	);
}

function AuthenticatedShare({ user, ...rest }) {
	return (
		<div {...rest}>
			<p>Shareable link: <a href="url">{ENDPOINTS.base + ENDPOINTS.handlers.import + user.id}</a></p>
			<p>Any tasks you marked as hidden will not be shared.</p>
		</div>);
}

function UnauthenticatedShare(props) {
	return (
		<div {...props}>
			<div className="mb-2">You must be signed in to create a shareable link</div>
			<div className="d-flex justify-content-end">
				<Link type="button" className="btn btn-outline-secondary m-1" push to="/signup">Sign up</Link>
				<Link type="button" className="btn btn-outline-primary m-1" push to="/signin">Sign in</Link>
			</div>
		</div>
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
			<button type="button" onClick={toggleOpen} {...rest}>{children}</button>
			<Popup open={open} closeOnDocumentClick onClose={closeModal} >
				<NewTaskForm onSubmit={handleSubmit} onCancel={closeModal} className="container bg-white border border-secondary p-5" />
			</Popup >
		</div >
	);
}

function NewTaskForm({ onSubmit, onCancel, ...rest }) {
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
			<TextInputGroup label="Name" name="name" onChange={handleTextInput} />
			<TextInputGroup label="Description" name="description" onChange={handleTextInput} />
			<CheckboxInputGroup label="Hidden" name="isHidden" onChange={handleCheckboxInput} />
			<CheckboxInputGroup label="Complete" name="isComplete" onChange={handleCheckboxInput} />
			<button type="button" className="btn btn-outline-secondary w-100 mt-2" onClick={onCancel}>Cancel</button>
			<Input className="btn btn-outline-primary mt-2" type="submit" value="Add task" />
		</Form>
	);
}

export function CheckboxInputGroup({ label, name, onChange, ...rest }) {
	return (
		<FormGroup check {...rest}>
			<Label check>
				<Input type="checkbox" name={name} onChange={onChange} />{' ' + label}
			</Label>
		</FormGroup>
	);
}
