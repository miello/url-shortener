<script lang="ts">
	import { v4 as uuidV4} from 'uuid'
	import { onMount } from "svelte";
	import { Router, Route } from 'svelte-navigator'
	import routes from './utils/routes';

	onMount(() => {
    if(!localStorage.getItem('uuid')) {
      localStorage.setItem('uuid', uuidV4())
    }
  })

</script>

<main class="min-h-screen bg flex">
	<div class="flex w-full">
		<Router>
			{#each routes as { path, Components}}
				<Route path={path}>
					<Components />
				</Route>
			{/each}
		</Router>
	</div>
</main>

<style lang="postcss" global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;

	@layer components {
		.bg {
			background: conic-gradient(from 180deg at 50% 50%, rgba(0, 255, 71, 0.4) -65.63deg, rgba(255, 77, 0, 0.4) 123.75deg, rgba(0, 255, 71, 0.4) 294.38deg, rgba(255, 77, 0, 0.4) 483.75deg), linear-gradient(0deg, #FFFFFF, #FFFFFF);
		}
	}
</style>