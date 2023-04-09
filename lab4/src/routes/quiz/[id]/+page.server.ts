import type { Answer, Quiz } from '../../../models/quiz';
import { TOKEN } from '$env/static/private';

export async function load({ cookies, fetch, params }) {
	const userId = cookies.get('userId');

	if (!userId) {
		return {
			authenticated: false
		};
	}

	const { id } = params;

	const quizState = cookies.get('quizState');

	if (quizState) {
		let quizStateObj = JSON.parse(quizState);
		if (quizStateObj[id] !== undefined) {
			return {
				authenticated: true,
				completed: true,
				score: quizStateObj[id]
			};
		}
	}

	const url = `https://late-glitter-4431.fly.dev/api/v54/quizzes/${id}`;
	const headers = {
		'X-Access-Token': TOKEN,
		'Content-Type': 'application/json'
	};

	const res = await fetch(url, { headers });
	const quiz: Quiz = await res.json();

	return {
		authenticated: true,
		quiz
	};
}

export const actions = {
	default: async ({ cookies, request, params }) => {
		const data = await request.formData();
		const userId = cookies.get('userId');

		if (!userId) {
			return;
		}

		let requests = [];
		for (let pair of data.entries()) {
			requests.push({
				data: {
					question_id: parseInt(pair[0]),
					answer: pair[1],
					user_id: parseInt(userId)
				}
			});
		}

		const { id } = params;
		const url = `https://late-glitter-4431.fly.dev/api/v54/quizzes/${id}/submit`;
		const headers = {
			'X-Access-Token': TOKEN,
			'Content-Type': 'application/json'
		};

		let score = 0;

		for (let request of requests) {
			const res = await fetch(url, {
				method: 'POST',
				headers,
				body: JSON.stringify(request)
			});

			if (res.status === 201) {
				const answer: Answer = await res.json();

				if (answer.correct) {
					score++;
				}
			}
		}

		let quizState = cookies.get('quizState');
		if (!quizState) {
			quizState = '{}';
		}

		let quizStateObj = JSON.parse(quizState);
		quizStateObj[id] = Math.round((score / requests.length) * 100);
		cookies.set('quizState', JSON.stringify(quizStateObj), { path: '/' });
	}
};
