<script lang="ts">
	export let data;
	let { authenticated, user } = data;
</script>

<svelte:head>
	<title>Register</title>
	<meta name="description" content="Register" />
</svelte:head>

<div class="text-column">
	{#if authenticated}
		<h1>Register</h1>
		<p>You are already logged in as <span>{user?.name} {user?.surname}</span>.</p>

		<form method="POST" action="?/delete">
			<input type="submit" value="Logout" />
		</form>
	{/if}
	{#if !authenticated}
		<h1>Register</h1>

		<form method="POST" action="?/create">
			<input type="text" name="name" autocomplete="off" placeholder="Name" required />
			<input type="text" name="surname" autocomplete="off" placeholder="Surname" required />
			<input type="submit" value="Submit" />
		</form>
	{/if}
</div>

<style>
	p {
		font-size: 1.5rem;
		margin-bottom: 2rem;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	input[type='text'] {
		padding: 0.5rem;
		font-size: 1.2rem;
		border: 2px solid #ccc;
		border-radius: 5px;
		margin-bottom: 1rem;
		transition: border-color 0.3s ease-in-out;
	}

	input[type='text']:focus {
		border-color: #ff3e00;
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
</style>
