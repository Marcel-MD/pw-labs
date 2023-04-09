import type { User } from '../../models/user';
import { TOKEN } from '$env/static/private';

export function load({ cookies }) {
	const id = cookies.get('userId');
	const name = cookies.get('userName');

	if (!id || !name) {
		return {
			authenticated: false,
			user: null
		};
	}

	let user = {
		id: parseInt(id),
		name: name
	};

	return {
		authenticated: true,
		user: user
	};
}

export const actions = {
	create: async ({ cookies, request }) => {
		const data = await request.formData();

		const url = `https://late-glitter-4431.fly.dev/api/v54/users`;
		const headers = {
			'X-Access-Token': TOKEN,
			'Content-Type': 'application/json'
		};
		const body = JSON.stringify({
			data: {
				name: data.get('name'),
				surname: data.get('surname')
			}
		});

		const res = await fetch(url, { headers, body, method: 'POST' });

		if (res.status === 201) {
			const user: User = await res.json();

			cookies.set('userId', user.id.toString());
			cookies.set('userName', user.name + ' ' + user.surname);
		}
	},

	delete: async ({ cookies, request }) => {
		const data = await request.formData();
		const id = cookies.get('userId');

		const url = `https://late-glitter-4431.fly.dev/api/v54/users/${id}`;

		const headers = {
			'X-Access-Token': TOKEN,
			'Content-Type': 'application/json'
		};

		const res = await fetch(url, { headers, method: 'DELETE' });

		if (res.status === 200) {
			cookies.delete('userId');
			cookies.delete('userName');
			cookies.delete('quizState', { path: '/' });
		}
	}
};
