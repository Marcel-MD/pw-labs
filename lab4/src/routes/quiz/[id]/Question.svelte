<script lang="ts">
	import type { Question } from '../../../models/quiz';
	export let question: Question;
</script>

<div class="question-card">
	<h2>{question.question}</h2>
	{#each question.answers as answer}
		<input type="radio" id={answer + question.id} name={question.question} value={answer} />
		<label for={answer + question.id}>{answer}</label>
	{/each}
</div>

<style>
	.question-card {
		margin: 1rem;
		background-color: rgba(255, 255, 255, 0.7);
		padding: 1rem;
		border-radius: 0.5rem;
		box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.2);
	}

	h2 {
		margin-top: 0;
		margin-bottom: 1rem;
		font-size: 1.5rem;
	}

	input[type='radio'] {
		display: none;
	}

	label {
		display: block;
		position: relative;
		padding-left: 2rem;
		margin-bottom: 0.5rem;
		cursor: pointer;
		font-size: 1.1rem;
	}

	label:before {
		content: '';
		display: block;
		position: absolute;
		left: 0;
		top: 0;
		width: 1.25rem;
		height: 1.25rem;
		border: 2px solid #ccc;
		border-radius: 50%;
		background-color: transparent;
		transition: all 0.3s ease-in-out;
	}

	input[type='radio']:checked + label:before {
		background-color: #ff3e00;
	}

	label:hover:before {
		border-color: #ff3e00;
	}

	label:active:before {
		transform: scale(0.8);
	}
</style>
