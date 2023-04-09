<script lang="ts">
	import type { QuizInfo } from '../../models/quiz';
	export let quizInfo: QuizInfo;
</script>

<div class="quiz-card">
	<h2>{quizInfo.title}</h2>
	<p>{quizInfo.questions_count} questions</p>
	<a href="/quiz/{quizInfo.id}">Start</a>
</div>

<style>
	.quiz-card {
		margin: 1rem;
		background-color: rgba(255, 255, 255, 0.7);
		padding: 1rem;
		border-radius: 0.5rem;
		box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.2);
	}

	h2 {
		margin-top: 0;
		margin-bottom: 0.5rem;
		font-size: 1.5rem;
	}

	p {
		margin-top: 0;
		margin-bottom: 0.5rem;
		font-size: 1.1rem;
	}

	a {
		display: inline-block;
		padding: 0.5rem 1rem;
		background-color: #ff3e00;
		color: #fff;
		text-decoration: none;
		border-radius: 5px;
	}

	a:hover {
		background-color: #ff6347;
	}
</style>
