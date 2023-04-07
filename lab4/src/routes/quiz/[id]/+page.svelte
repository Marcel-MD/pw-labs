<script lang="ts">
	import type { Quiz } from '../../../models/quiz';
	import Question from './Question.svelte';
	export let data: { quiz: Quiz };
</script>

<svelte:head>
	<title>Quiz</title>
	<meta name="description" content="Quiz" />
</svelte:head>

<div class="text-column">
	<h1>{data.quiz.title}</h1>
	{#each data.quiz.questions as question}
		<Question {question} />
	{/each}
</div>
