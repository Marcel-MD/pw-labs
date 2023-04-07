import type { QuizInfo } from '../../models/quiz';
import { TOKEN } from '$env/static/private';

export async function load({ fetch, params }) {
	const url = `https://late-glitter-4431.fly.dev/api/v54/quizzes/`;
	const headers = {
		'X-Access-Token': TOKEN,
		'Content-Type': 'application/json'
	};

	const res = await fetch(url, { headers });
	const quizzes: QuizInfo[] = await res.json();

	return { quizzes };
}
