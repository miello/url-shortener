<script lang="ts">
  import { Router, Route } from 'svelte-navigator'
  import routes from './utils/routes'
  import Alert from './components/common/Alert.svelte'
  import Navbar from './modules/Navbar.svelte'
</script>

<main class="min-h-screen bg flex">
  <div class="flex w-full">
    <Router>
      <Navbar />
      {#each routes as { path, Components }}
        <Route {path}>
          <Components />
        </Route>
      {/each}
    </Router>
    <Alert />
  </div>
</main>

<style lang="postcss" global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  @layer components {
    .bg {
      background: conic-gradient(
          from 180deg at 50% 50%,
          rgba(0, 255, 71, 0.4) -65.63deg,
          rgba(255, 77, 0, 0.4) 123.75deg,
          rgba(0, 255, 71, 0.4) 294.38deg,
          rgba(255, 77, 0, 0.4) 483.75deg
        ),
        linear-gradient(0deg, #ffffff, #ffffff);
    }
  }
</style>
