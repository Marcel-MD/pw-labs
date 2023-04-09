<script lang="ts">
	import Question from './Question.svelte';
	export let data;
	let { quiz, authenticated, completed, score } = data;
</script>

<svelte:head>
	<title>Quiz</title>
	<meta name="description" content="Quiz" />
</svelte:head>

<div class="text-column">
	{#if !authenticated && !quiz}
		<h1>Quiz</h1>
		<p>You have to register first.</p>
	{/if}
	{#if authenticated && completed}
		<h1>Quiz Completed</h1>
		<p>Your score is <span>{score}%</span></p>
	{/if}
	{#if authenticated && quiz && !completed}
		<h1>{quiz.title}</h1>
		<form method="post">
			{#each quiz.questions as question}
				<Question {question} />
			{/each}
			<div class="button-container">
				<input type="submit" value="Submit" />
			</div>
		</form>
	{/if}
</div>

<style>
	p {
		font-size: 1.5rem;
		margin: 0 auto;
	}

	span {
		color: #ff3e00;
	}

	input[type='submit'] {
		padding: 0.5rem 1.5rem;
		font-size: 1.2rem;
		background-color: #ff3e00;
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
	}

	input[type='submit']:hover {
		background-color: #ff6347;
	}

	.button-container {
		display: flex;
		justify-content: center;
		margin-top: 2rem;
	}
</style>
