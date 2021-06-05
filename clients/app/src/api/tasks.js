import ENDPOINTS from '../constants/api-endpoints'
import HTTP from '../constants/http-requests'

export async function requestDeleteTask(taskID) {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.specificTask + taskID, {
			method: HTTP.methods.delete,
			headers: {
				[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
			}
		});
		if (response.status !== 200) {
			throw await response.text();
		}
	} catch (error) {
		alert(error);
	}
}

export async function requestGetImportTasks(url) {
	try {
		const response = await fetch(url, {
			method: HTTP.methods.get,
			headers: {
				[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
			}
		});
		if (response.status !== 201) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestGetTodoList() {
	try {
		let response
		if (!localStorage.getItem(HTTP.headers.auth)) {
			response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.tasks, {
				method: HTTP.methods.get
			});
			localStorage.setItem(HTTP.headers.auth, response.headers.get(HTTP.headers.auth));
		} else {
			response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.tasks, {
				method: HTTP.methods.get,
				headers: {
					[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
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

export async function requestPatchTask(modifiedTask) {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.specificTask + modifiedTask.id, {
			method: HTTP.methods.patch,
			headers: {
				[HTTP.headers.contentType]: HTTP.headers.jsonType,
				[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
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

export async function requestPostTask(newTask) {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.tasks, {
			method: HTTP.methods.post,
			headers: {
				[HTTP.headers.contentType]: HTTP.headers.jsonType,
				[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
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
