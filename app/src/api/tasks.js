export async function requestGetTodoList() {
	try {
		let response
		if (!localStorage.getItem('Authorization')) {
			response = await fetch('https://api.thenightbeforeitsdue.de/tasks', {
				method: 'GET'
			});
			localStorage.setItem('Authorization', response.headers.get('Authorization'));
		} else {
			response = await fetch('https://api.thenightbeforeitsdue.de/tasks', {
				method: 'GET',
				headers: {
					'Authorization': localStorage.getItem('Authorization')
				}
			});
		}
		if (response.status !== 200) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestPostTask(newTask) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/tasks', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('Authorization')
			},
			body: JSON.stringify(newTask)
		});
		if (response.status !== 201) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestDeleteTask(taskID) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/tasks/' + taskID, {
			method: 'DELETE',
			headers: {
				'Authorization': localStorage.getItem('Authorization')
			}
		});
		if (response.status !== 200) {
			throw await response.text();
		}
	} catch (error) {
		alert(error);
	}
}

export async function requestUpdateTask(modifiedTask) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/tasks/' + modifiedTask.id, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('Authorization')
			},
			body: JSON.stringify(modifiedTask)
		});
		if (response.status !== 200) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}
