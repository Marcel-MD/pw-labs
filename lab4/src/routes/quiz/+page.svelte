<script lang="ts">
	import type { QuizInfo } from '../../models/quiz';
	import Quiz from './Quiz.svelte';
	export let data: { quizzes: QuizInfo[] };
</script>

<svelte:head>
	<title>Quizzes</title>
	<meta name="description" content="Quizzes" />
</svelte:head>

<div class="text-column">
	{#each data.quizzes as quizInfo}
		<Quiz {quizInfo} />
	{/each}
</div>
