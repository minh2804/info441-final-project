export async function requestPostUser(newUser) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/users', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(newUser)
		});
		if (response.status !== 201) {
			throw await response.text();
		}
		localStorage.setItem('Authorization', response.headers.get('Authorization'));
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestUpdateUser(updates) {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/users/me', {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('Authorization')
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
		const response = await fetch('https://api.thenightbeforeitsdue.de/sessions', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(credentials)
		});
		if (response.status !== 201) {
			throw await response.text();
		}
		localStorage.setItem('Authorization', response.headers.get('Authorization'));
		return response.json();
	} catch (error) {
		alert(error);
	}
}

export async function requestDeleteSession() {
	try {
		const response = await fetch('https://api.thenightbeforeitsdue.de/sessions/mine', {
			method: 'DELETE',
			headers: {
				'Authorization': localStorage.getItem('Authorization')
			}
		});
		if (response.status !== 200) {
			throw await response.text();
		}
		localStorage.removeItem('Authorization');
	} catch (error) {
		alert(error);
	}
}
