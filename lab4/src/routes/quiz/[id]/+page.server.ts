import type { Quiz } from '../../../models/quiz';
import { TOKEN } from '$env/static/private';

export async function load({ fetch, params }) {
	const { id } = params;
	const url = `https://late-glitter-4431.fly.dev/api/v54/quizzes/${id}`;
	const headers = {
		'X-Access-Token': TOKEN,
		'Content-Type': 'application/json'
	};

	const res = await fetch(url, { headers });
	const quiz: Quiz = await res.json();

	return { quiz };
}
