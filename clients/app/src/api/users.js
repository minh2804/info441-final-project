import ENDPOINTS from '../constants/api-endpoints'
import HTTP from '../constants/http-requests'

export async function requestDeleteSession() {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.sessionsMine, {
			method: HTTP.methods.delete,
			headers: {
				[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
			}
		});
		if (response.status !== 200) {
			throw await response.text();
		}
		localStorage.removeItem(HTTP.headers.auth);
	} catch (error) {
		alert(error);
	}
}

export async function requestPatchUser(updates) {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.userMe, {
			method: HTTP.methods.patch,
			headers: {
				[HTTP.headers.contentType]: HTTP.headers.jsonType,
				[HTTP.headers.auth]: localStorage.getItem(HTTP.headers.auth)
			},
			body: JSON.stringify(updates)
		});
		if (response.status !== 200) {
			throw await response.text();
		}
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestPostSession(credentials) {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.sessions, {
			method: HTTP.methods.post,
			headers: {
				[HTTP.headers.contentType]: HTTP.headers.jsonType
			},
			body: JSON.stringify(credentials)
		});
		if (response.status !== 201) {
			throw await response.text();
		}
		localStorage.setItem(HTTP.headers.auth, response.headers.get(HTTP.headers.auth));
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestPostUser(newUser) {
	try {
		const response = await fetch(ENDPOINTS.base + ENDPOINTS.handlers.users, {
			method: HTTP.methods.post,
			headers: {
				[HTTP.headers.contentType]: HTTP.headers.jsonType
			},
			body: JSON.stringify(newUser)
		});
		if (response.status !== 201) {
			throw await response.text();
		}
		localStorage.setItem(HTTP.headers.auth, response.headers.get(HTTP.headers.auth));
		return response.json();
	} catch (error) {
		alert(error);
	}
}
