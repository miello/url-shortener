<script lang="ts">
  import { navigate } from 'svelte-navigator'
  import { UpdateAlert } from '../../stores/AlertStores'

  import { UserStores } from '../../stores/UserStores'
  import Loading from '../common/Loading.svelte'

  let isReady = false
  let handle = null

  UserStores.subscribe((val) => {
    if (val.ready && !val.handle) {
      UpdateAlert({
        status: 'error',
        message: 'Please Login First',
      })
      navigate('/login')
    }
    isReady = val.ready
    handle = val.handle
  })
</script>

{#if !isReady}
  <div class="flex justify-center items-center w-full">
    <Loading className="w-20 h-20 border-8" />
  </div>
{:else}
  <slot />
{/if}
