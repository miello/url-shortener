<script lang="ts">
  import { onMount } from "svelte";
  import { Link } from "svelte-navigator";
  import { GetUser } from "../stores/UserStores";
  import Logo from "../components/icons/Logo.svelte";
  import Container from "../components/common/Container.svelte";
  import { UserStores } from '../stores/UserStores'
  import Loading from "../components/common/Loading.svelte";
import Avatar from "../components/common/Avatar.svelte";
  
  onMount(() => {
    GetUser()
  })
</script>

<Container rounded={false} className="flex justify-between items-center w-full absolute py-1">
  <Link to="/">
    <Logo />
  </Link>
  {#if $UserStores.ready}
    {#if !$UserStores.handle}
      <div class="flex items-center">
        <Link to="/login">
          <h6 class="mr-5 font-display hover:underline underline-offset-2">Login</h6>
        </Link>
        <Link to="/register">
          <h6 class="mr-5 font-display hover:underline underline-offset-2">Register</h6>
        </Link>
      </div>
    {:else}
      <Avatar name={$UserStores.handle} />
    {/if}
  {:else}
      <Loading />
  {/if}

</Container>